package ember

import (
	"github.com/brevis-network/brevis-sdk/sdk"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type AppCircuit struct {
	// custom input fields
	// UserAddr sdk.Uint248
}

var _ sdk.AppCircuit = &AppCircuit{}


var EventIdLowVolatilitySwap = sdk.ParseEventID(
	hexutil.MustDecode("0x104cd836999a92d13674bebdf353619d8fb00747e034b2c9afcde14d57b1c68b"))

// read the two string addresses from a file name currencies.csv
// the two strings should be assigned to two variables
// these two variables will be assigned to the currencies below


// func readCurrenciesFromFile(filename string) (string, string, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var lines []string
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 		if len(lines) == 2 {
// 			break
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return "", "", err
// 	}

// 	if len(lines) < 2 {
// 		return "", "", fmt.Errorf("file does not contain enough lines")
// 	}

// 	return strings.TrimSpace(lines[0]), strings.TrimSpace(lines[1]), nil
// }

// var currency0Addr, currency1Addr, err = readCurrenciesFromFile("currencies.csv")
// if err != nil {
// 	log.Fatalf("failed to read currencies: %v", err)
// }

// var currency0 = sdk.ConstUint248(
// 	common.HexToAddress(currency0Addr))
// var currency1 = sdk.ConstUint248(
// 	common.HexToAddress(currency1Addr))

func (c *AppCircuit) Allocate() (maxReceipts, maxSlots, maxTransactions int) {
	return 32, 0, 0
}

func (c *AppCircuit) Define(api *sdk.CircuitAPI, in sdk.DataInput) error {
	u248 := api.Uint248
	i248 := api.Int248

	receipts := sdk.NewDataStream(api, in.Receipts)

	sdk.AssertEach(receipts, func(l sdk.Receipt) sdk.Uint248 {
		// assertionPassed := u248.And(
			// 2. Check that the contract address of each log field is the expected contract
			// u248.IsEqual(l.Fields[0].Contract, currency0),
			// u248.IsEqual(l.Fields[1].Contract, currency1),
			// 3. Check the EventID of the fields are as expected
		// 	u248.IsEqual(l.Fields[0].EventID, EventIdLowVolatilitySwap),
		// )
		assertionPassed := u248.IsEqual(l.Fields[0].EventID, EventIdLowVolatilitySwap)
		return assertionPassed
	})

	// collect vol difference as points
	// Sum up the points
	volPoints := sdk.Map(receipts, func(cur sdk.Receipt) sdk.Uint248 {
		preVol := api.ToInt248(cur.Fields[2].Value)
		postVol := api.ToInt248(cur.Fields[3].Value)
		diffVal := u248.Sub(api.ToUint248(postVol), api.ToUint248(preVol))
		return i248.ABS(api.ToInt248(diffVal))
	})
	sumPoints := sdk.Sum(volPoints)

	blockNums := sdk.Map(receipts, func(cur sdk.Receipt) sdk.Uint248 { return api.ToUint248(cur.BlockNum) })
	minBlockNum := sdk.Min(blockNums)
	maxBlockNum := sdk.Max(blockNums)

	// Output will be reflected in app contract's callback in the form of
	// _circuitOutput: abi.encodePacked(uint256,uint248,uint64,address)
	api.OutputUint(248, sumPoints)
	api.OutputUint(64, minBlockNum)
	api.OutputUint(64, maxBlockNum)

	return nil
}
