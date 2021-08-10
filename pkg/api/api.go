package api

const apiEndpoint = "https://pokeapi.co/api/v2/"

func FindByNameOrID(nameOrID string) (*Pokemon, error) {
	res := &Pokemon{}
	err := fetch(api("pokemon/"+nameOrID), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindEncountersByNameOrID(nameOrID string) ([]*Encounter, error) {
	var res []*Encounter
	err := fetch(api("pokemon/"+nameOrID+"/encounters"), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func api(path string) string {
	return apiEndpoint + path
}
