package logic

import (
	"errors"
	"math/big"
)

// CheckMultiplicity -
func CheckMultiplicity(num int, divider int) (bool, error) {
	if divider == 0 {
		return false, errors.New("divider is zero")
	}

	return num%divider == 0, nil
}

// ForEachFibonacci -
func ForEachFibonacci(last uint, handler func(big.Int, uint)) {
	cache := [2]big.Int{
		*big.NewInt(0),
		*big.NewInt(1),
	}

	for i := uint(0); i < last; i++ {
		if i <= 1 {
			handler(*big.NewInt(int64(i)), i)
			continue
		}

		next := *new(big.Int).Add(&cache[0], &cache[1])
		cache[0] = cache[1]
		cache[1] = next

		handler(cache[1], i)
	}
}

// CalculateSimpleNumbersUpTo -
func CalculateSimpleNumbersUpTo(num uint) []int {
	list := []int{}
	idx := 1

	for {
		isStartNumber := idx == 2 || idx == 3 || idx == 5 || idx == 7
		isSuitableNumber := idx%2 != 0 && idx%3 != 0 && idx%5 != 0 && idx%7 != 0

		if isStartNumber || isSuitableNumber {
			list = append(list, idx)
		}

		if len(list) == 100 {
			break
		}

		idx++
	}

	return list
}
