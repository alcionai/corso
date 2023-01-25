package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusDetailsable 
type StatusDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    StatusBaseable
    GetAdditionalDetails()(*string)
    GetErrorCategory()(*ProvisioningStatusErrorCategory)
    GetErrorCode()(*string)
    GetReason()(*string)
    GetRecommendedAction()(*string)
    SetAdditionalDetails(value *string)()
    SetErrorCategory(value *ProvisioningStatusErrorCategory)()
    SetErrorCode(value *string)()
    SetReason(value *string)()
    SetRecommendedAction(value *string)()
}
