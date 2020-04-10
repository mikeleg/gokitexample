package customer

import "time"

type Customer struct {
	ID         int       `json:"id"`
	Legalname  string    `json:"legalname"`
	Createdate time.Time `json:"craetedate"`
	Modifydate time.Time `json:"modifydate"`
}
