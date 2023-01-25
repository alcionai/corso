package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyPortalBlockedActionable 
type CompanyPortalBlockedActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*CompanyPortalAction)
    GetOdataType()(*string)
    GetOwnerType()(*OwnerType)
    GetPlatform()(*DevicePlatformType)
    SetAction(value *CompanyPortalAction)()
    SetOdataType(value *string)()
    SetOwnerType(value *OwnerType)()
    SetPlatform(value *DevicePlatformType)()
}
