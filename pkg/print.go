package pokefinder

import (
	"fmt"
	"strings"

	"github.com/santazaaa/pokefinder/pkg/api"
)

func printPokemon(p *api.Pokemon) {
	fmt.Printf("ID: %v\nName: %v\n", p.ID, p.Name)

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Println("- " + t.Type.Name)
	}

	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Println("- " + stat.Stat.Name + ": " + fmt.Sprint(stat.BaseStat))
	}
}

func printEncountersInKanto(ens []*api.Encounter) {
	if len(ens) == 0 {
		fmt.Println("-")
		return
	}

	for i, enc := range ens {
		// we only interest in Kanto area's encounters (starts with kanto-)
		if !strings.HasPrefix(enc.LocationArea.Name, "kanto") {
			continue
		}
		fmt.Printf("%d. %v\n", i+1, enc.LocationArea.Name)

		// print unique methods for each location
		methodsMap := make(map[string]bool)
		for _, v := range enc.VersionDetails {
			for _, ed := range v.EncounterDetails {
				methodsMap[ed.Method.Name] = true
			}
		}
		for method := range methodsMap {
			fmt.Println("- " + method)
		}
	}
}
