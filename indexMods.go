package main

import (
	"encoding/json"
	"fmt"
	"my-brain/lib"
	"net/http"
)

func ReIndexMods(w http.ResponseWriter, r *http.Request) {

	Modules.AddModules()

	Indexes = make(map[string]lib.Doc)
	//Mods should be the list of Modules in
	fmt.Println(len(Modules.Mods))
	for modName, mod := range Modules.Mods {
		//MODULE SHOULD IMPLEMENT LINES BELOW
		fpaths, err := mod.CollectLocations()
		if err != nil {
			fmt.Println(err)
		}
		for indx, fpath := range fpaths {
			// TODO get content mod.GetContent (this is just what ever is at the location given)
			id := fmt.Sprintf("%s-%d", modName, indx)

			doc := mod.TokenizeDoc(id, fpath)
			Indexes[id] = doc
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
