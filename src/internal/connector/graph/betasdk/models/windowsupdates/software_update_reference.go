package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareUpdateReference 
type SoftwareUpdateReference struct {
    DeployableContent
}
// NewSoftwareUpdateReference instantiates a new SoftwareUpdateReference and sets the default values.
func NewSoftwareUpdateReference()(*SoftwareUpdateReference) {
    m := &SoftwareUpdateReference{
        DeployableContent: *NewDeployableContent(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.softwareUpdateReference";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSoftwareUpdateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSoftwareUpdateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.windowsUpdateReference":
                        return NewWindowsUpdateReference(), nil
                }
            }
        }
    }
    return NewSoftwareUpdateReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareUpdateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeployableContent.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SoftwareUpdateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeployableContent.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
