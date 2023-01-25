package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppProvisioningConfigGroupAssignment contains the properties used to assign an App provisioning configuration to a group.
type MobileAppProvisioningConfigGroupAssignment struct {
    Entity
    // The ID of the AAD group in which the app provisioning configuration is being targeted.
    targetGroupId *string
}
// NewMobileAppProvisioningConfigGroupAssignment instantiates a new mobileAppProvisioningConfigGroupAssignment and sets the default values.
func NewMobileAppProvisioningConfigGroupAssignment()(*MobileAppProvisioningConfigGroupAssignment) {
    m := &MobileAppProvisioningConfigGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppProvisioningConfigGroupAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppProvisioningConfigGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
// GetTargetGroupId gets the targetGroupId property value. The ID of the AAD group in which the app provisioning configuration is being targeted.
func (m *MobileAppProvisioningConfigGroupAssignment) GetTargetGroupId()(*string) {
    return m.targetGroupId
}
// Serialize serializes information the current object
func (m *MobileAppProvisioningConfigGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetGroupId sets the targetGroupId property value. The ID of the AAD group in which the app provisioning configuration is being targeted.
func (m *MobileAppProvisioningConfigGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
