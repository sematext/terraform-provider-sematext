package sematext

type GenericApiResponse struct {
	data    []byte // TODO - freeform conversion {...} ?
	errors  []Error
	message string
	success bool
}
