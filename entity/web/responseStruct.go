package web

import (
	"html/template"
)

type ResponseUser struct {
	ID uint `json:"id"`
	Name     string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Role      string `json:"role"`
}

type ResponseRole struct {
	ID uint `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	User        []ResponseUser `json:"user"`
}

type ResponseCreateRole struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	User        []ResponseUser `json:"user"`
}

type ResponseArtikel struct {
	ID uint `json:"id"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

type ResponseOrder struct {
	UUID       string `json:"uuid"`
	Name      string `json:"name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Telephone  string `json:"telephone"`
	EventName  string `json:"event_name"`
	EventPrice int    `json:"event_price"`
}

type ResponseLogin struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type ResponseEvent struct {
	ID uint `json:"id" form:"id"`
	Name        string               `json:"name" form:"name"`
	Description string               `json:"description" form:"description"`
	Price       int                  `json:"price" form:"price"`
	Participant []ListParticipant    `json:"participant" form:"participant"`
	Submission  []ResponseSubmission `json:"submission" form:"submission"`
}

type ResponseEventRekarda struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Price       int               `json:"price"`
	Participant []ListParticipant `json:"participant"`
	Delegasi    []ListDelegasi    `json:"delegasi"`
}

type ResponseSubmission struct {
	Name     string `json:"name"`
	Event    string `json:"event"`
	Filename string `json:"filename"`
}

type ResponseOrderRekarda struct {
	UUID       string         `json:"uuid"`
	Name      string         `json:"name"`
	Email      string         `json:"email"`
	Address    string         `json:"address"`
	Telephone  string         `json:"telephone"`
	EventName  string         `json:"event_name"`
	EventPrice int            `json:"event_price"`
	Delegasi   []ListDelegasi `json:"delegasi"`
}

type ListParticipant struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type ListDelegasi struct {
	Name  string `json:"name"`
	Gender string `json:"gender"`
}

type ResponsePayment struct {
	SnapURL string `json:"snap_url"`
	Token   string `json:"token"`
	Name    string `json:"name"`
	Event   string `json:"event"`
	Price   int    `json:"price"`
}

type ResponseForPayment struct {
	Name   string `json:"name"`
	Event  string `json:"event"`
	Price  int    `json:"price"`
	Status string `json:"status"`
}
type Error struct {
	Message string
	Code int
}

type ValidationList struct {
	Field string      `json:"field"`
	Error interface{} `json:"error"`
}

type ValidationError struct {
	Message string           `json:"message"`
	Errors  []ValidationList `json:"error"`
}

func (err ValidationError) Error() string {
	return err.Message
}

type WebValidationError struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
}

func (err Error) Error() string {
	return err.Message
}



type StatusPayment struct {
	Currency string `json:"currency"`
	FraudStatus string `json:"fraud_status"`
	GrossAmount string `json:"gross_amount"`
	OrderID string `json:"order_id"`
	PaymentType string `json:"payment_type"`
	StatusCode string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	TransactionID string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	TransactionTime string `json:"transaction_time"`

}