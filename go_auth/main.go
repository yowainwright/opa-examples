package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const OPAURL = "http://localhost:8181/v1/data/authz/allow"

type OPAInput struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type OPAResponse struct {
	Result bool `json:"result"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/protected", ProtectedEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	opaInput := OPAInput{
		Token: token,
		Role:  "admin",
	}

	allowed, err := isAllowed(opaInput)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if allowed {
		fmt.Fprintln(w, "Access granted")
	} else {
		http.Error(w, "Access denied", http.StatusForbidden)
	}
}

func isAllowed(input OPAInput) (bool, error) {
	client := &http.Client{}
	data, err := json.Marshal(map[string]interface{}{"input": input})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", OPAURL, bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var opaResp OPAResponse
	if err := json.NewDecoder(resp.Body).Decode(&opaResp); err != nil {
		return false, err
	}

	return opaResp.Result, nil
}
