package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
	"github.com/buger/jsonparser"
)

func main() {
	defer func() {
		common.PrintMemUsage()

		f, err := os.Create("mem.prof")
		if err != nil {
			fmt.Print("could not create memory profile: ", err)
			return
		}

		defer f.Close() // error handling omitted for example

		runtime.GC() // get up-to-date statistics

		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Print("could not write memory profile: ", err)
			return
		}
	}()

	d, err := readFile()
	if err != nil {
		return
	}

	parseData(d)
}

func readFile() ([]byte, error) {
	common.PrintMemUsage()

	data, err := ioutil.ReadFile(common.FileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	return data, nil
}

func parseData(data []byte) {
	common.PrintMemUsage()

	output := common.FooArray{
		Entries: []*common.Foo{},
	}

	_ = output

	// var handler func([]byte, []byte, jsonparser.ValueType, int) error
	// handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	// 	fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
	// 	return nil
	// }

	//nolint:errcheck
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		e, errInner := getEntry(value)
		if errInner != nil {
			fmt.Printf("Error decoding input2: %v\n", errInner)
			return
		}

		output.Entries = append(output.Entries, e)
	}, "entries")

	common.PrintMemUsage()

	fmt.Printf("Decoded %d entries\n", len(output.Entries))
}

func getEntry(data []byte) (*common.Foo, error) {
	e := &common.Foo{}

	//nolint:errcheck
	jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		switch string(key) {
		case "id":
			e.ID = string(value)
		case "labels":
			err := json.Unmarshal(value, &e.Labels)
			if err != nil {
				return fmt.Errorf("unmarshalling labels: %w", err)
			}
		case "modified":
			err := json.Unmarshal(value, &e.ModTime)
			if err != nil {
				return fmt.Errorf("unmarshalling modtime: %w", err)
			}
		case "deleted":
			err := json.Unmarshal(value, &e.Deleted)
			if err != nil {
				return fmt.Errorf("unmarshalling deleted: %w", err)
			}
		case "data":
			cpBuf := make([]byte, len(value))
			_ = copy(cpBuf, value)
			e.Content = cpBuf
		default:
			fmt.Printf("Unexpected Input: %v\n", key)
			return errors.New("Unexpected Input: " + string(key))
		}

		return nil
	})

	return e, nil
}
