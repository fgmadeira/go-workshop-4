package main

import (
	"fmt"
	"net/http"
)

type trainerHandler struct {
	region string
}

func (ph trainerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello %s", ph.region)))
}

func main() {
	ph := trainerHandler{region: "Kanto"}
	http.Handle("/trainer", &ph)

	http.ListenAndServe(":8000", nil)
}