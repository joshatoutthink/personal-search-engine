package lib

import (
	"regexp"
	"strings"
	//"bytes"
)

var (
	//Regexes
	MdTitle      = regexp.MustCompile("^title:")
	AlphaNumeric = regexp.MustCompile("[^A-Za-z0-9]")

	stopWords = []string{"a"}
)

func MarkdownTitle(path string) string {
	lines := LinesOfFile(path)

	heading := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if MdTitle.MatchString(line) {
			MdTitle.ReplaceAllString(line, "")
			return line
		}
	}
	return heading
}

func Tokenize(content string) TokenList {
	splitString := regexp.MustCompile("[^\\s]+").FindAllString(content, -1)
	words := make(TokenList)

	for _, piece := range splitString {
		mixCaseWord := AlphaNumeric.ReplaceAllString(piece, "")
		word := strings.ToLower(mixCaseWord)
		for _, stopWord := range stopWords {
			if word == stopWord {
				continue
			}
		}

		count := float64(0)
		_, ok := words[word]
		if ok {
			count = words[word]
		}
		words[word] = count + float64(1)
	}
	return words
}
