package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultDeviceCompliancePolicy 
type DefaultDeviceCompliancePolicy struct {
    DeviceCompliancePolicy
}
// NewDefaultDeviceCompliancePolicy instantiates a new DefaultDeviceCompliancePolicy and sets the default values.
func NewDefaultDeviceCompliancePolicy()(*DefaultDeviceCompliancePolicy) {
    m := &DefaultDeviceCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.defaultDeviceCompliancePolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDefaultDeviceCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultDeviceCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultDeviceCompliancePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultDeviceCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DefaultDeviceCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
