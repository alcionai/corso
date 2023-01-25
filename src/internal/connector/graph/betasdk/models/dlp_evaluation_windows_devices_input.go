package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpEvaluationWindowsDevicesInput 
type DlpEvaluationWindowsDevicesInput struct {
    DlpEvaluationInput
    // The contentProperties property
    contentProperties ContentPropertiesable
    // The sharedBy property
    sharedBy *string
}
// NewDlpEvaluationWindowsDevicesInput instantiates a new DlpEvaluationWindowsDevicesInput and sets the default values.
func NewDlpEvaluationWindowsDevicesInput()(*DlpEvaluationWindowsDevicesInput) {
    m := &DlpEvaluationWindowsDevicesInput{
        DlpEvaluationInput: *NewDlpEvaluationInput(),
    }
    odataTypeValue := "#microsoft.graph.dlpEvaluationWindowsDevicesInput";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDlpEvaluationWindowsDevicesInputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpEvaluationWindowsDevicesInputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDlpEvaluationWindowsDevicesInput(), nil
}
// GetContentProperties gets the contentProperties property value. The contentProperties property
func (m *DlpEvaluationWindowsDevicesInput) GetContentProperties()(ContentPropertiesable) {
    return m.contentProperties
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpEvaluationWindowsDevicesInput) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpEvaluationInput.GetFieldDeserializers()
    res["contentProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentPropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentProperties(val.(ContentPropertiesable))
        }
        return nil
    }
    res["sharedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedBy(val)
        }
        return nil
    }
    return res
}
// GetSharedBy gets the sharedBy property value. The sharedBy property
func (m *DlpEvaluationWindowsDevicesInput) GetSharedBy()(*string) {
    return m.sharedBy
}
// Serialize serializes information the current object
func (m *DlpEvaluationWindowsDevicesInput) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpEvaluationInput.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contentProperties", m.GetContentProperties())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharedBy", m.GetSharedBy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentProperties sets the contentProperties property value. The contentProperties property
func (m *DlpEvaluationWindowsDevicesInput) SetContentProperties(value ContentPropertiesable)() {
    m.contentProperties = value
}
// SetSharedBy sets the sharedBy property value. The sharedBy property
func (m *DlpEvaluationWindowsDevicesInput) SetSharedBy(value *string)() {
    m.sharedBy = value
}
