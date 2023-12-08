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

type ResponseCreateRole struct {
	Name string `json:"name"`
	Description string `json:"description"`
	User []ResponseUser `json:"user"`
}

type ResponseArtikel struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type ResponseOrder struct {
	UUID string `json:"uuid"`
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	TTL       time.Time `json:"tanggal_lahir"`
}

type ResponseLogin struct {
	FName     string `json:"first_name"`
	LName     string `json:"last_name"`
	Email     string `json:"email"`
	Token 	  string `json:"token"`
}