package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDeviceable 
type ComanagementEligibleDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientRegistrationStatus()(*DeviceRegistrationState)
    GetDeviceName()(*string)
    GetDeviceType()(*DeviceType)
    GetEntitySource()(*int32)
    GetManagementAgents()(*ManagementAgentType)
    GetManagementState()(*ManagementState)
    GetManufacturer()(*string)
    GetMdmStatus()(*string)
    GetModel()(*string)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetOwnerType()(*OwnerType)
    GetReferenceId()(*string)
    GetSerialNumber()(*string)
    GetStatus()(*ComanagementEligibleType)
    GetUpn()(*string)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    SetClientRegistrationStatus(value *DeviceRegistrationState)()
    SetDeviceName(value *string)()
    SetDeviceType(value *DeviceType)()
    SetEntitySource(value *int32)()
    SetManagementAgents(value *ManagementAgentType)()
    SetManagementState(value *ManagementState)()
    SetManufacturer(value *string)()
    SetMdmStatus(value *string)()
    SetModel(value *string)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetOwnerType(value *OwnerType)()
    SetReferenceId(value *string)()
    SetSerialNumber(value *string)()
    SetStatus(value *ComanagementEligibleType)()
    SetUpn(value *string)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
}
