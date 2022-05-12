package response

type ApiResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
}
