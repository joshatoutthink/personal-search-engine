package main

// MAIN.go
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"my-brain/lib"
	"net/http"
)

var Indexes map[string]lib.Doc
var Mods []string

//Starts server
// place all routes here
func handleRequests() {
	// creates a new instance of a mux router

	Mods = append(Mods, "/Users/joshkennedy00/sites/joshs/sandbox/content/brain/")
	myRouter := mux.NewRouter().StrictSlash(true)

	// routes
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllIndexes)
	myRouter.HandleFunc("/reindex", reIndexMods)
	myRouter.HandleFunc("/search", searchIndexes)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func reIndexMods(w http.ResponseWriter, r *http.Request) {
	Indexes = make(map[string]lib.Doc)
	for modIndx, mod := range Mods {
		fmt.Println(mod, "is a folder")
		fpaths := lib.PathsFromDir(mod)

		for indx, fpath := range fpaths {
			id := modIndx + indx
			content := string(lib.ReadFile(fpath))

			doc := lib.Doc{
				Tokens:  lib.Tokenize(content),
				Content: content,
				Heading: lib.MarkdownTitle(fpath),
				Path:    fpath,
				Id:      fmt.Sprint(id),
			}
			Indexes[fmt.Sprint(id)] = doc
		}
	}

	_, err := fmt.Println("50", Indexes)
	if err != nil {
		fmt.Println(err)
	}
	byteVal, err := json.Marshal(Indexes)
	if err != nil {
		fmt.Println(err)
	}
	newContent := string(byteVal)
	lib.Write("./indexes.json", &newContent)
	fmt.Fprintf(w, "Reindexed")
}

func searchIndexes(w http.ResponseWriter, r *http.Request) {

}

func returnAllIndexes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllIndexes")
	b, err := json.MarshalIndent(&Indexes, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
	json.NewEncoder(w).Encode(Indexes)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func unMarshalIndexFile(filePath string) {
	byteVal := lib.ReadFile(filePath)
	unmarshalErr := json.Unmarshal(byteVal, &Indexes)
	if unmarshalErr != nil {
		fmt.Printf("%+v\n", unmarshalErr)
	}
}

func run() {
	unMarshalIndexFile("./indexes.json")
	//START OUR WEB SERVER
	handleRequests()
}

func main() {
	run()
}
