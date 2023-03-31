package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
	"github.com/google/uuid"
)

func main() {
	buf := make([]byte, common.ItemSize)
	data := &common.FooArray{
		Entries: make([]*common.Foo, 0, common.NumItems),
	}

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

		content := common.Content{
			ID:   uuid.NewString(),
			Data: buf,
		}
		payload, _ := json.Marshal(content)

		item := common.Foo{
			ID:      uuid.NewString(),
			Labels:  map[string]string{"foo": "bar"},
			ModTime: time.Now(),
			Content: payload,
		}
		data.Entries = append(data.Entries, &item)
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
