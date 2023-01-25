package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskProfile 
type WindowsKioskProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The app base class used to identify the application info for the kiosk configuration
    appConfiguration WindowsKioskAppConfigurationable
    // The OdataType property
    odataType *string
    // Key of the entity.
    profileId *string
    // This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
    profileName *string
    // The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100 elements.
    userAccountsConfiguration []WindowsKioskUserable
}
// NewWindowsKioskProfile instantiates a new windowsKioskProfile and sets the default values.
func NewWindowsKioskProfile()(*WindowsKioskProfile) {
    m := &WindowsKioskProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsKioskProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskProfile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppConfiguration gets the appConfiguration property value. The app base class used to identify the application info for the kiosk configuration
func (m *WindowsKioskProfile) GetAppConfiguration()(WindowsKioskAppConfigurationable) {
    return m.appConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsKioskAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConfiguration(val.(WindowsKioskAppConfigurationable))
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
    res["profileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileId(val)
        }
        return nil
    }
    res["profileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
        }
        return nil
    }
    res["userAccountsConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsKioskUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsKioskUserable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsKioskUserable)
            }
            m.SetUserAccountsConfiguration(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsKioskProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetProfileId gets the profileId property value. Key of the entity.
func (m *WindowsKioskProfile) GetProfileId()(*string) {
    return m.profileId
}
// GetProfileName gets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsKioskProfile) GetProfileName()(*string) {
    return m.profileName
}
// GetUserAccountsConfiguration gets the userAccountsConfiguration property value. The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100 elements.
func (m *WindowsKioskProfile) GetUserAccountsConfiguration()([]WindowsKioskUserable) {
    return m.userAccountsConfiguration
}
// Serialize serializes information the current object
func (m *WindowsKioskProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appConfiguration", m.GetAppConfiguration())
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
        err := writer.WriteStringValue("profileId", m.GetProfileId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("profileName", m.GetProfileName())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountsConfiguration() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserAccountsConfiguration()))
        for i, v := range m.GetUserAccountsConfiguration() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userAccountsConfiguration", cast)
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
func (m *WindowsKioskProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppConfiguration sets the appConfiguration property value. The app base class used to identify the application info for the kiosk configuration
func (m *WindowsKioskProfile) SetAppConfiguration(value WindowsKioskAppConfigurationable)() {
    m.appConfiguration = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsKioskProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProfileId sets the profileId property value. Key of the entity.
func (m *WindowsKioskProfile) SetProfileId(value *string)() {
    m.profileId = value
}
// SetProfileName sets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsKioskProfile) SetProfileName(value *string)() {
    m.profileName = value
}
// SetUserAccountsConfiguration sets the userAccountsConfiguration property value. The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100 elements.
func (m *WindowsKioskProfile) SetUserAccountsConfiguration(value []WindowsKioskUserable)() {
    m.userAccountsConfiguration = value
}
