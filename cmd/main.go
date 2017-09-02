package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	h := newHandler()
	http.ListenAndServe(":"+port, h)
}
