package main

import (
	"flag"
	"fmt"
	"github.com/marmelab/cowsay/cowsayType"
	"github.com/marmelab/cowsay/util/balloon"
	"github.com/marmelab/cowsay/util/cow"
	"os"
	"unicode/utf8"
)

var thoughts = "\\"
var text = ""
var maxWidth int

func main() {
	var eyes, tongue, cowfile string
	var list, borg, dead, greedy, paranoia, stoned, tired, wired, youthful, think bool
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
	flag.BoolVar(&think, "think", false, "thinking cow")

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
		cow.List()
		return
	}

	if utf8.RuneCountInString(tongue) < 2 {
		tongue += " "
	}
	if len(os.Args) != 0 {
		text = os.Args[len(os.Args)-1]
	}

	d := cowsayType.Delimiters{
		First:  [2]rune{'/', '\\'},
		Middle: [2]rune{'|', '|'},
		Last:   [2]rune{'\\', '/'},
		Only:   [2]rune{'<', '>'},
	}

	if think == true {
		thoughts = "o"
		d.First = [2]rune{'(', ')'}
		d.Middle = [2]rune{'(', ')'}
		d.Last = [2]rune{'(', ')'}
		d.Only = [2]rune{'(', ')'}
	}

	cow := cow.Load(cowfile, eyes, tongue, thoughts)

	say := balloon.Say(text, maxWidth, d)

	fmt.Printf("%s\n%s", say, cow)
}
