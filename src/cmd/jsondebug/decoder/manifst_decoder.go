package decoder

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

var (
	_ common.ManifestDecoder = Array{}
)

type Array struct{}

func (d Array) Decode(r io.Reader, gcStats bool) error {
	_, err := DecodeManifestArray(r)
	return err
}

func DecodeManifestArray(r io.Reader) (common.Manifest, error) {
	var (
		dec = json.NewDecoder(r)
		res = common.Manifest{}
	)

	if err := expectDelimToken(dec, objectOpen); err != nil {
		return res, err
	}

	// Need to manually decode fields here since we can't reuse the stdlib
	// decoder due to memory issues.
	if err := parseManifestFields(dec, &res); err != nil {
		return res, err
	}

	// Consumes closing object curly brace after we're done. Don't need to check
	// for EOF because json.Decode only guarantees decoding the next JSON item in
	// the stream so this follows that.
	return res, expectDelimToken(dec, objectClose)
}

func parseManifestFields(dec *json.Decoder, res *common.Manifest) error {
	var seen bool

	for dec.More() {
		l, err := stringToken(dec)
		if err != nil {
			return err
		}

		// Only have `entries` field right now. This is stricter than the current
		// JSON decoder in the stdlib.
		if l != "entries" {
			return errors.Errorf("unexpected field name %s", l)
		} else if seen {
			return errors.New("repeated Entries field")
		}

		seen = true

		if err := decodeArray(dec, &res.Entries); err != nil {
			return err
		}
	}

	return nil
}
