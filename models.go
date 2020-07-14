package main

type Pokedex struct {
	Generation string        `json:"generation"`
	Identifier string        `json:"identifier"`
	Name string              `json:"name"`
	Games []string           `json:"games"`
	Region string            `json:"region"`
	MaxNationalDex int       `json:"-"`
	Pokemon []Pokemon        `json:"pokemon"`
}

type Pokemon struct {	
	DexNumber int            `json:"dex_number"`
	NationalDexNumber int    `json:"national_dex_number"`
	Identifier string        `json:"identifier"`
	Name string              `json:"name"`
}


type Box struct {
	Title string
	Pokemon []Pokemon
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
	Pokedex{"VIII", "sword-shield-galar", "Galar dex", []string{ "Sword", "Shield"}, "Galar", 0, []Pokemon{} },
	Pokedex{"VIII", "sword-shield-isle-of-armor", "Isle of Armor dex", []string{ "Sword", "Shield"}, "Galar", 0, []Pokemon{} },
}
