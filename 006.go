package main

import (
	"fmt"
)

func main() {
	sumsqrt, sqrtsum := 0, 0
	for i := 1; i <= 100; i++ {
		sumsqrt += i * i
		sqrtsum += i
	}
	sqrtsum = sqrtsum * sqrtsum
	fmt.Println(sqrtsum - sumsqrt)
}
