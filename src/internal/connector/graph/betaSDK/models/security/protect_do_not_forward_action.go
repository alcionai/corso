package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectDoNotForwardAction 
type ProtectDoNotForwardAction struct {
    InformationProtectionAction
}
// NewProtectDoNotForwardAction instantiates a new ProtectDoNotForwardAction and sets the default values.
func NewProtectDoNotForwardAction()(*ProtectDoNotForwardAction) {
    m := &ProtectDoNotForwardAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.protectDoNotForwardAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateProtectDoNotForwardActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectDoNotForwardActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectDoNotForwardAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectDoNotForwardAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ProtectDoNotForwardAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
