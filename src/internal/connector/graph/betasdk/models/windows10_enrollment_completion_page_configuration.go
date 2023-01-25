package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnrollmentCompletionPageConfiguration 
type Windows10EnrollmentCompletionPageConfiguration struct {
    DeviceEnrollmentConfiguration
    // Allow or block device reset on installation failure
    allowDeviceResetOnInstallFailure *bool
    // Allow the user to continue using the device on installation failure
    allowDeviceUseOnInstallFailure *bool
    // Allow or block log collection on installation failure
    allowLogCollectionOnInstallFailure *bool
    // Install all required apps as non blocking apps during white glove
    allowNonBlockingAppInstallation *bool
    // Allow the user to retry the setup on installation failure
    blockDeviceSetupRetryByUser *bool
    // Set custom error message to show upon installation failure
    customErrorMessage *string
    // Only show installation progress for first user post enrollment
    disableUserStatusTrackingAfterFirstUser *bool
    // Set installation progress timeout in minutes
    installProgressTimeoutInMinutes *int32
    // Allows quality updates installation during OOBE
    installQualityUpdates *bool
    // Selected applications to track the installation status
    selectedMobileAppIds []string
    // Show or hide installation progress to user
    showInstallationProgress *bool
    // Only show installation progress for Autopilot enrollment scenarios
    trackInstallProgressForAutopilotOnly *bool
}
// NewWindows10EnrollmentCompletionPageConfiguration instantiates a new Windows10EnrollmentCompletionPageConfiguration and sets the default values.
func NewWindows10EnrollmentCompletionPageConfiguration()(*Windows10EnrollmentCompletionPageConfiguration) {
    m := &Windows10EnrollmentCompletionPageConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10EnrollmentCompletionPageConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10EnrollmentCompletionPageConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EnrollmentCompletionPageConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EnrollmentCompletionPageConfiguration(), nil
}
// GetAllowDeviceResetOnInstallFailure gets the allowDeviceResetOnInstallFailure property value. Allow or block device reset on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowDeviceResetOnInstallFailure()(*bool) {
    return m.allowDeviceResetOnInstallFailure
}
// GetAllowDeviceUseOnInstallFailure gets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowDeviceUseOnInstallFailure()(*bool) {
    return m.allowDeviceUseOnInstallFailure
}
// GetAllowLogCollectionOnInstallFailure gets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowLogCollectionOnInstallFailure()(*bool) {
    return m.allowLogCollectionOnInstallFailure
}
// GetAllowNonBlockingAppInstallation gets the allowNonBlockingAppInstallation property value. Install all required apps as non blocking apps during white glove
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowNonBlockingAppInstallation()(*bool) {
    return m.allowNonBlockingAppInstallation
}
// GetBlockDeviceSetupRetryByUser gets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) GetBlockDeviceSetupRetryByUser()(*bool) {
    return m.blockDeviceSetupRetryByUser
}
// GetCustomErrorMessage gets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) GetCustomErrorMessage()(*string) {
    return m.customErrorMessage
}
// GetDisableUserStatusTrackingAfterFirstUser gets the disableUserStatusTrackingAfterFirstUser property value. Only show installation progress for first user post enrollment
func (m *Windows10EnrollmentCompletionPageConfiguration) GetDisableUserStatusTrackingAfterFirstUser()(*bool) {
    return m.disableUserStatusTrackingAfterFirstUser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EnrollmentCompletionPageConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["allowDeviceResetOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceResetOnInstallFailure(val)
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
    res["allowNonBlockingAppInstallation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNonBlockingAppInstallation(val)
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
    res["disableUserStatusTrackingAfterFirstUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserStatusTrackingAfterFirstUser(val)
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
    res["installQualityUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallQualityUpdates(val)
        }
        return nil
    }
    res["selectedMobileAppIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelectedMobileAppIds(res)
        }
        return nil
    }
    res["showInstallationProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInstallationProgress(val)
        }
        return nil
    }
    res["trackInstallProgressForAutopilotOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackInstallProgressForAutopilotOnly(val)
        }
        return nil
    }
    return res
}
// GetInstallProgressTimeoutInMinutes gets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *Windows10EnrollmentCompletionPageConfiguration) GetInstallProgressTimeoutInMinutes()(*int32) {
    return m.installProgressTimeoutInMinutes
}
// GetInstallQualityUpdates gets the installQualityUpdates property value. Allows quality updates installation during OOBE
func (m *Windows10EnrollmentCompletionPageConfiguration) GetInstallQualityUpdates()(*bool) {
    return m.installQualityUpdates
}
// GetSelectedMobileAppIds gets the selectedMobileAppIds property value. Selected applications to track the installation status
func (m *Windows10EnrollmentCompletionPageConfiguration) GetSelectedMobileAppIds()([]string) {
    return m.selectedMobileAppIds
}
// GetShowInstallationProgress gets the showInstallationProgress property value. Show or hide installation progress to user
func (m *Windows10EnrollmentCompletionPageConfiguration) GetShowInstallationProgress()(*bool) {
    return m.showInstallationProgress
}
// GetTrackInstallProgressForAutopilotOnly gets the trackInstallProgressForAutopilotOnly property value. Only show installation progress for Autopilot enrollment scenarios
func (m *Windows10EnrollmentCompletionPageConfiguration) GetTrackInstallProgressForAutopilotOnly()(*bool) {
    return m.trackInstallProgressForAutopilotOnly
}
// Serialize serializes information the current object
func (m *Windows10EnrollmentCompletionPageConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowDeviceResetOnInstallFailure", m.GetAllowDeviceResetOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowDeviceUseOnInstallFailure", m.GetAllowDeviceUseOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowLogCollectionOnInstallFailure", m.GetAllowLogCollectionOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowNonBlockingAppInstallation", m.GetAllowNonBlockingAppInstallation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockDeviceSetupRetryByUser", m.GetBlockDeviceSetupRetryByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customErrorMessage", m.GetCustomErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUserStatusTrackingAfterFirstUser", m.GetDisableUserStatusTrackingAfterFirstUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installProgressTimeoutInMinutes", m.GetInstallProgressTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("installQualityUpdates", m.GetInstallQualityUpdates())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedMobileAppIds() != nil {
        err = writer.WriteCollectionOfStringValues("selectedMobileAppIds", m.GetSelectedMobileAppIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showInstallationProgress", m.GetShowInstallationProgress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("trackInstallProgressForAutopilotOnly", m.GetTrackInstallProgressForAutopilotOnly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDeviceResetOnInstallFailure sets the allowDeviceResetOnInstallFailure property value. Allow or block device reset on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowDeviceResetOnInstallFailure(value *bool)() {
    m.allowDeviceResetOnInstallFailure = value
}
// SetAllowDeviceUseOnInstallFailure sets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    m.allowDeviceUseOnInstallFailure = value
}
// SetAllowLogCollectionOnInstallFailure sets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    m.allowLogCollectionOnInstallFailure = value
}
// SetAllowNonBlockingAppInstallation sets the allowNonBlockingAppInstallation property value. Install all required apps as non blocking apps during white glove
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowNonBlockingAppInstallation(value *bool)() {
    m.allowNonBlockingAppInstallation = value
}
// SetBlockDeviceSetupRetryByUser sets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) SetBlockDeviceSetupRetryByUser(value *bool)() {
    m.blockDeviceSetupRetryByUser = value
}
// SetCustomErrorMessage sets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *Windows10EnrollmentCompletionPageConfiguration) SetCustomErrorMessage(value *string)() {
    m.customErrorMessage = value
}
// SetDisableUserStatusTrackingAfterFirstUser sets the disableUserStatusTrackingAfterFirstUser property value. Only show installation progress for first user post enrollment
func (m *Windows10EnrollmentCompletionPageConfiguration) SetDisableUserStatusTrackingAfterFirstUser(value *bool)() {
    m.disableUserStatusTrackingAfterFirstUser = value
}
// SetInstallProgressTimeoutInMinutes sets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *Windows10EnrollmentCompletionPageConfiguration) SetInstallProgressTimeoutInMinutes(value *int32)() {
    m.installProgressTimeoutInMinutes = value
}
// SetInstallQualityUpdates sets the installQualityUpdates property value. Allows quality updates installation during OOBE
func (m *Windows10EnrollmentCompletionPageConfiguration) SetInstallQualityUpdates(value *bool)() {
    m.installQualityUpdates = value
}
// SetSelectedMobileAppIds sets the selectedMobileAppIds property value. Selected applications to track the installation status
func (m *Windows10EnrollmentCompletionPageConfiguration) SetSelectedMobileAppIds(value []string)() {
    m.selectedMobileAppIds = value
}
// SetShowInstallationProgress sets the showInstallationProgress property value. Show or hide installation progress to user
func (m *Windows10EnrollmentCompletionPageConfiguration) SetShowInstallationProgress(value *bool)() {
    m.showInstallationProgress = value
}
// SetTrackInstallProgressForAutopilotOnly sets the trackInstallProgressForAutopilotOnly property value. Only show installation progress for Autopilot enrollment scenarios
func (m *Windows10EnrollmentCompletionPageConfiguration) SetTrackInstallProgressForAutopilotOnly(value *bool)() {
    m.trackInstallProgressForAutopilotOnly = value
}
