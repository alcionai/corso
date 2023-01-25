package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppSupersedenceable 
type MobileAppSupersedenceable interface {
    MobileAppRelationshipable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSupersededAppCount()(*int32)
    GetSupersedenceType()(*MobileAppSupersedenceType)
    GetSupersedingAppCount()(*int32)
    SetSupersededAppCount(value *int32)()
    SetSupersedenceType(value *MobileAppSupersedenceType)()
    SetSupersedingAppCount(value *int32)()
}
