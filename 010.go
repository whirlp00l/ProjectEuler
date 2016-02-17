package main

import (
	"fmt"
)

func genPrime(num int) int {
	nums := make([]bool, num)
	sum := 0
	for i := 2; i < num; i++ {
		if nums[i] == false {
			sum += i
			for j := i * i; j < num; j += i {
				nums[j] = true
			}
		}
	}
	return sum
}
func main() {
	val := genPrime(2000000)
	fmt.Println(val)
}
