package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotifyUserAction 
type NotifyUserAction struct {
    DlpActionInfo
    // The actionLastModifiedDateTime property
    actionLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The emailText property
    emailText *string
    // The overrideOption property
    overrideOption *OverrideOption
    // The policyTip property
    policyTip *string
    // The recipients property
    recipients []string
}
// NewNotifyUserAction instantiates a new NotifyUserAction and sets the default values.
func NewNotifyUserAction()(*NotifyUserAction) {
    m := &NotifyUserAction{
        DlpActionInfo: *NewDlpActionInfo(),
    }
    return m
}
// CreateNotifyUserActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotifyUserActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotifyUserAction(), nil
}
// GetActionLastModifiedDateTime gets the actionLastModifiedDateTime property value. The actionLastModifiedDateTime property
func (m *NotifyUserAction) GetActionLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.actionLastModifiedDateTime
}
// GetEmailText gets the emailText property value. The emailText property
func (m *NotifyUserAction) GetEmailText()(*string) {
    return m.emailText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotifyUserAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpActionInfo.GetFieldDeserializers()
    res["actionLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionLastModifiedDateTime(val)
        }
        return nil
    }
    res["emailText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailText(val)
        }
        return nil
    }
    res["overrideOption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOverrideOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverrideOption(val.(*OverrideOption))
        }
        return nil
    }
    res["policyTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTip(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRecipients(res)
        }
        return nil
    }
    return res
}
// GetOverrideOption gets the overrideOption property value. The overrideOption property
func (m *NotifyUserAction) GetOverrideOption()(*OverrideOption) {
    return m.overrideOption
}
// GetPolicyTip gets the policyTip property value. The policyTip property
func (m *NotifyUserAction) GetPolicyTip()(*string) {
    return m.policyTip
}
// GetRecipients gets the recipients property value. The recipients property
func (m *NotifyUserAction) GetRecipients()([]string) {
    return m.recipients
}
// Serialize serializes information the current object
func (m *NotifyUserAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpActionInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("actionLastModifiedDateTime", m.GetActionLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailText", m.GetEmailText())
        if err != nil {
            return err
        }
    }
    if m.GetOverrideOption() != nil {
        cast := (*m.GetOverrideOption()).String()
        err = writer.WriteStringValue("overrideOption", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyTip", m.GetPolicyTip())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        err = writer.WriteCollectionOfStringValues("recipients", m.GetRecipients())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionLastModifiedDateTime sets the actionLastModifiedDateTime property value. The actionLastModifiedDateTime property
func (m *NotifyUserAction) SetActionLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.actionLastModifiedDateTime = value
}
// SetEmailText sets the emailText property value. The emailText property
func (m *NotifyUserAction) SetEmailText(value *string)() {
    m.emailText = value
}
// SetOverrideOption sets the overrideOption property value. The overrideOption property
func (m *NotifyUserAction) SetOverrideOption(value *OverrideOption)() {
    m.overrideOption = value
}
// SetPolicyTip sets the policyTip property value. The policyTip property
func (m *NotifyUserAction) SetPolicyTip(value *string)() {
    m.policyTip = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *NotifyUserAction) SetRecipients(value []string)() {
    m.recipients = value
}
