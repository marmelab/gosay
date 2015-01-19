package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var eyes = "oo"
var tongue = "  "
var thoughts = "\\"
var text = ""

type Delimiters struct {
	first, middle, last, only [2]rune
}

func split(text string) []string {
	var result []string
	length := len(text)
	lineCount := int(length/38) + 1
	result = make([]string, lineCount)
	for i := 0; i < lineCount; i++ {
		start := i * 38
		end := start + 38
		if end > length {
			end = length
		}
		result[i] = text[start:end]
	}

	return result
}

func say(text string) string {
	d := Delimiters{
		first:  [2]rune{'/', '\\'},
		middle: [2]rune{'|', '|'},
		last:   [2]rune{'\\', '/'},
		only:   [2]rune{'<', '>'},
	}
	lines := split(text)
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
				for spaceCount := 38 - len(line); spaceCount > 0; spaceCount-- {
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
	file, _ := ioutil.ReadFile("cows/default.cow")
	if len(os.Args) != 0 {
		text = os.Args[1]
	}
	cow := string(file)
	cow = strings.Replace(cow, "$the_cow = <<\"EOC\";\n", "", 1)
	cow = strings.Replace(cow, "\\\\", "\\", -1)
	cow = strings.Replace(cow, "$eyes", eyes, 1)
	cow = strings.Replace(cow, "$tongue", tongue, 1)
	cow = strings.Replace(cow, "$thoughts", thoughts, 2)
	cow = strings.Replace(cow, "\nEOC", "", 1)

	cow = say(text) + "\n" + cow

	fmt.Printf(cow)
}
