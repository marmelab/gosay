package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("cows/default.cow")

	fmt.Printf(string(file))
}
