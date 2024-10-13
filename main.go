package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("POST /execute", executeCommand)

	fmt.Println("server on port 8000")
	err := http.ListenAndServe(":8000", CorsMiddleware(router))

	if err != nil {
		fmt.Println("error starting server", err)
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving route /")
	io.WriteString(w, "Stack machine API")
}

func executeCommand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving route /execute")

	var command Command

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result, machineErr := StackMachine(command.Text)
	if machineErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Result{Status: http.StatusBadRequest, ErrorMsg: machineErr.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Result{Status: http.StatusOK, Data: result})
}
