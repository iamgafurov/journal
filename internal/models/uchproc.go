package models

import "time"

type Topic struct {
	Id        int64     `json:"id"`
	Kvd       int64     `json:"-"`
	Kst       int64     `json:"-"`
	Pnn       string    `json:"-"`
	Chruchgod string    `json:"-"`
	VchAccess string    `json:"-"`
	DtActive  time.Time `json:"-"`
	IsuSotId  int64     `json:"-"`
	ChActive  int       `json:"-"`
	Cnzap     string    `json:"cnzap"`
	Dtzap     time.Time `json:"dtzap"`
	Tema      string    `json:"tema"`
	KolLek    int       `json:"kol_lek"`
	KolSem    int       `json:"kol_sem"`
	KolPrak   int       `json:"kol_prak"`
	KolLab    int       `json:"kol_lab"`
	KolKmd    int       `json:"kol_kmd"`
	KolObsh   int       `json:"kol_obsh"`
	Editable  bool      `json:"editable"`
}

type Statement struct {
	Id        int64  `json:"id"`
	Kst       int64  `json:"kst"`
	Kpr       int64  `json:"kpr"`
	Kas       int64  `json:"kas"`
	VchAccess string `json:"vch_access"`
	Chruchgod string `json:"chruchgod"`
}
