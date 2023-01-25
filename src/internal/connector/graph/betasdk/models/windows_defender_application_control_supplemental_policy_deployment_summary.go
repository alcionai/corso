package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary 
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary struct {
    Entity
    // Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
    deployedDeviceCount *int32
    // Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
    failedDeviceCount *int32
}
// NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary instantiates a new windowsDefenderApplicationControlSupplementalPolicyDeploymentSummary and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary(), nil
}
// GetDeployedDeviceCount gets the deployedDeviceCount property value. Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetDeployedDeviceCount()(*int32) {
    return m.deployedDeviceCount
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetFailedDeviceCount()(*int32) {
    return m.failedDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deployedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedDeviceCount(val)
        }
        return nil
    }
    res["failedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deployedDeviceCount", m.GetDeployedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeployedDeviceCount sets the deployedDeviceCount property value. Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) SetDeployedDeviceCount(value *int32)() {
    m.deployedDeviceCount = value
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
