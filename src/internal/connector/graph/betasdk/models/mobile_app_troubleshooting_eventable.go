package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingEventable 
type MobileAppTroubleshootingEventable interface {
    DeviceManagementTroubleshootingEventable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationId()(*string)
    GetAppLogCollectionRequests()([]AppLogCollectionRequestable)
    GetHistory()([]MobileAppTroubleshootingHistoryItemable)
    GetManagedDeviceIdentifier()(*string)
    GetUserId()(*string)
    SetApplicationId(value *string)()
    SetAppLogCollectionRequests(value []AppLogCollectionRequestable)()
    SetHistory(value []MobileAppTroubleshootingHistoryItemable)()
    SetManagedDeviceIdentifier(value *string)()
    SetUserId(value *string)()
}
