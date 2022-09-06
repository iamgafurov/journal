package models

import "time"

type Topic struct {
	Id       int64     `json:"id"`
	Cnzap    string    `json:"cnzap"`
	Dtzap    time.Time `json:"dtzap"`
	Tema     string    `json:"tema"`
	KolLek   int       `json:"kolLek"`
	KolSem   int       `json:"kolSem"`
	KolPrak  int       `json:"kolPrak"`
	KolLab   int       `json:"kolLab"`
	KolKmd   int       `json:"kolKmd"`
	KolObsh  int       `json:"kolObsh"`
	Editable bool      `json:"editable"`
}
