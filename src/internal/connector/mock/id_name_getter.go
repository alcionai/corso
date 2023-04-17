package mock

import "context"

type IDNameGetter struct {
	ID, Name string
	Err      error
}

func (ing IDNameGetter) GetIDAndName(
	_ context.Context,
	_ string,
) (string, string, error) {
	return ing.ID, ing.Name, ing.Err
}
