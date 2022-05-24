package model

type Pokemon struct {
	Id         int    `csv:"Id"`
	Name       string `csv:"Name"`
	Type1      string `csv:"Type 1"`
	Type2      string `csv:"Type 2"`
	Total      int    `csv:"Total"`
	HP         int    `csv:"HP"`
	Attack     int    `csv:"Attack"`
	Defense    int    `csv:"Defense"`
	SpAtk      int    `csv:"Sp. Atk"`
	SpDef      int    `csv:"Sp. Def"`
	Speed      int    `csv:"Speed"`
	Generation int    `csv:"Generation"`
	Legendary  bool   `csv:"Legendary"`
}
