package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnTokenIssuanceStartListener 
type OnTokenIssuanceStartListener struct {
    AuthenticationEventListener
    // The handler property
    handler OnTokenIssuanceStartHandlerable
}
// NewOnTokenIssuanceStartListener instantiates a new OnTokenIssuanceStartListener and sets the default values.
func NewOnTokenIssuanceStartListener()(*OnTokenIssuanceStartListener) {
    m := &OnTokenIssuanceStartListener{
        AuthenticationEventListener: *NewAuthenticationEventListener(),
    }
    odataTypeValue := "#microsoft.graph.onTokenIssuanceStartListener";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnTokenIssuanceStartListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnTokenIssuanceStartListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnTokenIssuanceStartListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationEventListener.GetFieldDeserializers()
    res["handler"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnTokenIssuanceStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandler(val.(OnTokenIssuanceStartHandlerable))
        }
        return nil
    }
    return res
}
// GetHandler gets the handler property value. The handler property
func (m *OnTokenIssuanceStartListener) GetHandler()(OnTokenIssuanceStartHandlerable) {
    return m.handler
}
// Serialize serializes information the current object
func (m *OnTokenIssuanceStartListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationEventListener.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("handler", m.GetHandler())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHandler sets the handler property value. The handler property
func (m *OnTokenIssuanceStartListener) SetHandler(value OnTokenIssuanceStartHandlerable)() {
    m.handler = value
}
