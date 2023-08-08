package api

import (
	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
)

// createFromBytes generates an m365 object form bytes.
func createFromBytes(
	bytes []byte,
	createFunc serialization.ParsableFactory,
) (serialization.Parsable, error) {
	parseNode, err := kjson.NewJsonParseNodeFactory().GetRootParseNode("application/json", bytes)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes into base m365 object").With("bytes_len", len(bytes))
	}

	v, err := parseNode.GetObjectValue(createFunc)
	if err != nil {
		return nil, clues.Wrap(err, "parsing m365 object factory")
	}

	return v, nil
}
