package connector

import (
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
)

// Contains is a helper method for verifying if element
// is contained within the slice
func Contains(elems []string, value string) bool {
	for _, s := range elems {
		if value == s {
			return true
		}
	}
	return false
}

func optionsForMessageSnapshot() *msmessage.MessagesRequestBuilderGetRequestConfiguration {
	selecting := []string{"id", "parentFolderId"}
	options := &msmessage.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &msmessage.MessagesRequestBuilderGetQueryParameters{
			Select: selecting,
		},
	}
	return options
}
