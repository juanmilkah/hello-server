package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main(){
  http.HandleFunc("/", HandleHomeRoute)
  http.HandleFunc("/hello", HandleHello)

  port := ":8000"
  fmt.Println("Server Listening at port 8000")
  err := http.ListenAndServe(port, nil)
  if err != nil{
    log.Fatal(err)
  }
}

func HandleHomeRoute(w http.ResponseWriter, r *http.Request){
  paths := [...]string{"/","/hello"}
  joinedPaths := strings.Join(paths[:], ", ")
  
  w.Header().Set("Content-Type", "text/plain")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(joinedPaths))
}

func HandleHello(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "text/plain")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("World"))
}
