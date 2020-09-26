package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome To the Land of Tacos</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("localhost:3000 has been started")
	http.ListenAndServe(":3000", nil)
}
