package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionCalloutResultable 
type CustomExtensionCalloutResultable interface {
    AuthenticationEventHandlerResultable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalloutDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomExtensionId()(*string)
    GetErrorCode()(*int32)
    GetHttpStatus()(*int32)
    GetNumberOfAttempts()(*int32)
    SetCalloutDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomExtensionId(value *string)()
    SetErrorCode(value *int32)()
    SetHttpStatus(value *int32)()
    SetNumberOfAttempts(value *int32)()
}
