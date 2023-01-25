package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InvokeUserFlowListener 
type InvokeUserFlowListener struct {
    AuthenticationListener
    // The user flow that is invoked when this action executes.
    userFlow B2xIdentityUserFlowable
}
// NewInvokeUserFlowListener instantiates a new InvokeUserFlowListener and sets the default values.
func NewInvokeUserFlowListener()(*InvokeUserFlowListener) {
    m := &InvokeUserFlowListener{
        AuthenticationListener: *NewAuthenticationListener(),
    }
    odataTypeValue := "#microsoft.graph.invokeUserFlowListener";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateInvokeUserFlowListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvokeUserFlowListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvokeUserFlowListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvokeUserFlowListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationListener.GetFieldDeserializers()
    res["userFlow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateB2xIdentityUserFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserFlow(val.(B2xIdentityUserFlowable))
        }
        return nil
    }
    return res
}
// GetUserFlow gets the userFlow property value. The user flow that is invoked when this action executes.
func (m *InvokeUserFlowListener) GetUserFlow()(B2xIdentityUserFlowable) {
    return m.userFlow
}
// Serialize serializes information the current object
func (m *InvokeUserFlowListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationListener.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("userFlow", m.GetUserFlow())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserFlow sets the userFlow property value. The user flow that is invoked when this action executes.
func (m *InvokeUserFlowListener) SetUserFlow(value B2xIdentityUserFlowable)() {
    m.userFlow = value
}
