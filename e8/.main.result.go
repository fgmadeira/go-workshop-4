package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"diconium.com/madeifra/go-workshop-4/e7/pokemon"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type trainerHandler struct {
	region string
}

func (ph trainerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello %s", ph.region)))
}

func pokemonsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	e.Encode(pokemon.Wild)
}

func createPokemon(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		p := &pokemon.Pokemon{}

		decodeErr := decoder.Decode(&p)
		if decodeErr != nil {
			http.Error(w, "Bad request - Could parse pokemon.", 400)
			return
		}

		err := db.QueryRow(`
			INSERT INTO pokemons (name, type, weight)
			VALUES($1,$2, $3)
			RETURNING id
		`, p.Name, p.Type, p.Weight).Scan(&p.ID)

		if err != nil {
			http.Error(w, "Bad request - Could not create pokemon.", 400)
		} else {
			e := json.NewEncoder(w)
			e.Encode(p)
		}

	}
}

func pokemonHandlerFunc(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Bad request - Not an Id!", 400)
		}

		p := &pokemon.Pokemon{}
		row := db.QueryRow(`
			SELECT id, name, type, weight
			FROM public.pokemons
			WHERE id = $1
		`, id)
		rerr := row.Scan(&p.ID, &p.Name, &p.Type, &p.Weight)
		if rerr != nil {
			http.Error(w, "Pokemon Not Found", 404)
		} else {
			e := json.NewEncoder(w)
			e.Encode(p)
		}
	}
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		h.ServeHTTP(w, r)
		fmt.Println("Handler Finished!")
	})
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to connect to database: %v", err))
	}
	defer db.Close()

	r := mux.NewRouter()
	th := trainerHandler{region: "Kanto"}
	r.Handle("/trainer", &th)

	pokemonHandler := http.HandlerFunc(pokemonHandlerFunc(db))
	pokemonsHandler := http.HandlerFunc(pokemonsHandlerFunc)

	r.Handle("/pokemon", loggingMiddleware(pokemonsHandler)).Methods("GET")
	r.Handle("/pokemon", loggingMiddleware(createPokemon(db))).Methods("POST")
	r.Handle("/pokemon/{id}", loggingMiddleware(pokemonHandler))

	http.ListenAndServe(":8000", r)
}
