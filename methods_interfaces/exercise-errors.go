package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	v := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: ", v)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return x, nil
	} else {
		return x, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}