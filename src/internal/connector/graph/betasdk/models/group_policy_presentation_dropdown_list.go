package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDropdownList 
type GroupPolicyPresentationDropdownList struct {
    GroupPolicyUploadedPresentation
    // Localized string value identifying the default choice of the list of items.
    defaultItem GroupPolicyPresentationDropdownListItemable
    // Represents a set of localized display names and their associated values.
    items []GroupPolicyPresentationDropdownListItemable
    // Requirement to enter a value in the parameter box. The default value is false.
    required *bool
}
// NewGroupPolicyPresentationDropdownList instantiates a new GroupPolicyPresentationDropdownList and sets the default values.
func NewGroupPolicyPresentationDropdownList()(*GroupPolicyPresentationDropdownList) {
    m := &GroupPolicyPresentationDropdownList{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationDropdownList";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyPresentationDropdownListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationDropdownListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationDropdownList(), nil
}
// GetDefaultItem gets the defaultItem property value. Localized string value identifying the default choice of the list of items.
func (m *GroupPolicyPresentationDropdownList) GetDefaultItem()(GroupPolicyPresentationDropdownListItemable) {
    return m.defaultItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationDropdownList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["defaultItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyPresentationDropdownListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultItem(val.(GroupPolicyPresentationDropdownListItemable))
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyPresentationDropdownListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationDropdownListItemable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyPresentationDropdownListItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. Represents a set of localized display names and their associated values.
func (m *GroupPolicyPresentationDropdownList) GetItems()([]GroupPolicyPresentationDropdownListItemable) {
    return m.items
}
// GetRequired gets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDropdownList) GetRequired()(*bool) {
    return m.required
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationDropdownList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultItem", m.GetDefaultItem())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultItem sets the defaultItem property value. Localized string value identifying the default choice of the list of items.
func (m *GroupPolicyPresentationDropdownList) SetDefaultItem(value GroupPolicyPresentationDropdownListItemable)() {
    m.defaultItem = value
}
// SetItems sets the items property value. Represents a set of localized display names and their associated values.
func (m *GroupPolicyPresentationDropdownList) SetItems(value []GroupPolicyPresentationDropdownListItemable)() {
    m.items = value
}
// SetRequired sets the required property value. Requirement to enter a value in the parameter box. The default value is false.
func (m *GroupPolicyPresentationDropdownList) SetRequired(value *bool)() {
    m.required = value
}
