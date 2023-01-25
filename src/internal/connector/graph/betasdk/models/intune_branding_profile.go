package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntuneBrandingProfile this entity contains data which is used in customizing the tenant level appearance of the Company Portal applications as well as the end user web portal.
type IntuneBrandingProfile struct {
    Entity
    // The list of group assignments for the branding profile
    assignments []IntuneBrandingProfileAssignmentable
    // Collection of blocked actions on the company portal as per platform and device ownership types.
    companyPortalBlockedActions []CompanyPortalBlockedActionable
    // E-mail address of the person/organization responsible for IT support
    contactITEmailAddress *string
    // Name of the person/organization responsible for IT support
    contactITName *string
    // Text comments regarding the person/organization responsible for IT support
    contactITNotes *string
    // Phone number of the person/organization responsible for IT support
    contactITPhoneNumber *string
    // Time when the BrandingProfile was created
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Text comments regarding what the admin has access to on the device
    customCanSeePrivacyMessage *string
    // Text comments regarding what the admin doesn't have access to on the device
    customCantSeePrivacyMessage *string
    // Text comments regarding what the admin doesn't have access to on the device
    customPrivacyMessage *string
    // Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
    disableClientTelemetry *bool
    // Company/organization name that is displayed to end users
    displayName *string
    // Options available for enrollment flow customization
    enrollmentAvailability *EnrollmentAvailabilityOptions
    // Boolean that represents whether the profile is used as default or not
    isDefaultProfile *bool
    // Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
    isFactoryResetDisabled *bool
    // Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
    isRemoveDeviceDisabled *bool
    // Customized image displayed in Company Portal apps landing page
    landingPageCustomizedImage MimeContentable
    // Time when the BrandingProfile was last modified
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Logo image displayed in Company Portal apps which have a light background behind the logo
    lightBackgroundLogo MimeContentable
    // Display name of the company/organization’s IT helpdesk site
    onlineSupportSiteName *string
    // URL to the company/organization’s IT helpdesk site
    onlineSupportSiteUrl *string
    // URL to the company/organization’s privacy policy
    privacyUrl *string
    // Description of the profile
    profileDescription *string
    // Name of the profile
    profileName *string
    // List of scope tags assigned to the branding profile
    roleScopeTagIds []string
    // Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
    sendDeviceOwnershipChangePushNotification *bool
    // Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
    showAzureADEnterpriseApps *bool
    // Boolean that indicates if Configuration Manager Apps will be shown in Company Portal
    showConfigurationManagerApps *bool
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
    showDisplayNameNextToLogo *bool
    // Boolean that represents whether the administrator-supplied logo images are shown or not
    showLogo *bool
    // Boolean that indicates if Office WebApps will be shown in Company Portal
    showOfficeWebApps *bool
    // Primary theme color used in the Company Portal applications and web portal
    themeColor RgbColorable
    // Logo image displayed in Company Portal apps which have a theme color background behind the logo
    themeColorLogo MimeContentable
}
// NewIntuneBrandingProfile instantiates a new intuneBrandingProfile and sets the default values.
func NewIntuneBrandingProfile()(*IntuneBrandingProfile) {
    m := &IntuneBrandingProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIntuneBrandingProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandingProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntuneBrandingProfile(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the branding profile
func (m *IntuneBrandingProfile) GetAssignments()([]IntuneBrandingProfileAssignmentable) {
    return m.assignments
}
// GetCompanyPortalBlockedActions gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable) {
    return m.companyPortalBlockedActions
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITEmailAddress()(*string) {
    return m.contactITEmailAddress
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITName()(*string) {
    return m.contactITName
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITNotes()(*string) {
    return m.contactITNotes
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITPhoneNumber()(*string) {
    return m.contactITPhoneNumber
}
// GetCreatedDateTime gets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomCanSeePrivacyMessage gets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) GetCustomCanSeePrivacyMessage()(*string) {
    return m.customCanSeePrivacyMessage
}
// GetCustomCantSeePrivacyMessage gets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomCantSeePrivacyMessage()(*string) {
    return m.customCantSeePrivacyMessage
}
// GetCustomPrivacyMessage gets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomPrivacyMessage()(*string) {
    return m.customPrivacyMessage
}
// GetDisableClientTelemetry gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) GetDisableClientTelemetry()(*bool) {
    return m.disableClientTelemetry
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrollmentAvailability gets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrandingProfile) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    return m.enrollmentAvailability
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrandingProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntuneBrandingProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntuneBrandingProfileAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(IntuneBrandingProfileAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
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
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["isDefaultProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultProfile(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["profileDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileDescription(val)
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
    res["themeColorLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColorLogo(val.(MimeContentable))
        }
        return nil
    }
    return res
}
// GetIsDefaultProfile gets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) GetIsDefaultProfile()(*bool) {
    return m.isDefaultProfile
}
// GetIsFactoryResetDisabled gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsFactoryResetDisabled()(*bool) {
    return m.isFactoryResetDisabled
}
// GetIsRemoveDeviceDisabled gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsRemoveDeviceDisabled()(*bool) {
    return m.isRemoveDeviceDisabled
}
// GetLandingPageCustomizedImage gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) GetLandingPageCustomizedImage()(MimeContentable) {
    return m.landingPageCustomizedImage
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) GetLightBackgroundLogo()(MimeContentable) {
    return m.lightBackgroundLogo
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteName()(*string) {
    return m.onlineSupportSiteName
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteUrl()(*string) {
    return m.onlineSupportSiteUrl
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) GetPrivacyUrl()(*string) {
    return m.privacyUrl
}
// GetProfileDescription gets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) GetProfileDescription()(*string) {
    return m.profileDescription
}
// GetProfileName gets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) GetProfileName()(*string) {
    return m.profileName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSendDeviceOwnershipChangePushNotification gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrandingProfile) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    return m.sendDeviceOwnershipChangePushNotification
}
// GetShowAzureADEnterpriseApps gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowAzureADEnterpriseApps()(*bool) {
    return m.showAzureADEnterpriseApps
}
// GetShowConfigurationManagerApps gets the showConfigurationManagerApps property value. Boolean that indicates if Configuration Manager Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowConfigurationManagerApps()(*bool) {
    return m.showConfigurationManagerApps
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) GetShowDisplayNameNextToLogo()(*bool) {
    return m.showDisplayNameNextToLogo
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) GetShowLogo()(*bool) {
    return m.showLogo
}
// GetShowOfficeWebApps gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowOfficeWebApps()(*bool) {
    return m.showOfficeWebApps
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) GetThemeColor()(RgbColorable) {
    return m.themeColor
}
// GetThemeColorLogo gets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) GetThemeColorLogo()(MimeContentable) {
    return m.themeColorLogo
}
// Serialize serializes information the current object
func (m *IntuneBrandingProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompanyPortalBlockedActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyPortalBlockedActions()))
        for i, v := range m.GetCompanyPortalBlockedActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("companyPortalBlockedActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITEmailAddress", m.GetContactITEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITName", m.GetContactITName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITNotes", m.GetContactITNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITPhoneNumber", m.GetContactITPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCanSeePrivacyMessage", m.GetCustomCanSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCantSeePrivacyMessage", m.GetCustomCantSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyMessage", m.GetCustomPrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableClientTelemetry", m.GetDisableClientTelemetry())
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
    if m.GetEnrollmentAvailability() != nil {
        cast := (*m.GetEnrollmentAvailability()).String()
        err = writer.WriteStringValue("enrollmentAvailability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultProfile", m.GetIsDefaultProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFactoryResetDisabled", m.GetIsFactoryResetDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRemoveDeviceDisabled", m.GetIsRemoveDeviceDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("landingPageCustomizedImage", m.GetLandingPageCustomizedImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lightBackgroundLogo", m.GetLightBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineSupportSiteName", m.GetOnlineSupportSiteName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineSupportSiteUrl", m.GetOnlineSupportSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privacyUrl", m.GetPrivacyUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileDescription", m.GetProfileDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileName", m.GetProfileName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sendDeviceOwnershipChangePushNotification", m.GetSendDeviceOwnershipChangePushNotification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showAzureADEnterpriseApps", m.GetShowAzureADEnterpriseApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showConfigurationManagerApps", m.GetShowConfigurationManagerApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showDisplayNameNextToLogo", m.GetShowDisplayNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showLogo", m.GetShowLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showOfficeWebApps", m.GetShowOfficeWebApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("themeColor", m.GetThemeColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("themeColorLogo", m.GetThemeColorLogo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the branding profile
func (m *IntuneBrandingProfile) SetAssignments(value []IntuneBrandingProfileAssignmentable)() {
    m.assignments = value
}
// SetCompanyPortalBlockedActions sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)() {
    m.companyPortalBlockedActions = value
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITName(value *string)() {
    m.contactITName = value
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
// SetCreatedDateTime sets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomCanSeePrivacyMessage sets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) SetCustomCanSeePrivacyMessage(value *string)() {
    m.customCanSeePrivacyMessage = value
}
// SetCustomCantSeePrivacyMessage sets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) SetCustomCantSeePrivacyMessage(value *string)() {
    m.customCantSeePrivacyMessage = value
}
// SetCustomPrivacyMessage sets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) SetCustomPrivacyMessage(value *string)() {
    m.customPrivacyMessage = value
}
// SetDisableClientTelemetry sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) SetDisableClientTelemetry(value *bool)() {
    m.disableClientTelemetry = value
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrollmentAvailability sets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrandingProfile) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    m.enrollmentAvailability = value
}
// SetIsDefaultProfile sets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) SetIsDefaultProfile(value *bool)() {
    m.isDefaultProfile = value
}
// SetIsFactoryResetDisabled sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) SetIsFactoryResetDisabled(value *bool)() {
    m.isFactoryResetDisabled = value
}
// SetIsRemoveDeviceDisabled sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrandingProfile) SetIsRemoveDeviceDisabled(value *bool)() {
    m.isRemoveDeviceDisabled = value
}
// SetLandingPageCustomizedImage sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) SetLandingPageCustomizedImage(value MimeContentable)() {
    m.landingPageCustomizedImage = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) SetLightBackgroundLogo(value MimeContentable)() {
    m.lightBackgroundLogo = value
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
// SetProfileDescription sets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) SetProfileDescription(value *string)() {
    m.profileDescription = value
}
// SetProfileName sets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) SetProfileName(value *string)() {
    m.profileName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSendDeviceOwnershipChangePushNotification sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrandingProfile) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    m.sendDeviceOwnershipChangePushNotification = value
}
// SetShowAzureADEnterpriseApps sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowAzureADEnterpriseApps(value *bool)() {
    m.showAzureADEnterpriseApps = value
}
// SetShowConfigurationManagerApps sets the showConfigurationManagerApps property value. Boolean that indicates if Configuration Manager Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowConfigurationManagerApps(value *bool)() {
    m.showConfigurationManagerApps = value
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) SetShowLogo(value *bool)() {
    m.showLogo = value
}
// SetShowOfficeWebApps sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowOfficeWebApps(value *bool)() {
    m.showOfficeWebApps = value
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) SetThemeColor(value RgbColorable)() {
    m.themeColor = value
}
// SetThemeColorLogo sets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) SetThemeColorLogo(value MimeContentable)() {
    m.themeColorLogo = value
}
