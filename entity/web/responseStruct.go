package web

import (
	"html/template"
	"time"
)

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
	Content template.HTML `json:"content"`
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

type ResponseEvent struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price int `json:"price"`
	Participant []ListParticipant `json:"participant"`
	Delegasi []ListDelegasi `json:"delegasi"`
}

type ListParticipant struct {
	FName string `json:"first_name"`
	LName string `json:"last_name"`
}

type ListDelegasi struct {
	FName string `json:"first_name"`
	LName string `json:"last_name"`
}