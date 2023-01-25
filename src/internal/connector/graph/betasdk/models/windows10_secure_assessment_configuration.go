package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10SecureAssessmentConfiguration 
type Windows10SecureAssessmentConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to allow the app from printing during the test.
    allowPrinting *bool
    // Indicates whether or not to allow screen capture capability during a test.
    allowScreenCapture *bool
    // Indicates whether or not to allow text suggestions during the test.
    allowTextSuggestion *bool
    // Specifies the application user model ID of the assessment app launched when a user signs in to a secure assessment with a local guest account. Important notice: this property must be set with localGuestAccountName in order to make the local guest account sign-in experience work properly for secure assessments.
    assessmentAppUserModelId *string
    // The account used to configure the Windows device for taking the test. The user can be a domain account (domain/user), an AAD account (username@tenant.com) or a local account (username).
    configurationAccount *string
    // Type of accounts that are allowed for Windows10SecureAssessment ConfigurationAccount.
    configurationAccountType *SecureAssessmentAccountType
    // Url link to an assessment that's automatically loaded when the secure assessment browser is launched. It has to be a valid Url (http[s]://msdn.microsoft.com/).
    launchUri *string
    // Specifies the display text for the local guest account shown on the sign-in screen. Typically is the name of an assessment. When the user clicks the local guest account on the sign-in screen, an assessment app is launched with a specified assessment URL. Secure assessments can only be configured with local guest account sign-in on devices running Windows 10, version 1903 or later. Important notice: this property must be set with assessmentAppUserModelID in order to make the local guest account sign-in experience work properly for secure assessments.
    localGuestAccountName *string
}
// NewWindows10SecureAssessmentConfiguration instantiates a new Windows10SecureAssessmentConfiguration and sets the default values.
func NewWindows10SecureAssessmentConfiguration()(*Windows10SecureAssessmentConfiguration) {
    m := &Windows10SecureAssessmentConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10SecureAssessmentConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10SecureAssessmentConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10SecureAssessmentConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10SecureAssessmentConfiguration(), nil
}
// GetAllowPrinting gets the allowPrinting property value. Indicates whether or not to allow the app from printing during the test.
func (m *Windows10SecureAssessmentConfiguration) GetAllowPrinting()(*bool) {
    return m.allowPrinting
}
// GetAllowScreenCapture gets the allowScreenCapture property value. Indicates whether or not to allow screen capture capability during a test.
func (m *Windows10SecureAssessmentConfiguration) GetAllowScreenCapture()(*bool) {
    return m.allowScreenCapture
}
// GetAllowTextSuggestion gets the allowTextSuggestion property value. Indicates whether or not to allow text suggestions during the test.
func (m *Windows10SecureAssessmentConfiguration) GetAllowTextSuggestion()(*bool) {
    return m.allowTextSuggestion
}
// GetAssessmentAppUserModelId gets the assessmentAppUserModelId property value. Specifies the application user model ID of the assessment app launched when a user signs in to a secure assessment with a local guest account. Important notice: this property must be set with localGuestAccountName in order to make the local guest account sign-in experience work properly for secure assessments.
func (m *Windows10SecureAssessmentConfiguration) GetAssessmentAppUserModelId()(*string) {
    return m.assessmentAppUserModelId
}
// GetConfigurationAccount gets the configurationAccount property value. The account used to configure the Windows device for taking the test. The user can be a domain account (domain/user), an AAD account (username@tenant.com) or a local account (username).
func (m *Windows10SecureAssessmentConfiguration) GetConfigurationAccount()(*string) {
    return m.configurationAccount
}
// GetConfigurationAccountType gets the configurationAccountType property value. Type of accounts that are allowed for Windows10SecureAssessment ConfigurationAccount.
func (m *Windows10SecureAssessmentConfiguration) GetConfigurationAccountType()(*SecureAssessmentAccountType) {
    return m.configurationAccountType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10SecureAssessmentConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allowPrinting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPrinting(val)
        }
        return nil
    }
    res["allowScreenCapture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowScreenCapture(val)
        }
        return nil
    }
    res["allowTextSuggestion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTextSuggestion(val)
        }
        return nil
    }
    res["assessmentAppUserModelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssessmentAppUserModelId(val)
        }
        return nil
    }
    res["configurationAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationAccount(val)
        }
        return nil
    }
    res["configurationAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecureAssessmentAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationAccountType(val.(*SecureAssessmentAccountType))
        }
        return nil
    }
    res["launchUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLaunchUri(val)
        }
        return nil
    }
    res["localGuestAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalGuestAccountName(val)
        }
        return nil
    }
    return res
}
// GetLaunchUri gets the launchUri property value. Url link to an assessment that's automatically loaded when the secure assessment browser is launched. It has to be a valid Url (http[s]://msdn.microsoft.com/).
func (m *Windows10SecureAssessmentConfiguration) GetLaunchUri()(*string) {
    return m.launchUri
}
// GetLocalGuestAccountName gets the localGuestAccountName property value. Specifies the display text for the local guest account shown on the sign-in screen. Typically is the name of an assessment. When the user clicks the local guest account on the sign-in screen, an assessment app is launched with a specified assessment URL. Secure assessments can only be configured with local guest account sign-in on devices running Windows 10, version 1903 or later. Important notice: this property must be set with assessmentAppUserModelID in order to make the local guest account sign-in experience work properly for secure assessments.
func (m *Windows10SecureAssessmentConfiguration) GetLocalGuestAccountName()(*string) {
    return m.localGuestAccountName
}
// Serialize serializes information the current object
func (m *Windows10SecureAssessmentConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowPrinting", m.GetAllowPrinting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowScreenCapture", m.GetAllowScreenCapture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowTextSuggestion", m.GetAllowTextSuggestion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assessmentAppUserModelId", m.GetAssessmentAppUserModelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("configurationAccount", m.GetConfigurationAccount())
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationAccountType() != nil {
        cast := (*m.GetConfigurationAccountType()).String()
        err = writer.WriteStringValue("configurationAccountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("launchUri", m.GetLaunchUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localGuestAccountName", m.GetLocalGuestAccountName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowPrinting sets the allowPrinting property value. Indicates whether or not to allow the app from printing during the test.
func (m *Windows10SecureAssessmentConfiguration) SetAllowPrinting(value *bool)() {
    m.allowPrinting = value
}
// SetAllowScreenCapture sets the allowScreenCapture property value. Indicates whether or not to allow screen capture capability during a test.
func (m *Windows10SecureAssessmentConfiguration) SetAllowScreenCapture(value *bool)() {
    m.allowScreenCapture = value
}
// SetAllowTextSuggestion sets the allowTextSuggestion property value. Indicates whether or not to allow text suggestions during the test.
func (m *Windows10SecureAssessmentConfiguration) SetAllowTextSuggestion(value *bool)() {
    m.allowTextSuggestion = value
}
// SetAssessmentAppUserModelId sets the assessmentAppUserModelId property value. Specifies the application user model ID of the assessment app launched when a user signs in to a secure assessment with a local guest account. Important notice: this property must be set with localGuestAccountName in order to make the local guest account sign-in experience work properly for secure assessments.
func (m *Windows10SecureAssessmentConfiguration) SetAssessmentAppUserModelId(value *string)() {
    m.assessmentAppUserModelId = value
}
// SetConfigurationAccount sets the configurationAccount property value. The account used to configure the Windows device for taking the test. The user can be a domain account (domain/user), an AAD account (username@tenant.com) or a local account (username).
func (m *Windows10SecureAssessmentConfiguration) SetConfigurationAccount(value *string)() {
    m.configurationAccount = value
}
// SetConfigurationAccountType sets the configurationAccountType property value. Type of accounts that are allowed for Windows10SecureAssessment ConfigurationAccount.
func (m *Windows10SecureAssessmentConfiguration) SetConfigurationAccountType(value *SecureAssessmentAccountType)() {
    m.configurationAccountType = value
}
// SetLaunchUri sets the launchUri property value. Url link to an assessment that's automatically loaded when the secure assessment browser is launched. It has to be a valid Url (http[s]://msdn.microsoft.com/).
func (m *Windows10SecureAssessmentConfiguration) SetLaunchUri(value *string)() {
    m.launchUri = value
}
// SetLocalGuestAccountName sets the localGuestAccountName property value. Specifies the display text for the local guest account shown on the sign-in screen. Typically is the name of an assessment. When the user clicks the local guest account on the sign-in screen, an assessment app is launched with a specified assessment URL. Secure assessments can only be configured with local guest account sign-in on devices running Windows 10, version 1903 or later. Important notice: this property must be set with assessmentAppUserModelID in order to make the local guest account sign-in experience work properly for secure assessments.
func (m *Windows10SecureAssessmentConfiguration) SetLocalGuestAccountName(value *string)() {
    m.localGuestAccountName = value
}
