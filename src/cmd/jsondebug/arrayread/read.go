package main

import (
	"fmt"
	"os"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
	"github.com/alcionai/corso/src/cmd/jsondebug/decoder"
)

func main() {
	readData()
}

func readData() {
	f, err := os.Open(common.FileName)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}

	defer f.Close()

	output, err := decoder.DecodeFooArray(f)
	if err != nil {
		fmt.Printf("Error decoding input: %v\n", err)
		return
	}

	common.PrintMemUsage()

	fmt.Printf("got array with %d items\n", len(output.Entries))
}
