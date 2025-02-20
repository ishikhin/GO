package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "strconv"
	// "strings"
)

type Note struct {
	Text string `json:"text"`
}

type Library struct {
	Library_ID int    `json:"library_id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Adress     string `json:"adress"`
	Phone      int    `json:"phone"`
	Email      string `json:"email"`
}

var libraries []Library

func main() {
	// Инициализируем тестовыми данными
	// notes = []Note{
	// 	{Text: "First note"},
	// 	{Text: "Second note"},
	// }

	libraries = []Library{
		{Library_ID: 12345678, Country: "Russia", City: "Moscow", Adress: "Konenkova", Phone: 88005553535, Email: "MSL@email.ru"}}

	http.HandleFunc("/libraries", libraryHandler)
	// http.HandleFunc("/notes/", noteByIdHandler)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func libraryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libraries)
	case "POST":
		var newLibrary Library
		err := json.NewDecoder(r.Body).Decode(&newLibrary)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newLibrary.Library_ID < 0 {
			http.Error(w, "Text field cannot be empty", http.StatusBadRequest)
			return
		}
		libraries = append(libraries, newLibrary)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newLibrary)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// func noteByIdHandler(w http.ResponseWriter, r *http.Request) {
// 	parts := strings.Split(r.URL.Path, "/")
// 	if len(parts) < 3 {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	idStr := parts[2]
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil || id < 0 || id >= len(notes) {
// 		http.Error(w, "Note not found", http.StatusNotFound)
// 		return
// 	}

// 	switch r.Method {
// 	case "GET":
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(notes[id])
// 	case "DELETE":
// 		notes = append(notes[:id], notes[id+1:]...)
// 		w.WriteHeader(http.StatusNoContent)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }
