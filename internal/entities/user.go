package entities

type User struct {
	BaseGorm
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
