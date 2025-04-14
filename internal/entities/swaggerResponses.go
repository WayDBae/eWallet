package entities

// Ping
//
// swagger:response ping
type swaggerPing struct {
	Data struct {
		// default: Успешно
		Message string `json:"message"`
		// default: Clean arch pong! 🥎
		Payload string `json:"payload"`
	}
}

// swagger:response passwordLen
type swaggerPasswordLen struct {
	Data struct {
		// default: Пароль должен содержать как минимум 8 символов
		Message string `json:"message"`
		// default: null
		Payload string `json:"payload"`
	}
}
