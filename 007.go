package main

import (
	"fmt"
)

func genPrime(num, th int) int {
	nums := make([]bool, num)
	count := 0
	for i := 2; i < num; i++ {
		if nums[i] == false {
			count++
			if count == th {
				return i
			}
			for j := i * i; j < num; j += i {
				nums[j] = true
			}
		}
	}
	return -1
}
func main() {
	val := genPrime(10000000, 10001)
	fmt.Println(val)
}
