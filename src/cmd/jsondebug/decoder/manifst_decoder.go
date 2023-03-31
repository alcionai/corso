package decoder

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

var (
	_ common.ManifestDecoder     = Array{}
	_ common.ByteManifestDecoder = Array{}
	_ common.ManifestDecoder     = ArrayFull{}
	_ common.ByteManifestDecoder = ArrayFull{}
	_ common.ManifestDecoder     = Map{}
	_ common.ByteManifestDecoder = Map{}
)

type Array struct{}

func (d Array) Decode(r io.Reader, gcStats bool) error {
	_, err := DecodeManifestArray(r)
	return err
}

func (d Array) DecodeBytes(data []byte, gcStats bool) error {
	r := bytes.NewReader(data)
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

type ArrayFull struct{}

func (d ArrayFull) Decode(r io.Reader, gcStats bool) error {
	_, err := d.decodeManifestArray(r)
	return err
}

func (d ArrayFull) DecodeBytes(data []byte, gcStats bool) error {
	r := bytes.NewReader(data)
	_, err := d.decodeManifestArray(r)

	return err
}

func (d ArrayFull) decodeManifestArray(r io.Reader) (common.Manifest, error) {
	var (
		dec = json.NewDecoder(r)
		res = common.Manifest{}
	)

	if err := expectDelimToken(dec, objectOpen); err != nil {
		return res, err
	}

	// Need to manually decode fields here since we can't reuse the stdlib
	// decoder due to memory issues.
	if err := d.parseManifestEntries(dec, &res); err != nil {
		return res, err
	}

	// Consumes closing object curly brace after we're done. Don't need to check
	// for EOF because json.Decode only guarantees decoding the next JSON item in
	// the stream so this follows that.
	return res, expectDelimToken(dec, objectClose)
}

func (d ArrayFull) parseManifestEntries(dec *json.Decoder, res *common.Manifest) error {
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

		if err := expectDelimToken(dec, arrayOpen); err != nil {
			return err
		}

		for dec.More() {
			ent, err := d.parseManifestFields(dec)
			if err != nil {
				return err
			}

			res.Entries = append(res.Entries, ent)
		}

		if err := expectDelimToken(dec, arrayClose); err != nil {
			return err
		}
	}

	return nil
}

func (d ArrayFull) parseManifestFields(dec *json.Decoder) (*common.ManifestEntry, error) {
	if err := expectDelimToken(dec, objectOpen); err != nil {
		return nil, err
	}

	var (
		seen = map[string]struct{}{}
		res  = &common.ManifestEntry{}
	)

	for dec.More() {
		l, err := stringToken(dec)
		if err != nil {
			return nil, err
		}

		if _, ok := seen[l]; ok {
			return nil, errors.Errorf("repeated field %s", l)
		}

		switch l {
		case "id":
			err = dec.Decode(&res.ID)

		case "labels":
			err = dec.Decode(&res.Labels)

		case "modified":
			err = dec.Decode(&res.ModTime)

		case "deleted":
			err = dec.Decode(&res.Deleted)

		case "data":
			err = dec.Decode(&res.Content)

		default:
			if _, err := dec.Token(); err != nil {
				return nil, errors.Wrapf(err, "consuming value for unexpected field %s", l)
			}

			continue
		}

		seen[l] = struct{}{}

		if err != nil {
			return nil, errors.Wrapf(err, "decoding value for field %s", l)
		}
	}

	if err := expectDelimToken(dec, objectClose); err != nil {
		return nil, err
	}

	return res, nil
}

type Map struct{}

func (d Map) Decode(r io.Reader, gcStats bool) error {
	_, err := d.decodeManifestArray(r)
	return err
}

func (d Map) DecodeBytes(data []byte, gcStats bool) error {
	r := bytes.NewReader(data)
	_, err := d.decodeManifestArray(r)

	return err
}

func (d Map) decodeManifestArray(r io.Reader) (common.Manifest, error) {
	var (
		dec = json.NewDecoder(r)
		res = common.Manifest{}
	)

	if err := expectDelimToken(dec, objectOpen); err != nil {
		return res, err
	}

	// Need to manually decode fields here since we can't reuse the stdlib
	// decoder due to memory issues.
	if err := d.parseManifestEntries(dec, &res); err != nil {
		return res, err
	}

	// Consumes closing object curly brace after we're done. Don't need to check
	// for EOF because json.Decode only guarantees decoding the next JSON item in
	// the stream so this follows that.
	return res, expectDelimToken(dec, objectClose)
}

func (d Map) parseManifestEntries(dec *json.Decoder, res *common.Manifest) error {
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

		if err := expectDelimToken(dec, arrayOpen); err != nil {
			return err
		}

		for dec.More() {
			ent := map[string]any{}

			if err := dec.Decode(&ent); err != nil {
				return err
			}

			// Give up here, just check how many bytes it needs during benchmarking.
			// fmt.Printf("%+v\n", ent)
			// return errors.New("exit early")

			// me := &common.ManifestEntry{
			//   ModTime:
			// }
		} //nolint: wsl

		if err := expectDelimToken(dec, arrayClose); err != nil {
			return err
		}
	}

	return nil
}
