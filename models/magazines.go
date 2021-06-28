package models

type Magazine struct {
	ID               int    `json:"id"`
	Holder           string `json:"holder"`
	IsReserved       bool   `json:"isReserved"`
	Price            string `json:"price"`
	DateTillReserved string `json:"dateTillReserves"`
	CompanyID        string `json:"compID"`
}
