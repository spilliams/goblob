package main

import "github.com/spilliams/goblob/internal/parse"

func main() {
	// look for string arg
	// or pipe in?

	// parse string
	p := parse.NewBlobParser()
	p.parse()

	// output in json
}
