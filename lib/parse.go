package lib

import (
	"regexp"
	"strings"
	"fmt"
)

var (
	MdTitle = regexp.MustCompile("^title:")
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


func Tokenize(content string)  TokenList{
//TODO
}
