package pokemon

import (
	"fmt"
)

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Weight int    `json:"weight"`
}

var Wild = []Pokemon{
	{
		ID:     1,
		Name:   "Bulbasaur",
		Type:   "grass",
		Weight: 69,
	},
	{
		ID:     2,
		Name:   "Charmander",
		Type:   "fire",
		Weight: 85,
	},
	{
		ID:     3,
		Name:   "Squirtle",
		Type:   "water",
		Weight: 90,
	},
	{
		ID:     4,
		Name:   "Caterpie",
		Type:   "bug",
		Weight: 29,
	},
	{
		ID:     5,
		Name:   "Pidgey",
		Type:   "flying",
		Weight: 18,
	},
	{
		ID:     6,
		Name:   "Rattata",
		Type:   "normal",
		Weight: 35,
	},
	{
		ID:     7,
		Name:   "Articuno",
		Type:   "ice",
		Weight: 554,
	},
	{
		ID:     8,
		Name:   "Zapdos",
		Type:   "electric",
		Weight: 526,
	},
	{
		ID:     9,
		Name:   "Moltres",
		Type:   "fire",
		Weight: 600,
	},
	{
		ID:     10,
		Name:   "Mewtwo",
		Type:   "psychic",
		Weight: 1220,
	},
}

var WildString = fmt.Sprint(Wild)
