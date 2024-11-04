package sdk

import (
	"fmt"
	"math/big"

	pgoldilocks "github.com/OpenAssetStandards/poseidon-goldilocks-go"
	"github.com/brevis-network/brevis-sdk/common/utils"
	zkhashutils "github.com/brevis-network/zk-hash/utils"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/std/algebra/emulated/sw_bn254"
	replonk "github.com/consensys/gnark/std/recursion/plonk"
)

var (
	ReceiptD     = new(big.Int).SetUint64(uint64(1))
	StorageD     = new(big.Int).SetUint64(uint64(2))
	TransactionD = new(big.Int).SetUint64(uint64(3))

	ReceiptVkHashHex     = "0x03e2e4a805080e6a792c9927ab14a83f427ec91d601ab463446231954b4a6d38"
	StorageVkHashHex     = "0x19efa435ccc1bf59b4e97281e7dcc612e2eda104987cfa9f6457435a90156297"
	TransactionVkHashHex = "0x051756e504e16aa2635181222db1d8a9310930599ce44c000414b02101431c89"
	MiddleNodeVkHashHex  = "0x15dc69eafcfd4546b82fecf468fd5878e2f7cb2270abee4e15abb638c77bbe52"
	AggAllVkHash         = "0x078ab850e8148fc412016972abf837fddbc8c7f87d049e337fcdfdc1a47caca2"

	ReceiptCircuitDigestHash = &pgoldilocks.HashOut256{12588371140186256321, 172593625752985613, 5743194239886726695, 16594510053085424922}
	StorageCircuitDigestHash = &pgoldilocks.HashOut256{4042215064603061775, 14063122104860698852, 844953845653799173, 12599854710728207088}
	TxCircuitDigestHash      = &pgoldilocks.HashOut256{6774922168907253952, 18184742135013640175, 7585912384159636816, 18115410258184238714}

	P2AggRecursionLeafCircuitDigestHash      = &pgoldilocks.HashOut256{6297162860691876658, 7207440660511781486, 956392925008767441, 15443083968980057808}
	P2Bn128WrapCircuitDigestHashForOnly2Leaf = utils.Hex2BigInt("0x1a608b7771b23c7972539bad7f395d7ad5c4a5e5be9dc891960179c1cacd8c78") // for from P2AggRecursionLeafCircuitDigestHash

	P2AggRecursionMiddleFormMiddleLeafCircuitDigestHash  = &pgoldilocks.HashOut256{14561383570925761150, 12116100132768392867, 2216989100987824959, 11720816772597981334}
	P2AggRecursionNoLeafCircuitDigestHash                = &pgoldilocks.HashOut256{14007309231803840793, 2325011900429631668, 6598512353030159473, 12456847712279341912}
	P2Bn128WrapCircuitDigestHashForOnlyFromLeafRecursion = utils.Hex2BigInt("0x08e1f454d096c46ebf2e1f40ff4858ce8188f86cc9623242ea605b774aee12aa") // for from P2AggRecursionMiddleFormMiddleLeafCircuitDigestHash
	P2Bn128WrapCircuitDigestHash                         = utils.Hex2BigInt("0x1E24794162210326BC751EB2FB4AFB6CB76B2CD94E6CEAFB9191A18B6E24A9D1") // for from P2AggRecursionNoLeafCircuitDigestHash

	ReceiptVkHash     = utils.Hex2BigInt(ReceiptVkHashHex)
	StorageVkHash     = utils.Hex2BigInt(StorageVkHashHex)
	TransactionVkHash = utils.Hex2BigInt(TransactionVkHashHex)
	MiddleNodeVkHash  = utils.Hex2BigInt(MiddleNodeVkHashHex)

	ReceiptNode = Hash2HashDigestNode{
		CircuitDigest: ReceiptD,
		VkHash:        ReceiptVkHash,
	}

	StorageNode = Hash2HashDigestNode{
		CircuitDigest: StorageD,
		VkHash:        StorageVkHash,
	}

	TransactionNode = Hash2HashDigestNode{
		CircuitDigest: TransactionD,
		VkHash:        TransactionVkHash,
	}

	ReceiptPlonky2Node = Plonky2DigestNode{
		CurCircuitDigest: ReceiptCircuitDigestHash,
		IsLeafNode:       true,
	}

	StoragePlonky2Node = Plonky2DigestNode{
		CurCircuitDigest: StorageCircuitDigestHash,
		IsLeafNode:       true,
	}

	TransactionPlonky2Node = Plonky2DigestNode{
		CurCircuitDigest: TxCircuitDigestHash,
		IsLeafNode:       true,
	}
)

