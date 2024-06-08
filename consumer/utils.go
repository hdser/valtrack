package consumer

import (
	"encoding/hex"
	"math"

	"github.com/prysmaticlabs/go-bitfield"
)

func stringToBitvector64(str string) (bitfield.Bitvector64, error) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return bitfield.Bitvector64{}, err
	}
	var bv bitfield.Bitvector64 = data
	copy(bv[:], data)
	return bv, nil
}

func indexesFromStrBitfield(bitVStr string) []int64 {
	bitV, _ := stringToBitvector64(bitVStr)

	indexes := make([]int64, 0, bitV.Len())

	for i := int64(0); i < 64; i++ {
		if bitV.BitAt(uint64(i)) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func indexesFromBitfield(bitV bitfield.Bitvector64) []int64 {
	indexes := make([]int64, 0, bitV.Len())

	for i := int64(0); i < 64; i++ {
		if bitV.BitAt(uint64(i)) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

// TODO: is this correct
func extractShortLivedSubnets(subscribed []int64, longLived []int64) []int64 {
	var shortLived []int64
	for i := 0; i < 64; i++ {
		if contains(subscribed, int64(i)) && !contains(longLived, int64(i)) {
			shortLived = append(shortLived, int64(i))
		}
	}

	return shortLived
}

func contains[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func ComputeNewAverage(prevAvg int32, prevCount uint64, currValidatorCount int) int32 {
	sum := int64(prevCount)*int64(prevAvg) + int64(currValidatorCount)
	newCount := int64(prevCount + 1)
	newAvg := float64(sum) / float64(newCount)
	return int32(math.Round(newAvg))
}
