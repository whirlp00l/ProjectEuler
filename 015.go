package main

import (
	"fmt"
)

func main() {
	var array [21][21]int
	for i := 0; i < 20; i++ {
		array[i][20] = 1
		array[20][i] = 1
	}
	for i := 19; i >= 0; i-- {
		for j := 19; j >= 0; j-- {
			array[i][j] = array[i+1][j] + array[i][j+1]
		}
	}
	fmt.Println(array[0][0])
}
