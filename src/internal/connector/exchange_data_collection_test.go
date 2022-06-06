package connector

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

const chunkSize = 1024

func deepCompare(t *testing.T, f1, f2 os.File) bool {
	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				t.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func TestConnector1(t *testing.T) {

	dc := new(ExchangeDataCollection)

	// read 1024 bytes at a time
	buf := make([]byte, 1024)

	err := dc.InitDataCollection(CORSO_USER)
	if err != nil {
		t.Fatal(err)
	}

	for !dc.isEmpty() {

		// Get the next time in Data Collection
		ds, err := dc.NextItem()
		if err != nil {
			fmt.Println(err)

			// there are no next items
			if err == io.EOF {
				t.Logf("End of Data Collection")
				break
			}

			//Error!
			dc.CloseDataCollection()
			t.Fatal(err)
		}

		fmt.Printf("ds UUID: %v\n", ds.UUID())

		//function to read from DataStream and compare with Sample Data
		func() {
			//Create temp file
			tempFile, err := os.CreateTemp("./", "*")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tempFile.Name())

			// Copy content to temp file
			for {
				readNBytes, err := ds.Read(buf)
				if err == io.EOF {
					// there is no more data to read
					break
				}

				writeNBytes, err := tempFile.Write(buf[:readNBytes])
				if err != nil {
					t.Fatal(err)
				}
				_ = writeNBytes
			}

			tempFile.Seek(0, io.SeekStart)

			sampleFile, err := os.Open(CORSO_SAMPLE_DATA_FILE)
			if err != nil {
				t.Fatal(err)
			}

			defer sampleFile.Close()

			if !deepCompare(t, *tempFile, *sampleFile) {
				t.Fatalf("%v and sample.data are not same", ds.UUID())
			} else {

				fmt.Printf("%v and sample.data are same\n", ds.UUID())
			}

		}()

	}

	err = dc.CloseDataCollection()
	if err != nil {
		t.Fatal(err)
	}
}
