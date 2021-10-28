package main

import (
	"fmt"
	"net/http"
	"strconv"

	"diconium.com/madeifra/go-workshop-4/e7/pokemon"
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

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		h.ServeHTTP(w, r)
		fmt.Println("Handler Finished!")
	})
}

func main() {
	r := mux.NewRouter()
	th := trainerHandler{region: "Kanto"}
	r.Handle("/trainer", &th)

	pokemonHandler := http.HandlerFunc(pokemonsHandlerFunc)
	pokemonsHandler := http.HandlerFunc(pokemonsHandlerFunc)

	r.Handle("/pokemon", loggingMiddleware(pokemonsHandler))
	r.Handle("/pokemon/{id}", loggingMiddleware(pokemonHandler))

	http.ListenAndServe(":8000", r)
}

