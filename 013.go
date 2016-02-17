package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func main() {
	fileBuf, err := ioutil.ReadFile("013.data")
	if err != nil {
		panic(err)
	}
	fileStr := string(fileBuf)
	s := strings.SplitN(fileStr, "\n", 100)
	n := new(big.Int).SetInt64(0)
	for _, v := range s {
		t := new(big.Int)
		t.SetString(v, 10)
		n.Add(n, t)
	}
	fmt.Println(n.String()[0:10])
}
