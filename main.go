package main
// MAIN.go
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"my-brain/lib"
)

var Indexes map[string]lib.Doc

//Starts server
// place all routes here
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	fmt.Println(lib.MarkdownTitle("./indexes.json"))
	// routes
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllIndexes)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
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
	unMarshalIndexFile("/Users/joshkennedy00/sites/joshs/my-brain/prototype/indexes.json")
	//START OUR WEB SERVER
	handleRequests()
}

func main() {
	run()
}
