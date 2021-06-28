package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome")
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
