package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupMembershipGovernanceCriteria 
type GroupMembershipGovernanceCriteria struct {
    GovernanceCriteria
    // The groupId property
    groupId *string
}
// NewGroupMembershipGovernanceCriteria instantiates a new GroupMembershipGovernanceCriteria and sets the default values.
func NewGroupMembershipGovernanceCriteria()(*GroupMembershipGovernanceCriteria) {
    m := &GroupMembershipGovernanceCriteria{
        GovernanceCriteria: *NewGovernanceCriteria(),
    }
    odataTypeValue := "#microsoft.graph.groupMembershipGovernanceCriteria";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupMembershipGovernanceCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupMembershipGovernanceCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupMembershipGovernanceCriteria(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupMembershipGovernanceCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GovernanceCriteria.GetFieldDeserializers()
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
// GetGroupId gets the groupId property value. The groupId property
func (m *GroupMembershipGovernanceCriteria) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *GroupMembershipGovernanceCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GovernanceCriteria.Serialize(writer)
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
// SetGroupId sets the groupId property value. The groupId property
func (m *GroupMembershipGovernanceCriteria) SetGroupId(value *string)() {
    m.groupId = value
}
