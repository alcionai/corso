package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftApplicationDataAccessSettings 
type MicrosoftApplicationDataAccessSettings struct {
    Entity
    // The ID of an Azure Active Directory (Azure AD) security group for which the members are allowed to access Microsoft 365 data using only Microsoft 365 apps, but not other Microsoft apps such as Edge.  This is only applicable if isEnabledForAllMicrosoftApplications is set to true.
    disabledForGroup *string
    // When set to true, all users in the organization can access in a Microsoft app any Microsoft 365 data that the user has been authorized to access. The Microsoft app can be a Microsoft 365 app (for example, Excel, Outlook) or non-Microsoft 365 app (for example, Edge). The default is true.  It is possible to disable this access for a subset of users in an Azure AD security group, by specifying the group in the disabledForGroup property.  When set to false, all users can access authorized Microsoft 365 data only in a Microsoft 365 app.
    isEnabledForAllMicrosoftApplications *bool
}
// NewMicrosoftApplicationDataAccessSettings instantiates a new microsoftApplicationDataAccessSettings and sets the default values.
func NewMicrosoftApplicationDataAccessSettings()(*MicrosoftApplicationDataAccessSettings) {
    m := &MicrosoftApplicationDataAccessSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftApplicationDataAccessSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftApplicationDataAccessSettings(), nil
}
// GetDisabledForGroup gets the disabledForGroup property value. The ID of an Azure Active Directory (Azure AD) security group for which the members are allowed to access Microsoft 365 data using only Microsoft 365 apps, but not other Microsoft apps such as Edge.  This is only applicable if isEnabledForAllMicrosoftApplications is set to true.
func (m *MicrosoftApplicationDataAccessSettings) GetDisabledForGroup()(*string) {
    return m.disabledForGroup
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftApplicationDataAccessSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["disabledForGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledForGroup(val)
        }
        return nil
    }
    res["isEnabledForAllMicrosoftApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForAllMicrosoftApplications(val)
        }
        return nil
    }
    return res
}
// GetIsEnabledForAllMicrosoftApplications gets the isEnabledForAllMicrosoftApplications property value. When set to true, all users in the organization can access in a Microsoft app any Microsoft 365 data that the user has been authorized to access. The Microsoft app can be a Microsoft 365 app (for example, Excel, Outlook) or non-Microsoft 365 app (for example, Edge). The default is true.  It is possible to disable this access for a subset of users in an Azure AD security group, by specifying the group in the disabledForGroup property.  When set to false, all users can access authorized Microsoft 365 data only in a Microsoft 365 app.
func (m *MicrosoftApplicationDataAccessSettings) GetIsEnabledForAllMicrosoftApplications()(*bool) {
    return m.isEnabledForAllMicrosoftApplications
}
// Serialize serializes information the current object
func (m *MicrosoftApplicationDataAccessSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("disabledForGroup", m.GetDisabledForGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForAllMicrosoftApplications", m.GetIsEnabledForAllMicrosoftApplications())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisabledForGroup sets the disabledForGroup property value. The ID of an Azure Active Directory (Azure AD) security group for which the members are allowed to access Microsoft 365 data using only Microsoft 365 apps, but not other Microsoft apps such as Edge.  This is only applicable if isEnabledForAllMicrosoftApplications is set to true.
func (m *MicrosoftApplicationDataAccessSettings) SetDisabledForGroup(value *string)() {
    m.disabledForGroup = value
}
// SetIsEnabledForAllMicrosoftApplications sets the isEnabledForAllMicrosoftApplications property value. When set to true, all users in the organization can access in a Microsoft app any Microsoft 365 data that the user has been authorized to access. The Microsoft app can be a Microsoft 365 app (for example, Excel, Outlook) or non-Microsoft 365 app (for example, Edge). The default is true.  It is possible to disable this access for a subset of users in an Azure AD security group, by specifying the group in the disabledForGroup property.  When set to false, all users can access authorized Microsoft 365 data only in a Microsoft 365 app.
func (m *MicrosoftApplicationDataAccessSettings) SetIsEnabledForAllMicrosoftApplications(value *bool)() {
    m.isEnabledForAllMicrosoftApplications = value
}
