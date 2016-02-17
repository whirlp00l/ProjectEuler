package main

import "fmt"

func gcd(a, b int) (ret int) {
	var temp int
	for b != 0 {
		temp = a % b
		a = b
		b = temp
	}
	return a
}

func main() {
	prev, c := 1, 1
	_gcd, lcm := 0, 1
	for c = 2; c <= 20; c++ {
		_gcd = gcd(prev, c)
		lcm *= (c / _gcd)
		prev = lcm
	}
	fmt.Println(lcm)
}
