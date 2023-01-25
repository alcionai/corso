package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicyGroupAssignment 
type DeviceCompliancePolicyGroupAssignment struct {
    Entity
    // The navigation link to the  device compliance polic targeted.
    deviceCompliancePolicy DeviceCompliancePolicyable
    // Indicates if this group is should be excluded. Defaults that the group should be included
    excludeGroup *bool
    // The Id of the AAD group we are targeting the device compliance policy to.
    targetGroupId *string
}
// NewDeviceCompliancePolicyGroupAssignment instantiates a new DeviceCompliancePolicyGroupAssignment and sets the default values.
func NewDeviceCompliancePolicyGroupAssignment()(*DeviceCompliancePolicyGroupAssignment) {
    m := &DeviceCompliancePolicyGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicyGroupAssignment(), nil
}
// GetDeviceCompliancePolicy gets the deviceCompliancePolicy property value. The navigation link to the  device compliance polic targeted.
func (m *DeviceCompliancePolicyGroupAssignment) GetDeviceCompliancePolicy()(DeviceCompliancePolicyable) {
    return m.deviceCompliancePolicy
}
// GetExcludeGroup gets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceCompliancePolicyGroupAssignment) GetExcludeGroup()(*bool) {
    return m.excludeGroup
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicyGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCompliancePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicy(val.(DeviceCompliancePolicyable))
        }
        return nil
    }
    res["excludeGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeGroup(val)
        }
        return nil
    }
    res["targetGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
// GetTargetGroupId gets the targetGroupId property value. The Id of the AAD group we are targeting the device compliance policy to.
func (m *DeviceCompliancePolicyGroupAssignment) GetTargetGroupId()(*string) {
    return m.targetGroupId
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicyGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicy", m.GetDeviceCompliancePolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("excludeGroup", m.GetExcludeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceCompliancePolicy sets the deviceCompliancePolicy property value. The navigation link to the  device compliance polic targeted.
func (m *DeviceCompliancePolicyGroupAssignment) SetDeviceCompliancePolicy(value DeviceCompliancePolicyable)() {
    m.deviceCompliancePolicy = value
}
// SetExcludeGroup sets the excludeGroup property value. Indicates if this group is should be excluded. Defaults that the group should be included
func (m *DeviceCompliancePolicyGroupAssignment) SetExcludeGroup(value *bool)() {
    m.excludeGroup = value
}
// SetTargetGroupId sets the targetGroupId property value. The Id of the AAD group we are targeting the device compliance policy to.
func (m *DeviceCompliancePolicyGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
