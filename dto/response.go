package dto

type Response[T any] struct {
	Status       int     `json:"status"`
	Message      string  `json:"message"`
	Data         T       `json:"data"`
	ErrorMessage *string `json:"error"`
}

type UserResponse Response[interface{}]
