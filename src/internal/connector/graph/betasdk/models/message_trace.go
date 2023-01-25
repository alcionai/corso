package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageTrace 
type MessageTrace struct {
    Entity
    // The destinationIPAddress property
    destinationIPAddress *string
    // The messageId property
    messageId *string
    // The receivedDateTime property
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The recipients property
    recipients []MessageRecipientable
    // The senderEmail property
    senderEmail *string
    // The size property
    size *int32
    // The sourceIPAddress property
    sourceIPAddress *string
    // The subject property
    subject *string
}
// NewMessageTrace instantiates a new MessageTrace and sets the default values.
func NewMessageTrace()(*MessageTrace) {
    m := &MessageTrace{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMessageTraceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageTraceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageTrace(), nil
}
// GetDestinationIPAddress gets the destinationIPAddress property value. The destinationIPAddress property
func (m *MessageTrace) GetDestinationIPAddress()(*string) {
    return m.destinationIPAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageTrace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["destinationIPAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationIPAddress(val)
        }
        return nil
    }
    res["messageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(MessageRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["senderEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderEmail(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sourceIPAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIPAddress(val)
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
// GetMessageId gets the messageId property value. The messageId property
func (m *MessageTrace) GetMessageId()(*string) {
    return m.messageId
}
// GetReceivedDateTime gets the receivedDateTime property value. The receivedDateTime property
func (m *MessageTrace) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
// GetRecipients gets the recipients property value. The recipients property
func (m *MessageTrace) GetRecipients()([]MessageRecipientable) {
    return m.recipients
}
// GetSenderEmail gets the senderEmail property value. The senderEmail property
func (m *MessageTrace) GetSenderEmail()(*string) {
    return m.senderEmail
}
// GetSize gets the size property value. The size property
func (m *MessageTrace) GetSize()(*int32) {
    return m.size
}
// GetSourceIPAddress gets the sourceIPAddress property value. The sourceIPAddress property
func (m *MessageTrace) GetSourceIPAddress()(*string) {
    return m.sourceIPAddress
}
// GetSubject gets the subject property value. The subject property
func (m *MessageTrace) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *MessageTrace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("destinationIPAddress", m.GetDestinationIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderEmail", m.GetSenderEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceIPAddress", m.GetSourceIPAddress())
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
// SetDestinationIPAddress sets the destinationIPAddress property value. The destinationIPAddress property
func (m *MessageTrace) SetDestinationIPAddress(value *string)() {
    m.destinationIPAddress = value
}
// SetMessageId sets the messageId property value. The messageId property
func (m *MessageTrace) SetMessageId(value *string)() {
    m.messageId = value
}
// SetReceivedDateTime sets the receivedDateTime property value. The receivedDateTime property
func (m *MessageTrace) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *MessageTrace) SetRecipients(value []MessageRecipientable)() {
    m.recipients = value
}
// SetSenderEmail sets the senderEmail property value. The senderEmail property
func (m *MessageTrace) SetSenderEmail(value *string)() {
    m.senderEmail = value
}
// SetSize sets the size property value. The size property
func (m *MessageTrace) SetSize(value *int32)() {
    m.size = value
}
// SetSourceIPAddress sets the sourceIPAddress property value. The sourceIPAddress property
func (m *MessageTrace) SetSourceIPAddress(value *string)() {
    m.sourceIPAddress = value
}
// SetSubject sets the subject property value. The subject property
func (m *MessageTrace) SetSubject(value *string)() {
    m.subject = value
}
