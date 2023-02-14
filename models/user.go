package models

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Friend struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	Id         string   `json:"id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Gender     string   `json:"gender"`
	Address    Address  `json:"address"`
	Friends    []Friend `json:"friends"`
	CardNumber string   `json:"card_number"`
	Birthday   string   `json:"birthday"`
	Profession string   `json:"profession"`
}
