package testing

import (
	"bufio"
	"os"
)

func LoadAFile(aFile string) ([]byte, error) {
	// Preserves '\n' of original file. Uses incremental version when file too large
	bytes, err := os.ReadFile(aFile)
	if err != nil {
		f, err := os.Open(aFile)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		buffer := make([]byte, 0)
		reader := bufio.NewScanner(f)
		for reader.Scan() {
			temp := reader.Bytes()
			buffer = append(buffer, temp...)
		}
		aErr := reader.Err()
		if aErr != nil {
			return nil, aErr
		}
		return buffer, nil
	}
	return bytes, nil
}
