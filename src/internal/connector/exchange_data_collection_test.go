package connector

import (
	"fmt"
	"io"
	"testing"
)

func TestConnector1(t *testing.T) {

	dc := new(ExchangeDataCollection)

	err := dc.InitDataCollection(CORSO_USER)
	if err != nil {
		t.Fatal(err)
	}

	for !dc.isEmpty() {
		ds, err := dc.NextItem()
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				fmt.Println("end of file")
				break
			}
			dc.CloseDataCollection()
			t.Fatal(err)
		}

		fmt.Printf("ds UUID: %v\n", ds.UUID())

		//TODO Data read and compare

	}

	err = dc.CloseDataCollection()
	if err != nil {
		t.Fatal(err)
	}
}
