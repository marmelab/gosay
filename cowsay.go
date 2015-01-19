package main

import (
	"flag"
	"fmt"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

var thoughts = "\\"
var text = ""
var maxWidth int

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
	var cowPath = os.Getenv("COWPATH")
	if cowPath == "" {
		cowPath = "./cows"
	}
	var eyes, tongue, cowfile string
	var list, borg, dead, greedy, paranoia, stoned, tired, wired, youthful bool
	flag.StringVar(&eyes, "e", "oo", "specify the eye")
	flag.StringVar(&tongue, "T", "  ", "specify the tongue")
	flag.StringVar(&cowfile, "f", "default", "specify the cow file to use")
	flag.IntVar(&maxWidth, "W", 40, "specify roughly where the word should be wrapped")
	flag.BoolVar(&list, "l", false, "list cow file in COWPATH")
	flag.BoolVar(&borg, "b", false, "borg mode")
	flag.BoolVar(&dead, "d", false, "dead mode")
	flag.BoolVar(&greedy, "g", false, "greedy mode")
	flag.BoolVar(&paranoia, "p", false, "paranoia mode")
	flag.BoolVar(&stoned, "s", false, "stoned mode")
	flag.BoolVar(&tired, "t", false, "tired mode")
	flag.BoolVar(&wired, "w", false, "wired mode")
	flag.BoolVar(&youthful, "y", false, "youthful mode")

	flag.Parse()

	switch {
	case borg == true:
		eyes = "=="
		tongue = "  "
	case dead == true:
		eyes = "xx"
		tongue = "U "
	case greedy == true:
		eyes = "$$"
		tongue = "  "
	case paranoia == true:
		eyes = "@@"
		tongue = "  "
	case stoned == true:
		eyes = "**"
		tongue = "U "
	case tired == true:
		eyes = "--"
		tongue = "  "
	case wired == true:
		eyes = "OO"
		tongue = "  "
	case youthful == true:
		eyes = ".."
		tongue = "  "
	}

	if list {
		files, error := ioutil.ReadDir(cowPath)
		if error != nil {
			log.Fatal(error)
			return
		}
		for _, f := range files {
			name := strings.Split(f.Name(), ".")
			if len(name) > 1 && name[1] == "cow" {
				fmt.Println(name[0])
			}
		}
		return
	}

	if utf8.RuneCountInString(tongue) < 2 {
		tongue += " "
	}

	var filePath string
	var absolute, _ = regexp.MatchString("/", cowfile)
	if absolute == true {
		filePath = fmt.Sprintf("%s.cow", cowfile)
	} else {
		filePath = fmt.Sprintf("%s/%s.cow", cowPath, cowfile)
	}

	file, error := ioutil.ReadFile(filePath)
	if error != nil {
		log.Fatal(error)
		return
	}
	if len(os.Args) != 0 {
		text = os.Args[len(os.Args)-1]
	}

	cow := string(file)

	r, error := regexp.Compile("##.*\n")
	if error != nil {
		log.Fatal(error)
		return
	}
	cow = r.ReplaceAllString(cow, "")

	cow = strings.Replace(cow, "$the_cow = <<EOC;\n", "", 1)
	cow = strings.Replace(cow, "$the_cow = <<\"EOC\";\n", "", 1)
	cow = strings.Replace(cow, "\\\\", "\\", -1)
	cow = strings.Replace(cow, "\\@", "@", -1)
	cow = strings.Replace(cow, "$eyes", eyes, -1)
	cow = strings.Replace(cow, "$tongue", tongue, -1)
	cow = strings.Replace(cow, "$thoughts", thoughts, -1)
	cow = strings.Replace(cow, "\nEOC", "", 1)

	cow = say(text) + "\n" + cow

	fmt.Printf(cow)
}
