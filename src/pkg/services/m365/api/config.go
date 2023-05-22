package api

import (
	"fmt"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	maxNonDeltaPageSize = int32(999)
	maxDeltaPageSize    = int32(500)
)

// selectable values, case insensitive
// not comprehensive - just adding ones that might
// get easily misspelled.
// eg: we don't need a const for "id"
const (
	parentFolderID    = "parentFolderId"
	displayName       = "displayName"
	userPrincipalName = "userPrincipalName"
)

// header keys
const (
	headerKeyConsistencyLevel = "ConsistencyLevel"
	headerKeyPrefer           = "Prefer"
)

// header values
const (
	idTypeImmutable = `IdType="ImmutableId"`
	eventual        = "eventual"
)

// ---------------------------------------------------------------------------
// not exported
// ---------------------------------------------------------------------------

func preferPageSize(size int32) string {
	return fmt.Sprintf("odata.maxpagesize=%d", size)
}

func preferImmutableIDs(t bool) string {
	if !t {
		return ""
	}

	return idTypeImmutable
}

func newPreferHeaders(values ...string) *abstractions.RequestHeaders {
	vs := []string{}

	for _, v := range values {
		if len(v) > 0 {
			vs = append(vs, v)
		}
	}

	headers := abstractions.NewRequestHeaders()
	headers.Add(headerKeyPrefer, strings.Join(vs, ","))

	return headers
}

func newEventualConsistencyHeaders() *abstractions.RequestHeaders {
	headers := abstractions.NewRequestHeaders()
	headers.Add(headerKeyConsistencyLevel, eventual)

	return headers
}

// makes a slice with []string{"id", s...}
func idAnd(ss ...string) []string {
	return append([]string{"id"}, ss...)
}

// ---------------------------------------------------------------------------
// exported
// ---------------------------------------------------------------------------

func DriveItemSelectDefault() []string {
	return idAnd(
		"content.downloadUrl",
		"createdBy",
		"createdDateTime",
		"file",
		"folder",
		"lastModifiedDateTime",
		"name",
		"package",
		"parentReference",
		"root",
		"sharepointIds",
		"size",
		"deleted",
		"malware",
		"shared")
}
