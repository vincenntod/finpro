package auth

import "time"

type Account struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type Transaction struct {
	Id        int       `json:"id"`
	OdaNumber int       `json:"oda_number"`
	Status    string    `json:"status"`
	Price     float32   `json:"price"`
	TotalData int       `json:"total_data"`
	CreatedAt time.Time `json:"created_at"`
}
