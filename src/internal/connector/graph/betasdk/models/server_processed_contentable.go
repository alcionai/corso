package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServerProcessedContentable
type ServerProcessedContentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetComponentDependencies() []MetaDataKeyStringPairable
	GetCustomMetadata() []MetaDataKeyValuePairable
	GetHtmlStrings() []MetaDataKeyStringPairable
	GetImageSources() []MetaDataKeyStringPairable
	GetLinks() []MetaDataKeyStringPairable
	GetOdataType() *string
	GetSearchablePlainTexts() []MetaDataKeyStringPairable
	SetComponentDependencies(value []MetaDataKeyStringPairable)
	SetCustomMetadata(value []MetaDataKeyValuePairable)
	SetHtmlStrings(value []MetaDataKeyStringPairable)
	SetImageSources(value []MetaDataKeyStringPairable)
	SetLinks(value []MetaDataKeyStringPairable)
	SetOdataType(value *string)
	SetSearchablePlainTexts(value []MetaDataKeyStringPairable)
}
