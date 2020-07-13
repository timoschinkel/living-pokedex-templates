package main

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func generate_json() {
	db, err := sql.Open("sqlite3", "./pokedex.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // close connection after method returns

	fmt.Printf("Connected to database (%T)\n", db)

	for _, pokedex := range pokedexes {
		if pokedex.MaxNationalDex == 0 {
			fmt.Printf("No maximal national dex number defined for '%s', skipping...\n", pokedex.Name)
			continue
		}

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
