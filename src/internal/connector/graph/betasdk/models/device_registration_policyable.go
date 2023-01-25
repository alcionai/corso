package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceRegistrationPolicyable 
type DeviceRegistrationPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureADJoin()(AzureAdJoinPolicyable)
    GetAzureADRegistration()(AzureADRegistrationPolicyable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMultiFactorAuthConfiguration()(*MultiFactorAuthConfiguration)
    GetUserDeviceQuota()(*int32)
    SetAzureADJoin(value AzureAdJoinPolicyable)()
    SetAzureADRegistration(value AzureADRegistrationPolicyable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMultiFactorAuthConfiguration(value *MultiFactorAuthConfiguration)()
    SetUserDeviceQuota(value *int32)()
}
