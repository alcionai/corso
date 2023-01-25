package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotifyUserActionable 
type NotifyUserActionable interface {
    DlpActionInfoable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmailText()(*string)
    GetOverrideOption()(*OverrideOption)
    GetPolicyTip()(*string)
    GetRecipients()([]string)
    SetActionLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmailText(value *string)()
    SetOverrideOption(value *OverrideOption)()
    SetPolicyTip(value *string)()
    SetRecipients(value []string)()
}
