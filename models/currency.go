package models

type Currency struct {
	ID           int     `db:"id" json:"Cur_ID"`
	Date         string  `db:"date" json:"Date"`
	Abbreviation string  `db:"abbreviation" json:"Cur_Abbreviation"`
	OfficialRate float64 `db:"official_rate" json:"Cur_OfficialRate"`
	Scale        int     `db:"scale" json:"Cur_Scale"`
	Name         string  `db:"name" json:"Cur_Name"`
}
