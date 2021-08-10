package pokefinder

import (
	"fmt"
	"strings"

	"github.com/santazaaa/pokefinder/pkg/api"
)

func SearchPokemon(nameOrID string) {
	defer api.SaveCacheFile()

	nameOrID = strings.ToLower(nameOrID)

	p, err := api.FindByNameOrID(nameOrID)
	if err != nil {
		if api.IsNotFoundError(err) {
			fmt.Println("Not Found")
		} else {
			fmt.Println("Error when trying to find a pokemon: ", err)
		}
		return
	}

	encs, err := api.FindEncountersByNameOrID(nameOrID)
	if err != nil {
		if api.IsNotFoundError(err) {
			fmt.Println("Not Found")
		} else {
			fmt.Println("Error when trying to find pokemon encounters: ", err)
		}
		return
	}

	// print pokemon data
	fmt.Println("Found Pokemon Data")
	printPokemon(p)

	// print encounters data
	// we're only interested in Kanto area
	encsInKanto := filterEncountersInLocation(encs, "kanto")
	fmt.Println("Encounter Location(s) and Method(s) in Kanto:")
	printEncountersInKanto(encsInKanto)
}

func filterEncountersInLocation(encs []*api.Encounter, locPrefix string) []*api.Encounter {
	var res []*api.Encounter
	for _, enc := range encs {
		if locPrefix != "" && !strings.HasPrefix(enc.LocationArea.Name, locPrefix) {
			continue
		}
		res = append(res, enc)
	}
	return res
}
