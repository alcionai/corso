package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewPolicy 
type AccessReviewPolicy struct {
    Entity
    // Description for this policy. Read-only.
    description *string
    // Display name for this policy. Read-only.
    displayName *string
    // If true, group owners can create and manage access reviews on groups they own.
    isGroupOwnerManagementEnabled *bool
}
// NewAccessReviewPolicy instantiates a new accessReviewPolicy and sets the default values.
func NewAccessReviewPolicy()(*AccessReviewPolicy) {
    m := &AccessReviewPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewPolicy(), nil
}
// GetDescription gets the description property value. Description for this policy. Read-only.
func (m *AccessReviewPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for this policy. Read-only.
func (m *AccessReviewPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isGroupOwnerManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGroupOwnerManagementEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsGroupOwnerManagementEnabled gets the isGroupOwnerManagementEnabled property value. If true, group owners can create and manage access reviews on groups they own.
func (m *AccessReviewPolicy) GetIsGroupOwnerManagementEnabled()(*bool) {
    return m.isGroupOwnerManagementEnabled
}
// Serialize serializes information the current object
func (m *AccessReviewPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isGroupOwnerManagementEnabled", m.GetIsGroupOwnerManagementEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description for this policy. Read-only.
func (m *AccessReviewPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for this policy. Read-only.
func (m *AccessReviewPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsGroupOwnerManagementEnabled sets the isGroupOwnerManagementEnabled property value. If true, group owners can create and manage access reviews on groups they own.
func (m *AccessReviewPolicy) SetIsGroupOwnerManagementEnabled(value *bool)() {
    m.isGroupOwnerManagementEnabled = value
}
