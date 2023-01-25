package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailPayloadDetail 
type EmailPayloadDetail struct {
    PayloadDetail
    // Email address of the user.
    fromEmail *string
    // Display name of the user.
    fromName *string
    // Indicates whether the sender is not from the user's organization.
    isExternalSender *bool
    // The subject of the email address sent to the user.
    subject *string
}
// NewEmailPayloadDetail instantiates a new EmailPayloadDetail and sets the default values.
func NewEmailPayloadDetail()(*EmailPayloadDetail) {
    m := &EmailPayloadDetail{
        PayloadDetail: *NewPayloadDetail(),
    }
    odataTypeValue := "#microsoft.graph.emailPayloadDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEmailPayloadDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailPayloadDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailPayloadDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailPayloadDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PayloadDetail.GetFieldDeserializers()
    res["fromEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromEmail(val)
        }
        return nil
    }
    res["fromName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromName(val)
        }
        return nil
    }
    res["isExternalSender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalSender(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetFromEmail gets the fromEmail property value. Email address of the user.
func (m *EmailPayloadDetail) GetFromEmail()(*string) {
    return m.fromEmail
}
// GetFromName gets the fromName property value. Display name of the user.
func (m *EmailPayloadDetail) GetFromName()(*string) {
    return m.fromName
}
// GetIsExternalSender gets the isExternalSender property value. Indicates whether the sender is not from the user's organization.
func (m *EmailPayloadDetail) GetIsExternalSender()(*bool) {
    return m.isExternalSender
}
// GetSubject gets the subject property value. The subject of the email address sent to the user.
func (m *EmailPayloadDetail) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *EmailPayloadDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PayloadDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fromEmail", m.GetFromEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fromName", m.GetFromName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExternalSender", m.GetIsExternalSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFromEmail sets the fromEmail property value. Email address of the user.
func (m *EmailPayloadDetail) SetFromEmail(value *string)() {
    m.fromEmail = value
}
// SetFromName sets the fromName property value. Display name of the user.
func (m *EmailPayloadDetail) SetFromName(value *string)() {
    m.fromName = value
}
// SetIsExternalSender sets the isExternalSender property value. Indicates whether the sender is not from the user's organization.
func (m *EmailPayloadDetail) SetIsExternalSender(value *bool)() {
    m.isExternalSender = value
}
// SetSubject sets the subject property value. The subject of the email address sent to the user.
func (m *EmailPayloadDetail) SetSubject(value *string)() {
    m.subject = value
}
