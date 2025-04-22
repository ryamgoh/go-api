package api

import (
	"encoding/json"
	"net/http"
)

// Credit Balance Params
type CreditBalanceParams struct {
	Username string
}

// Credit Balance Resp
type CreditBalanceResponse struct {
	// Code
	Code int

	// Account Credit Balance
	Balance float64
}

// Error message
type Error struct {
	// Error Code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	// we use this if want a generic error message, the client doesn't need to know
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
)
