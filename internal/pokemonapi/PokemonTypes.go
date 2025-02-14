package pokemonapi

type PokemonTypes struct {
	Slot int         `json:"slot"`
	Type PokemonType `json:"type"`
}
