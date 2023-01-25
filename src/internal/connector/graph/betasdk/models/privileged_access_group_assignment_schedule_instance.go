package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupAssignmentScheduleInstance 
type PrivilegedAccessGroupAssignmentScheduleInstance struct {
    PrivilegedAccessScheduleInstance
    // The accessId property
    accessId *PrivilegedAccessGroupRelationships
    // The activatedUsing property
    activatedUsing PrivilegedAccessGroupEligibilityScheduleInstanceable
    // The assignmentScheduleId property
    assignmentScheduleId *string
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
// NewPrivilegedAccessGroupAssignmentScheduleInstance instantiates a new PrivilegedAccessGroupAssignmentScheduleInstance and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleInstance()(*PrivilegedAccessGroupAssignmentScheduleInstance) {
    m := &PrivilegedAccessGroupAssignmentScheduleInstance{
        PrivilegedAccessScheduleInstance: *NewPrivilegedAccessScheduleInstance(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstance(), nil
}
// GetAccessId gets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetActivatedUsing gets the activatedUsing property value. The activatedUsing property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetActivatedUsing()(PrivilegedAccessGroupEligibilityScheduleInstanceable) {
    return m.activatedUsing
}
// GetAssignmentScheduleId gets the assignmentScheduleId property value. The assignmentScheduleId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetAssignmentScheduleId()(*string) {
    return m.assignmentScheduleId
}
// GetAssignmentType gets the assignmentType property value. The assignmentType property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetAssignmentType()(*PrivilegedAccessGroupAssignmentType) {
    return m.assignmentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessScheduleInstance.GetFieldDeserializers()
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
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(PrivilegedAccessGroupEligibilityScheduleInstanceable))
        }
        return nil
    }
    res["assignmentScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentScheduleId(val)
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
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetGroupId()(*string) {
    return m.groupId
}
// GetMemberType gets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetMemberType()(*PrivilegedAccessGroupMemberType) {
    return m.memberType
}
// GetPrincipal gets the principal property value. The principal property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) GetPrincipalId()(*string) {
    return m.principalId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessScheduleInstance.Serialize(writer)
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
    {
        err = writer.WriteStringValue("assignmentScheduleId", m.GetAssignmentScheduleId())
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
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetActivatedUsing sets the activatedUsing property value. The activatedUsing property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetActivatedUsing(value PrivilegedAccessGroupEligibilityScheduleInstanceable)() {
    m.activatedUsing = value
}
// SetAssignmentScheduleId sets the assignmentScheduleId property value. The assignmentScheduleId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetAssignmentScheduleId(value *string)() {
    m.assignmentScheduleId = value
}
// SetAssignmentType sets the assignmentType property value. The assignmentType property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetAssignmentType(value *PrivilegedAccessGroupAssignmentType)() {
    m.assignmentType = value
}
// SetGroup sets the group property value. The group property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetGroupId(value *string)() {
    m.groupId = value
}
// SetMemberType sets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetMemberType(value *PrivilegedAccessGroupMemberType)() {
    m.memberType = value
}
// SetPrincipal sets the principal property value. The principal property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupAssignmentScheduleInstance) SetPrincipalId(value *string)() {
    m.principalId = value
}
