package main

import (
	"fmt"
	"net/http"
	"strconv"

	"diconium.com/madeifra/go-workshop-4/pokemon"
	"github.com/gorilla/mux"
)

type trainerHandler struct {
	region string
}

func (ph trainerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello %s", ph.region)))
}

func pokemonsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(pokemon.WildString))
}

func pokemonHandlerFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad request - Not an Id!", 400)
	}

	w.Write([]byte(fmt.Sprint(pokemon.Wild[id])))
}

func main() {
	r := mux.NewRouter()
	ph := trainerHandler{region: "Kanto"}
	r.Handle("/trainer", &ph)

	r.HandleFunc("/pokemon", pokemonsHandlerFunc)
	r.HandleFunc("/pokemon/{id}", pokemonHandlerFunc)

	http.ListenAndServe(":8000", r)
}
