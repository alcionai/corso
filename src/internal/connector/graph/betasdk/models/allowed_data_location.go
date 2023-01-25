package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AllowedDataLocation 
type AllowedDataLocation struct {
    Entity
    // The appId property
    appId *string
    // The domain property
    domain *string
    // The isDefault property
    isDefault *bool
    // The location property
    location *string
}
// NewAllowedDataLocation instantiates a new AllowedDataLocation and sets the default values.
func NewAllowedDataLocation()(*AllowedDataLocation) {
    m := &AllowedDataLocation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAllowedDataLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAllowedDataLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllowedDataLocation(), nil
}
// GetAppId gets the appId property value. The appId property
func (m *AllowedDataLocation) GetAppId()(*string) {
    return m.appId
}
// GetDomain gets the domain property value. The domain property
func (m *AllowedDataLocation) GetDomain()(*string) {
    return m.domain
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AllowedDataLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *AllowedDataLocation) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetLocation gets the location property value. The location property
func (m *AllowedDataLocation) GetLocation()(*string) {
    return m.location
}
// Serialize serializes information the current object
func (m *AllowedDataLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The appId property
func (m *AllowedDataLocation) SetAppId(value *string)() {
    m.appId = value
}
// SetDomain sets the domain property value. The domain property
func (m *AllowedDataLocation) SetDomain(value *string)() {
    m.domain = value
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *AllowedDataLocation) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetLocation sets the location property value. The location property
func (m *AllowedDataLocation) SetLocation(value *string)() {
    m.location = value
}
