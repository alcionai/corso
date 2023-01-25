package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcAuditResourceable 
type CloudPcAuditResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetModifiedProperties()([]CloudPcAuditPropertyable)
    GetOdataType()(*string)
    GetResourceId()(*string)
    GetType()(*string)
    SetDisplayName(value *string)()
    SetModifiedProperties(value []CloudPcAuditPropertyable)()
    SetOdataType(value *string)()
    SetResourceId(value *string)()
    SetType(value *string)()
}
