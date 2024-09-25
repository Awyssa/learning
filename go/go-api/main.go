package main

import (
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	users = make(map[string]User)
	mu    sync.Mutex
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/register", registerHandler)
	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
