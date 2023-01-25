package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectorySettingTemplate 
type DirectorySettingTemplate struct {
    DirectoryObject
    // Description of the template. Read-only.
    description *string
    // Display name of the template. Read-only.
    displayName *string
    // Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
    values []SettingTemplateValueable
}
// NewDirectorySettingTemplate instantiates a new DirectorySettingTemplate and sets the default values.
func NewDirectorySettingTemplate()(*DirectorySettingTemplate) {
    m := &DirectorySettingTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.directorySettingTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDirectorySettingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectorySettingTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectorySettingTemplate(), nil
}
// GetDescription gets the description property value. Description of the template. Read-only.
func (m *DirectorySettingTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the template. Read-only.
func (m *DirectorySettingTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySettingTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingTemplateValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingTemplateValueable, len(val))
            for i, v := range val {
                res[i] = v.(SettingTemplateValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
func (m *DirectorySettingTemplate) GetValues()([]SettingTemplateValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *DirectorySettingTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
// SetDescription sets the description property value. Description of the template. Read-only.
func (m *DirectorySettingTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the template. Read-only.
func (m *DirectorySettingTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetValues sets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
func (m *DirectorySettingTemplate) SetValues(value []SettingTemplateValueable)() {
    m.values = value
}
