package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageTraceable 
type MessageTraceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestinationIPAddress()(*string)
    GetMessageId()(*string)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipients()([]MessageRecipientable)
    GetSenderEmail()(*string)
    GetSize()(*int32)
    GetSourceIPAddress()(*string)
    GetSubject()(*string)
    SetDestinationIPAddress(value *string)()
    SetMessageId(value *string)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipients(value []MessageRecipientable)()
    SetSenderEmail(value *string)()
    SetSize(value *int32)()
    SetSourceIPAddress(value *string)()
    SetSubject(value *string)()
}
