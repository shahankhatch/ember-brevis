package sdk

import (
	"fmt"
	"github.com/consensys/gnark/frontend"
	"math/big"
)

type DataStream[T CircuitVariable] struct {
	api        *CircuitAPI
	underlying []T
	toggles    []frontend.Variable
}

func NewDataStream[T CircuitVariable](api *CircuitAPI, in DataPoints[T]) *DataStream[T] {
	return &DataStream[T]{
		api:        api,
		underlying: in.Raw,
		toggles:    in.Toggles,
	}
}

func newDataStream[T CircuitVariable](api *CircuitAPI, in []T, toggles []frontend.Variable) *DataStream[T] {
	return &DataStream[T]{
		api:        api,
		underlying: in,
		toggles:    toggles,
	}
}

// GetUnderlying gets an element from the DataStream. Performed on the underlying data
// directly. It also requires the underlying data slot is valid
func GetUnderlying[T CircuitVariable](ds *DataStream[T], index int) T {
	v := ds.underlying[index]
	t := ds.toggles[index]
	ds.api.g.AssertIsEqual(t, 1)
	return v
}

// RangeUnderlying selects a range of the data stream. Performed on the underlying data directly.
func RangeUnderlying[T CircuitVariable](ds *DataStream[T], start, end int) *DataStream[T] {
	return newDataStream(ds.api, ds.underlying[start:end], ds.toggles[start:end])
}

// WindowUnderlying splits a DataStream into many equal sized List. Performed on the
// underlying data directly. Panics if `size` does not divide the length of the
// underlying list. Use Range to cut the list length into a multiple of `size`
// first
func WindowUnderlying[T CircuitVariable](ds *DataStream[T], size int) *DataStream[List[T]] {
	l := len(ds.underlying)
	if l%size != 0 {
		panic(fmt.Errorf("cannot Window on DataStream of size %d: %d mod %d != 0", l, l, size))
	}
	var toggles []frontend.Variable
	var ret []List[T]
	for i := 0; i < l-size; i += size {
		start := i
		end := start + size
		ret = append(ret, ds.underlying[start:end])
		var toggle frontend.Variable = 0
		for _, t := range ds.toggles[start:end] {
			toggle = ds.api.g.And(toggle, t)
		}
		toggles = append(toggles, toggle)
	}
	return newDataStream(ds.api, ret, toggles)
}

type AssertFunc[T CircuitVariable] func(current T) Uint248

func AssertEach[T CircuitVariable](ds *DataStream[T], assertFunc AssertFunc[T]) {
	for i, data := range ds.underlying {
		pass := assertFunc(data).Val
		valid := ds.api.isEqual(ds.toggles[i], 1)
		pass = ds.api.g.Select(valid, pass, 1)
		ds.api.g.AssertIsEqual(pass, 1)
	}
}

// SortFunc returns 1 if a, b are sorted, 0 if not.
type SortFunc[T CircuitVariable] func(a, b T) Uint248

// IsSorted returns 1 if the data stream is sorted to the criteria of sortFunc, 0 if not.
func IsSorted[T CircuitVariable](ds *DataStream[T], sortFunc SortFunc[T]) Uint248 {
	// The following code uses prev and prevValid to pass the signal of last known
	// valid element of the data stream. This is needed because the stream could have
	// already been filtered, meaning we could have "gaps" between valid elements
	//
	//TODO:
	// we could use a bool in ds to indicate whether the toggles this ds has been
	// touched (the stream has been filtered) before this part of the user circuit
	// where this method is called. if it has not been touched, we probably don't
	// need to use prev and prevValid signals.
	api := ds.api.g
	var sorted frontend.Variable
	prev := ds.underlying[0]
	prevValid := ds.toggles[0]

	for i := 1; i < len(ds.underlying); i++ {
		curr := ds.underlying[i]
		currValid := ds.toggles[i]

		sorted = sortFunc(prev, curr).Val
		sorted = api.Select(api.And(prevValid, currValid), sorted, 1)

		prev = Select(ds.api, newU248(currValid), curr, prev)
		prevValid = currValid
	}
	return newU248(sorted)
}

// AssertSorted Performs the sortFunc on each valid pair of data points and assert the result to be 1.
func AssertSorted[T CircuitVariable](ds *DataStream[T], sortFunc SortFunc[T]) {
	ds.api.Uint248.AssertIsEqual(IsSorted(ds, sortFunc), newU248(1))
}

// Count returns the number of valid elements (i.e. toggled on) in the data stream.
func Count[T CircuitVariable](ds *DataStream[T]) Uint248 {
	t := ds.toggles
	count := ds.api.g.Add(t[0], t[1], t[2:]...)
	return newU248(count)
}

type GetValueFunc[T any] func(current T) Uint248

