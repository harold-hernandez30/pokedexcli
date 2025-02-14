package pokemonapi

import "fmt"

type Pokemon struct {
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Weight         int            `json:"weight"`
	Height         int            `json:"height"`
	Stats          []PokemonStat  `json:"stats"`
	Types          []PokemonTypes `json:"types"`
}

func (pokemon Pokemon) Inspect() {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	pokemon.PrintStats()
	pokemon.PrintTypes()
}

func (pokemon Pokemon) PrintStats() {
	fmt.Print("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v\n", stat.String())
	}
}

func (pokemon Pokemon) PrintTypes() {
	fmt.Print("Types:\n")
	for _, Type := range pokemon.Types {
		fmt.Printf("  - %v\n", Type.Type.Name)
	}
}
