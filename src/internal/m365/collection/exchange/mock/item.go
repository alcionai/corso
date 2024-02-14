package mock

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// get and serialize item mock
// ---------------------------------------------------------------------------

type ItemGetSerialize struct {
	GetData        serialization.Parsable
	GetCount       int
	GetErr         error
	SerializeCount int
	SerializeErr   error
}

func (m *ItemGetSerialize) GetItem(
	context.Context,
	string, string,
	*fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	m.GetCount++
	return m.GetData, &details.ExchangeInfo{}, m.GetErr
}

func (m *ItemGetSerialize) Serialize(
	ctx context.Context,
	p serialization.Parsable,
	_ string, _ string,
) ([]byte, error) {
	m.SerializeCount++

	if p == nil || m.SerializeErr != nil {
		return nil, m.SerializeErr
	}

	return api.Mail{}.Serialize(ctx, p, "", "")
}

func DefaultItemGetSerialize() *ItemGetSerialize {
	return &ItemGetSerialize{}
}

// ---------------------------------------------------------------------------
// can skip item failure mock
// ---------------------------------------------------------------------------

type canSkipFailChecker struct {
	canSkip bool
}

func (m canSkipFailChecker) CanSkipItemFailure(
	error,
	string,
	string,
	control.Options,
) (fault.SkipCause, bool) {
	return fault.SkipCause("testing"), m.canSkip
}

func NeverCanSkipFailChecker() *canSkipFailChecker {
	return &canSkipFailChecker{}
}
