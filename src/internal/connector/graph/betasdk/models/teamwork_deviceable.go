package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceable 
type TeamworkDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(TeamworkDeviceActivityable)
    GetActivityState()(*TeamworkDeviceActivityState)
    GetCompanyAssetTag()(*string)
    GetConfiguration()(TeamworkDeviceConfigurationable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentUser()(TeamworkUserIdentityable)
    GetDeviceType()(*TeamworkDeviceType)
    GetHardwareDetail()(TeamworkHardwareDetailable)
    GetHealth()(TeamworkDeviceHealthable)
    GetHealthStatus()(*TeamworkDeviceHealthStatus)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()(*string)
    GetOperations()([]TeamworkDeviceOperationable)
    SetActivity(value TeamworkDeviceActivityable)()
    SetActivityState(value *TeamworkDeviceActivityState)()
    SetCompanyAssetTag(value *string)()
    SetConfiguration(value TeamworkDeviceConfigurationable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentUser(value TeamworkUserIdentityable)()
    SetDeviceType(value *TeamworkDeviceType)()
    SetHardwareDetail(value TeamworkHardwareDetailable)()
    SetHealth(value TeamworkDeviceHealthable)()
    SetHealthStatus(value *TeamworkDeviceHealthStatus)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value *string)()
    SetOperations(value []TeamworkDeviceOperationable)()
}