func GroupBy[T, R CircuitVariable](ds *DataStream[T], f ReduceFunc[T, R], reduceInit R, getValue GetValueFunc[T]) (*DataStream[R], error) {
	g := ds.api.g
	values := make([]frontend.Variable, len(ds.underlying))
	for i, v := range ds.underlying {
		values[i] = getValue(v)
	}
	groupValues, err := computeGroupValuesHint(g, values, ds.toggles)
	if err != nil {
		return nil, err
	}
	aggResults := make([]R, len(groupValues))
	aggResultToggles := make([]frontend.Variable, len(aggResults))
	for i, p := range groupValues {
		group := Filter(ds, func(current T) Uint248 {
			v := getValue(current)
			return newU248(g.IsZero(g.Sub(v, p)))
		})
		aggResults[i] = Reduce(group, reduceInit, f)
		aggResultToggles[i] = g.Sub(1, g.IsZero(p))
	}
	return newDataStream(ds.api, aggResults, aggResultToggles), nil
}

func computeGroupValuesHint(api frontend.API, values, toggles []frontend.Variable) ([]frontend.Variable, error) {
	inputs := []frontend.Variable{len(values)}
	inputs = append(inputs, values...)
	inputs = append(inputs, toggles...)
	return api.Compiler().NewHint(GroupValuesHint, len(values), inputs...)
}

func GroupValuesHint(_ *big.Int, inputs []*big.Int, outputs []*big.Int) error {
	numValues := inputs[0].Int64()
	values := inputs[1 : 1+numValues]
	toggles := inputs[1+numValues:]

	var uniqueValues []*big.Int
	for i, v := range values {
		if toggles[i].Sign() == 0 {
			continue
		}
		found := false
		for _, uv := range uniqueValues {
			if uv.Cmp(v) == 1 {
				found = true
			}
		}
		if !found {
			uniqueValues = append(uniqueValues, v)
		}
	}
	for i := len(uniqueValues); i < int(numValues); i++ {
		uniqueValues = append(uniqueValues, big.NewInt(0))
	}
	copy(outputs, uniqueValues)
	return nil
}

type MapFunc[T, R CircuitVariable] func(current T) R

func Map[T, R CircuitVariable](ds *DataStream[T], mapFunc MapFunc[T, R]) *DataStream[R] {
	res := make([]R, len(ds.underlying))
	for i, data := range ds.underlying {
		res[i] = mapFunc(data)
	}
	return newDataStream(ds.api, res, ds.toggles)
}

type ReduceFunc[T, R CircuitVariable] func(accumulator R, current T) (newAccumulator R)

// Reduce reduces the data stream to another CircuitVariable
func Reduce[T, R CircuitVariable](ds *DataStream[T], initial R, reduceFunc ReduceFunc[T, R]) R {
	var acc = initial
	for i, data := range ds.underlying {
		newAcc := reduceFunc(acc, data)
		oldAccVals := acc.Values()
		values := make([]frontend.Variable, len(oldAccVals))
		for j, newAccV := range newAcc.Values() {
			values[j] = Select(ds.api, newU248(ds.toggles[i]), newU248(newAccV), newU248(oldAccVals[j]))
		}
		acc.SetValues(values...)
	}
	return acc
}

// FilterFunc must return 1/0 to include/exclude `current` in the filter result
type FilterFunc[T CircuitVariable] func(current T) Uint248

func Filter[T CircuitVariable](ds *DataStream[T], filterFunc FilterFunc[T]) *DataStream[T] {
	api := ds.api.g
	newToggles := make([]frontend.Variable, len(ds.underlying))
	for i, data := range ds.underlying {
		toggle := filterFunc(data).Val
		valid := ds.api.isEqual(ds.toggles[i], 1)
		newToggles[i] = api.Select(api.And(toggle, valid), 1, 0)
	}
	return newDataStream(ds.api, ds.underlying, newToggles)
}

// MinGeneric finds out the minimum value of the selected field from the data stream. Uses Reduce under the hood.
func MinGeneric[T CircuitVariable](ds *DataStream[T], initialMin T, lt SortFunc[T]) T {
	return Reduce(ds, initialMin, func(min, current T) (newMin T) {
		curLtMin := lt(current, min)
		return Select(ds.api, curLtMin, current, min)
	})
}

// MaxGeneric finds out the maximum value of the selected field from the data stream. Uses Reduce under the hood.
func MaxGeneric[T CircuitVariable](ds *DataStream[T], initialMax T, gt SortFunc[T]) T {
	return Reduce(ds, initialMax, func(max, current T) (newMax T) {
		curLtMin := gt(current, max)
		return Select(ds.api, curLtMin, current, max)
	})
}

func Min(ds *DataStream[Uint248]) Uint248 {
	return MinGeneric(ds, newU248(0), func(a, b Uint248) Uint248 {
		return ds.api.Uint248.IsLessThan(a, b)
	})
}

func Max(ds *DataStream[Uint248]) Uint248 {
	return MaxGeneric(ds, newU248(MaxUint248), func(a, b Uint248) Uint248 {
		return ds.api.Uint248.IsGreaterThan(a, b)
	})
}

// Sum sums values of the selected field in the data stream. Uses Reduce.
func Sum(ds *DataStream[Uint248]) Uint248 {
	return Reduce(ds, newU248(0), func(sum Uint248, curr Uint248) (newSum Uint248) {
		return ds.api.Uint248.Add(sum, curr)
	})
}

// Mean calculates the arithmetic mean over the selected fields of the data stream. Uses Sum.
func Mean(ds *DataStream[Uint248]) Uint248 {
	sum := Sum(ds)
	quo, _ := ds.api.Uint248.Div(sum, Count(ds))
	return quo
}
