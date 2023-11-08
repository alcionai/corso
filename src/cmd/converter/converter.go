package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alcionai/corso/src/internal/converters/eml"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: converter <source-format> <target-format> <filename>")
		os.Exit(1)
	}

	from := os.Args[1]
	to := os.Args[2]
	filename := os.Args[3]

	body, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var out string

	switch from {
	case "msg":
		switch to {
		case "eml":
			msg, err := api.BytesToMessageable(body)
			if err != nil {
				log.Fatal(err)
			}

			out, err = eml.ToEml(msg)
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatal("Unknown target format", to)
		}
	default:
		log.Fatal("Unknown source format", from)
	}

	fmt.Print(out)
}
