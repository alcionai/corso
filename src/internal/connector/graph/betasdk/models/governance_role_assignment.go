package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleAssignment provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type GovernanceRoleAssignment struct {
    Entity
    // The state of the assignment. The value can be Eligible for eligible assignment or Active if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
    assignmentState *string
    // For a non-permanent role assignment, this is the time when the role assignment will be expired. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The external ID the resource that is used to identify the role assignment in the provider.
    externalId *string
    // Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
    linkedEligibleRoleAssignment GovernanceRoleAssignmentable
    // If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that eligible assignment; Otherwise, the value is null.
    linkedEligibleRoleAssignmentId *string
    // The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope), Group (if the role assignment is not inherited, but comes from the membership of a group assignment), or User (if the role assignment is neither inherited nor from a group assignment).
    memberType *string
    // Read-only. The resource associated with the role assignment.
    resource GovernanceResourceable
    // Required. The ID of the resource which the role assignment is associated with.
    resourceId *string
    // Read-only. The role definition associated with the role assignment.
    roleDefinition GovernanceRoleDefinitionable
    // Required. The ID of the role definition which the role assignment is associated with.
    roleDefinitionId *string
    // The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *string
    // Read-only. The subject associated with the role assignment.
    subject GovernanceSubjectable
    // Required. The ID of the subject which the role assignment is associated with.
    subjectId *string
}
// NewGovernanceRoleAssignment instantiates a new governanceRoleAssignment and sets the default values.
func NewGovernanceRoleAssignment()(*GovernanceRoleAssignment) {
    m := &GovernanceRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRoleAssignment(), nil
}
// GetAssignmentState gets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment or Active if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
func (m *GovernanceRoleAssignment) GetAssignmentState()(*string) {
    return m.assignmentState
}
// GetEndDateTime gets the endDateTime property value. For a non-permanent role assignment, this is the time when the role assignment will be expired. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignment) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExternalId gets the externalId property value. The external ID the resource that is used to identify the role assignment in the provider.
func (m *GovernanceRoleAssignment) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentState(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["linkedEligibleRoleAssignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkedEligibleRoleAssignment(val.(GovernanceRoleAssignmentable))
        }
        return nil
    }
    res["linkedEligibleRoleAssignmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkedEligibleRoleAssignmentId(val)
        }
        return nil
    }
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(GovernanceResourceable))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(GovernanceRoleDefinitionable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(GovernanceSubjectable))
        }
        return nil
    }
    res["subjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectId(val)
        }
        return nil
    }
    return res
}
// GetLinkedEligibleRoleAssignment gets the linkedEligibleRoleAssignment property value. Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
func (m *GovernanceRoleAssignment) GetLinkedEligibleRoleAssignment()(GovernanceRoleAssignmentable) {
    return m.linkedEligibleRoleAssignment
}
// GetLinkedEligibleRoleAssignmentId gets the linkedEligibleRoleAssignmentId property value. If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that eligible assignment; Otherwise, the value is null.
func (m *GovernanceRoleAssignment) GetLinkedEligibleRoleAssignmentId()(*string) {
    return m.linkedEligibleRoleAssignmentId
}
// GetMemberType gets the memberType property value. The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope), Group (if the role assignment is not inherited, but comes from the membership of a group assignment), or User (if the role assignment is neither inherited nor from a group assignment).
func (m *GovernanceRoleAssignment) GetMemberType()(*string) {
    return m.memberType
}
// GetResource gets the resource property value. Read-only. The resource associated with the role assignment.
func (m *GovernanceRoleAssignment) GetResource()(GovernanceResourceable) {
    return m.resource
}
// GetResourceId gets the resourceId property value. Required. The ID of the resource which the role assignment is associated with.
func (m *GovernanceRoleAssignment) GetResourceId()(*string) {
    return m.resourceId
}
// GetRoleDefinition gets the roleDefinition property value. Read-only. The role definition associated with the role assignment.
func (m *GovernanceRoleAssignment) GetRoleDefinition()(GovernanceRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Required. The ID of the role definition which the role assignment is associated with.
func (m *GovernanceRoleAssignment) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// GetStartDateTime gets the startDateTime property value. The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignment) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
func (m *GovernanceRoleAssignment) GetStatus()(*string) {
    return m.status
}
// GetSubject gets the subject property value. Read-only. The subject associated with the role assignment.
func (m *GovernanceRoleAssignment) GetSubject()(GovernanceSubjectable) {
    return m.subject
}
// GetSubjectId gets the subjectId property value. Required. The ID of the subject which the role assignment is associated with.
func (m *GovernanceRoleAssignment) GetSubjectId()(*string) {
    return m.subjectId
}
// Serialize serializes information the current object
func (m *GovernanceRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignmentState", m.GetAssignmentState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("linkedEligibleRoleAssignment", m.GetLinkedEligibleRoleAssignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("linkedEligibleRoleAssignmentId", m.GetLinkedEligibleRoleAssignmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectId", m.GetSubjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentState sets the assignmentState property value. The state of the assignment. The value can be Eligible for eligible assignment or Active if it is directly assigned Active by administrators, or activated on an eligible assignment by the users.
func (m *GovernanceRoleAssignment) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
// SetEndDateTime sets the endDateTime property value. For a non-permanent role assignment, this is the time when the role assignment will be expired. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignment) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExternalId sets the externalId property value. The external ID the resource that is used to identify the role assignment in the provider.
func (m *GovernanceRoleAssignment) SetExternalId(value *string)() {
    m.externalId = value
}
// SetLinkedEligibleRoleAssignment sets the linkedEligibleRoleAssignment property value. Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the object of that eligible assignment; Otherwise, the value is null.
func (m *GovernanceRoleAssignment) SetLinkedEligibleRoleAssignment(value GovernanceRoleAssignmentable)() {
    m.linkedEligibleRoleAssignment = value
}
// SetLinkedEligibleRoleAssignmentId sets the linkedEligibleRoleAssignmentId property value. If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that eligible assignment; Otherwise, the value is null.
func (m *GovernanceRoleAssignment) SetLinkedEligibleRoleAssignmentId(value *string)() {
    m.linkedEligibleRoleAssignmentId = value
}
// SetMemberType sets the memberType property value. The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope), Group (if the role assignment is not inherited, but comes from the membership of a group assignment), or User (if the role assignment is neither inherited nor from a group assignment).
func (m *GovernanceRoleAssignment) SetMemberType(value *string)() {
    m.memberType = value
}
// SetResource sets the resource property value. Read-only. The resource associated with the role assignment.
func (m *GovernanceRoleAssignment) SetResource(value GovernanceResourceable)() {
    m.resource = value
}
// SetResourceId sets the resourceId property value. Required. The ID of the resource which the role assignment is associated with.
func (m *GovernanceRoleAssignment) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetRoleDefinition sets the roleDefinition property value. Read-only. The role definition associated with the role assignment.
func (m *GovernanceRoleAssignment) SetRoleDefinition(value GovernanceRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Required. The ID of the role definition which the role assignment is associated with.
func (m *GovernanceRoleAssignment) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// SetStartDateTime sets the startDateTime property value. The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleAssignment) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *GovernanceRoleAssignment) SetStatus(value *string)() {
    m.status = value
}
// SetSubject sets the subject property value. Read-only. The subject associated with the role assignment.
func (m *GovernanceRoleAssignment) SetSubject(value GovernanceSubjectable)() {
    m.subject = value
}
// SetSubjectId sets the subjectId property value. Required. The ID of the subject which the role assignment is associated with.
func (m *GovernanceRoleAssignment) SetSubjectId(value *string)() {
    m.subjectId = value
}
