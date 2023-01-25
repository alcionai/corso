package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JustifyAction 
type JustifyAction struct {
    InformationProtectionAction
}
// NewJustifyAction instantiates a new JustifyAction and sets the default values.
func NewJustifyAction()(*JustifyAction) {
    m := &JustifyAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.justifyAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateJustifyActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJustifyActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJustifyAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JustifyAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *JustifyAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
