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

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		buf, _, _, errInner := jsonparser.Get(value, "A")
		if errInner != nil {
			fmt.Printf("Error decoding input: %v\n", err)
			return
		}
		_ = value
		cpBuf := make([]byte, len(buf))
		_ = copy(cpBuf, buf)
		e := &common.Foo{
			A: cpBuf,
		}
		output.Entries = append(output.Entries, e)
	}, "entries")

	common.PrintMemUsage()
}
