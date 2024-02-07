package mock

import (
	"context"

	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

type IDNameGetter struct {
	ID, Name string
	Err      error
}

func (ing IDNameGetter) GetIDAndName(
	_ context.Context,
	_ string,
	_ api.CallConfig,
) (string, string, error) {
	return ing.ID, ing.Name, ing.Err
}
