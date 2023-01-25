package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MembershipOutlierInsight 
type MembershipOutlierInsight struct {
    GovernanceInsight
    // Navigation link to the container directory object. For example, to a group.
    container DirectoryObjectable
    // Indicates the identifier of the container, for example, a group ID.
    containerId *string
    // Navigation link to a member object. For example, to a user.
    member DirectoryObjectable
    // Indicates the identifier of the user.
    memberId *string
    // The outlierContainerType property
    outlierContainerType *OutlierContainerType
    // The outlierMemberType property
    outlierMemberType *OutlierMemberType
}
// NewMembershipOutlierInsight instantiates a new MembershipOutlierInsight and sets the default values.
func NewMembershipOutlierInsight()(*MembershipOutlierInsight) {
    m := &MembershipOutlierInsight{
        GovernanceInsight: *NewGovernanceInsight(),
    }
    odataTypeValue := "#microsoft.graph.membershipOutlierInsight";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMembershipOutlierInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMembershipOutlierInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMembershipOutlierInsight(), nil
}
// GetContainer gets the container property value. Navigation link to the container directory object. For example, to a group.
func (m *MembershipOutlierInsight) GetContainer()(DirectoryObjectable) {
    return m.container
}
// GetContainerId gets the containerId property value. Indicates the identifier of the container, for example, a group ID.
func (m *MembershipOutlierInsight) GetContainerId()(*string) {
    return m.containerId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MembershipOutlierInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GovernanceInsight.GetFieldDeserializers()
    res["container"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainer(val.(DirectoryObjectable))
        }
        return nil
    }
    res["containerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerId(val)
        }
        return nil
    }
    res["member"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMember(val.(DirectoryObjectable))
        }
        return nil
    }
    res["memberId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberId(val)
        }
        return nil
    }
    res["outlierContainerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOutlierContainerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlierContainerType(val.(*OutlierContainerType))
        }
        return nil
    }
    res["outlierMemberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOutlierMemberType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlierMemberType(val.(*OutlierMemberType))
        }
        return nil
    }
    return res
}
// GetMember gets the member property value. Navigation link to a member object. For example, to a user.
func (m *MembershipOutlierInsight) GetMember()(DirectoryObjectable) {
    return m.member
}
// GetMemberId gets the memberId property value. Indicates the identifier of the user.
func (m *MembershipOutlierInsight) GetMemberId()(*string) {
    return m.memberId
}
// GetOutlierContainerType gets the outlierContainerType property value. The outlierContainerType property
func (m *MembershipOutlierInsight) GetOutlierContainerType()(*OutlierContainerType) {
    return m.outlierContainerType
}
// GetOutlierMemberType gets the outlierMemberType property value. The outlierMemberType property
func (m *MembershipOutlierInsight) GetOutlierMemberType()(*OutlierMemberType) {
    return m.outlierMemberType
}
// Serialize serializes information the current object
func (m *MembershipOutlierInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GovernanceInsight.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("container", m.GetContainer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("containerId", m.GetContainerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("member", m.GetMember())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("memberId", m.GetMemberId())
        if err != nil {
            return err
        }
    }
    if m.GetOutlierContainerType() != nil {
        cast := (*m.GetOutlierContainerType()).String()
        err = writer.WriteStringValue("outlierContainerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOutlierMemberType() != nil {
        cast := (*m.GetOutlierMemberType()).String()
        err = writer.WriteStringValue("outlierMemberType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainer sets the container property value. Navigation link to the container directory object. For example, to a group.
func (m *MembershipOutlierInsight) SetContainer(value DirectoryObjectable)() {
    m.container = value
}
// SetContainerId sets the containerId property value. Indicates the identifier of the container, for example, a group ID.
func (m *MembershipOutlierInsight) SetContainerId(value *string)() {
    m.containerId = value
}
// SetMember sets the member property value. Navigation link to a member object. For example, to a user.
func (m *MembershipOutlierInsight) SetMember(value DirectoryObjectable)() {
    m.member = value
}
// SetMemberId sets the memberId property value. Indicates the identifier of the user.
func (m *MembershipOutlierInsight) SetMemberId(value *string)() {
    m.memberId = value
}
// SetOutlierContainerType sets the outlierContainerType property value. The outlierContainerType property
func (m *MembershipOutlierInsight) SetOutlierContainerType(value *OutlierContainerType)() {
    m.outlierContainerType = value
}
// SetOutlierMemberType sets the outlierMemberType property value. The outlierMemberType property
func (m *MembershipOutlierInsight) SetOutlierMemberType(value *OutlierMemberType)() {
    m.outlierMemberType = value
}
