package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 1000001)
	steps, maxSteps, cur, max, notFound := 0, 0, 0, 0, true
	for i := 1; i < len(arr); i++ {
		notFound = true
		cur = i
		steps = 0
		for notFound {
			if cur == 1 {
				notFound = false
				arr[i] = steps
				if steps > maxSteps {
					maxSteps = steps
					max = i
				}
			} else if cur < i {
				notFound = false
				steps = steps + arr[cur]
				arr[i] = steps
				if steps > maxSteps {
					maxSteps = steps
					max = i
				}
			} else {
				if cur%2 == 0 {
					cur /= 2
				} else {
					cur = 3*cur + 1
				}
			}
			steps++
		}
	}
	fmt.Println(max)
}
