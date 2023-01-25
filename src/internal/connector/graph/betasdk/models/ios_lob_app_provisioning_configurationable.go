package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosLobAppProvisioningConfigurationable 
type IosLobAppProvisioningConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]IosLobAppProvisioningConfigurationAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatusable)
    GetDisplayName()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupAssignments()([]MobileAppProvisioningConfigGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayload()([]byte)
    GetPayloadFileName()(*string)
    GetRoleScopeTagIds()([]string)
    GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatusable)
    GetVersion()(*int32)
    SetAssignments(value []IosLobAppProvisioningConfigurationAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatusable)()
    SetDisplayName(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupAssignments(value []MobileAppProvisioningConfigGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayload(value []byte)()
    SetPayloadFileName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatusable)()
    SetVersion(value *int32)()
}
