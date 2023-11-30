package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alcionai/corso/src/internal/converters/eml"
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
			out, err = eml.FromJSON(context.Background(), body)
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
