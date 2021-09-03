package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type TokenList map[string]float64

type Doc struct {
	Content string    `json:"content"`
	Heading string    `json:"heading"`
	Id      string    `json:"id"`
	Path    string    `json:"path"`
	Tokens  TokenList `json:"tokens"`
}

var Indexes map[string]Doc

//Starts server
// place all routes here
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

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
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("couldnt read index file")
		fmt.Println(err)
	}
	fmt.Println("success")

	defer file.Close()
	byteVal, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

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
