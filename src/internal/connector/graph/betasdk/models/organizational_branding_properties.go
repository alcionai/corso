package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalBrandingProperties 
type OrganizationalBrandingProperties struct {
    Entity
    // Color that appears in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
    backgroundColor *string
    // Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
    backgroundImage []byte
    // A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    backgroundImageRelativeUrl *string
    // A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
    bannerLogo []byte
    // A relative URL for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
    bannerLogoRelativeUrl *string
    // A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
    cdnList []string
    // A custom URL for resetting account credentials. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
    customAccountResetCredentialsUrl *string
    // A string to replace the default 'Can't access your account?' self-service password reset (SSPR) hyperlink text on the sign-in page. This text must be in Unicode format and not exceed 256 characters.
    customCannotAccessYourAccountText *string
    // A custom URL to replace the default URL of the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in page. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters. DO NOT USE. Use customAccountResetCredentialsUrl instead.
    customCannotAccessYourAccountUrl *string
    // CSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
    customCSS []byte
    // A relative URL for the customCSS property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    customCSSRelativeUrl *string
    // A string to replace the default 'Forgot my password' hyperlink text on the sign-in form. This text must be in Unicode format and not exceed 256 characters.
    customForgotMyPasswordText *string
    // A string to replace the default 'Privacy and Cookies' hyperlink text in the footer. This text must be in Unicode format and not exceed 256 characters.
    customPrivacyAndCookiesText *string
    // A custom URL to replace the default URL of the 'Privacy and Cookies' hyperlink in the footer. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
    customPrivacyAndCookiesUrl *string
    // A string to replace the default 'reset it now' hyperlink text on the sign-in form. This text must be in Unicode format and not exceed 256 characters. DO NOT USE: Customization of the 'reset it now' hyperlink text is currently not supported.
    customResetItNowText *string
    // A string to replace the the default 'Terms of Use' hyperlink text in the footer. This text must be in Unicode format and not exceed 256 characters.
    customTermsOfUseText *string
    // A custom URL to replace the default URL of the 'Terms of Use' hyperlink in the footer. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128characters.
    customTermsOfUseUrl *string
    // A custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
    favicon []byte
    // A relative url for the favicon above that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    faviconRelativeUrl *string
    // The RGB color to apply to customize the color of the header.
    headerBackgroundColor *string
    // A company logo that appears in the header of the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
    headerLogo []byte
    // A relative URL for the headerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
    headerLogoRelativeUrl *string
    // Represents the layout configuration to be displayed on the login page for a tenant.
    loginPageLayoutConfiguration LoginPageLayoutConfigurationable
    // Represents the various texts that can be hidden on the login page for a tenant.
    loginPageTextVisibilitySettings LoginPageTextVisibilitySettingsable
    // Text that appears at the bottom of the sign-in box. Use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be in Unicode format and not exceed 1024 characters.
    signInPageText *string
    // A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
    squareLogo []byte
    // A square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
    squareLogoDark []byte
    // A relative URL for the squareLogoDark property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    squareLogoDarkRelativeUrl *string
    // A relative URL for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    squareLogoRelativeUrl *string
    // A string that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
    usernameHintText *string
}
// NewOrganizationalBrandingProperties instantiates a new organizationalBrandingProperties and sets the default values.
func NewOrganizationalBrandingProperties()(*OrganizationalBrandingProperties) {
    m := &OrganizationalBrandingProperties{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOrganizationalBrandingPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalBrandingPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.organizationalBranding":
                        return NewOrganizationalBranding(), nil
                    case "#microsoft.graph.organizationalBrandingLocalization":
                        return NewOrganizationalBrandingLocalization(), nil
                }
            }
        }
    }
    return NewOrganizationalBrandingProperties(), nil
}
// GetBackgroundColor gets the backgroundColor property value. Color that appears in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) GetBackgroundColor()(*string) {
    return m.backgroundColor
}
// GetBackgroundImage gets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) GetBackgroundImage()([]byte) {
    return m.backgroundImage
}
// GetBackgroundImageRelativeUrl gets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrl()(*string) {
    return m.backgroundImageRelativeUrl
}
// GetBannerLogo gets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetBannerLogo()([]byte) {
    return m.bannerLogo
}
// GetBannerLogoRelativeUrl gets the bannerLogoRelativeUrl property value. A relative URL for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBannerLogoRelativeUrl()(*string) {
    return m.bannerLogoRelativeUrl
}
// GetCdnList gets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) GetCdnList()([]string) {
    return m.cdnList
}
// GetCustomAccountResetCredentialsUrl gets the customAccountResetCredentialsUrl property value. A custom URL for resetting account credentials. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
func (m *OrganizationalBrandingProperties) GetCustomAccountResetCredentialsUrl()(*string) {
    return m.customAccountResetCredentialsUrl
}
// GetCustomCannotAccessYourAccountText gets the customCannotAccessYourAccountText property value. A string to replace the default 'Can't access your account?' self-service password reset (SSPR) hyperlink text on the sign-in page. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) GetCustomCannotAccessYourAccountText()(*string) {
    return m.customCannotAccessYourAccountText
}
// GetCustomCannotAccessYourAccountUrl gets the customCannotAccessYourAccountUrl property value. A custom URL to replace the default URL of the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in page. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters. DO NOT USE. Use customAccountResetCredentialsUrl instead.
func (m *OrganizationalBrandingProperties) GetCustomCannotAccessYourAccountUrl()(*string) {
    return m.customCannotAccessYourAccountUrl
}
// GetCustomCSS gets the customCSS property value. CSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
func (m *OrganizationalBrandingProperties) GetCustomCSS()([]byte) {
    return m.customCSS
}
// GetCustomCSSRelativeUrl gets the customCSSRelativeUrl property value. A relative URL for the customCSS property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetCustomCSSRelativeUrl()(*string) {
    return m.customCSSRelativeUrl
}
// GetCustomForgotMyPasswordText gets the customForgotMyPasswordText property value. A string to replace the default 'Forgot my password' hyperlink text on the sign-in form. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) GetCustomForgotMyPasswordText()(*string) {
    return m.customForgotMyPasswordText
}
// GetCustomPrivacyAndCookiesText gets the customPrivacyAndCookiesText property value. A string to replace the default 'Privacy and Cookies' hyperlink text in the footer. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) GetCustomPrivacyAndCookiesText()(*string) {
    return m.customPrivacyAndCookiesText
}
// GetCustomPrivacyAndCookiesUrl gets the customPrivacyAndCookiesUrl property value. A custom URL to replace the default URL of the 'Privacy and Cookies' hyperlink in the footer. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
func (m *OrganizationalBrandingProperties) GetCustomPrivacyAndCookiesUrl()(*string) {
    return m.customPrivacyAndCookiesUrl
}
// GetCustomResetItNowText gets the customResetItNowText property value. A string to replace the default 'reset it now' hyperlink text on the sign-in form. This text must be in Unicode format and not exceed 256 characters. DO NOT USE: Customization of the 'reset it now' hyperlink text is currently not supported.
func (m *OrganizationalBrandingProperties) GetCustomResetItNowText()(*string) {
    return m.customResetItNowText
}
// GetCustomTermsOfUseText gets the customTermsOfUseText property value. A string to replace the the default 'Terms of Use' hyperlink text in the footer. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) GetCustomTermsOfUseText()(*string) {
    return m.customTermsOfUseText
}
// GetCustomTermsOfUseUrl gets the customTermsOfUseUrl property value. A custom URL to replace the default URL of the 'Terms of Use' hyperlink in the footer. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128characters.
func (m *OrganizationalBrandingProperties) GetCustomTermsOfUseUrl()(*string) {
    return m.customTermsOfUseUrl
}
// GetFavicon gets the favicon property value. A custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
func (m *OrganizationalBrandingProperties) GetFavicon()([]byte) {
    return m.favicon
}
// GetFaviconRelativeUrl gets the faviconRelativeUrl property value. A relative url for the favicon above that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetFaviconRelativeUrl()(*string) {
    return m.faviconRelativeUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalBrandingProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["backgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["backgroundImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImage(val)
        }
        return nil
    }
    res["backgroundImageRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImageRelativeUrl(val)
        }
        return nil
    }
    res["bannerLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogo(val)
        }
        return nil
    }
    res["bannerLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogoRelativeUrl(val)
        }
        return nil
    }
    res["cdnList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCdnList(res)
        }
        return nil
    }
    res["customAccountResetCredentialsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomAccountResetCredentialsUrl(val)
        }
        return nil
    }
    res["customCannotAccessYourAccountText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCannotAccessYourAccountText(val)
        }
        return nil
    }
    res["customCannotAccessYourAccountUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCannotAccessYourAccountUrl(val)
        }
        return nil
    }
    res["customCSS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCSS(val)
        }
        return nil
    }
    res["customCSSRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCSSRelativeUrl(val)
        }
        return nil
    }
    res["customForgotMyPasswordText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomForgotMyPasswordText(val)
        }
        return nil
    }
    res["customPrivacyAndCookiesText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyAndCookiesText(val)
        }
        return nil
    }
    res["customPrivacyAndCookiesUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyAndCookiesUrl(val)
        }
        return nil
    }
    res["customResetItNowText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomResetItNowText(val)
        }
        return nil
    }
    res["customTermsOfUseText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomTermsOfUseText(val)
        }
        return nil
    }
    res["customTermsOfUseUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomTermsOfUseUrl(val)
        }
        return nil
    }
    res["favicon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFavicon(val)
        }
        return nil
    }
    res["faviconRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaviconRelativeUrl(val)
        }
        return nil
    }
    res["headerBackgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderBackgroundColor(val)
        }
        return nil
    }
    res["headerLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderLogo(val)
        }
        return nil
    }
    res["headerLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderLogoRelativeUrl(val)
        }
        return nil
    }
    res["loginPageLayoutConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLoginPageLayoutConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginPageLayoutConfiguration(val.(LoginPageLayoutConfigurationable))
        }
        return nil
    }
    res["loginPageTextVisibilitySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLoginPageTextVisibilitySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginPageTextVisibilitySettings(val.(LoginPageTextVisibilitySettingsable))
        }
        return nil
    }
    res["signInPageText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInPageText(val)
        }
        return nil
    }
    res["squareLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogo(val)
        }
        return nil
    }
    res["squareLogoDark"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoDark(val)
        }
        return nil
    }
    res["squareLogoDarkRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoDarkRelativeUrl(val)
        }
        return nil
    }
    res["squareLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoRelativeUrl(val)
        }
        return nil
    }
    res["usernameHintText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameHintText(val)
        }
        return nil
    }
    return res
}
// GetHeaderBackgroundColor gets the headerBackgroundColor property value. The RGB color to apply to customize the color of the header.
func (m *OrganizationalBrandingProperties) GetHeaderBackgroundColor()(*string) {
    return m.headerBackgroundColor
}
// GetHeaderLogo gets the headerLogo property value. A company logo that appears in the header of the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetHeaderLogo()([]byte) {
    return m.headerLogo
}
// GetHeaderLogoRelativeUrl gets the headerLogoRelativeUrl property value. A relative URL for the headerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetHeaderLogoRelativeUrl()(*string) {
    return m.headerLogoRelativeUrl
}
// GetLoginPageLayoutConfiguration gets the loginPageLayoutConfiguration property value. Represents the layout configuration to be displayed on the login page for a tenant.
func (m *OrganizationalBrandingProperties) GetLoginPageLayoutConfiguration()(LoginPageLayoutConfigurationable) {
    return m.loginPageLayoutConfiguration
}
// GetLoginPageTextVisibilitySettings gets the loginPageTextVisibilitySettings property value. Represents the various texts that can be hidden on the login page for a tenant.
func (m *OrganizationalBrandingProperties) GetLoginPageTextVisibilitySettings()(LoginPageTextVisibilitySettingsable) {
    return m.loginPageTextVisibilitySettings
}
// GetSignInPageText gets the signInPageText property value. Text that appears at the bottom of the sign-in box. Use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be in Unicode format and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) GetSignInPageText()(*string) {
    return m.signInPageText
}
// GetSquareLogo gets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetSquareLogo()([]byte) {
    return m.squareLogo
}
// GetSquareLogoDark gets the squareLogoDark property value. A square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetSquareLogoDark()([]byte) {
    return m.squareLogoDark
}
// GetSquareLogoDarkRelativeUrl gets the squareLogoDarkRelativeUrl property value. A relative URL for the squareLogoDark property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetSquareLogoDarkRelativeUrl()(*string) {
    return m.squareLogoDarkRelativeUrl
}
// GetSquareLogoRelativeUrl gets the squareLogoRelativeUrl property value. A relative URL for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetSquareLogoRelativeUrl()(*string) {
    return m.squareLogoRelativeUrl
}
// GetUsernameHintText gets the usernameHintText property value. A string that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) GetUsernameHintText()(*string) {
    return m.usernameHintText
}
// Serialize serializes information the current object
func (m *OrganizationalBrandingProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("backgroundImage", m.GetBackgroundImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("backgroundImageRelativeUrl", m.GetBackgroundImageRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("bannerLogo", m.GetBannerLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bannerLogoRelativeUrl", m.GetBannerLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCdnList() != nil {
        err = writer.WriteCollectionOfStringValues("cdnList", m.GetCdnList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customAccountResetCredentialsUrl", m.GetCustomAccountResetCredentialsUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCannotAccessYourAccountText", m.GetCustomCannotAccessYourAccountText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCannotAccessYourAccountUrl", m.GetCustomCannotAccessYourAccountUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("customCSS", m.GetCustomCSS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCSSRelativeUrl", m.GetCustomCSSRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customForgotMyPasswordText", m.GetCustomForgotMyPasswordText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyAndCookiesText", m.GetCustomPrivacyAndCookiesText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyAndCookiesUrl", m.GetCustomPrivacyAndCookiesUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customResetItNowText", m.GetCustomResetItNowText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customTermsOfUseText", m.GetCustomTermsOfUseText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customTermsOfUseUrl", m.GetCustomTermsOfUseUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("favicon", m.GetFavicon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faviconRelativeUrl", m.GetFaviconRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerBackgroundColor", m.GetHeaderBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("headerLogo", m.GetHeaderLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerLogoRelativeUrl", m.GetHeaderLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("loginPageLayoutConfiguration", m.GetLoginPageLayoutConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("loginPageTextVisibilitySettings", m.GetLoginPageTextVisibilitySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInPageText", m.GetSignInPageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogo", m.GetSquareLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogoDark", m.GetSquareLogoDark())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoDarkRelativeUrl", m.GetSquareLogoDarkRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoRelativeUrl", m.GetSquareLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameHintText", m.GetUsernameHintText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackgroundColor sets the backgroundColor property value. Color that appears in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// SetBackgroundImage sets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) SetBackgroundImage(value []byte)() {
    m.backgroundImage = value
}
// SetBackgroundImageRelativeUrl sets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(value *string)() {
    m.backgroundImageRelativeUrl = value
}
// SetBannerLogo sets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetBannerLogo(value []byte)() {
    m.bannerLogo = value
}
// SetBannerLogoRelativeUrl sets the bannerLogoRelativeUrl property value. A relative URL for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetBannerLogoRelativeUrl(value *string)() {
    m.bannerLogoRelativeUrl = value
}
// SetCdnList sets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) SetCdnList(value []string)() {
    m.cdnList = value
}
// SetCustomAccountResetCredentialsUrl sets the customAccountResetCredentialsUrl property value. A custom URL for resetting account credentials. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
func (m *OrganizationalBrandingProperties) SetCustomAccountResetCredentialsUrl(value *string)() {
    m.customAccountResetCredentialsUrl = value
}
// SetCustomCannotAccessYourAccountText sets the customCannotAccessYourAccountText property value. A string to replace the default 'Can't access your account?' self-service password reset (SSPR) hyperlink text on the sign-in page. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) SetCustomCannotAccessYourAccountText(value *string)() {
    m.customCannotAccessYourAccountText = value
}
// SetCustomCannotAccessYourAccountUrl sets the customCannotAccessYourAccountUrl property value. A custom URL to replace the default URL of the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in page. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters. DO NOT USE. Use customAccountResetCredentialsUrl instead.
func (m *OrganizationalBrandingProperties) SetCustomCannotAccessYourAccountUrl(value *string)() {
    m.customCannotAccessYourAccountUrl = value
}
// SetCustomCSS sets the customCSS property value. CSS styling that appears on the sign-in page. The allowed format is .css format only and not larger than 25 KB.
func (m *OrganizationalBrandingProperties) SetCustomCSS(value []byte)() {
    m.customCSS = value
}
// SetCustomCSSRelativeUrl sets the customCSSRelativeUrl property value. A relative URL for the customCSS property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetCustomCSSRelativeUrl(value *string)() {
    m.customCSSRelativeUrl = value
}
// SetCustomForgotMyPasswordText sets the customForgotMyPasswordText property value. A string to replace the default 'Forgot my password' hyperlink text on the sign-in form. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) SetCustomForgotMyPasswordText(value *string)() {
    m.customForgotMyPasswordText = value
}
// SetCustomPrivacyAndCookiesText sets the customPrivacyAndCookiesText property value. A string to replace the default 'Privacy and Cookies' hyperlink text in the footer. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) SetCustomPrivacyAndCookiesText(value *string)() {
    m.customPrivacyAndCookiesText = value
}
// SetCustomPrivacyAndCookiesUrl sets the customPrivacyAndCookiesUrl property value. A custom URL to replace the default URL of the 'Privacy and Cookies' hyperlink in the footer. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128 characters.
func (m *OrganizationalBrandingProperties) SetCustomPrivacyAndCookiesUrl(value *string)() {
    m.customPrivacyAndCookiesUrl = value
}
// SetCustomResetItNowText sets the customResetItNowText property value. A string to replace the default 'reset it now' hyperlink text on the sign-in form. This text must be in Unicode format and not exceed 256 characters. DO NOT USE: Customization of the 'reset it now' hyperlink text is currently not supported.
func (m *OrganizationalBrandingProperties) SetCustomResetItNowText(value *string)() {
    m.customResetItNowText = value
}
// SetCustomTermsOfUseText sets the customTermsOfUseText property value. A string to replace the the default 'Terms of Use' hyperlink text in the footer. This text must be in Unicode format and not exceed 256 characters.
func (m *OrganizationalBrandingProperties) SetCustomTermsOfUseText(value *string)() {
    m.customTermsOfUseText = value
}
// SetCustomTermsOfUseUrl sets the customTermsOfUseUrl property value. A custom URL to replace the default URL of the 'Terms of Use' hyperlink in the footer. This URL must be in ASCII format or non-ASCII characters must be URL encoded, and not exceed 128characters.
func (m *OrganizationalBrandingProperties) SetCustomTermsOfUseUrl(value *string)() {
    m.customTermsOfUseUrl = value
}
// SetFavicon sets the favicon property value. A custom icon (favicon) to replace a default Microsoft product favicon on an Azure AD tenant.
func (m *OrganizationalBrandingProperties) SetFavicon(value []byte)() {
    m.favicon = value
}
// SetFaviconRelativeUrl sets the faviconRelativeUrl property value. A relative url for the favicon above that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetFaviconRelativeUrl(value *string)() {
    m.faviconRelativeUrl = value
}
// SetHeaderBackgroundColor sets the headerBackgroundColor property value. The RGB color to apply to customize the color of the header.
func (m *OrganizationalBrandingProperties) SetHeaderBackgroundColor(value *string)() {
    m.headerBackgroundColor = value
}
// SetHeaderLogo sets the headerLogo property value. A company logo that appears in the header of the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetHeaderLogo(value []byte)() {
    m.headerLogo = value
}
// SetHeaderLogoRelativeUrl sets the headerLogoRelativeUrl property value. A relative URL for the headerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetHeaderLogoRelativeUrl(value *string)() {
    m.headerLogoRelativeUrl = value
}
// SetLoginPageLayoutConfiguration sets the loginPageLayoutConfiguration property value. Represents the layout configuration to be displayed on the login page for a tenant.
func (m *OrganizationalBrandingProperties) SetLoginPageLayoutConfiguration(value LoginPageLayoutConfigurationable)() {
    m.loginPageLayoutConfiguration = value
}
// SetLoginPageTextVisibilitySettings sets the loginPageTextVisibilitySettings property value. Represents the various texts that can be hidden on the login page for a tenant.
func (m *OrganizationalBrandingProperties) SetLoginPageTextVisibilitySettings(value LoginPageTextVisibilitySettingsable)() {
    m.loginPageTextVisibilitySettings = value
}
// SetSignInPageText sets the signInPageText property value. Text that appears at the bottom of the sign-in box. Use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be in Unicode format and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) SetSignInPageText(value *string)() {
    m.signInPageText = value
}
// SetSquareLogo sets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetSquareLogo(value []byte)() {
    m.squareLogo = value
}
// SetSquareLogoDark sets the squareLogoDark property value. A square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetSquareLogoDark(value []byte)() {
    m.squareLogoDark = value
}
// SetSquareLogoDarkRelativeUrl sets the squareLogoDarkRelativeUrl property value. A relative URL for the squareLogoDark property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetSquareLogoDarkRelativeUrl(value *string)() {
    m.squareLogoDarkRelativeUrl = value
}
// SetSquareLogoRelativeUrl sets the squareLogoRelativeUrl property value. A relative URL for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetSquareLogoRelativeUrl(value *string)() {
    m.squareLogoRelativeUrl = value
}
// SetUsernameHintText sets the usernameHintText property value. A string that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) SetUsernameHintText(value *string)() {
    m.usernameHintText = value
}
