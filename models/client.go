package models

type Client struct {
	Name  string `json:"name"`
	Login string `json:"login"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
