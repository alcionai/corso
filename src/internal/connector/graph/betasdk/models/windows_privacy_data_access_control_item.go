package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPrivacyDataAccessControlItem specify access control level per privacy data category
type WindowsPrivacyDataAccessControlItem struct {
    Entity
    // Determine the access level to specific Windows privacy data category.
    accessLevel *WindowsPrivacyDataAccessLevel
    // The Package Family Name of a Windows app. When set, the access level applies to the specified application.
    appDisplayName *string
    // The Package Family Name of a Windows app. When set, the access level applies to the specified application.
    appPackageFamilyName *string
    // Windows privacy data category specifier for privacy data access.
    dataCategory *WindowsPrivacyDataCategory
}
// NewWindowsPrivacyDataAccessControlItem instantiates a new windowsPrivacyDataAccessControlItem and sets the default values.
func NewWindowsPrivacyDataAccessControlItem()(*WindowsPrivacyDataAccessControlItem) {
    m := &WindowsPrivacyDataAccessControlItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPrivacyDataAccessControlItem(), nil
}
// GetAccessLevel gets the accessLevel property value. Determine the access level to specific Windows privacy data category.
func (m *WindowsPrivacyDataAccessControlItem) GetAccessLevel()(*WindowsPrivacyDataAccessLevel) {
    return m.accessLevel
}
// GetAppDisplayName gets the appDisplayName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppPackageFamilyName gets the appPackageFamilyName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) GetAppPackageFamilyName()(*string) {
    return m.appPackageFamilyName
}
// GetDataCategory gets the dataCategory property value. Windows privacy data category specifier for privacy data access.
func (m *WindowsPrivacyDataAccessControlItem) GetDataCategory()(*WindowsPrivacyDataCategory) {
    return m.dataCategory
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPrivacyDataAccessControlItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsPrivacyDataAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*WindowsPrivacyDataAccessLevel))
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appPackageFamilyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPackageFamilyName(val)
        }
        return nil
    }
    res["dataCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsPrivacyDataCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataCategory(val.(*WindowsPrivacyDataCategory))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsPrivacyDataAccessControlItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
        err = writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPackageFamilyName", m.GetAppPackageFamilyName())
        if err != nil {
            return err
        }
    }
    if m.GetDataCategory() != nil {
        cast := (*m.GetDataCategory()).String()
        err = writer.WriteStringValue("dataCategory", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessLevel sets the accessLevel property value. Determine the access level to specific Windows privacy data category.
func (m *WindowsPrivacyDataAccessControlItem) SetAccessLevel(value *WindowsPrivacyDataAccessLevel)() {
    m.accessLevel = value
}
// SetAppDisplayName sets the appDisplayName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppPackageFamilyName sets the appPackageFamilyName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) SetAppPackageFamilyName(value *string)() {
    m.appPackageFamilyName = value
}
// SetDataCategory sets the dataCategory property value. Windows privacy data category specifier for privacy data access.
func (m *WindowsPrivacyDataAccessControlItem) SetDataCategory(value *WindowsPrivacyDataCategory)() {
    m.dataCategory = value
}
