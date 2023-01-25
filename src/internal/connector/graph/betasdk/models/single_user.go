package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SingleUser 
type SingleUser struct {
    UserSet
    // The name of the user in Azure AD. Read only.
    description *string
    // The ID of the user in Azure AD.
    id *string
}
// NewSingleUser instantiates a new SingleUser and sets the default values.
func NewSingleUser()(*SingleUser) {
    m := &SingleUser{
        UserSet: *NewUserSet(),
    }
    odataTypeValue := "#microsoft.graph.singleUser";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSingleUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSingleUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSingleUser(), nil
}
// GetDescription gets the description property value. The name of the user in Azure AD. Read only.
func (m *SingleUser) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SingleUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserSet.GetFieldDeserializers()
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the user in Azure AD.
func (m *SingleUser) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *SingleUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserSet.Serialize(writer)
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
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The name of the user in Azure AD. Read only.
func (m *SingleUser) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The ID of the user in Azure AD.
func (m *SingleUser) SetId(value *string)() {
    m.id = value
}
