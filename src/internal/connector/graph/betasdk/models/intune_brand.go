package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntuneBrand intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
type IntuneBrand struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of blocked actions on the company portal as per platform and device ownership types.
    companyPortalBlockedActions []CompanyPortalBlockedActionable
    // Email address of the person/organization responsible for IT support.
    contactITEmailAddress *string
    // Name of the person/organization responsible for IT support.
    contactITName *string
    // Text comments regarding the person/organization responsible for IT support.
    contactITNotes *string
    // Phone number of the person/organization responsible for IT support.
    contactITPhoneNumber *string
    // The custom privacy message used to explain what the organization can see and do on managed devices.
    customCanSeePrivacyMessage *string
    // The custom privacy message used to explain what the organization can’t see or do on managed devices.
    customCantSeePrivacyMessage *string
    // The custom privacy message used to explain what the organization can’t see or do on managed devices.
    customPrivacyMessage *string
    // Logo image displayed in Company Portal apps which have a dark background behind the logo.
    darkBackgroundLogo MimeContentable
    // Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
    disableClientTelemetry *bool
    // Company/organization name that is displayed to end users.
    displayName *string
    // Options available for enrollment flow customization
    enrollmentAvailability *EnrollmentAvailabilityOptions
    // Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
    isFactoryResetDisabled *bool
    // Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
    isRemoveDeviceDisabled *bool
    // Customized image displayed in Company Portal app landing page
    landingPageCustomizedImage MimeContentable
    // Logo image displayed in Company Portal apps which have a light background behind the logo.
    lightBackgroundLogo MimeContentable
    // The OdataType property
    odataType *string
    // Display name of the company/organization’s IT helpdesk site.
    onlineSupportSiteName *string
    // URL to the company/organization’s IT helpdesk site.
    onlineSupportSiteUrl *string
    // URL to the company/organization’s privacy policy.
    privacyUrl *string
    // List of scope tags assigned to the default branding profile
    roleScopeTagIds []string
    // Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
    sendDeviceOwnershipChangePushNotification *bool
    // Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
    showAzureADEnterpriseApps *bool
    // Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
    showConfigurationManagerApps *bool
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showDisplayNameNextToLogo *bool
    // Boolean that represents whether the administrator-supplied logo images are shown or not shown.
    showLogo *bool
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showNameNextToLogo *bool
    // Boolean that indicates if Office WebApps will be shown in Company Portal
    showOfficeWebApps *bool
    // Primary theme color used in the Company Portal applications and web portal.
    themeColor RgbColorable
}
// NewIntuneBrand instantiates a new intuneBrand and sets the default values.
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIntuneBrandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntuneBrand(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompanyPortalBlockedActions gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable) {
    return m.companyPortalBlockedActions
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    return m.contactITEmailAddress
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITName()(*string) {
    return m.contactITName
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITNotes()(*string) {
    return m.contactITNotes
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    return m.contactITPhoneNumber
}
// GetCustomCanSeePrivacyMessage gets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) GetCustomCanSeePrivacyMessage()(*string) {
    return m.customCanSeePrivacyMessage
}
// GetCustomCantSeePrivacyMessage gets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomCantSeePrivacyMessage()(*string) {
    return m.customCantSeePrivacyMessage
}
// GetCustomPrivacyMessage gets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomPrivacyMessage()(*string) {
    return m.customPrivacyMessage
}
// GetDarkBackgroundLogo gets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) GetDarkBackgroundLogo()(MimeContentable) {
    return m.darkBackgroundLogo
}
// GetDisableClientTelemetry gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) GetDisableClientTelemetry()(*bool) {
    return m.disableClientTelemetry
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrollmentAvailability gets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrand) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    return m.enrollmentAvailability
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrand) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["companyPortalBlockedActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanyPortalBlockedActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyPortalBlockedActionable, len(val))
            for i, v := range val {
                res[i] = v.(CompanyPortalBlockedActionable)
            }
            m.SetCompanyPortalBlockedActions(res)
        }
        return nil
    }
    res["contactITEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITEmailAddress(val)
        }
        return nil
    }
    res["contactITName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITName(val)
        }
        return nil
    }
    res["contactITNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITNotes(val)
        }
        return nil
    }
    res["contactITPhoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITPhoneNumber(val)
        }
        return nil
    }
    res["customCanSeePrivacyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCanSeePrivacyMessage(val)
        }
        return nil
    }
    res["customCantSeePrivacyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCantSeePrivacyMessage(val)
        }
        return nil
    }
    res["customPrivacyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyMessage(val)
        }
        return nil
    }
    res["darkBackgroundLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDarkBackgroundLogo(val.(MimeContentable))
        }
        return nil
    }
    res["disableClientTelemetry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableClientTelemetry(val)
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
    res["enrollmentAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentAvailabilityOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentAvailability(val.(*EnrollmentAvailabilityOptions))
        }
        return nil
    }
    res["isFactoryResetDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFactoryResetDisabled(val)
        }
        return nil
    }
    res["isRemoveDeviceDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemoveDeviceDisabled(val)
        }
        return nil
    }
    res["landingPageCustomizedImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLandingPageCustomizedImage(val.(MimeContentable))
        }
        return nil
    }
    res["lightBackgroundLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLightBackgroundLogo(val.(MimeContentable))
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
    res["onlineSupportSiteName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteName(val)
        }
        return nil
    }
    res["onlineSupportSiteUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteUrl(val)
        }
        return nil
    }
    res["privacyUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyUrl(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["sendDeviceOwnershipChangePushNotification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendDeviceOwnershipChangePushNotification(val)
        }
        return nil
    }
    res["showAzureADEnterpriseApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowAzureADEnterpriseApps(val)
        }
        return nil
    }
    res["showConfigurationManagerApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowConfigurationManagerApps(val)
        }
        return nil
    }
    res["showDisplayNameNextToLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowDisplayNameNextToLogo(val)
        }
        return nil
    }
    res["showLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLogo(val)
        }
        return nil
    }
    res["showNameNextToLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowNameNextToLogo(val)
        }
        return nil
    }
    res["showOfficeWebApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowOfficeWebApps(val)
        }
        return nil
    }
    res["themeColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRgbColorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColor(val.(RgbColorable))
        }
        return nil
    }
    return res
}
// GetIsFactoryResetDisabled gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) GetIsFactoryResetDisabled()(*bool) {
    return m.isFactoryResetDisabled
}
// GetIsRemoveDeviceDisabled gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) GetIsRemoveDeviceDisabled()(*bool) {
    return m.isRemoveDeviceDisabled
}
// GetLandingPageCustomizedImage gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) GetLandingPageCustomizedImage()(MimeContentable) {
    return m.landingPageCustomizedImage
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) GetLightBackgroundLogo()(MimeContentable) {
    return m.lightBackgroundLogo
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IntuneBrand) GetOdataType()(*string) {
    return m.odataType
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    return m.onlineSupportSiteName
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    return m.onlineSupportSiteUrl
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    return m.privacyUrl
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSendDeviceOwnershipChangePushNotification gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    return m.sendDeviceOwnershipChangePushNotification
}
// GetShowAzureADEnterpriseApps gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) GetShowAzureADEnterpriseApps()(*bool) {
    return m.showAzureADEnterpriseApps
}
// GetShowConfigurationManagerApps gets the showConfigurationManagerApps property value. Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
func (m *IntuneBrand) GetShowConfigurationManagerApps()(*bool) {
    return m.showConfigurationManagerApps
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    return m.showDisplayNameNextToLogo
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) GetShowLogo()(*bool) {
    return m.showLogo
}
// GetShowNameNextToLogo gets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    return m.showNameNextToLogo
}
// GetShowOfficeWebApps gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) GetShowOfficeWebApps()(*bool) {
    return m.showOfficeWebApps
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) GetThemeColor()(RgbColorable) {
    return m.themeColor
}
// Serialize serializes information the current object
func (m *IntuneBrand) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCompanyPortalBlockedActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyPortalBlockedActions()))
        for i, v := range m.GetCompanyPortalBlockedActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("companyPortalBlockedActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITEmailAddress", m.GetContactITEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITName", m.GetContactITName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITNotes", m.GetContactITNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITPhoneNumber", m.GetContactITPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customCanSeePrivacyMessage", m.GetCustomCanSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customCantSeePrivacyMessage", m.GetCustomCantSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customPrivacyMessage", m.GetCustomPrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("darkBackgroundLogo", m.GetDarkBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("disableClientTelemetry", m.GetDisableClientTelemetry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentAvailability() != nil {
        cast := (*m.GetEnrollmentAvailability()).String()
        err := writer.WriteStringValue("enrollmentAvailability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isFactoryResetDisabled", m.GetIsFactoryResetDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRemoveDeviceDisabled", m.GetIsRemoveDeviceDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("landingPageCustomizedImage", m.GetLandingPageCustomizedImage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lightBackgroundLogo", m.GetLightBackgroundLogo())
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
        err := writer.WriteStringValue("onlineSupportSiteName", m.GetOnlineSupportSiteName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onlineSupportSiteUrl", m.GetOnlineSupportSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privacyUrl", m.GetPrivacyUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendDeviceOwnershipChangePushNotification", m.GetSendDeviceOwnershipChangePushNotification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showAzureADEnterpriseApps", m.GetShowAzureADEnterpriseApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showConfigurationManagerApps", m.GetShowConfigurationManagerApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showDisplayNameNextToLogo", m.GetShowDisplayNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showLogo", m.GetShowLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showNameNextToLogo", m.GetShowNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showOfficeWebApps", m.GetShowOfficeWebApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("themeColor", m.GetThemeColor())
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
func (m *IntuneBrand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompanyPortalBlockedActions sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)() {
    m.companyPortalBlockedActions = value
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITName(value *string)() {
    m.contactITName = value
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
// SetCustomCanSeePrivacyMessage sets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) SetCustomCanSeePrivacyMessage(value *string)() {
    m.customCanSeePrivacyMessage = value
}
// SetCustomCantSeePrivacyMessage sets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) SetCustomCantSeePrivacyMessage(value *string)() {
    m.customCantSeePrivacyMessage = value
}
// SetCustomPrivacyMessage sets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) SetCustomPrivacyMessage(value *string)() {
    m.customPrivacyMessage = value
}
// SetDarkBackgroundLogo sets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) SetDarkBackgroundLogo(value MimeContentable)() {
    m.darkBackgroundLogo = value
}
// SetDisableClientTelemetry sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) SetDisableClientTelemetry(value *bool)() {
    m.disableClientTelemetry = value
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrollmentAvailability sets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrand) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    m.enrollmentAvailability = value
}
// SetIsFactoryResetDisabled sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) SetIsFactoryResetDisabled(value *bool)() {
    m.isFactoryResetDisabled = value
}
// SetIsRemoveDeviceDisabled sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) SetIsRemoveDeviceDisabled(value *bool)() {
    m.isRemoveDeviceDisabled = value
}
// SetLandingPageCustomizedImage sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) SetLandingPageCustomizedImage(value MimeContentable)() {
    m.landingPageCustomizedImage = value
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) SetLightBackgroundLogo(value MimeContentable)() {
    m.lightBackgroundLogo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IntuneBrand) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSendDeviceOwnershipChangePushNotification sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    m.sendDeviceOwnershipChangePushNotification = value
}
// SetShowAzureADEnterpriseApps sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) SetShowAzureADEnterpriseApps(value *bool)() {
    m.showAzureADEnterpriseApps = value
}
// SetShowConfigurationManagerApps sets the showConfigurationManagerApps property value. Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
func (m *IntuneBrand) SetShowConfigurationManagerApps(value *bool)() {
    m.showConfigurationManagerApps = value
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    m.showLogo = value
}
// SetShowNameNextToLogo sets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    m.showNameNextToLogo = value
}
// SetShowOfficeWebApps sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) SetShowOfficeWebApps(value *bool)() {
    m.showOfficeWebApps = value
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) SetThemeColor(value RgbColorable)() {
    m.themeColor = value
}
