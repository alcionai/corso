package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcManagementGroupAssignmentTarget 
type CloudPcManagementGroupAssignmentTarget struct {
    CloudPcManagementAssignmentTarget
    // The id of the assignment's target group
    groupId *string
    // The servicePlanId property
    servicePlanId *string
}
// NewCloudPcManagementGroupAssignmentTarget instantiates a new CloudPcManagementGroupAssignmentTarget and sets the default values.
func NewCloudPcManagementGroupAssignmentTarget()(*CloudPcManagementGroupAssignmentTarget) {
    m := &CloudPcManagementGroupAssignmentTarget{
        CloudPcManagementAssignmentTarget: *NewCloudPcManagementAssignmentTarget(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcManagementGroupAssignmentTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcManagementGroupAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcManagementGroupAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcManagementGroupAssignmentTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcManagementGroupAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcManagementAssignmentTarget.GetFieldDeserializers()
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
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The id of the assignment's target group
func (m *CloudPcManagementGroupAssignmentTarget) GetGroupId()(*string) {
    return m.groupId
}
// GetServicePlanId gets the servicePlanId property value. The servicePlanId property
func (m *CloudPcManagementGroupAssignmentTarget) GetServicePlanId()(*string) {
    return m.servicePlanId
}
// Serialize serializes information the current object
func (m *CloudPcManagementGroupAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupId sets the groupId property value. The id of the assignment's target group
func (m *CloudPcManagementGroupAssignmentTarget) SetGroupId(value *string)() {
    m.groupId = value
}
// SetServicePlanId sets the servicePlanId property value. The servicePlanId property
func (m *CloudPcManagementGroupAssignmentTarget) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
