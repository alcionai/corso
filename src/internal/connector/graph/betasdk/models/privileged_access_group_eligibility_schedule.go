package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupEligibilitySchedule 
type PrivilegedAccessGroupEligibilitySchedule struct {
    PrivilegedAccessSchedule
    // The accessId property
    accessId *PrivilegedAccessGroupRelationships
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
// NewPrivilegedAccessGroupEligibilitySchedule instantiates a new privilegedAccessGroupEligibilitySchedule and sets the default values.
func NewPrivilegedAccessGroupEligibilitySchedule()(*PrivilegedAccessGroupEligibilitySchedule) {
    m := &PrivilegedAccessGroupEligibilitySchedule{
        PrivilegedAccessSchedule: *NewPrivilegedAccessSchedule(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupEligibilitySchedule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilitySchedule(), nil
}
// GetAccessId gets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupEligibilitySchedule) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroupEligibilitySchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *PrivilegedAccessGroupEligibilitySchedule) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupEligibilitySchedule) GetGroupId()(*string) {
    return m.groupId
}
// GetMemberType gets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupEligibilitySchedule) GetMemberType()(*PrivilegedAccessGroupMemberType) {
    return m.memberType
}
// GetPrincipal gets the principal property value. The principal property
func (m *PrivilegedAccessGroupEligibilitySchedule) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupEligibilitySchedule) GetPrincipalId()(*string) {
    return m.principalId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupEligibilitySchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrivilegedAccessGroupEligibilitySchedule) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetGroup sets the group property value. The group property
func (m *PrivilegedAccessGroupEligibilitySchedule) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupEligibilitySchedule) SetGroupId(value *string)() {
    m.groupId = value
}
// SetMemberType sets the memberType property value. The memberType property
func (m *PrivilegedAccessGroupEligibilitySchedule) SetMemberType(value *PrivilegedAccessGroupMemberType)() {
    m.memberType = value
}
// SetPrincipal sets the principal property value. The principal property
func (m *PrivilegedAccessGroupEligibilitySchedule) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupEligibilitySchedule) SetPrincipalId(value *string)() {
    m.principalId = value
}
