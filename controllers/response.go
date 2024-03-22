package controllers

import (
	"encoding/json"
	model "martini/models"
	"net/http"
)

// Error Response
func sendErrorResponse(w http.ResponseWriter, status int, message string) {
	var response model.ErrorResponse
	response.Status = status
	response.Message = message

	// Mengirimkan response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Success Response
func sendSuccessResponse(w http.ResponseWriter, status int, message string) {
	var response model.SuccessResponse
	response.Status = status
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Get All User
func sendGetUsersResponse(w http.ResponseWriter, status int, message string, data []model.Users) {
	var response model.UsersResponse
	response.Data = data
	response.Status = status
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
