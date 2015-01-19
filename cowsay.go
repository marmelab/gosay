package main

import (
	"flag"
	"fmt"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

var thoughts = "\\"
var text = ""
var maxWidth = 40

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

	text = wordwrap.WrapString(text, uint(maxWidth))

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		length := utf8.RuneCountInString(line)
		if length > maxWidth {
			maxWidth = length
		}
	}

	nbLines := len(lines)
	upper := " "
	lower := " "
	for l := maxWidth; l >= 0; l-- {
		upper += "_"
		lower += "-"
	}

	if nbLines > 1 {
		newText := ""
		for index, line := range lines {
			for spaceCount := maxWidth - utf8.RuneCountInString(line); spaceCount > 0; spaceCount-- {
				line += " "
			}
			if index == 0 {
				newText = fmt.Sprintf("%c %s %c\n", d.first[0], line, d.first[1])
			} else if index == nbLines-1 {
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
	var eyes, tongue string
	flag.StringVar(&eyes, "e", "oo", "specify the eye")
	flag.StringVar(&tongue, "T", "  ", "specify the tongue")
	flag.Parse()
	if len(tongue) < 2 {
		tongue += " "
	}
	file, _ := ioutil.ReadFile("cows/default.cow")
	if len(os.Args) != 0 {
		text = os.Args[len(os.Args)-1]
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