func CalBrevisCircuitDigest(receiptCount, storageCount, transactionCount int, appVk plonk.VerifyingKey) (*big.Int, error) {
	reVk, err := replonk.ValueOfVerifyingKey[sw_bn254.ScalarField, sw_bn254.G1Affine, sw_bn254.G2Affine](appVk)
	if err != nil {
		return nil, err
	}

	appVkHashBytes := utils.CalculateAppVkHashForBn254(reVk)
	appVkHashBigInt := new(big.Int).SetBytes(appVkHashBytes)

	hash2HashDigest, err := GetHash2HashCircuitDigest(receiptCount, storageCount, transactionCount)
	if err != nil {
		return nil, err
	}

	plonky2RootFromBn128Digest, isRecursionOnLeaf, isRecursionRecursionOfLeaf, err := GetPlonky2CircuitDigestFromRootNodeSelf(receiptCount, storageCount, transactionCount)
	if err != nil {
		return nil, err
	}

	hasher := zkhashutils.NewPoseidonBn254()
	hasher.Write(new(big.Int).SetUint64(plonky2RootFromBn128Digest[0]))
	hasher.Write(new(big.Int).SetUint64(plonky2RootFromBn128Digest[1]))
	hasher.Write(new(big.Int).SetUint64(plonky2RootFromBn128Digest[2]))
	hasher.Write(new(big.Int).SetUint64(plonky2RootFromBn128Digest[3]))

	if isRecursionOnLeaf {
		hasher.Write(P2Bn128WrapCircuitDigestHashForOnly2Leaf)
	} else if isRecursionRecursionOfLeaf {
		hasher.Write(P2Bn128WrapCircuitDigestHashForOnlyFromLeafRecursion)
	} else {
		hasher.Write(P2Bn128WrapCircuitDigestHash)
	}

	hasher.Write(hash2HashDigest)
	hasher.Write(utils.Hex2BigInt(MiddleNodeVkHashHex))

	hasher.Write(appVkHashBigInt)

	return hasher.Sum()
}

type Hash2HashDigestNode struct {
	CircuitDigest *big.Int
	VkHash        *big.Int
}

func GetHash2HashCircuitDigest(receiptCount, storageCount, transactionCount int) (*big.Int, error) {
	receiptLeafCount, storageLeafCount, transactionLeafCount, totalLeafCount, err := GetAndCheckLeafCount(receiptCount, storageCount, transactionCount)
	if err != nil {
		return nil, err
	}

	var totalLeafs []Hash2HashDigestNode
	for i := 0; i < receiptLeafCount; i++ {
		totalLeafs = append(totalLeafs, ReceiptNode)
	}
	for i := 0; i < storageLeafCount; i++ {
		totalLeafs = append(totalLeafs, StorageNode)
	}
	for i := 0; i < transactionLeafCount; i++ {
		totalLeafs = append(totalLeafs, TransactionNode)
	}
	if len(totalLeafs) != totalLeafCount {
		return nil, fmt.Errorf("len(totalLeafs) != totalLeafCount, %d %d", totalLeafs, totalLeafCount)
	}
	elementCount := totalLeafCount
	for {
		if elementCount == 1 {
			return totalLeafs[0].CircuitDigest, nil
		}
		for i := 0; i < elementCount/2; i++ {
			parent, hashErr := CalOneHash2HashNodeDigest(totalLeafs[2*i], totalLeafs[2*i+1])
			if hashErr != nil {
				return nil, fmt.Errorf("fail to hash in CalOneHash2HashNodeDigest, %d %d -> %d err: %v", 2*i, 2*i+1, i, hashErr)
			}
			totalLeafs[i] = Hash2HashDigestNode{
				CircuitDigest: parent,
				VkHash:        MiddleNodeVkHash,
			}
		}
		elementCount = elementCount / 2
	}
}

func CalOneHash2HashNodeDigest(left, right Hash2HashDigestNode) (*big.Int, error) {
	poseidonHasher := zkhashutils.NewPoseidonBn254()
	poseidonHasher.Write(left.CircuitDigest)
	poseidonHasher.Write(right.CircuitDigest)
	poseidonHasher.Write(left.VkHash)
	poseidonHasher.Write(right.VkHash)

	return poseidonHasher.Sum()
}

type Plonky2DigestNode struct {
	PubCircuitDigest           *pgoldilocks.HashOut256 // in pub
	CurCircuitDigest           *pgoldilocks.HashOut256 // in json
	IsLeafNode                 bool
	IsRecursionOfLeaf          bool
	IsRecursionRecursionOfLeaf bool
}

func GetPlonky2CircuitDigestFromRootNodeSelf(receiptCount, storageCount, transactionCount int) (*pgoldilocks.HashOut256, bool, bool, error) {
	plonky2RootNode, err := GetPlonky2CircuitDigest(receiptCount, storageCount, transactionCount)
	if err != nil {
		return nil, false, false, err
	}

	var data []uint64
	data = append(data, plonky2RootNode.PubCircuitDigest[:]...)
	data = append(data, plonky2RootNode.CurCircuitDigest[:]...)

	hashRes, err := pgoldilocks.HashNoPadU64Array(data[:])
	if err != nil {
		return nil, false, false, err
	}

	return hashRes, plonky2RootNode.IsRecursionOfLeaf, plonky2RootNode.IsRecursionRecursionOfLeaf, nil
}

