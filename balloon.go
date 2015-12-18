package cowsay

import (
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"strings"
	"unicode/utf8"
)

type balloonBorder struct {
	First, Middle, Last, Only [2]rune
}

type balloonTrail string

var normalBorder = &balloonBorder{
	[2]rune{'/', '\\'},
	[2]rune{'|', '|'},
	[2]rune{'\\', '/'},
	[2]rune{'<', '>'},
}

var thoughtfulBorder = &balloonBorder{
	[2]rune{'(', ')'},
	[2]rune{'(', ')'},
	[2]rune{'(', ')'},
	[2]rune{'(', ')'},
}

var normalTrail balloonTrail = "\\"
var thoughtfulTrail balloonTrail = "o"

func buildBalloon(text string, opts *Options) (string, balloonTrail) {
	maxWidth := opts.MaxWidth
	border, trail := balloonSettings(opts.Think)
	text = wordwrap.WrapString(text, uint(maxWidth))
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		length := utf8.RuneCountInString(line)
		if length > maxWidth {
			maxWidth = length
		}
	}

	numLines := len(lines)
	upper := " "
	lower := " "

	if numLines == 1 {
		maxWidth = len(lines[0])
		upper += " "
		lower += " "
	}

	for l := maxWidth + 1; l >= 0; l-- {
		upper += "_"
		lower += "-"
	}

	if numLines > 1 {
		newText := ""
		for index, line := range lines {
			for spaceCount := maxWidth - utf8.RuneCountInString(line); spaceCount > 0; spaceCount-- {
				line += " "
			}

			if index == 0 {
				newText = fmt.Sprintf("%c %s %c\n", border.First[0], line, border.First[1])
			} else if index == numLines-1 {
				newText += fmt.Sprintf("%c %s %c", border.Last[0], line, border.Last[1])
			} else {
				newText += fmt.Sprintf("%c %s %c\n", border.Middle[0], line, border.Middle[1])
			}
		}

		return fmt.Sprintf("%s\n%s \n%s", upper, newText, lower), trail
	}

	return fmt.Sprintf("%s\n %c %s %c \n%s", upper, border.Only[0], lines[0], border.Only[1], lower), trail
}

func balloonSettings(thoughtful bool) (*balloonBorder, balloonTrail) {
	if thoughtful {
		return thoughtfulBorder, thoughtfulTrail
	}

	return normalBorder, normalTrail
}
