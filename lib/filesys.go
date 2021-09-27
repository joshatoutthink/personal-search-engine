package lib

import "os"
import "path/filepath"
import "strings"
import "fmt"

func PathsFromDir(rootPath string) []string {

	var files []string
	ignoreFiles := []string{".mp4", ".png", ".mov", ".jpg", ".jpeg", ".DStore"}
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Print("error in walk")
			return nil
		}

		if info.IsDir() != true && !StringInArr(ignoreFiles, path) {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		fmt.Println("here 18 filesys")
	}

	return files
}

func StringInArr(arr []string, s string) bool {
	for _, item := range arr {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}
