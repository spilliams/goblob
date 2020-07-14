package main

import (
	"fmt"
	"os"

	"github.com/spilliams/goblob/cmd/goblob"
)

func main() {
	cmd := goblob.NewCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
