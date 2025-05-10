package entities

type Currency struct {
	BaseGorm
	Code      string `json:"code"`
	Character string `json:"character"`
}
