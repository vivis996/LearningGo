package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
