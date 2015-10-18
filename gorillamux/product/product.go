package product

import "encoding/json"

type Product struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Qty  int16  `json:"qty"`
}

func (p Product) String() string {
	buffer, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(buffer)
}
