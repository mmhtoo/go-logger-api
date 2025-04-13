package helpers

import "time"

type apiBaseResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message string `json:"message"`
}

type apiDataResponse[T any] struct {
	*apiBaseResponse
	Data T `json:"data"`
}

type apiErrorResponse[E any] struct {
	*apiBaseResponse
	Error E `json:"error"`
}

func NewAPIBaseResponse(message string) *apiBaseResponse {
	return &apiBaseResponse{
		Timestamp: time.Now(),
		Message: message,
	}
}

func NewAPIDataResponse[T any](data T, message string) *apiDataResponse[T] {
	return &apiDataResponse[T]{
		apiBaseResponse: NewAPIBaseResponse(message),
		Data: data,
	}
}

func NewAPIErrorResponse[E any](error E, message string) *apiErrorResponse[E] {
	return &apiErrorResponse[E]{
		apiBaseResponse: NewAPIBaseResponse(message),
		Error: error,
	}
}