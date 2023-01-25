package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnTokenIssuanceStartCustomExtensionHandler 
type OnTokenIssuanceStartCustomExtensionHandler struct {
    OnTokenIssuanceStartHandler
    // The customExtension property
    customExtension OnTokenIssuanceStartCustomExtensionable
}
// NewOnTokenIssuanceStartCustomExtensionHandler instantiates a new OnTokenIssuanceStartCustomExtensionHandler and sets the default values.
func NewOnTokenIssuanceStartCustomExtensionHandler()(*OnTokenIssuanceStartCustomExtensionHandler) {
    m := &OnTokenIssuanceStartCustomExtensionHandler{
        OnTokenIssuanceStartHandler: *NewOnTokenIssuanceStartHandler(),
    }
    odataTypeValue := "#microsoft.graph.onTokenIssuanceStartCustomExtensionHandler";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnTokenIssuanceStartCustomExtensionHandlerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnTokenIssuanceStartCustomExtensionHandlerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartCustomExtensionHandler(), nil
}
// GetCustomExtension gets the customExtension property value. The customExtension property
func (m *OnTokenIssuanceStartCustomExtensionHandler) GetCustomExtension()(OnTokenIssuanceStartCustomExtensionable) {
    return m.customExtension
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnTokenIssuanceStartCustomExtensionHandler) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnTokenIssuanceStartHandler.GetFieldDeserializers()
    res["customExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnTokenIssuanceStartCustomExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(OnTokenIssuanceStartCustomExtensionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OnTokenIssuanceStartCustomExtensionHandler) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnTokenIssuanceStartHandler.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. The customExtension property
func (m *OnTokenIssuanceStartCustomExtensionHandler) SetCustomExtension(value OnTokenIssuanceStartCustomExtensionable)() {
    m.customExtension = value
}
