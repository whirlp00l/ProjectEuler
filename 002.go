package main

import "fmt"

func main() {
	prev, cur, next, sum := 1, 1, 0, 0
	for cur < 4e6 {
		if cur%2 == 0 {
			sum += cur
		}
		next = prev + cur
		prev = cur
		cur = next
	}
	fmt.Printf("Even Fibonnacii num sum is %d", sum)
}
