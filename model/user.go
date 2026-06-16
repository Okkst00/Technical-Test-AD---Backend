package model

type User struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	Status       bool   `json:"status"`
	MembershipID int    `json:"membership_id"`
}

type RegisterRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}