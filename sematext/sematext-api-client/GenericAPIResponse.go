package sematext

// GenericAPIResponse TODO Doc Comment
type GenericAPIResponse struct {
	Data    []byte  `json:"data"` // TODO - freeform conversion {...} ?
	Errors  []Error `json:"errors"`
	Message string  `json:"message"`
	Success bool    `json:"success"`
}
