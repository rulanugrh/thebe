package web

import "time"

type ResponseUser struct {
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Role      string `json:"role"`
	TTL       time.Time `json:"tanggal_lahir"`
}

type ResponseRole struct {
	Name string `json:"name"`
	Description string `json:"description"`
	User []ResponseUser `json:"user"`
}

type ResponseOrder struct {
	UUID string `json:"uuid"`
	User ResponseUser `json:"user"`
}