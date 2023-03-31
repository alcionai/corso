package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
	"github.com/buger/jsonparser"
)

func main() {
	f, err := os.Open(common.FileName)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	output := common.FooArray{
		Entries: []*common.Foo{},
	}

	_ = output

	common.PrintMemUsage()

	// var handler func([]byte, []byte, jsonparser.ValueType, int) error
	// handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	// 	fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
	// 	return nil
	// }

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		id, _ := jsonparser.GetString(value, "id")
		content, _, _, errInner := jsonparser.Get(value, "data")
		if errInner != nil {
			fmt.Printf("Error decoding input: %v\n", errInner)
			return
		}

		cpBuf := make([]byte, len(content))
		_ = copy(cpBuf, content)
		e := &common.Foo{
			ID:      id,
			Content: cpBuf,
		}
		output.Entries = append(output.Entries, e)

	}, "entries")

	common.PrintMemUsage()

	// for _, e := range output.Entries {
	// 	fmt.Printf("ID: '%s'\n Content: %s \n", e.ID, e.Content)
	// }
}
