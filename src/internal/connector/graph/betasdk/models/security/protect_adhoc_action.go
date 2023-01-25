package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectAdhocAction 
type ProtectAdhocAction struct {
    InformationProtectionAction
}
// NewProtectAdhocAction instantiates a new ProtectAdhocAction and sets the default values.
func NewProtectAdhocAction()(*ProtectAdhocAction) {
    m := &ProtectAdhocAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.protectAdhocAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateProtectAdhocActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectAdhocActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectAdhocAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectAdhocAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ProtectAdhocAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
