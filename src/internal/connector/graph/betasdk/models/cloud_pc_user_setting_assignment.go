package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcUserSettingAssignment provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcUserSettingAssignment struct {
    Entity
    // The date and time this assignment was created. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The assignment target for the user setting. Currently, the only target supported for this user setting is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
    target CloudPcManagementAssignmentTargetable
}
// NewCloudPcUserSettingAssignment instantiates a new cloudPcUserSettingAssignment and sets the default values.
func NewCloudPcUserSettingAssignment()(*CloudPcUserSettingAssignment) {
    m := &CloudPcUserSettingAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcUserSettingAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcUserSettingAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcUserSettingAssignment(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time this assignment was created. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
func (m *CloudPcUserSettingAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcUserSettingAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(CloudPcManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The assignment target for the user setting. Currently, the only target supported for this user setting is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcUserSettingAssignment) GetTarget()(CloudPcManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *CloudPcUserSettingAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time this assignment was created. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
func (m *CloudPcUserSettingAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetTarget sets the target property value. The assignment target for the user setting. Currently, the only target supported for this user setting is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcUserSettingAssignment) SetTarget(value CloudPcManagementAssignmentTargetable)() {
    m.target = value
}
