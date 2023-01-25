package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Commandable 
type Commandable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppServiceName()(*string)
    GetError()(*string)
    GetPackageFamilyName()(*string)
    GetPayload()(PayloadRequestable)
    GetPermissionTicket()(*string)
    GetPostBackUri()(*string)
    GetResponsepayload()(PayloadResponseable)
    GetStatus()(*string)
    GetType()(*string)
    SetAppServiceName(value *string)()
    SetError(value *string)()
    SetPackageFamilyName(value *string)()
    SetPayload(value PayloadRequestable)()
    SetPermissionTicket(value *string)()
    SetPostBackUri(value *string)()
    SetResponsepayload(value PayloadResponseable)()
    SetStatus(value *string)()
    SetType(value *string)()
}
