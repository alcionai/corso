package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeConfigurationGroupAssignmentTarget 
type OfficeConfigurationGroupAssignmentTarget struct {
    OfficeConfigurationAssignmentTarget
    // The Id of the AAD group we are targeting the device configuration to.
    groupId *string
}
// NewOfficeConfigurationGroupAssignmentTarget instantiates a new OfficeConfigurationGroupAssignmentTarget and sets the default values.
func NewOfficeConfigurationGroupAssignmentTarget()(*OfficeConfigurationGroupAssignmentTarget) {
    m := &OfficeConfigurationGroupAssignmentTarget{
        OfficeConfigurationAssignmentTarget: *NewOfficeConfigurationAssignmentTarget(),
    }
    odataTypeValue := "#microsoft.graph.officeConfigurationGroupAssignmentTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOfficeConfigurationGroupAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeConfigurationGroupAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeConfigurationGroupAssignmentTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeConfigurationGroupAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OfficeConfigurationAssignmentTarget.GetFieldDeserializers()
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *OfficeConfigurationGroupAssignmentTarget) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *OfficeConfigurationGroupAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OfficeConfigurationAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupId sets the groupId property value. The Id of the AAD group we are targeting the device configuration to.
func (m *OfficeConfigurationGroupAssignmentTarget) SetGroupId(value *string)() {
    m.groupId = value
}
