package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectorySetting provides operations to manage the collection of site entities.
type DirectorySetting struct {
    Entity
    // Display name of this group of settings, which comes from the associated template. Read-only.
    displayName *string
    // Unique identifier for the template used to create this group of settings. Read-only.
    templateId *string
    // Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced directorySettingTemplates object.
    values []SettingValueable
}
// NewDirectorySetting instantiates a new directorySetting and sets the default values.
func NewDirectorySetting()(*DirectorySetting) {
    m := &DirectorySetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectorySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectorySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectorySetting(), nil
}
// GetDisplayName gets the displayName property value. Display name of this group of settings, which comes from the associated template. Read-only.
func (m *DirectorySetting) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(SettingValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. Unique identifier for the template used to create this group of settings. Read-only.
func (m *DirectorySetting) GetTemplateId()(*string) {
    return m.templateId
}
// GetValues gets the values property value. Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced directorySettingTemplates object.
func (m *DirectorySetting) GetValues()([]SettingValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *DirectorySetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name of this group of settings, which comes from the associated template. Read-only.
func (m *DirectorySetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTemplateId sets the templateId property value. Unique identifier for the template used to create this group of settings. Read-only.
func (m *DirectorySetting) SetTemplateId(value *string)() {
    m.templateId = value
}
// SetValues sets the values property value. Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced directorySettingTemplates object.
func (m *DirectorySetting) SetValues(value []SettingValueable)() {
    m.values = value
}
