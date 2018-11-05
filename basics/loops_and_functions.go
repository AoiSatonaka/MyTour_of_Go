package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = float64(1)
	for i:=float64(0); float32(z) != float32(i);{
		i = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z,i)
	}
	return
}

func main()  {
	fmt.Println(Sqrt(3))
	//比較用
	fmt.Println(math.Sqrt(2))
}