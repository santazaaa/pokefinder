package main

import (
	"fmt"
	"os"

	pokefinder "github.com/santazaaa/pokefinder/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please specify a pokemon name or id")
		return
	}

	nameOrID := args[0]
	pokefinder.SearchPokemon(nameOrID)
}
