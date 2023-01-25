package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailUrlThreatSubmission 
type EmailUrlThreatSubmission struct {
    EmailThreatSubmission
    // Specifies the url of the message to be submitted.
    messageUrl *string
}
// NewEmailUrlThreatSubmission instantiates a new EmailUrlThreatSubmission and sets the default values.
func NewEmailUrlThreatSubmission()(*EmailUrlThreatSubmission) {
    m := &EmailUrlThreatSubmission{
        EmailThreatSubmission: *NewEmailThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.emailUrlThreatSubmission";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEmailUrlThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailUrlThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailUrlThreatSubmission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailUrlThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EmailThreatSubmission.GetFieldDeserializers()
    res["messageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageUrl(val)
        }
        return nil
    }
    return res
}
// GetMessageUrl gets the messageUrl property value. Specifies the url of the message to be submitted.
func (m *EmailUrlThreatSubmission) GetMessageUrl()(*string) {
    return m.messageUrl
}
// Serialize serializes information the current object
func (m *EmailUrlThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EmailThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("messageUrl", m.GetMessageUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMessageUrl sets the messageUrl property value. Specifies the url of the message to be submitted.
func (m *EmailUrlThreatSubmission) SetMessageUrl(value *string)() {
    m.messageUrl = value
}
