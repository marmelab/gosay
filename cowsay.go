package main

import (
	"flag"
	"fmt"
	"github.com/marmelab/cowsay/util/balloon"
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

	cow = balloon.Say(text, maxWidth) + "\n" + cow

	fmt.Printf(cow)
}
