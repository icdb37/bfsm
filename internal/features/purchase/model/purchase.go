package model

type ExtraCost struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Cost int64  `json:"cost"` //额外费用，分
}
