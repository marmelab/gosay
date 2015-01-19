package cow

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var cowPath = os.Getenv("COWPATH")

func Load(cowfile, eyes, tongue, thoughts string) string {
	if cowPath == "" {
		cowPath = "./cows"
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
	}

	cow := string(file)

	r, error := regexp.Compile("##.*\n")
	if error != nil {
		log.Fatal(error)
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

	return cow
}

func List() {
	if cowPath == "" {
		cowPath = "./cows"
	}

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
}
