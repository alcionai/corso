package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeOnPremisesPolicyable 
type DeviceManagementExchangeOnPremisesPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessRules()([]DeviceManagementExchangeAccessRuleable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetDefaultAccessLevel()(*DeviceManagementExchangeAccessLevel)
    GetKnownDeviceClasses()([]DeviceManagementExchangeDeviceClassable)
    GetNotificationContent()([]byte)
    SetAccessRules(value []DeviceManagementExchangeAccessRuleable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetDefaultAccessLevel(value *DeviceManagementExchangeAccessLevel)()
    SetKnownDeviceClasses(value []DeviceManagementExchangeDeviceClassable)()
    SetNotificationContent(value []byte)()
}
