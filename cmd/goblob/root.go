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
	var flags struct {
		concise bool
	}

	cmd := &cobra.Command{
		Use:   "goblob",
		Short: "Convert SQL blob strings into JSON",
		RunE: func(cmd *cobra.Command, args []string) error {
			var input string
			if len(args) == 0 {
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
			var out []byte
			var err error
			if flags.concise {
				out, err = json.Marshal(p.ParsedObject())
			} else {
				out, err = json.MarshalIndent(p.ParsedObject(), "", "  ")
			}
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}

	cmd.Flags().BoolVarP(&flags.concise, "concise", "c", false, "Make output more concise")

	return cmd
}
