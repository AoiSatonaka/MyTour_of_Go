package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = float64(1)
	for i := float64(1); i*i > 1e-10; z -= i{
		i = (z*z - x) / (2*z)
	}
	return
}

func main()  {
	fmt.Println(Sqrt(2))
	//比較用
	fmt.Println(math.Sqrt(2))
}