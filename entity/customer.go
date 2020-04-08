package entity

type Customer struct {
	ID        uint   `json:"id,omitempty"`
	Legalname string `json:"legalname"`
}
