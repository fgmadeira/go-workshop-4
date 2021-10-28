package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"diconium.com/madeifra/go-workshop-4/pokemon"
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
	pr, _ := regexp.Compile(`/pokemon/(\d+)`)
	id, err := strconv.Atoi(pr.FindStringSubmatch(r.URL.Path)[1])
	if err != nil {
		http.Error(w, "Bad request - Not an Id!", 400)
	}

	w.Write([]byte(fmt.Sprint(pokemon.Wild[id])))
}

func main() {
	ph := trainerHandler{region: "Kanto"}
	http.Handle("/trainer", &ph)

	http.HandleFunc("/pokemon", pokemonsHandlerFunc)
	http.HandleFunc("/pokemon/", pokemonHandlerFunc)

	http.ListenAndServe(":8000", nil)
}
