package decoder

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

var _ common.ManifestDecoder = Stdlib{}

type Stdlib struct{}

func (d Stdlib) Decode(r io.Reader, gcStats bool) error {
	dec := json.NewDecoder(r)
	output := common.Manifest{}

	if err := dec.Decode(&output); err != nil {
		return errors.Wrap(err, "decoding input")
	}

	return nil
}
