package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupEligibilityScheduleRequest 
type PrivilegedAccessGroupEligibilityScheduleRequest struct {
    PrivilegedAccessScheduleRequest
    // The accessId property
    accessId *PrivilegedAccessGroupRelationships
    // The group property
    group Groupable
    // The groupId property
    groupId *string
    // The principal property
    principal DirectoryObjectable
    // The principalId property
    principalId *string
    // The targetSchedule property
    targetSchedule PrivilegedAccessGroupEligibilityScheduleable
    // The targetScheduleId property
    targetScheduleId *string
}
// NewPrivilegedAccessGroupEligibilityScheduleRequest instantiates a new PrivilegedAccessGroupEligibilityScheduleRequest and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequest()(*PrivilegedAccessGroupEligibilityScheduleRequest) {
    m := &PrivilegedAccessGroupEligibilityScheduleRequest{
        PrivilegedAccessScheduleRequest: *NewPrivilegedAccessScheduleRequest(),
    }
    odataTypeValue := "#microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequest(), nil
}
// GetAccessId gets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetAccessId()(*PrivilegedAccessGroupRelationships) {
    return m.accessId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrivilegedAccessScheduleRequest.GetFieldDeserializers()
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
    res["targetSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchedule(val.(PrivilegedAccessGroupEligibilityScheduleable))
        }
        return nil
    }
    res["targetScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetScheduleId(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetGroup()(Groupable) {
    return m.group
}
// GetGroupId gets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetGroupId()(*string) {
    return m.groupId
}
// GetPrincipal gets the principal property value. The principal property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetPrincipalId()(*string) {
    return m.principalId
}
// GetTargetSchedule gets the targetSchedule property value. The targetSchedule property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable) {
    return m.targetSchedule
}
// GetTargetScheduleId gets the targetScheduleId property value. The targetScheduleId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) GetTargetScheduleId()(*string) {
    return m.targetScheduleId
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrivilegedAccessScheduleRequest.Serialize(writer)
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
    {
        err = writer.WriteObjectValue("targetSchedule", m.GetTargetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetScheduleId", m.GetTargetScheduleId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessId sets the accessId property value. The accessId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetAccessId(value *PrivilegedAccessGroupRelationships)() {
    m.accessId = value
}
// SetGroup sets the group property value. The group property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetGroup(value Groupable)() {
    m.group = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetGroupId(value *string)() {
    m.groupId = value
}
// SetPrincipal sets the principal property value. The principal property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetTargetSchedule sets the targetSchedule property value. The targetSchedule property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)() {
    m.targetSchedule = value
}
// SetTargetScheduleId sets the targetScheduleId property value. The targetScheduleId property
func (m *PrivilegedAccessGroupEligibilityScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
