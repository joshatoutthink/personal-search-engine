package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Write(path string, content *string) {

	f, err := os.Create(path)
	if err != nil {
		fmt.Print("error: Couldnt Create ", path, " because: %v", err)
	}
	defer f.Close()

	b, err := f.WriteString(*content)
	if err != nil {
		fmt.Println("error: Couldnt write to ", path, "  because:", err)
	}

	f.Sync()

	fmt.Print(b)
}

func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("couldnt read index file")
		fmt.Println(err)
	}
	defer file.Close()
	byteVal, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return byteVal
}

func LinesOfFile(filePath string) []string {
	lines := make([]string, 100)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("couldnt read index file")
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
