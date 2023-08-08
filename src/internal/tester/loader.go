package tester

import (
	"bufio"
	"os"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
)

func LoadAFile(t *testing.T, fileName string) []byte {
	// Preserves '\n' of original file. Uses incremental version when file too large
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		f, err := os.Open(fileName)
		require.NoError(t, err, "opening file:", fileName, clues.ToCore(err))

		defer f.Close()

		buffer := make([]byte, 0)
		reader := bufio.NewScanner(f)

		for reader.Scan() {
			temp := reader.Bytes()
			buffer = append(buffer, temp...)
		}

		require.NoError(t, reader.Err(), "reading file:", fileName, clues.ToCore(err))

		return buffer
	}

	return bytes
}
