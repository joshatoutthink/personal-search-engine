package main

import (
	"encoding/json"
	"fmt"
	"my-brain/lib"
	"net/http"
)

func ReIndexMods(w http.ResponseWriter, r *http.Request) {
	Indexes = make(map[string]lib.Doc)
	//Mods should be the list of Modules in
	for modIndx, mod := range Mods {
		//MODULE SHOULD IMPLEMENT LINES BELOW
		fpaths := lib.PathsFromDir(mod) //TODO: will be mod.CollectLocations

		for indx, fpath := range fpaths {
			// TODO get content mod.GetContent (this is just what ever is at the location given)
			id := modIndx + indx                   //module specific
			content := string(lib.ReadFile(fpath)) //module.specific

			//TODO mod.TokenizeDoc
			doc := lib.Doc{
				Tokens:  lib.Tokenize(content),
				Content: content,
				Heading: lib.MarkdownTitle(fpath),
				Path:    fpath,
				Id:      fmt.Sprint(id),
			}

			Indexes[fmt.Sprint(id)] = doc
		}

		//END MODULE SPECIFIC
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
