package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnsupportedGroupPolicyExtension unsupported Group Policy Extension.
type UnsupportedGroupPolicyExtension struct {
    Entity
    // ExtensionType of the unsupported extension.
    extensionType *string
    // Namespace Url of the unsupported extension.
    namespaceUrl *string
    // Node name of the unsupported extension.
    nodeName *string
    // Scope of the group policy setting.
    settingScope *GroupPolicySettingScope
}
// NewUnsupportedGroupPolicyExtension instantiates a new unsupportedGroupPolicyExtension and sets the default values.
func NewUnsupportedGroupPolicyExtension()(*UnsupportedGroupPolicyExtension) {
    m := &UnsupportedGroupPolicyExtension{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnsupportedGroupPolicyExtension(), nil
}
// GetExtensionType gets the extensionType property value. ExtensionType of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetExtensionType()(*string) {
    return m.extensionType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnsupportedGroupPolicyExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["extensionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionType(val)
        }
        return nil
    }
    res["namespaceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespaceUrl(val)
        }
        return nil
    }
    res["nodeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeName(val)
        }
        return nil
    }
    res["settingScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingScope(val.(*GroupPolicySettingScope))
        }
        return nil
    }
    return res
}
// GetNamespaceUrl gets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNamespaceUrl()(*string) {
    return m.namespaceUrl
}
// GetNodeName gets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNodeName()(*string) {
    return m.nodeName
}
// GetSettingScope gets the settingScope property value. Scope of the group policy setting.
func (m *UnsupportedGroupPolicyExtension) GetSettingScope()(*GroupPolicySettingScope) {
    return m.settingScope
}
// Serialize serializes information the current object
func (m *UnsupportedGroupPolicyExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("extensionType", m.GetExtensionType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("namespaceUrl", m.GetNamespaceUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nodeName", m.GetNodeName())
        if err != nil {
            return err
        }
    }
    if m.GetSettingScope() != nil {
        cast := (*m.GetSettingScope()).String()
        err = writer.WriteStringValue("settingScope", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExtensionType sets the extensionType property value. ExtensionType of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetExtensionType(value *string)() {
    m.extensionType = value
}
// SetNamespaceUrl sets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetNamespaceUrl(value *string)() {
    m.namespaceUrl = value
}
// SetNodeName sets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetNodeName(value *string)() {
    m.nodeName = value
}
// SetSettingScope sets the settingScope property value. Scope of the group policy setting.
func (m *UnsupportedGroupPolicyExtension) SetSettingScope(value *GroupPolicySettingScope)() {
    m.settingScope = value
}
