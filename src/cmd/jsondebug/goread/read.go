package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

func main() {
	f, err := os.Open(common.FileName)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}

	defer f.Close()

	dec := json.NewDecoder(f)

	output := common.FooArray{}

	if err := dec.Decode(&output); err != nil {
		fmt.Printf("Error decoding input: %v\n", err)
		return
	}

	common.PrintMemUsage()
}
