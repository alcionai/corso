package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRecipient 
type MessageRecipient struct {
    Entity
    // The deliveryStatus property
    deliveryStatus *MessageStatus
    // The events property
    events []MessageEventable
    // The recipientEmail property
    recipientEmail *string
}
// NewMessageRecipient instantiates a new MessageRecipient and sets the default values.
func NewMessageRecipient()(*MessageRecipient) {
    m := &MessageRecipient{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMessageRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRecipient(), nil
}
// GetDeliveryStatus gets the deliveryStatus property value. The deliveryStatus property
func (m *MessageRecipient) GetDeliveryStatus()(*MessageStatus) {
    return m.deliveryStatus
}
// GetEvents gets the events property value. The events property
func (m *MessageRecipient) GetEvents()([]MessageEventable) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deliveryStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryStatus(val.(*MessageStatus))
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMessageEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MessageEventable, len(val))
            for i, v := range val {
                res[i] = v.(MessageEventable)
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["recipientEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientEmail(val)
        }
        return nil
    }
    return res
}
// GetRecipientEmail gets the recipientEmail property value. The recipientEmail property
func (m *MessageRecipient) GetRecipientEmail()(*string) {
    return m.recipientEmail
}
// Serialize serializes information the current object
func (m *MessageRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeliveryStatus() != nil {
        cast := (*m.GetDeliveryStatus()).String()
        err = writer.WriteStringValue("deliveryStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmail", m.GetRecipientEmail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeliveryStatus sets the deliveryStatus property value. The deliveryStatus property
func (m *MessageRecipient) SetDeliveryStatus(value *MessageStatus)() {
    m.deliveryStatus = value
}
// SetEvents sets the events property value. The events property
func (m *MessageRecipient) SetEvents(value []MessageEventable)() {
    m.events = value
}
// SetRecipientEmail sets the recipientEmail property value. The recipientEmail property
func (m *MessageRecipient) SetRecipientEmail(value *string)() {
    m.recipientEmail = value
}
