package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureUpdateReference 
type FeatureUpdateReference struct {
    WindowsUpdateReference
    // Specifies a feature update by version.
    version *string
}
// NewFeatureUpdateReference instantiates a new FeatureUpdateReference and sets the default values.
func NewFeatureUpdateReference()(*FeatureUpdateReference) {
    m := &FeatureUpdateReference{
        WindowsUpdateReference: *NewWindowsUpdateReference(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.featureUpdateReference";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFeatureUpdateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureUpdateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureUpdateReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureUpdateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateReference.GetFieldDeserializers()
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetVersion gets the version property value. Specifies a feature update by version.
func (m *FeatureUpdateReference) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *FeatureUpdateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateReference.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVersion sets the version property value. Specifies a feature update by version.
func (m *FeatureUpdateReference) SetVersion(value *string)() {
    m.version = value
}
