package pgl

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](vals ...T) T {

	if len(vals) == 0 {
		var zeroVal T
		return zeroVal
	}

	min := vals[0]

	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}

	return min
}

func Max[T constraints.Ordered](vals ...T) T {

	if len(vals) == 0 {
		var zeroVal T
		return zeroVal
	}

	max := vals[0]

	for i := 1; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}

	return max
}
