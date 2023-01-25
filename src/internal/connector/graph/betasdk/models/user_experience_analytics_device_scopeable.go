package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScopeable 
type UserExperienceAnalyticsDeviceScopeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceScopeName()(*string)
    GetEnabled()(*bool)
    GetIsBuiltIn()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOperator()(*DeviceScopeOperator)
    GetOwnerId()(*string)
    GetParameter()(*DeviceScopeParameter)
    GetStatus()(*DeviceScopeStatus)
    GetValue()(*string)
    GetValueObjectId()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceScopeName(value *string)()
    SetEnabled(value *bool)()
    SetIsBuiltIn(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOperator(value *DeviceScopeOperator)()
    SetOwnerId(value *string)()
    SetParameter(value *DeviceScopeParameter)()
    SetStatus(value *DeviceScopeStatus)()
    SetValue(value *string)()
    SetValueObjectId(value *string)()
}
