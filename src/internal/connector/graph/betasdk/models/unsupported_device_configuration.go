package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnsupportedDeviceConfiguration 
type UnsupportedDeviceConfiguration struct {
    DeviceConfiguration
    // Details describing why the entity is unsupported. This collection can contain a maximum of 1000 elements.
    details []UnsupportedDeviceConfigurationDetailable
    // The type of entity that would be returned otherwise.
    originalEntityTypeName *string
}
// NewUnsupportedDeviceConfiguration instantiates a new UnsupportedDeviceConfiguration and sets the default values.
func NewUnsupportedDeviceConfiguration()(*UnsupportedDeviceConfiguration) {
    m := &UnsupportedDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.unsupportedDeviceConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUnsupportedDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnsupportedDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnsupportedDeviceConfiguration(), nil
}
// GetDetails gets the details property value. Details describing why the entity is unsupported. This collection can contain a maximum of 1000 elements.
func (m *UnsupportedDeviceConfiguration) GetDetails()([]UnsupportedDeviceConfigurationDetailable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnsupportedDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnsupportedDeviceConfigurationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnsupportedDeviceConfigurationDetailable, len(val))
            for i, v := range val {
                res[i] = v.(UnsupportedDeviceConfigurationDetailable)
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["originalEntityTypeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEntityTypeName(val)
        }
        return nil
    }
    return res
}
// GetOriginalEntityTypeName gets the originalEntityTypeName property value. The type of entity that would be returned otherwise.
func (m *UnsupportedDeviceConfiguration) GetOriginalEntityTypeName()(*string) {
    return m.originalEntityTypeName
}
// Serialize serializes information the current object
func (m *UnsupportedDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalEntityTypeName", m.GetOriginalEntityTypeName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. Details describing why the entity is unsupported. This collection can contain a maximum of 1000 elements.
func (m *UnsupportedDeviceConfiguration) SetDetails(value []UnsupportedDeviceConfigurationDetailable)() {
    m.details = value
}
// SetOriginalEntityTypeName sets the originalEntityTypeName property value. The type of entity that would be returned otherwise.
func (m *UnsupportedDeviceConfiguration) SetOriginalEntityTypeName(value *string)() {
    m.originalEntityTypeName = value
}
