package main

import (
    "fmt"
    "os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command specified. Available commands are 'json' and 'html'")
		os.Exit(1)
	}

	var command = os.Args[1];

	if command == "json" {
		generate_json()
	} else if command == "html" {
		generate_html()
	} else {
		fmt.Printf("Received unknown command '%s'. Available commands are 'json' and 'html'\n", command)
		os.Exit(1)
	}
}