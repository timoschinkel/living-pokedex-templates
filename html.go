package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
)

func (pokedex Pokedex) Boxes () []Box {
	var box_size = 30;
	var num_of_boxes = int(math.Ceil(float64(len(pokedex.Pokemon)) / float64(box_size)))

	boxes := make([]Box, num_of_boxes)
	for index := 0; index < num_of_boxes; index++ {
		var start = int(index * box_size + 1)
		var end = int(math.Min(float64(start + box_size - 1), float64(len(pokedex.Pokemon))))

		boxes[index] = Box{fmt.Sprintf("%03d - %03d", start, end), pokedex.Pokemon[start - 1:end]}
	}

	return boxes
}

func (pokemon Pokemon) Url (pokedex Pokedex) string {
	mapping := map[string]string{
		"red-blue-yellow": "",
		"gold-silver-crystal": "gs",
		"ruby-sapphire-emerald": "rs",
		"firered-leafgreen": "rs",
		"diamond-pearl-platinum": "dp",
		"heartgold-soulsilver": "dp",
		"black-white-black-2-white-2": "bw",
		"x-y": "xy",
		"omega-ruby-alpha-sapphire": "xy",
		"sun-moon": "sm",
		"ultra-sun-ultra-moon": "sm",
		"sword-shield-galar": "swsh",
		"sword-shield-isle-of-armor": "swsh",
	}
	
	separator := "-"
	if mapping[pokedex.Identifier] == "" {
		separator = ""
	}

	if mapping[pokedex.Identifier] == "swsh" {
		// Serebii has changed their url structure starting with generation VIII
		return fmt.Sprintf("https://www.serebii.net/pokedex%s%s/%s/", separator, mapping[pokedex.Identifier], pokemon.Identifier)
	} else {
		return fmt.Sprintf("https://www.serebii.net/pokedex%s%s/%03d.shtml", separator, mapping[pokedex.Identifier], pokemon.NationalDexNumber)		
	}
}

func generate_html() {
	generate_pokedex_html()
	generate_index_html()
}

func generate_pokedex_html() {
	// Read template
	tpl, err := template.ParseFiles("./templates/pokedex.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over json files and parse template with json data
	files, _ := filepath.Glob("data/*.json")
	for _, file := range files {
		reader, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}		

		pokedex := Pokedex{}
		err = json.Unmarshal([]byte(reader), &pokedex)

		filename := fmt.Sprintf("docs/%s.html", pokedex.Identifier)
		writer, err := os.Create(filename)
		err = tpl.Execute(writer, pokedex)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Written \"%s\" to %s\n", pokedex.Name, filename)
	}
}

func generate_index_html() {
	// Read template
	tpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	generations := make(map[string][]Pokedex)
	for _, pokedex := range pokedexes {
		if _, exists := generations[pokedex.Generation]; ! exists {
			generations[pokedex.Generation] = []Pokedex{}
		}
		
		generations[pokedex.Generation] = append(generations[pokedex.Generation], pokedex)
	}

	filename := "docs/index.html"
	writer, err := os.Create(filename)
	err = tpl.Execute(writer, generations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Written \"index\" to %s\n", filename)
}
