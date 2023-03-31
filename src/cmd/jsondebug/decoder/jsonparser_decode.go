package decoder

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/buger/jsonparser"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

var _ common.ManifestDecoder = JsonParser{}

//revive:disable-next-line:var-naming
type JsonParser struct{}

func (d JsonParser) Decode(r io.Reader, gcStats bool) error {
	if gcStats {
		common.PrintMemUsage()
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "reading data")
	}

	return parseManifestData(data, gcStats)
}

func (d JsonParser) DecodeBytes(data []byte, gcStats bool) error {
	if gcStats {
		common.PrintMemUsage()
	}

	return parseManifestData(data, gcStats)
}

func parseManifestData(data []byte, gcStats bool) error {
	if gcStats {
		common.PrintMemUsage()
	}

	var (
		errs   *multierror.Error
		output = common.Manifest{
			Entries: []*common.ManifestEntry{},
		}
	)

	_ = output

	// var handler func([]byte, []byte, jsonparser.ValueType, int) error
	// handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	// 	fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
	// 	return nil
	// }

	//nolint:errcheck
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		e, errInner := getManifestEntry(value)
		if errInner != nil {
			errs = multierror.Append(errs, err)
		}

		output.Entries = append(output.Entries, e)
	}, "entries")

	if gcStats {
		common.PrintMemUsage()

		fmt.Printf("Decoded %d entries\n", len(output.Entries))
	}

	return errs.ErrorOrNil()
}

func getManifestEntry(data []byte) (*common.ManifestEntry, error) {
	var (
		errs  *multierror.Error
		err   error
		e     = &common.ManifestEntry{}
		paths = [][]string{
			{"id"},
			{"labels"},
			{"modified"},
			{"deleted"},
			{"data"},
		}
	)

	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, iterErr error) {
		switch idx {
		case 0:
			e.ID = string(value)

		case 1:
			err = json.Unmarshal(value, &e.Labels)
			if err != nil {
				err = errors.Wrap(err, "unmarshalling labels")
			}

		case 2:
			e.ModTime, err = time.Parse(time.RFC3339, string(value))
			if err != nil {
				err = errors.Wrap(err, "unmarshalling modtime")
			}

		case 3:
			err = json.Unmarshal(value, &e.Deleted)
			if err != nil {
				err = errors.Wrap(err, "unmarshalling deleted")
			}

		case 4:
			e.Content = make([]byte, len(value))
			n := copy(e.Content, value)
			if n != len(value) {
				err = errors.Errorf("failed to copy content; got %d bytes", n)
			}

		default:
			err = errors.Errorf("unexpected input %v", idx)
		}

		errs = multierror.Append(errs, err)
	}, paths...)

	return e, errs.ErrorOrNil()
}
