package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomAction 
type CustomAction struct {
    InformationProtectionAction
    // Name of the custom action.
    name *string
    // Properties, in key value pair format, of the action.
    properties []KeyValuePairable
}
// NewCustomAction instantiates a new CustomAction and sets the default values.
func NewCustomAction()(*CustomAction) {
    m := &CustomAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.customAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the custom action.
func (m *CustomAction) GetName()(*string) {
    return m.name
}
// GetProperties gets the properties property value. Properties, in key value pair format, of the action.
func (m *CustomAction) GetProperties()([]KeyValuePairable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *CustomAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name of the custom action.
func (m *CustomAction) SetName(value *string)() {
    m.name = value
}
// SetProperties sets the properties property value. Properties, in key value pair format, of the action.
func (m *CustomAction) SetProperties(value []KeyValuePairable)() {
    m.properties = value
}
