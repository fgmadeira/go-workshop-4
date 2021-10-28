package main

import (
	"fmt"
	"net/http"

	"diconium.com/madeifra/go-workshop-4/pokemon"
)

type trainerHandler struct {
	region string
}

func (ph trainerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello %s", ph.region)))
}

func pokemonHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(pokemon.WildString))
}

func main() {
	ph := trainerHandler{region: "Kanto"}
	http.Handle("/trainer", &ph)

	http.HandleFunc("/pokemon", pokemonHandlerFunc)

	http.ListenAndServe(":8000", nil)
}