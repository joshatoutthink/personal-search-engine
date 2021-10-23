package contentType

import (
	"my-brain/lib"
)

type digitalGardenContent struct{}

func (*digitalGardenContent) CollectLocations() (map[string]string, error) {
	return make(map[string]string), nil
}

func (*digitalGardenContent) GetContent(location string) (interface{}, error) {
	return make(map[string]string), nil
}

func (*digitalGardenContent) TokenizeDoc() lib.Doc {
	return lib.Doc{
		Content: " ",
		Heading: "",
		Id:      "",
		Path:    "",
		Tokens:  make(lib.TokenList),
	}
}
