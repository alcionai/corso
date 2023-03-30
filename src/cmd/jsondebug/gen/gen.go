package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

func main() {
	buf := make([]byte, common.ItemSize)
	data := make([]common.Foo, 0, common.NumItems)

	for i := 0; i < common.NumItems; i++ {
		n, err := rand.Read(buf)
		if err != nil {
			fmt.Printf("Error reading random data: %v\n", err)
			return
		} else if n != common.ItemSize {
			fmt.Printf(
				"Short read for item data: wanted %d, got %d\n",
				common.ItemSize,
				n,
			)
			return
		}

		item := common.Foo{A: buf}
		data = append(data, item)
	}

	f, err := os.Create(common.FileName)
	if err != nil {
		fmt.Printf("Error making output file: %v\n", err)
		return
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	if err := enc.Encode(data); err != nil {
		fmt.Printf("Error writing json to file: %v\n", err)
		return
	}
}
