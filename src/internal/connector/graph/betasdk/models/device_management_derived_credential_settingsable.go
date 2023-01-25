package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDerivedCredentialSettingsable 
type DeviceManagementDerivedCredentialSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetHelpUrl()(*string)
    GetIssuer()(*DeviceManagementDerivedCredentialIssuer)
    GetNotificationType()(*DeviceManagementDerivedCredentialNotificationType)
    GetRenewalThresholdPercentage()(*int32)
    SetDisplayName(value *string)()
    SetHelpUrl(value *string)()
    SetIssuer(value *DeviceManagementDerivedCredentialIssuer)()
    SetNotificationType(value *DeviceManagementDerivedCredentialNotificationType)()
    SetRenewalThresholdPercentage(value *int32)()
}
