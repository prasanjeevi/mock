package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server listening on " + port)
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	status := os.Getenv("STATUS")
	code := r.URL.Query().Get("code")
	if code == "400" || status == "400" {
		w.WriteHeader(http.StatusBadRequest)
	} else if code == "500" || status == "500" {
		w.WriteHeader(http.StatusInternalServerError)
	} else if code != "" {
		statusCode, err := strconv.Atoi(code)
		if err == nil && statusCode > 90 && statusCode < 512 {
			w.WriteHeader(statusCode)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	fmt.Println("Handling request Status: " + status + " Code: " + code)
	return
}
