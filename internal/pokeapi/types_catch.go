package pokeapi

type Pokemon struct {
	LocationAreaEncounters string `json:"location_area_encounters"`
	BaseExperience         int    `json:"base_experience"`
	Weight                 int    `json:"weight"`
	Height                 int    `json:"height"`
	Name                   string `json:"name"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}