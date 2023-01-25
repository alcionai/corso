package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PresenceStatusMessage 
type PresenceStatusMessage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Time in which the status message expires.If not provided, the status message does not expire.expiryDateTime.dateTime should not include time zone.expiryDateTime is not available when requesting presence of another user.
    expiryDateTime DateTimeTimeZoneable
    // Status message item. The only supported format currently is message.contentType = 'text'.
    message ItemBodyable
    // The OdataType property
    odataType *string
    // Time in which the status message was published.Read-only.publishedDateTime is not available when requesting presence of another user.
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewPresenceStatusMessage instantiates a new presenceStatusMessage and sets the default values.
func NewPresenceStatusMessage()(*PresenceStatusMessage) {
    m := &PresenceStatusMessage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePresenceStatusMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePresenceStatusMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPresenceStatusMessage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PresenceStatusMessage) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpiryDateTime gets the expiryDateTime property value. Time in which the status message expires.If not provided, the status message does not expire.expiryDateTime.dateTime should not include time zone.expiryDateTime is not available when requesting presence of another user.
func (m *PresenceStatusMessage) GetExpiryDateTime()(DateTimeTimeZoneable) {
    return m.expiryDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PresenceStatusMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(ItemBodyable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["publishedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Status message item. The only supported format currently is message.contentType = 'text'.
func (m *PresenceStatusMessage) GetMessage()(ItemBodyable) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PresenceStatusMessage) GetOdataType()(*string) {
    return m.odataType
}
// GetPublishedDateTime gets the publishedDateTime property value. Time in which the status message was published.Read-only.publishedDateTime is not available when requesting presence of another user.
func (m *PresenceStatusMessage) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.publishedDateTime
}
// Serialize serializes information the current object
func (m *PresenceStatusMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("expiryDateTime", m.GetExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PresenceStatusMessage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpiryDateTime sets the expiryDateTime property value. Time in which the status message expires.If not provided, the status message does not expire.expiryDateTime.dateTime should not include time zone.expiryDateTime is not available when requesting presence of another user.
func (m *PresenceStatusMessage) SetExpiryDateTime(value DateTimeTimeZoneable)() {
    m.expiryDateTime = value
}
// SetMessage sets the message property value. Status message item. The only supported format currently is message.contentType = 'text'.
func (m *PresenceStatusMessage) SetMessage(value ItemBodyable)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PresenceStatusMessage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPublishedDateTime sets the publishedDateTime property value. Time in which the status message was published.Read-only.publishedDateTime is not available when requesting presence of another user.
func (m *PresenceStatusMessage) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
