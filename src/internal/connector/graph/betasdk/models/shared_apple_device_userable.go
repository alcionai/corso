package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedAppleDeviceUserable 
type SharedAppleDeviceUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataQuota()(*int64)
    GetDataToSync()(*bool)
    GetDataUsed()(*int64)
    GetOdataType()(*string)
    GetUserPrincipalName()(*string)
    SetDataQuota(value *int64)()
    SetDataToSync(value *bool)()
    SetDataUsed(value *int64)()
    SetOdataType(value *string)()
    SetUserPrincipalName(value *string)()
}
