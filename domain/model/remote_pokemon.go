package model

type RemotePokemon struct {
	Id                     uint64 `json:"id"`
	Name                   string `json:"name"`
	Height                 int    `json:"height"`
	IsDefault              bool   `json:"is_default"`
	Order                  int    `json:"order"`
	Weight                 int    `json:"weight"`
	BaseExperience         int    `json:"base_experience"`
	LocationAreaEncounters string `json:"location_area_encounters"`
}
