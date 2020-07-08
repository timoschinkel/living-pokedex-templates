package main

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Pokedex struct {
	Generation string 		`json:"generation"`
	Identifier string		`json:"identifier"`
	Name string				`json:"name"`
	Games []string			`json:"games"`
	Region string			`json:"region"`
	MaxNationalDex int		`json:"-"`
	Pokemon []Pokemon		`json:"pokemon"`
}

type Pokemon struct {	
	DexNumber int			`json:"dex_number"`
	NationalDexNumber int	`json:"national_dex_number"`
	Identifier string		`json:"identifier"`
	Name string				`json:"name"`
}

var pokedexes = []Pokedex{ 
	Pokedex{"I", "red-blue-yellow", "Kanto dex", []string{ "Red", "Blue", "Yellow" }, "Kanto", 151, []Pokemon{} },
	Pokedex{"II", "gold-silver-crystal", "National Johto dex", []string{ "Gold", "Silver", "Crystal" }, "Johto", 251, []Pokemon{} },
	Pokedex{"III", "ruby-sapphire-emerald", "National Hoenn dex", []string{ "Ruby", "Sapphire", "Emerald" }, "Hoenn", 386, []Pokemon{} },
	Pokedex{"III", "firered-leafgreen", "National updated Kanto dex", []string{ "FireRed", "Leafgreen" }, "Kanto", 386, []Pokemon{} },
	Pokedex{"IV", "diamond-pearl-platinum", "National Sinnoh dex", []string{ "Diamond", "Pearl", "Platinum" }, "Sinnoh", 490, []Pokemon{} },
	Pokedex{"IV", "heartgold-soulsilver", "National updated Johto dex", []string{ "HeartGold", "SoulSilver" }, "Johto", 490, []Pokemon{} },
	Pokedex{"V", "black-white-black-2-white-2", "National Unova dex", []string{ "Black", "White", "Black 2", "White 2" }, "Unova", 649, []Pokemon{} },
	Pokedex{"VI", "x-y", "National Kalos dex", []string{ "X", "Y" }, "Kalos", 718, []Pokemon{} },
	Pokedex{"VI", "omega-ruby-alpha-sapphire", "Nation updated Hoenn dex", []string{ "Omega Ruby", "Alpha Sapphire" }, "Hoenn", 718, []Pokemon{} },
	Pokedex{"VII", "sun-moon", "Nation Alola dex", []string{ "Sun", "Moon" }, "Alola", 802, []Pokemon{} },
	Pokedex{"VII", "ultra-sun-ultra-moon", "National updated Alola dex", []string{ "Ultra Sun", "Ultra Moon" }, "Alola", 807, []Pokemon{} },
	// generation VIII is not yet available in Veekun
}

func main() {
	fmt.Printf("Living Pokédex Templates generator\n")

	db, err := sql.Open("sqlite3", "./pokedex.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // close connection after `main` returns

	fmt.Printf("Connected to database (%T)\n", db)

	for _, pokedex := range pokedexes {
		rows, err := db.Query(`
			select
				pokemon_species.id as id,
				pokemon_species.identifier as identifier,
				pokemon_species_names.name
			from 
				pokemon_species
				join pokemon_species_names on pokemon_species.id = pokemon_species_names.pokemon_species_id and pokemon_species_names.local_language_id = 9
			where 
				pokemon_species.id <= ?
			order by
				pokemon_species.id asc`, pokedex.MaxNationalDex)
		if err != nil {
			log.Fatal(err)
		}

		pokedex.Pokemon = make([]Pokemon, pokedex.MaxNationalDex)

		defer rows.Close() // close result set after operations
		for rows.Next() {
			var id int
			var identifier, name string

			err = rows.Scan(&id, &identifier, &name)
			if err != nil {
				log.Fatal(err)
			}

			pokedex.Pokemon[id - 1] = Pokemon{id, id, identifier, name}
		}

		json, err := json.MarshalIndent(pokedex, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("data/" + pokedex.Identifier + ".json", json, 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Written \"%s\" to data/%s.json\n", pokedex.Name, pokedex.Identifier)
	}
}