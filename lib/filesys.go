package lib
import "os"
import "path/filepath"

func PathsFromDir(rootPath string) []string {
	var files []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
