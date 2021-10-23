package contentType

import (
	"fmt"
	"my-brain/lib"
)

type DigitalGardenContent struct {
	dir     string
	content string
}

var DigitalGarden = DigitalGardenContent{
	dir:     "/Users/joshkennedy00/sites/joshs/sandbox/content/brain/",
	content: "",
}

func (D *DigitalGardenContent) CollectLocations() ([]string, error) {
	return lib.PathsFromDir(D.dir), nil
}

func (D *DigitalGardenContent) GetContent(location string) (string, error) {
	return string(lib.ReadFile(location)), nil
}

func (D *DigitalGardenContent) TokenizeDoc(id string, fpath string) lib.Doc {
	content, err := D.GetContent(fpath)
	if err != nil {
		fmt.Println(err)
	}
	return lib.Doc{
		Tokens:  lib.Tokenize(content),
		Content: content,
		Heading: lib.MarkdownTitle(fpath),
		Path:    fpath,
		Id:      fmt.Sprint(id),
	}
}
