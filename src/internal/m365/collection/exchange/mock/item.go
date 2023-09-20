package mock

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

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
	bool,
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
