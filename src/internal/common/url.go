package common

import (
	"net/url"

	"github.com/alcionai/clues"
)

// GetQueryParamFromURL parses an URL and returns value of the specified
// query parameter. In case of multiple occurrences, first one is returned.
func GetQueryParamFromURL(
	rawURL, queryParam string,
) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", clues.Wrap(err, "parsing url")
	}

	qp := u.Query()

	val := qp.Get(queryParam)
	if len(val) == 0 {
		return "", clues.New("query param not found").With("query_param", queryParam)
	}

	return val, nil
}
