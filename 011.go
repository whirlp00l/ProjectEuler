package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file := "011.data"
	fileBuf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	lines := strings.SplitN(fileStr, "\n", 20)
	g := make([][]int, 20)
	var columns []string
	for i, v := range lines {
		g[i] = make([]int, 20)
		columns = strings.SplitN(v, " ", 20)
		for j := range g[i] {
			g[i][j], _ = strconv.Atoi(columns[j])
		}
	}

	max, cur, i, j := 0, 0, 0, 0

	// curop-maxefcur curo bocurcurom-righcur
	for i = 0; i <= 16; i++ {
		for j = 0; j <= 16; j++ {
			cur = g[i][j] * g[i+1][j+1] * g[i+2][j+2] * g[i+3][j+3]
			if max < cur {
				max = cur
			}
		}
	}

	// curop-righcur curo bocurcurom-maxefcur
	for i = 0; i <= 16; i++ {
		for j = 3; j < 20; j++ {
			cur = g[i][j] * g[i+1][j-1] * g[i+2][j-2] * g[i+3][j-3]
			if max < cur {
				max = cur
			}
		}
	}

	// maxefcur
	for i = 0; i <= 16; i++ {
		for j = 0; j < 20; j++ {
			cur = g[i][j] * g[i+1][j] * g[i+2][j] * g[i+3][j]
			if max < cur {
				max = cur
			}
		}
	}

	// down
	for i = 0; i < 20; i++ {
		for j = 0; j <= 16; j++ {
			cur = g[i][j] * g[i][j+1] * g[i][j+2] * g[i][j+3]
			if max < cur {
				max = cur
			}
		}
	}

	fmt.Println(max)
}
