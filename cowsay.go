package main

import (
	"flag"
	"fmt"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"io/ioutil"
	"os"
	"strings"
)

var tongue = "  "
var thoughts = "\\"
var text = ""

type Delimiters struct {
	first, middle, last, only [2]rune
}

func say(text string) string {
	d := Delimiters{
		first:  [2]rune{'/', '\\'},
		middle: [2]rune{'|', '|'},
		last:   [2]rune{'\\', '/'},
		only:   [2]rune{'<', '>'},
	}

	text = wordwrap.WrapString(text, 38)
	lines := strings.Split(text, "\n")
	nbLines := len(lines)
	upper := " "
	lower := " "
	for l := len(lines[0]); l >= 0; l-- {
		upper += "_"
		lower += "-"
	}

	if nbLines > 1 {
		newText := ""
		for index, line := range lines {
			if index == 0 {
				newText = fmt.Sprintf("%c %s %c\n", d.first[0], line, d.first[1])
			} else if index == nbLines-1 {
				for spaceCount := 40 - len(line); spaceCount > 0; spaceCount-- {
					line += " "
				}
				newText += fmt.Sprintf("%c %s %c", d.last[0], line, d.last[1])
			} else {
				newText += fmt.Sprintf("%c %s %c\n", d.middle[0], line, d.middle[1])
			}
		}
		return fmt.Sprintf("%s\n%s \n%s", upper, newText, lower)
	} else {
		return fmt.Sprintf("%s\n %s \n%s", upper, d.only[0], lines[0], d.only[1], lower)
	}
}

func main() {
	var e string
	flag.StringVar(&e, "e", "oo", "specify the eye")
	flag.Parse()
	file, _ := ioutil.ReadFile("cows/default.cow")
	if len(os.Args) != 0 {
		text = os.Args[len(os.Args)-1]
	}

	cow := string(file)
	cow = strings.Replace(cow, "$the_cow = <<\"EOC\";\n", "", 1)
	cow = strings.Replace(cow, "\\\\", "\\", -1)
	cow = strings.Replace(cow, "$eyes", e, 1)
	cow = strings.Replace(cow, "$tongue", tongue, 1)
	cow = strings.Replace(cow, "$thoughts", thoughts, 2)
	cow = strings.Replace(cow, "\nEOC", "", 1)

	cow = say(text) + "\n" + cow

	fmt.Printf(cow)
}
