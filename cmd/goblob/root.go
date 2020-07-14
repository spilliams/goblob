package goblob

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spilliams/goblob/internal/parse"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "goblob",
		Short: "Convert SQL blob strings into JSON",
		RunE: func(cmd *cobra.Command, args []string) error {
			var input string
			if len(args) == 0 {
				fmt.Println("(reading stdin)")
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					input = scanner.Text()
				}
			} else {
				input = args[0]
			}
			input = strings.TrimSpace(input)

			p := parse.NewBlobParser()
			if err := p.Parse(input); err != nil {
				return err
			}

			// output in json
			out, err := json.Marshal(p.ParsedObject())
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}
}
