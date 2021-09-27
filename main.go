package main

// MAIN.go
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"my-brain/lib"
	"net/http"
	"regexp"
	"sort"
	"strings"
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

			//ignore files

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
	query := r.URL.Query().Get("q")

	splitString := regexp.MustCompile("[^\\s]+").FindAllString(query, -1)
	var docListScores = make(map[string]float64)

	for _, peice := range splitString {
		mixCaseWord := lib.AlphaNumeric.ReplaceAllString(peice, "")
		word := strings.ToLower(mixCaseWord)

		//Do we really need to look through all the
		//indexes every single time ??
		for id, doc := range Indexes {

			_, ok := docListScores[id]
			if !ok {
				docListScores[id] = float64(0)
			}

			timesInDoc := float64(0)
			_, ok = doc.Tokens[word]
			if ok {
				timesInDoc = doc.Tokens[word]
			}

			docListScores[id] += timesInDoc
			if len(doc.Heading) > 1 {
				if strings.Contains(strings.ToLower(doc.Heading), word) && !lib.StringInArr(lib.StopWords, word) {
					docListScores[id] += float64(5)
				}
			}

		}
	}

	//Sort the scores
	docListScoresArr := mapToArray(docListScores)
	sort.Sort(byDocListScore(docListScoresArr))

	sortedDocList := make([]lib.Doc, 0)
	for _, scoreDoc := range docListScoresArr {
		if scoreDoc.score > float64(0) {
			sortedDocList = append(sortedDocList, Indexes[scoreDoc.id])
		}
	}

	_, err := json.MarshalIndent(&sortedDocList, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(sortedDocList)

}

type DocListScore struct {
	id    string
	score float64
}

func mapToArray(m map[string]float64) []DocListScore {
	var Arr []DocListScore
	for id, score := range m {
		Arr = append(Arr, DocListScore{id, score})
	}
	return Arr
}

//Some Go Sorting Magic
type byDocListScore []DocListScore

func (s byDocListScore) Len() int {
	return len(s)
}
func (s byDocListScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byDocListScore) Less(i, j int) bool {
	docI := s[i]
	docJ := s[j]
	return docI.score > docJ.score
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
