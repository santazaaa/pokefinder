# Pokefinder
Pokefinder is a simple go terminal app that can be used to search for pokemon data from https://pokeapi.co/.

## Requirements
- Go v1.15 or above (Download at https://golang.org/dl/)

## Installation
Install go modules by running
```
go install
```

## Development
### Run from code
```
go run . <pokemon name or id>
```

### Build
This will compile the code to an executable file depending on the current platform.
```
go build
```

## How to Use
### Search for pokemon data by name or id
```
./pokefinder <pokemon name or id>
```
Example
```
./pokefinder ditto
```
Result
```
Found Pokemon Data
ID: 132
Name: ditto
Types:
- normal
Stats:
- hp: 48
- attack: 48
- defense: 48
- special-attack: 48
- special-defense: 48
- speed: 48
Encounter Location(s) and Method(s) in Kanto:
1. kanto-route-13-area
- walk
2. kanto-route-14-area
- walk
3. kanto-route-15-area
- walk
4. kanto-route-23-area
- walk
```

* `Not Found` text will be display if pokemon data is not returned from the api
### Clearing API cache
* The cache will expire automatically every 7 days.
* Remove `pokeapi-cache.txt` next to the executable to clear the cache manually.
