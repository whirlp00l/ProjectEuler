package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of all multiples of 3 or 5 below 1000 = %d\n", sum)
}
