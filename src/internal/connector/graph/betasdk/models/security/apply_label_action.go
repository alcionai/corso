package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplyLabelAction 
type ApplyLabelAction struct {
    InformationProtectionAction
    // The collection of actions that should be implemented by the caller.
    actions []InformationProtectionActionable
    // The actionSource property
    actionSource *ActionSource
    // If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
    responsibleSensitiveTypeIds []string
    // The sensitivityLabelId property
    sensitivityLabelId *string
}
// NewApplyLabelAction instantiates a new ApplyLabelAction and sets the default values.
func NewApplyLabelAction()(*ApplyLabelAction) {
    m := &ApplyLabelAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.applyLabelAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateApplyLabelActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplyLabelActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplyLabelAction(), nil
}
// GetActions gets the actions property value. The collection of actions that should be implemented by the caller.
func (m *ApplyLabelAction) GetActions()([]InformationProtectionActionable) {
    return m.actions
}
// GetActionSource gets the actionSource property value. The actionSource property
func (m *ApplyLabelAction) GetActionSource()(*ActionSource) {
    return m.actionSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyLabelAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInformationProtectionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionActionable, len(val))
            for i, v := range val {
                res[i] = v.(InformationProtectionActionable)
            }
            m.SetActions(res)
        }
        return nil
    }
    res["actionSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSource(val.(*ActionSource))
        }
        return nil
    }
    res["responsibleSensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResponsibleSensitiveTypeIds(res)
        }
        return nil
    }
    res["sensitivityLabelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabelId(val)
        }
        return nil
    }
    return res
}
// GetResponsibleSensitiveTypeIds gets the responsibleSensitiveTypeIds property value. If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
func (m *ApplyLabelAction) GetResponsibleSensitiveTypeIds()([]string) {
    return m.responsibleSensitiveTypeIds
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *ApplyLabelAction) GetSensitivityLabelId()(*string) {
    return m.sensitivityLabelId
}
// Serialize serializes information the current object
func (m *ApplyLabelAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActions()))
        for i, v := range m.GetActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("actions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionSource() != nil {
        cast := (*m.GetActionSource()).String()
        err = writer.WriteStringValue("actionSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponsibleSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("responsibleSensitiveTypeIds", m.GetResponsibleSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sensitivityLabelId", m.GetSensitivityLabelId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The collection of actions that should be implemented by the caller.
func (m *ApplyLabelAction) SetActions(value []InformationProtectionActionable)() {
    m.actions = value
}
// SetActionSource sets the actionSource property value. The actionSource property
func (m *ApplyLabelAction) SetActionSource(value *ActionSource)() {
    m.actionSource = value
}
// SetResponsibleSensitiveTypeIds sets the responsibleSensitiveTypeIds property value. If the label was the result of an automatic classification, supply the list of sensitive info type GUIDs that resulted in the returned label.
func (m *ApplyLabelAction) SetResponsibleSensitiveTypeIds(value []string)() {
    m.responsibleSensitiveTypeIds = value
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *ApplyLabelAction) SetSensitivityLabelId(value *string)() {
    m.sensitivityLabelId = value
}
