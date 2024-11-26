package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Path struct {
	URLPath string
}

func main() {
	http.HandleFunc("/", HandleHomeRoute)
	http.HandleFunc("/hello", HandleHello)

	port := ":8000"
	fmt.Println("Server Listening at port 8000")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleHomeRoute(w http.ResponseWriter, r *http.Request) {
	paths := map[string]Path{
		"Home": {
			URLPath: "/",
		},
		"Hello": {
			URLPath: "/hello",
		},
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paths)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("World"))
}
