package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// HorizontalSectionColumnable
type HorizontalSectionColumnable interface {
	msmodel.Entityable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetWebparts() []WebPartable
	GetWidth() *int32
	SetWebparts(value []WebPartable)
	SetWidth(value *int32)
}
