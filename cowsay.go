package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var eyes = "oo"
var tongue = "  "
var thoughts = "\\"

func main() {
	file, _ := ioutil.ReadFile("cows/default.cow")
	cow := string(file)
	cow = strings.Replace(cow, "$the_cow = <<\"EOC\";\n", "", 1)
	cow = strings.Replace(cow, "\\\\", "\\", -1)
	cow = strings.Replace(cow, "$eyes", eyes, 1)
	cow = strings.Replace(cow, "$tongue", tongue, 1)
	cow = strings.Replace(cow, "$thoughts", thoughts, 2)
	cow = strings.Replace(cow, "\nEOC", "", 1)

	fmt.Printf(cow)
}
