package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// VerticalSectionable
type VerticalSectionable interface {
	msmodel.Entityable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmphasis() *SectionEmphasisType
	GetWebparts() []WebPartable
	SetEmphasis(value *SectionEmphasisType)
	SetWebparts(value []WebPartable)
}
