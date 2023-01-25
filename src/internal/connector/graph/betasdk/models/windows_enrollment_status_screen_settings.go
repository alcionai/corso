package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsEnrollmentStatusScreenSettings enrollment status screen setting
type WindowsEnrollmentStatusScreenSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Allow or block user to use device before profile and app installation complete
    allowDeviceUseBeforeProfileAndAppInstallComplete *bool
    // Allow the user to continue using the device on installation failure
    allowDeviceUseOnInstallFailure *bool
    // Allow or block log collection on installation failure
    allowLogCollectionOnInstallFailure *bool
    // Allow the user to retry the setup on installation failure
    blockDeviceSetupRetryByUser *bool
    // Set custom error message to show upon installation failure
    customErrorMessage *string
    // Show or hide installation progress to user
    hideInstallationProgress *bool
    // Set installation progress timeout in minutes
    installProgressTimeoutInMinutes *int32
    // The OdataType property
    odataType *string
}
// NewWindowsEnrollmentStatusScreenSettings instantiates a new windowsEnrollmentStatusScreenSettings and sets the default values.
func NewWindowsEnrollmentStatusScreenSettings()(*WindowsEnrollmentStatusScreenSettings) {
    m := &WindowsEnrollmentStatusScreenSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsEnrollmentStatusScreenSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsEnrollmentStatusScreenSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsEnrollmentStatusScreenSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsEnrollmentStatusScreenSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowDeviceUseBeforeProfileAndAppInstallComplete gets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool) {
    return m.allowDeviceUseBeforeProfileAndAppInstallComplete
}
// GetAllowDeviceUseOnInstallFailure gets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseOnInstallFailure()(*bool) {
    return m.allowDeviceUseOnInstallFailure
}
// GetAllowLogCollectionOnInstallFailure gets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowLogCollectionOnInstallFailure()(*bool) {
    return m.allowLogCollectionOnInstallFailure
}
// GetBlockDeviceSetupRetryByUser gets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetBlockDeviceSetupRetryByUser()(*bool) {
    return m.blockDeviceSetupRetryByUser
}
// GetCustomErrorMessage gets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetCustomErrorMessage()(*string) {
    return m.customErrorMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsEnrollmentStatusScreenSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowDeviceUseBeforeProfileAndAppInstallComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseBeforeProfileAndAppInstallComplete(val)
        }
        return nil
    }
    res["allowDeviceUseOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseOnInstallFailure(val)
        }
        return nil
    }
    res["allowLogCollectionOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLogCollectionOnInstallFailure(val)
        }
        return nil
    }
    res["blockDeviceSetupRetryByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDeviceSetupRetryByUser(val)
        }
        return nil
    }
    res["customErrorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomErrorMessage(val)
        }
        return nil
    }
    res["hideInstallationProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideInstallationProgress(val)
        }
        return nil
    }
    res["installProgressTimeoutInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallProgressTimeoutInMinutes(val)
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
    return res
}
// GetHideInstallationProgress gets the hideInstallationProgress property value. Show or hide installation progress to user
func (m *WindowsEnrollmentStatusScreenSettings) GetHideInstallationProgress()(*bool) {
    return m.hideInstallationProgress
}
// GetInstallProgressTimeoutInMinutes gets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *WindowsEnrollmentStatusScreenSettings) GetInstallProgressTimeoutInMinutes()(*int32) {
    return m.installProgressTimeoutInMinutes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsEnrollmentStatusScreenSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WindowsEnrollmentStatusScreenSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowDeviceUseBeforeProfileAndAppInstallComplete", m.GetAllowDeviceUseBeforeProfileAndAppInstallComplete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeviceUseOnInstallFailure", m.GetAllowDeviceUseOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowLogCollectionOnInstallFailure", m.GetAllowLogCollectionOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockDeviceSetupRetryByUser", m.GetBlockDeviceSetupRetryByUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customErrorMessage", m.GetCustomErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideInstallationProgress", m.GetHideInstallationProgress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("installProgressTimeoutInMinutes", m.GetInstallProgressTimeoutInMinutes())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsEnrollmentStatusScreenSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowDeviceUseBeforeProfileAndAppInstallComplete sets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)() {
    m.allowDeviceUseBeforeProfileAndAppInstallComplete = value
}
// SetAllowDeviceUseOnInstallFailure sets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    m.allowDeviceUseOnInstallFailure = value
}
// SetAllowLogCollectionOnInstallFailure sets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    m.allowLogCollectionOnInstallFailure = value
}
// SetBlockDeviceSetupRetryByUser sets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetBlockDeviceSetupRetryByUser(value *bool)() {
    m.blockDeviceSetupRetryByUser = value
}
// SetCustomErrorMessage sets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetCustomErrorMessage(value *string)() {
    m.customErrorMessage = value
}
// SetHideInstallationProgress sets the hideInstallationProgress property value. Show or hide installation progress to user
func (m *WindowsEnrollmentStatusScreenSettings) SetHideInstallationProgress(value *bool)() {
    m.hideInstallationProgress = value
}
// SetInstallProgressTimeoutInMinutes sets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *WindowsEnrollmentStatusScreenSettings) SetInstallProgressTimeoutInMinutes(value *int32)() {
    m.installProgressTimeoutInMinutes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsEnrollmentStatusScreenSettings) SetOdataType(value *string)() {
    m.odataType = value
}
