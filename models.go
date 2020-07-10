package main

type Pokedex struct {
	Generation string     `json:"generation"`
	Identifier string     `json:"identifier"`
	Name string           `json:"name"`
	Games []string        `json:"games"`
	Region string         `json:"region"`
	MaxNationalDex int    `json:"-"`
	Pokemon []Pokemon     `json:"pokemon"`
}

type Pokemon struct {	
	DexNumber int         `json:"dex_number"`
	NationalDexNumber int `json:"national_dex_number"`
	Identifier string     `json:"identifier"`
	Name string           `json:"name"`
}