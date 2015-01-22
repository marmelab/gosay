package balloon

import (
	"fmt"
	"github.com/marmelab/cowsay/cowsayType"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"strings"
	"unicode/utf8"
)

func Say(text string, maxWidth int, d cowsayType.Delimiters) string {

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
	if nbLines == 1 {
		maxWidth = len(lines[0])
		upper += " "
		lower += " "
	}

	for l := maxWidth + 1; l >= 0; l-- {
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
				newText = fmt.Sprintf("%c %s %c\n", d.First[0], line, d.First[1])
			} else if index == nbLines-1 {
				newText += fmt.Sprintf("%c %s %c", d.Last[0], line, d.Last[1])
			} else {
				newText += fmt.Sprintf("%c %s %c\n", d.Middle[0], line, d.Middle[1])
			}
		}

		return fmt.Sprintf("%s\n%s \n%s", upper, newText, lower)
	} else {
		return fmt.Sprintf("%s\n %c %s %c \n%s", upper, d.Only[0], lines[0], d.Only[1], lower)
	}
}
