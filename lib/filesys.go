package lib

import "os"
import "path/filepath"
import "fmt"

func PathsFromDir(rootPath string) []string {
	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Print("error in walk")
			return nil
		}
		if info.IsDir() != true {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		fmt.Println("here 18 filesys")
	}

	return files
}