func GetPlonky2CircuitDigest(receiptCount, storageCount, transactionCount int) (*Plonky2DigestNode, error) {
	receiptLeafCount, storageLeafCount, transactionLeafCount, totalLeafCount, err := GetAndCheckLeafCount(receiptCount, storageCount, transactionCount)
	if err != nil {
		return nil, err
	}

	var totalLeafs []Plonky2DigestNode
	for i := 0; i < receiptLeafCount; i++ {
		totalLeafs = append(totalLeafs, ReceiptPlonky2Node)
	}
	for i := 0; i < storageLeafCount; i++ {
		totalLeafs = append(totalLeafs, StoragePlonky2Node)
	}
	for i := 0; i < transactionLeafCount; i++ {
		totalLeafs = append(totalLeafs, TransactionPlonky2Node)
	}
	if len(totalLeafs) != totalLeafCount {
		return nil, fmt.Errorf("len(totalLeafs) != totalLeafCount, %d %d", len(totalLeafs), totalLeafCount)
	}
	elementCount := totalLeafCount
	for {
		if elementCount == 1 {
			return &totalLeafs[0], nil
		}
		for i := 0; i < elementCount/2; i++ {
			parent, hashErr := CalOnePlonky2NodeDigest(totalLeafs[2*i], totalLeafs[2*i+1])
			if hashErr != nil {
				return nil, fmt.Errorf("fail to hash in CalOneHash2HashNodeDigest, %d %d -> %d err: %v", 2*i, 2*i+1, i, hashErr)
			}
			if totalLeafs[2*i].IsLeafNode {
				totalLeafs[i] = Plonky2DigestNode{
					PubCircuitDigest:  parent,
					CurCircuitDigest:  P2AggRecursionLeafCircuitDigestHash,
					IsRecursionOfLeaf: true,
				}
			} else if totalLeafs[2*i].IsRecursionOfLeaf {
				totalLeafs[i] = Plonky2DigestNode{
					PubCircuitDigest:           parent,
					CurCircuitDigest:           P2AggRecursionMiddleFormMiddleLeafCircuitDigestHash,
					IsRecursionRecursionOfLeaf: true,
				}
			} else {
				totalLeafs[i] = Plonky2DigestNode{
					PubCircuitDigest: parent,
					CurCircuitDigest: P2AggRecursionNoLeafCircuitDigestHash,
				}
			}
		}
		elementCount = elementCount / 2
	}
}

func CalOnePlonky2NodeDigest(left, right Plonky2DigestNode) (*pgoldilocks.HashOut256, error) {
	if left.IsLeafNode != right.IsLeafNode {
		return nil, fmt.Errorf("left leaf not equal to right leaf, left: %+v, right: %+v", left, right)
	}

	if left.IsLeafNode {
		var preimage []uint64

		preimage = append(preimage, left.CurCircuitDigest[:]...)
		preimage = append(preimage, right.CurCircuitDigest[:]...)

		return pgoldilocks.HashNoPadU64Array(preimage)
	} else {
		var preimage1, preimage2, preimage3 []uint64
		preimage1 = append(preimage1, left.PubCircuitDigest[:]...)
		preimage1 = append(preimage1, right.PubCircuitDigest[:]...)
		hash1, err := pgoldilocks.HashNoPadU64Array(preimage1)
		if err != nil {
			return nil, fmt.Errorf("fail to hash data preimage1, left:%+v, right: %+v, err:%v", left, right, err)
		}
		preimage2 = append(preimage2, left.CurCircuitDigest[:]...)
		preimage2 = append(preimage2, right.CurCircuitDigest[:]...)
		hash2, err := pgoldilocks.HashNoPadU64Array(preimage2)
		if err != nil {
			return nil, fmt.Errorf("fail to hash data preimage2, left:%+v, right: %+v, err:%v", left, right, err)
		}
		preimage3 = append(preimage3, hash1[:]...)
		preimage3 = append(preimage3, hash2[:]...)

		return pgoldilocks.HashNoPadU64Array(preimage3)
	}
}

func GetAndCheckLeafCount(receiptCount, storageCount, transactionCount int) (receiptLeafCount int, storageLeafCount int, transactionLeafCount int, totalLeafCount int, err error) {
	if receiptCount%32 != 0 {
		return 0, 0, 0, 0, fmt.Errorf("receipt count is not n * 32")
	}
	receiptLeafCount = receiptCount / 32

	if storageCount%32 != 0 {
		return 0, 0, 0, 0, fmt.Errorf("storage count is not n * 32")
	}
	storageLeafCount = storageCount / 32

	if transactionCount%32 != 0 {
		return 0, 0, 0, 0, fmt.Errorf("transaction count is not n * 32")
	}
	transactionLeafCount = transactionCount / 32

	totalLeafCount = receiptLeafCount + storageLeafCount + transactionLeafCount
	if !CheckNumberPowerOfTwo(totalLeafCount) {
		return 0, 0, 0, 0, fmt.Errorf("leaf count n is not power of 2, totalLeafCount: %d", totalLeafCount)
	}

	return receiptLeafCount, storageLeafCount, transactionLeafCount, totalLeafCount, nil
}
