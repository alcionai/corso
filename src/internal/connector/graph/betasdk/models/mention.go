package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Mention provides operations to manage the collection of site entities.
type Mention struct {
    Entity
    // The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
    application *string
    // A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
    clientReference *string
    // The email information of the user who made the mention.
    createdBy EmailAddressable
    // The date and time that the mention is created on the client.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
    deepLink *string
    // The mentioned property
    mentioned EmailAddressable
    // Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
    mentionText *string
    // The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
    serverCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewMention instantiates a new mention and sets the default values.
func NewMention()(*Mention) {
    m := &Mention{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMentionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMentionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMention(), nil
}
// GetApplication gets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
func (m *Mention) GetApplication()(*string) {
    return m.application
}
// GetClientReference gets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) GetClientReference()(*string) {
    return m.clientReference
}
// GetCreatedBy gets the createdBy property value. The email information of the user who made the mention.
func (m *Mention) GetCreatedBy()(EmailAddressable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that the mention is created on the client.
func (m *Mention) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeepLink gets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) GetDeepLink()(*string) {
    return m.deepLink
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Mention) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val)
        }
        return nil
    }
    res["clientReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientReference(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailAddressable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deepLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeepLink(val)
        }
        return nil
    }
    res["mentioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentioned(val.(EmailAddressable))
        }
        return nil
    }
    res["mentionText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionText(val)
        }
        return nil
    }
    res["serverCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCreatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetMentioned gets the mentioned property value. The mentioned property
func (m *Mention) GetMentioned()(EmailAddressable) {
    return m.mentioned
}
// GetMentionText gets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
func (m *Mention) GetMentionText()(*string) {
    return m.mentionText
}
// GetServerCreatedDateTime gets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
func (m *Mention) GetServerCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.serverCreatedDateTime
}
// Serialize serializes information the current object
func (m *Mention) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientReference", m.GetClientReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deepLink", m.GetDeepLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mentioned", m.GetMentioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mentionText", m.GetMentionText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("serverCreatedDateTime", m.GetServerCreatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
func (m *Mention) SetApplication(value *string)() {
    m.application = value
}
// SetClientReference sets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) SetClientReference(value *string)() {
    m.clientReference = value
}
// SetCreatedBy sets the createdBy property value. The email information of the user who made the mention.
func (m *Mention) SetCreatedBy(value EmailAddressable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that the mention is created on the client.
func (m *Mention) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeepLink sets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) SetDeepLink(value *string)() {
    m.deepLink = value
}
// SetMentioned sets the mentioned property value. The mentioned property
func (m *Mention) SetMentioned(value EmailAddressable)() {
    m.mentioned = value
}
// SetMentionText sets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
func (m *Mention) SetMentionText(value *string)() {
    m.mentionText = value
}
// SetServerCreatedDateTime sets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
func (m *Mention) SetServerCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.serverCreatedDateTime = value
}
