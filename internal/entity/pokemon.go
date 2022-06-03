package entity

type Pokemon struct {
	Id                     uint64 `json:"id" csv:"id"`
	Name                   string `json:"name" csv:"name"`
	Height                 int    `json:"height" csv:"height"`
	IsDefault              bool   `json:"is_default" csv:"is_default"`
	Order                  int    `json:"order" csv:"order"`
	Weight                 int    `json:"weight" csv:"weight"`
	BaseExperience         int    `json:"base_experience" csv:"base_experience"`
	LocationAreaEncounters string `json:"location_area_encounters" csv:"location_area_encounters"`
}
