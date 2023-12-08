package web

type ResponseSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseFailure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}