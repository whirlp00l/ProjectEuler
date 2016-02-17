package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 998; i++ {
		for j := 1; j+i < 999; j++ {
			k := 1000 - j - i
			if i*i+j*j == k*k {
				fmt.Println(i * j * k)
			}
		}
	}
}
