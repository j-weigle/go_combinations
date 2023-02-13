package combinations

import "math/bits"

// GetSubset gets the nth subset of all possible combination subsets
// that could be gotten from set when using the All function
func GetSubset[T any](set *[]T, pos uint64) []T {
	var subset []T
	for i := 0; i < len(*set); i++ {
		if (pos>>i)&1 == 1 {
			subset = append(subset, (*set)[i])
		}
	}
	return subset
}

// All gets every combination subset that is possible for the set
func All[T any](set *[]T) (subsets [][]T) {
	for pos := uint64(1); pos < (1 << uint(len(*set))); pos++ {
		subsets = append(subsets, GetSubset(set, pos))
	}
	return subsets
}

// AllQualifying gets the combination subsets that qualify according to the
// criteria defined in the qualifies callback function
func AllQualifying[T any](set *[]T, qualifies func(*[]T) bool) (subsets [][]T) {
	for pos := uint64(1); pos < (1 << uint(len(*set))); pos++ {
		subset := GetSubset(set, pos)
		if qualifies(&subset) {
			subsets = append(subsets, subset)
		}
	}
	return subsets
}

// AllQualifyingPositions gets the positions of the combination subsets that
// would be generated on the nth iteration from the All function if that
// position's subset qualifies according to the criteria defined in the
// qualifies callback function
func AllQualifyingPositions[T any](set *[]T, qualifies func(*[]T) bool) (positions []uint64) {
	for pos := uint64(1); pos < (1 << uint(len(*set))); pos++ {
		subset := GetSubset(set, pos)
		if qualifies(&subset) {
			positions = append(positions, pos)
		}
	}
	return positions
}

// Combinations gets the combination subsets that are possible for the set
// that are the length of the sample size r
func Combinations[T any](set *[]T, r int) (subsets [][]T) {
	for pos := uint64(1); pos < (1 << uint(len(*set))); pos++ {
		if bits.OnesCount(uint(pos)) != r {
			continue
		}
		subsets = append(subsets, GetSubset(set, pos))
	}
	return subsets
}

// CombinationsPositions gets the positions of the combination subsets that
// are possible for the set that are the length of the sample size r where
// the position is the nth iteration from the All function
func CombinationsPositions[T any](set *[]T, r int) (positions []uint64) {
	for pos := uint64(1); pos < (1 << uint(len(*set))); pos++ {
		if bits.OnesCount(uint(pos)) != r {
			continue
		}
		positions = append(positions, pos)
	}
	return positions
}

// CombinationsQualifyingPositions gets the positions of the combination
// subsets that are possible for the set that are the length of the sample
// size r where the position is the nth iteration from the all function and
// the position's subset qualifies according to the criteria defined in the
// qualifies callback function
func CombinationsQualifyingPositions[T any](set *[]T, r int, qualifies func(*[]T) bool) (positions []uint64) {
	for pos := uint64(1); pos < (1 << uint(len(*set))); pos++ {
		if bits.OnesCount(uint(pos)) != r {
			continue
		}
		subset := GetSubset(set, pos)
		if qualifies(&subset) {
			positions = append(positions, pos)
		}
	}
	return positions
}
