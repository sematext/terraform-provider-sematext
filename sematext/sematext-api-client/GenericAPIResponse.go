package sematext

// GenericAPIResponse TODO Doc Comment
type GenericAPIResponse struct {
	Data    interface{} `json:"data"`
	Errors  []Error     `json:"errors"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}
