package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupAssignmentSchedule 
type PrivilegedAccessGroupAssignmentSchedule struct {
    PrivilegedAccessSchedule
    // The accessId property
    accessId *PrivilegedAccessGroupRelationships
    // The activatedUsing property
    activatedUsing PrivilegedAccessGroupEligibilityScheduleable
    // The assignmentType property
    assignmentType *PrivilegedAccessGroupAssignmentType
    // The group property
    group Groupable
    // The groupId property
    groupId *string
    // The memberType property
    memberType *PrivilegedAccessGroupMemberType
    // The principal property
    principal DirectoryObjectable
    // The principalId property
    principalId *string
}
// NewPrivilegedAccessGroupAssignmentSchedule instantiates a new PrivilegedAccessGroupAssignmentSchedule and sets the default values.
func NewPrivilegedAccessGroupAssignmentSchedule()(*PrivilegedAccessGroupAssignmentSchedule) {
    m := &PrivilegedAccessGroupAssignmentSchedule{
        PrivilegedAccessSchedule: *NewPrivilegedAccessSchedule(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupAssignmentSchedule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrivilegedAccessGroupAssignmentScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupAssignmentScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupAssignmentSchedule(), nil
}
// GetAccessId gets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetActivatedUsing gets the activatedUsing property value. The activatedUsing property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetActivatedUsing()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.activatedUsing
}
// GetAssignmentType gets the assignmentType property value. The assignmentType property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetAssignmentType()(*PrivilegedAccessGroupAssignmentType) {
    return m.assignmentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroupAssignmentSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessSchedule.GetFieldDeserializers()
    res["accessId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupRelationships)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessId(val.(*PrivilegedAccessGroupRelationships))
        }
        return nil
    }
    res["activatedUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(PrivilegedAccessGroupEligibilityScheduleable))
        }
        return nil
    }
    res["assignmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupAssignmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentType(val.(*PrivilegedAccessGroupAssignmentType))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
        }
        return nil
    }
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
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrivilegedAccessGroupMemberType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val.(*PrivilegedAccessGroupMemberType))
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(DirectoryObjectable))
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetGroupId()(*string) {
    return m.groupId
}
// GetMemberType gets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetMemberType()(*PrivilegedAccessGroupMemberType) {
    return m.memberType
}
// GetPrincipal gets the principal property value. The principal property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupAssignmentSchedule) GetPrincipalId()(*string) {
    return m.principalId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupAssignmentSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessSchedule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessId() != nil {
        cast := (*m.GetAccessId()).String()
        err = writer.WriteStringValue("accessId", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentType() != nil {
        cast := (*m.GetAssignmentType()).String()
        err = writer.WriteStringValue("assignmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    if m.GetMemberType() != nil {
        cast := (*m.GetMemberType()).String()
        err = writer.WriteStringValue("memberType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessId sets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetActivatedUsing sets the activatedUsing property value. The activatedUsing property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetActivatedUsing(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.activatedUsing = value
}
// SetAssignmentType sets the assignmentType property value. The assignmentType property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetAssignmentType(value *PrivilegedAccessGroupAssignmentType)() {
    m.assignmentType = value
}
// SetGroup sets the group property value. The group property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetGroupId(value *string)() {
    m.groupId = value
}
// SetMemberType sets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetMemberType(value *PrivilegedAccessGroupMemberType)() {
    m.memberType = value
}
// SetPrincipal sets the principal property value. The principal property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupAssignmentSchedule) SetPrincipalId(value *string)() {
    m.principalId = value
}
