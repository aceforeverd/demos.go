package main

import (
	"fmt"
	"math"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("negative float %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, errNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
