package main

// MAIN.go
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"my-brain/contentType"
	"my-brain/lib"
	"net/http"
)

var Indexes map[string]lib.Doc
var Modules contentType.ModuleCollection

//Starts server
// place all routes here
func handleRequests() {
	// creates a new instance of a mux router
	// todo replace with module loader and module system

	myRouter := mux.NewRouter().StrictSlash(true)

	// routeshttp.Handle("/", http.FileServer(http.Dir("./static")))
	myRouter.Handle("/", http.FileServer(http.Dir("./client/build")))
	myRouter.HandleFunc("/all", returnAllIndexes)
	myRouter.HandleFunc("/reindex", ReIndexMods)
	myRouter.HandleFunc("/search", SearchIndexes)

	originsOk := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk)(myRouter)))
}

type DocListScore struct {
	id    string
	score float64
}

func returnAllIndexes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: returnallindexes")
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
