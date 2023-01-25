package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAllowBlockListEntryResultable 
type TenantAllowBlockListEntryResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEntryType()(*TenantAllowBlockListEntryType)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIdentity()(*string)
    GetOdataType()(*string)
    GetStatus()(*LongRunningOperationStatus)
    GetValue()(*string)
    SetEntryType(value *TenantAllowBlockListEntryType)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIdentity(value *string)()
    SetOdataType(value *string)()
    SetStatus(value *LongRunningOperationStatus)()
    SetValue(value *string)()
}
