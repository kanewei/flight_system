package entity

type Airport struct {
	Code      string `json:"code"`
	Terminal  string `json:"terminal"`
	City      string `json:"city"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}
