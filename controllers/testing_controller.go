package controllers

import (
	"net/http"
	"encoding/json"
	"log"
)

func ShowSuccess(w http.ResponseWriter, r *http.Request)  {
	log.Println("ShowSuccess...")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	text := "\"status\": \"OK\""
	result, _ := json.Marshal(text)

	w.Write(result)

}

func ShowError(w http.ResponseWriter, r *http.Request)  {
	log.Println("ShowError...")
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
}

func ShowInternalError(w http.ResponseWriter, r *http.Request) {
	log.Println("ShowInternalError...")
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
}