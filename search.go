package main

import (
	"encoding/json"
	"fmt"
	"my-brain/lib"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func SearchIndexes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	splitString := regexp.MustCompile(`[^\s]+`).FindAllString(query, -1)

	var docListScores = make(map[string]float64)

	for _, peice := range splitString {
		mixCaseWord := lib.AlphaNumeric.ReplaceAllString(peice, "") // removes
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
			//TODO check if any words are in the docs module name -> add 5 if so
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
