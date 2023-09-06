package dto

type Response[T any] struct {
	Data         T       `json:"data"`
	ErrorMessage *string `json:"error"`
}

type UserResponse Response[interface{}]
