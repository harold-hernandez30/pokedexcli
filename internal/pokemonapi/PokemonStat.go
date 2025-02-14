package pokemonapi

import "fmt"

type Stat struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	BaseStat  int  `json:"base_stat"`
	Statistic Stat `json:"stat"`
}

func (stat PokemonStat) String() string {
	return fmt.Sprintf("%v: %d", stat.Statistic.Name, stat.BaseStat)
}
