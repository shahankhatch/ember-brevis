package sdk

import (
	"bytes"
	"fmt"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
)

// returns little endian bits of data
func decomposeBits(data *big.Int, length uint) []uint {
	return decompose[uint](data, 1, length)
}

func recompose[T uint | byte](data []T, bitSize int) *big.Int {
	d := make([]*big.Int, len(data))
	for i := 0; i < len(data); i++ {
		d[i] = big.NewInt(int64(data[i]))
	}
	return recomposeBig(d, bitSize)
}

func recomposeBig(data []*big.Int, bitSize int) *big.Int {
	r := big.NewInt(0)
	for i := 0; i < len(data); i++ {
		r.Add(r, new(big.Int).Lsh(data[i], uint(i*bitSize)))
		r.Mod(r, ecc.BLS12_377.ScalarField())
	}
	return r
}

func decompose[T uint | byte](data *big.Int, bitSize uint, length uint) []T {
	res := decomposeBig(data, bitSize, length)
	ret := make([]T, length)
	for i, limb := range res {
		ret[i] = T(limb.Uint64())
	}
	return ret
}

func decomposeBig(data *big.Int, bitSize, length uint) []*big.Int {
	if uint(data.BitLen()) > length*bitSize {
		panic(fmt.Errorf("decomposed integer (bit len %d) does not fit into output (bit len %d, length %d)",
			data.BitLen(), bitSize, length))
	}
	decomposed := make([]*big.Int, length)
	base := new(big.Int).Lsh(big.NewInt(1), bitSize)
	d := new(big.Int).Set(data)
	for i := 0; i < int(length); i++ {
		rem := new(big.Int)
		d.DivMod(d, base, rem)
		decomposed[i] = rem
	}
	return decomposed
}

func packBitsToInt(bits []uint, bitSize int) []*big.Int {
	var r []*big.Int
	for i := 0; i < len(bits); i += bitSize {
		end := i + bitSize
		if end > len(bits) {
			end = len(bits)
		}
		bs := bits[i:end]
		z := recompose(bs, 1)
		r = append(r, z)
	}
	return r
}

// flips the order of the groups of groupSize. e.g. [1,2,3,4,5,6] with groupSize 2 is flipped to [5,6,3,4,1,2]
func flipByGroups[T any](in []T, groupSize int) []T {
	res := make([]T, len(in))
	copy(res, in)
	for i := 0; i < len(res)/groupSize/2; i++ {
		for j := 0; j < groupSize; j++ {
			a := i*groupSize + j
			b := len(res) - (i+1)*groupSize + j
			res[a], res[b] = res[b], res[a]
		}
	}
	return res
}

func newVars[T any](vs []T) []frontend.Variable {
	ret := make([]frontend.Variable, len(vs))
	for i, v := range vs {
		ret[i] = v
	}
	return ret
}

// copied from
// https://github.com/Consensys/gnark/blob/5711c4ae475535ce2a0febdeade86ff98914a378/internal/utils/convert.go#L39C1-L39C1
// with minor changes
func fromInterface(input interface{}) *big.Int {
	if input == nil {
		return big.NewInt(0)
	}
	in := input.(interface{})
	var r big.Int
	switch v := in.(type) {
	case Uint248:
		r.Set(fromInterface(v.Val))
	case big.Int:
		r.Set(&v)
	case *big.Int:
		r.Set(v)
	case uint8:
		r.SetUint64(uint64(v))
	case uint16:
		r.SetUint64(uint64(v))
	case uint32:
		r.SetUint64(uint64(v))
	case uint64:
		r.SetUint64(v)
	case uint:
		r.SetUint64(uint64(v))
	case int8:
		r.SetInt64(int64(v))
	case int16:
		r.SetInt64(int64(v))
	case int32:
		r.SetInt64(int64(v))
	case int64:
		r.SetInt64(v)
	case int:
		r.SetInt64(int64(v))
	case bool:
		var b uint64
		if v {
			b = 1
		}
		r.SetUint64(b)
	case string:
		if _, ok := r.SetString(v, 0); !ok {
			panic("unable to set big.Int from string " + v)
		}
	case common.Address:
		r.SetBytes(v[:])
	case []byte:
		r.SetBytes(v)
	}
	return &r
}

func mustWriteToBytes(w io.WriterTo) []byte {
	b := bytes.NewBuffer([]byte{})
	_, err := w.WriteTo(b)
	if err != nil {
		panic(fmt.Errorf("failed to write vk to bytes stream %s", err.Error()))
	}
	return b.Bytes()
}

func parseBitStr(s string) []frontend.Variable {
	ret := make([]frontend.Variable, len(s))
	for i, c := range s {
		if c == '0' {
			ret[i] = 0
		} else {
			ret[i] = 1
		}
	}
	return ret
}
