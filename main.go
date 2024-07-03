package main

import (
	"log"
	"net/http"
)

func Hellohandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi from handler"))
}
func main() {
	http.HandleFunc("/", Hellohandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
