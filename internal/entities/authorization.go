package entities

type Registration struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	PhoneNumber string `json:"phone_number" validate:"required,len=12"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
}

type OTPVerify struct {
	PhoneNumber string `json:"phone_number"`
	OtpCode     string `json:"otp_code" validate:"required,len=4"`
}

// OtpSession — structure is designed to form a payload,
// which will later be unloaded into a redis
type OtpSession struct {
	CreatedAt     int64  `json:"created_at"`
	SendAttempts  int    `json:"send_attempts"`
	CheckAttempts int    `json:"check_attempts"`
	Message       string `json:"-"`
	Phone         string `json:"phone"`
	Code          string `json:"code" validate:"required,len=4"`
	Password      string `json:"password" validate:"required,min=8,max=100"`
}
