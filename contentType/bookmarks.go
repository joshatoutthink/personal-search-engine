package contentType

// import (
// "encoding/json"
// "fmt"
// "my-brain/lib"
// )

// type BookmarkContent struct {
// file string
// }

// var Bookmark = BookmarkContent{
// file: "/Users/joshkennedy00/Library/Application Support/Google/Chrome/Befault/Bookmarks",
// }

// func (B *BookmarkContent) CollectLocations() ([]string, error) {
// BookmarkFile := make(map[string]interface{})
// byteVal := lib.ReadFile(B.file)
// unmarshalErr := json.Unmarshal(byteVal, &BookmarkFile)
// if unmarshalErr != nil {
// fmt.Println(unmarshalErr)
// }

// return bookmarks, nil
// }

// func (B *BookmarkContent) TokenizeDoc(id string, fpath string) lib.Doc {
// return lib.Doc{
// Tokens:  lib.Tokenize(content),
// Content: content,
// Heading: lib.MarkdownTitle(fpath),
// Path:    fpath,
// Id:      fmt.Sprint(id),
// }
// }
