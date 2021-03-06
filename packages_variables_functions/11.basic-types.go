package main

import (
	"fmt"
	"math/cmplx"
)

var (
	tobe bool = false
	maxint uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main()  {
	fmt.Printf("Type: %T Value: %v\n", tobe, tobe)
	fmt.Printf("Type: %T Value: %v\n", maxint, maxint)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
