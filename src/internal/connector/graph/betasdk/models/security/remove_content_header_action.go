package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveContentHeaderAction 
type RemoveContentHeaderAction struct {
    InformationProtectionAction
    // The name of the UI element of the header to be removed.
    uiElementNames []string
}
// NewRemoveContentHeaderAction instantiates a new RemoveContentHeaderAction and sets the default values.
func NewRemoveContentHeaderAction()(*RemoveContentHeaderAction) {
    m := &RemoveContentHeaderAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.removeContentHeaderAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRemoveContentHeaderActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveContentHeaderActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveContentHeaderAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveContentHeaderAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["uiElementNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUiElementNames(res)
        }
        return nil
    }
    return res
}
// GetUiElementNames gets the uiElementNames property value. The name of the UI element of the header to be removed.
func (m *RemoveContentHeaderAction) GetUiElementNames()([]string) {
    return m.uiElementNames
}
// Serialize serializes information the current object
func (m *RemoveContentHeaderAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUiElementNames() != nil {
        err = writer.WriteCollectionOfStringValues("uiElementNames", m.GetUiElementNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUiElementNames sets the uiElementNames property value. The name of the UI element of the header to be removed.
func (m *RemoveContentHeaderAction) SetUiElementNames(value []string)() {
    m.uiElementNames = value
}
