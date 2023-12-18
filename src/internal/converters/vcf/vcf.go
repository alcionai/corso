package vcf

import (
	"bytes"
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/emersion/go-vcard"
)

// This package helps convert the json response backed up from graph
// API to a vCard file.
// Ref: https://datatracker.ietf.org/doc/html/rfc6350

func FromJSON(ctx context.Context, body []byte) (string, error) {
	vc := vcard.Card{}
	vcard.ToV4(vc)

	data, err := api.BytesToContactable(body)
	if err != nil {
		return "", clues.Wrap(err, "converting to contactable")
	}

	name := vcard.Name{
		GivenName:  ptr.Val(data.GetGivenName()),
		FamilyName: ptr.Val(data.GetSurname()),
	}
	vc.SetName(&name)

	out := bytes.NewBuffer(nil)
	enc := vcard.NewEncoder(out)

	err = enc.Encode(vc)
	if err != nil {
		return "", clues.Wrap(err, "encoding vcard")
	}

	outc := out.String()
	return outc, nil
}
