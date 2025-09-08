package utils

import "math/rand"

type Number interface {
	int | int64 | int32 | int16 | int8 | float32 | float64 | byte
}

func Next[N Number](min, max int) N {
	return N(min + rand.Intn(max-min+1))
}

func NextInt1(min, max int) int {
	return Next[int](min, max)
}

func NextInt2(max int) int {
	return Next[int](0, max)
}
