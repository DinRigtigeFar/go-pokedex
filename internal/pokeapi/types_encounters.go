package pokeapi

type RespEncounters struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
	LocationName string `json:"name"`
}

type RespFoundAtLocations []struct {
	LocationArea struct {
		Name string `json:"name"`
	} `json:"location_area"`
}