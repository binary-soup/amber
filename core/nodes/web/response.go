package web

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Error:   err,
	}
}

type DataResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func NewSuccessDataResponse(data any) DataResponse {
	return DataResponse{
		Success: true,
		Data:    data,
	}
}

func sendJSONResponse(w http.ResponseWriter, status int, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

func sendErrorResponse(w http.ResponseWriter, err error) {
	sendJSONResponse(w, http.StatusBadRequest, NewErrorResponse(err.Error()))
}
