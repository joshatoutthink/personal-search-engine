package contentType

import (
	"my-brain/lib"
)

type ModuleInterface interface {
	CollectLocations() ([]string, error)
	TokenizeDoc(id string, fpath string) lib.Doc
}

type modList map[string]ModuleInterface

type ModuleCollection struct {
	Mods modList
}

func (M *ModuleCollection) add(name string, module ModuleInterface) {
	M.Mods[name] = module
}

func (M *ModuleCollection) AddModules() {
	M.Mods = make(modList) // clears
	M.add("digital-garden", &DigitalGarden)
}
