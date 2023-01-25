package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleVppTokenTroubleshootingEvent 
type AppleVppTokenTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
    // Apple Volume Purchase Program Token Identifier.
    tokenId *string
}
// NewAppleVppTokenTroubleshootingEvent instantiates a new AppleVppTokenTroubleshootingEvent and sets the default values.
func NewAppleVppTokenTroubleshootingEvent()(*AppleVppTokenTroubleshootingEvent) {
    m := &AppleVppTokenTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    return m
}
// CreateAppleVppTokenTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleVppTokenTroubleshootingEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleVppTokenTroubleshootingEvent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleVppTokenTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTroubleshootingEvent.GetFieldDeserializers()
    res["tokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenId(val)
        }
        return nil
    }
    return res
}
// GetTokenId gets the tokenId property value. Apple Volume Purchase Program Token Identifier.
func (m *AppleVppTokenTroubleshootingEvent) GetTokenId()(*string) {
    return m.tokenId
}
// Serialize serializes information the current object
func (m *AppleVppTokenTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTroubleshootingEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tokenId", m.GetTokenId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTokenId sets the tokenId property value. Apple Volume Purchase Program Token Identifier.
func (m *AppleVppTokenTroubleshootingEvent) SetTokenId(value *string)() {
    m.tokenId = value
}
