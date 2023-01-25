package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcServicePlanable 
type CloudPcServicePlanable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetRamInGB()(*int32)
    GetStorageInGB()(*int32)
    GetType()(*CloudPcServicePlanType)
    GetUserProfileInGB()(*int32)
    GetVCpuCount()(*int32)
    SetDisplayName(value *string)()
    SetRamInGB(value *int32)()
    SetStorageInGB(value *int32)()
    SetType(value *CloudPcServicePlanType)()
    SetUserProfileInGB(value *int32)()
    SetVCpuCount(value *int32)()
}
