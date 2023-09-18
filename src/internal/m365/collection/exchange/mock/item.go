package mock

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

type ItemGetSerialize struct {
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
	return nil, &details.ExchangeInfo{}, m.GetErr
}

func (m *ItemGetSerialize) Serialize(
	context.Context,
	serialization.Parsable,
	string, string,
) ([]byte, error) {
	m.SerializeCount++
	return nil, m.SerializeErr
}

func DefaultItemGetSerialize() *ItemGetSerialize {
	return &ItemGetSerialize{}
}
