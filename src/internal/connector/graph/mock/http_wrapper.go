package mock

import (
	"context"
	"io"
	"net/http"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

type requester struct{}

func (r requester) Request(
	_ context.Context,
	_, _ string,
	_ io.Reader,
	_ map[string]string,
) (*http.Response, error) {
	return nil, clues.New("not implemented")
}

func NewRequester() graph.Requester {
	return requester{}
}
