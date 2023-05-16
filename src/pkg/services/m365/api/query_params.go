package api

import (
	"fmt"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const deltaMaxPageSize = 200

// buildPreferHeaders returns the headers we add to item delta page requests.
func buildPreferHeaders(pageSize, immutableID bool) *abstractions.RequestHeaders {
	var allHeaders []string

	if pageSize {
		allHeaders = append(allHeaders, fmt.Sprintf("odata.maxpagesize=%d", deltaMaxPageSize))
	}

	if immutableID {
		allHeaders = append(allHeaders, `IdType="ImmutableId"`)
	}

	headers := abstractions.NewRequestHeaders()
	headers.Add("Prefer", strings.Join(allHeaders, ","))

	return headers
}
