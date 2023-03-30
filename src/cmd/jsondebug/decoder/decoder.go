package decoder

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/alcionai/clues"
)

const (
	arrayOpen  = "["
	arrayClose = "]"
)

func DecodeArray[T any](r io.Reader, output *[]T) error {
	dec := json.NewDecoder(r)
	arrayStarted := false

	for {
		t, err := dec.Token()
		if err == io.EOF {
			// Done processing input.
			break
		} else if err != nil {
			return clues.Wrap(err, "reading JSON token")
		}

		d, ok := t.(json.Delim)
		if !ok {
			return clues.New(fmt.Sprintf("unexpected token: (%T) %v", t, t))
		} else if arrayStarted && d.String() == arrayClose {
			break
		} else if d.String() != arrayOpen {
			return clues.New(fmt.Sprintf("expected array start but found %s", d))
		}

		arrayStarted = true

		for dec.More() {
			tmp := *new(T)
			if err := dec.Decode(&tmp); err != nil {
				return clues.Wrap(err, "decoding array element")
			}

			*output = append(*output, tmp)
		}
	}

	return nil
}
