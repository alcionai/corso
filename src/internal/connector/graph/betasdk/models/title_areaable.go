package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TitleAreaable
type TitleAreaable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAlternativeText() *string
	GetEnableGradientEffect() *bool
	GetImageWebUrl() *string
	GetLayout() *TitleAreaLayoutType
	GetOdataType() *string
	GetServerProcessedContent() ServerProcessedContentable
	GetShowAuthor() *bool
	GetShowPublishedDate() *bool
	GetShowTextBlockAboveTitle() *bool
	GetTextAboveTitle() *string
	GetTextAlignment() *TitleAreaTextAlignmentType
	SetAlternativeText(value *string)
	SetEnableGradientEffect(value *bool)
	SetImageWebUrl(value *string)
	SetLayout(value *TitleAreaLayoutType)
	SetOdataType(value *string)
	SetServerProcessedContent(value ServerProcessedContentable)
	SetShowAuthor(value *bool)
	SetShowPublishedDate(value *bool)
	SetShowTextBlockAboveTitle(value *bool)
	SetTextAboveTitle(value *string)
	SetTextAlignment(value *TitleAreaTextAlignmentType)
}
