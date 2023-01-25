package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateReference 
type WindowsUpdateReference struct {
    SoftwareUpdateReference
}
// NewWindowsUpdateReference instantiates a new WindowsUpdateReference and sets the default values.
func NewWindowsUpdateReference()(*WindowsUpdateReference) {
    m := &WindowsUpdateReference{
        SoftwareUpdateReference: *NewSoftwareUpdateReference(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.windowsUpdateReference";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsUpdateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsUpdates.expeditedQualityUpdateReference":
                        return NewExpeditedQualityUpdateReference(), nil
                    case "#microsoft.graph.windowsUpdates.featureUpdateReference":
                        return NewFeatureUpdateReference(), nil
                    case "#microsoft.graph.windowsUpdates.qualityUpdateReference":
                        return NewQualityUpdateReference(), nil
                }
            }
        }
    }
    return NewWindowsUpdateReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SoftwareUpdateReference.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsUpdateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SoftwareUpdateReference.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
