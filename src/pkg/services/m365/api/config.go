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
	bccRecipients     = "bccRecipients"
	ccRecipients      = "ccRecipients"
	createdDateTime   = "createdDateTime"
	displayName       = "displayName"
	emailAddresses    = "emailAddresses"
	givenName         = "givenName"
	isCancelled       = "isCancelled"
	isDraft           = "isDraft"
	mobilePhone       = "mobilePhone"
	parentFolderID    = "parentFolderId"
	receivedDateTime  = "receivedDateTime"
	recurrence        = "recurrence"
	sentDateTime      = "sentDateTime"
	surname           = "surname"
	toRecipients      = "toRecipients"
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
	id := []string{"id"}

	if len(ss) == 0 {
		return id
	}

	return append(id, ss...)
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

// URL cache only needs a subset of item properties
func DriveItemSelectURLCache() []string {
	return idAnd(
		"content.downloadUrl",
		"deleted",
		"file",
		"folder")
}
