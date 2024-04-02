package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"token-pagination/utils"
)

const pageSize = 2 // Number of items per page

func main() {
	http.HandleFunc("/data", handleData)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// Get token from query parameters
	token := r.URL.Query().Get("token")

	// Parse token
	lastId, err := utils.DecodeToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	// Query data from database
	data, nextID, totalCount := utils.QueryData(lastId, pageSize)
	if err != nil {
		http.Error(w, "Failed to query data", http.StatusInternalServerError)
		return
	}

	// Encode nextToken
	var encodedNextToken string
	if nextID != -1 {
		encodedNextToken = utils.EncodeToken(nextID)
	}

	// Create paginated response
	response := utils.PaginatedResponse{
		Data:       data,
		NextToken:  encodedNextToken,
		TotalCount: totalCount,
	}

	// Convert response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set content type and send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
