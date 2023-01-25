package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutOfBoxExperienceSettings out of box experience setting
type OutOfBoxExperienceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceUsageType property
    deviceUsageType *WindowsDeviceUsageType
    // If set to true, then the user can't start over with different account, on company sign-in
    hideEscapeLink *bool
    // Show or hide EULA to user
    hideEULA *bool
    // Show or hide privacy settings to user
    hidePrivacySettings *bool
    // The OdataType property
    odataType *string
    // If set, then skip the keyboard selection page if Language and Region are set
    skipKeyboardSelectionPage *bool
    // The userType property
    userType *WindowsUserType
}
// NewOutOfBoxExperienceSettings instantiates a new outOfBoxExperienceSettings and sets the default values.
func NewOutOfBoxExperienceSettings()(*OutOfBoxExperienceSettings) {
    m := &OutOfBoxExperienceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOutOfBoxExperienceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutOfBoxExperienceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutOfBoxExperienceSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutOfBoxExperienceSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceUsageType gets the deviceUsageType property value. The deviceUsageType property
func (m *OutOfBoxExperienceSettings) GetDeviceUsageType()(*WindowsDeviceUsageType) {
    return m.deviceUsageType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutOfBoxExperienceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceUsageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceUsageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceUsageType(val.(*WindowsDeviceUsageType))
        }
        return nil
    }
    res["hideEscapeLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideEscapeLink(val)
        }
        return nil
    }
    res["hideEULA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideEULA(val)
        }
        return nil
    }
    res["hidePrivacySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidePrivacySettings(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["skipKeyboardSelectionPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipKeyboardSelectionPage(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*WindowsUserType))
        }
        return nil
    }
    return res
}
// GetHideEscapeLink gets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) GetHideEscapeLink()(*bool) {
    return m.hideEscapeLink
}
// GetHideEULA gets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) GetHideEULA()(*bool) {
    return m.hideEULA
}
// GetHidePrivacySettings gets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) GetHidePrivacySettings()(*bool) {
    return m.hidePrivacySettings
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OutOfBoxExperienceSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetSkipKeyboardSelectionPage gets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) GetSkipKeyboardSelectionPage()(*bool) {
    return m.skipKeyboardSelectionPage
}
// GetUserType gets the userType property value. The userType property
func (m *OutOfBoxExperienceSettings) GetUserType()(*WindowsUserType) {
    return m.userType
}
// Serialize serializes information the current object
func (m *OutOfBoxExperienceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceUsageType() != nil {
        cast := (*m.GetDeviceUsageType()).String()
        err := writer.WriteStringValue("deviceUsageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideEscapeLink", m.GetHideEscapeLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideEULA", m.GetHideEULA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidePrivacySettings", m.GetHidePrivacySettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("skipKeyboardSelectionPage", m.GetSkipKeyboardSelectionPage())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err := writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutOfBoxExperienceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceUsageType sets the deviceUsageType property value. The deviceUsageType property
func (m *OutOfBoxExperienceSettings) SetDeviceUsageType(value *WindowsDeviceUsageType)() {
    m.deviceUsageType = value
}
// SetHideEscapeLink sets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) SetHideEscapeLink(value *bool)() {
    m.hideEscapeLink = value
}
// SetHideEULA sets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) SetHideEULA(value *bool)() {
    m.hideEULA = value
}
// SetHidePrivacySettings sets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) SetHidePrivacySettings(value *bool)() {
    m.hidePrivacySettings = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OutOfBoxExperienceSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSkipKeyboardSelectionPage sets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) SetSkipKeyboardSelectionPage(value *bool)() {
    m.skipKeyboardSelectionPage = value
}
// SetUserType sets the userType property value. The userType property
func (m *OutOfBoxExperienceSettings) SetUserType(value *WindowsUserType)() {
    m.userType = value
}
