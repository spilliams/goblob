package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spilliams/goblob/internal/parse"
)

func main() {
	// look for string arg
	// or pipe in?

	// parse string
	p := parse.NewBlobParser()
	err := p.Parse("s:4:\"1234\";")
	if err != nil {
		log.Fatal(err)
	}

	// output in json
	out, err := json.Marshal(p.ParsedObject())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
