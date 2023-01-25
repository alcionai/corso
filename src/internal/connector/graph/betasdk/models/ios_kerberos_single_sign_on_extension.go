package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosKerberosSingleSignOnExtension 
type IosKerberosSingleSignOnExtension struct {
    IosSingleSignOnExtension
    // Gets or sets the Active Directory site.
    activeDirectorySiteCode *string
    // Enables or disables whether the Kerberos extension can automatically determine its site name.
    blockActiveDirectorySiteAutoDiscovery *bool
    // Enables or disables Keychain usage.
    blockAutomaticLogin *bool
    // Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
    cacheName *string
    // Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
    credentialBundleIdAccessControlList []string
    // Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
    domainRealms []string
    // Gets or sets a list of hosts or domain names for which the app extension performs SSO.
    domains []string
    // When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
    isDefaultRealm *bool
    // When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
    managedAppsInBundleIdACLIncluded *bool
    // Enables or disables password changes.
    passwordBlockModification *bool
    // Gets or sets the URL that the user will be sent to when they initiate a password change.
    passwordChangeUrl *string
    // Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
    passwordEnableLocalSync *bool
    // Overrides the default password expiration in days. For most domains, this value is calculated automatically.
    passwordExpirationDays *int32
    // Gets or sets the number of days until the user is notified that their password will expire (default is 15).
    passwordExpirationNotificationDays *int32
    // Gets or sets the minimum number of days until a user can change their password again.
    passwordMinimumAgeDays *int32
    // Gets or sets the minimum length of a password.
    passwordMinimumLength *int32
    // Gets or sets the number of previous passwords to block.
    passwordPreviousPasswordBlockCount *int32
    // Enables or disables whether passwords must meet Active Directory's complexity requirements.
    passwordRequireActiveDirectoryComplexity *bool
    // Gets or sets a description of the password complexity requirements.
    passwordRequirementsDescription *string
    // Gets or sets the case-sensitive realm name for this profile.
    realm *string
    // Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
    requireUserPresence *bool
    // Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
    signInHelpText *string
    // Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
    userPrincipalName *string
}
// NewIosKerberosSingleSignOnExtension instantiates a new IosKerberosSingleSignOnExtension and sets the default values.
func NewIosKerberosSingleSignOnExtension()(*IosKerberosSingleSignOnExtension) {
    m := &IosKerberosSingleSignOnExtension{
        IosSingleSignOnExtension: *NewIosSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.iosKerberosSingleSignOnExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosKerberosSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosKerberosSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosKerberosSingleSignOnExtension(), nil
}
// GetActiveDirectorySiteCode gets the activeDirectorySiteCode property value. Gets or sets the Active Directory site.
func (m *IosKerberosSingleSignOnExtension) GetActiveDirectorySiteCode()(*string) {
    return m.activeDirectorySiteCode
}
// GetBlockActiveDirectorySiteAutoDiscovery gets the blockActiveDirectorySiteAutoDiscovery property value. Enables or disables whether the Kerberos extension can automatically determine its site name.
func (m *IosKerberosSingleSignOnExtension) GetBlockActiveDirectorySiteAutoDiscovery()(*bool) {
    return m.blockActiveDirectorySiteAutoDiscovery
}
// GetBlockAutomaticLogin gets the blockAutomaticLogin property value. Enables or disables Keychain usage.
func (m *IosKerberosSingleSignOnExtension) GetBlockAutomaticLogin()(*bool) {
    return m.blockAutomaticLogin
}
// GetCacheName gets the cacheName property value. Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
func (m *IosKerberosSingleSignOnExtension) GetCacheName()(*string) {
    return m.cacheName
}
// GetCredentialBundleIdAccessControlList gets the credentialBundleIdAccessControlList property value. Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
func (m *IosKerberosSingleSignOnExtension) GetCredentialBundleIdAccessControlList()([]string) {
    return m.credentialBundleIdAccessControlList
}
// GetDomainRealms gets the domainRealms property value. Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
func (m *IosKerberosSingleSignOnExtension) GetDomainRealms()([]string) {
    return m.domainRealms
}
// GetDomains gets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *IosKerberosSingleSignOnExtension) GetDomains()([]string) {
    return m.domains
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosKerberosSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosSingleSignOnExtension.GetFieldDeserializers()
    res["activeDirectorySiteCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDirectorySiteCode(val)
        }
        return nil
    }
    res["blockActiveDirectorySiteAutoDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockActiveDirectorySiteAutoDiscovery(val)
        }
        return nil
    }
    res["blockAutomaticLogin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockAutomaticLogin(val)
        }
        return nil
    }
    res["cacheName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheName(val)
        }
        return nil
    }
    res["credentialBundleIdAccessControlList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCredentialBundleIdAccessControlList(res)
        }
        return nil
    }
    res["domainRealms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDomainRealms(res)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["isDefaultRealm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultRealm(val)
        }
        return nil
    }
    res["managedAppsInBundleIdACLIncluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedAppsInBundleIdACLIncluded(val)
        }
        return nil
    }
    res["passwordBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockModification(val)
        }
        return nil
    }
    res["passwordChangeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordChangeUrl(val)
        }
        return nil
    }
    res["passwordEnableLocalSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordEnableLocalSync(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordExpirationNotificationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationNotificationDays(val)
        }
        return nil
    }
    res["passwordMinimumAgeDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumAgeDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequireActiveDirectoryComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequireActiveDirectoryComplexity(val)
        }
        return nil
    }
    res["passwordRequirementsDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequirementsDescription(val)
        }
        return nil
    }
    res["realm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealm(val)
        }
        return nil
    }
    res["requireUserPresence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireUserPresence(val)
        }
        return nil
    }
    res["signInHelpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInHelpText(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultRealm gets the isDefaultRealm property value. When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
func (m *IosKerberosSingleSignOnExtension) GetIsDefaultRealm()(*bool) {
    return m.isDefaultRealm
}
// GetManagedAppsInBundleIdACLIncluded gets the managedAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
func (m *IosKerberosSingleSignOnExtension) GetManagedAppsInBundleIdACLIncluded()(*bool) {
    return m.managedAppsInBundleIdACLIncluded
}
// GetPasswordBlockModification gets the passwordBlockModification property value. Enables or disables password changes.
func (m *IosKerberosSingleSignOnExtension) GetPasswordBlockModification()(*bool) {
    return m.passwordBlockModification
}
// GetPasswordChangeUrl gets the passwordChangeUrl property value. Gets or sets the URL that the user will be sent to when they initiate a password change.
func (m *IosKerberosSingleSignOnExtension) GetPasswordChangeUrl()(*string) {
    return m.passwordChangeUrl
}
// GetPasswordEnableLocalSync gets the passwordEnableLocalSync property value. Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
func (m *IosKerberosSingleSignOnExtension) GetPasswordEnableLocalSync()(*bool) {
    return m.passwordEnableLocalSync
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Overrides the default password expiration in days. For most domains, this value is calculated automatically.
func (m *IosKerberosSingleSignOnExtension) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordExpirationNotificationDays gets the passwordExpirationNotificationDays property value. Gets or sets the number of days until the user is notified that their password will expire (default is 15).
func (m *IosKerberosSingleSignOnExtension) GetPasswordExpirationNotificationDays()(*int32) {
    return m.passwordExpirationNotificationDays
}
// GetPasswordMinimumAgeDays gets the passwordMinimumAgeDays property value. Gets or sets the minimum number of days until a user can change their password again.
func (m *IosKerberosSingleSignOnExtension) GetPasswordMinimumAgeDays()(*int32) {
    return m.passwordMinimumAgeDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Gets or sets the minimum length of a password.
func (m *IosKerberosSingleSignOnExtension) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Gets or sets the number of previous passwords to block.
func (m *IosKerberosSingleSignOnExtension) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequireActiveDirectoryComplexity gets the passwordRequireActiveDirectoryComplexity property value. Enables or disables whether passwords must meet Active Directory's complexity requirements.
func (m *IosKerberosSingleSignOnExtension) GetPasswordRequireActiveDirectoryComplexity()(*bool) {
    return m.passwordRequireActiveDirectoryComplexity
}
// GetPasswordRequirementsDescription gets the passwordRequirementsDescription property value. Gets or sets a description of the password complexity requirements.
func (m *IosKerberosSingleSignOnExtension) GetPasswordRequirementsDescription()(*string) {
    return m.passwordRequirementsDescription
}
// GetRealm gets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *IosKerberosSingleSignOnExtension) GetRealm()(*string) {
    return m.realm
}
// GetRequireUserPresence gets the requireUserPresence property value. Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
func (m *IosKerberosSingleSignOnExtension) GetRequireUserPresence()(*bool) {
    return m.requireUserPresence
}
// GetSignInHelpText gets the signInHelpText property value. Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
func (m *IosKerberosSingleSignOnExtension) GetSignInHelpText()(*string) {
    return m.signInHelpText
}
// GetUserPrincipalName gets the userPrincipalName property value. Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
func (m *IosKerberosSingleSignOnExtension) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *IosKerberosSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activeDirectorySiteCode", m.GetActiveDirectorySiteCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockActiveDirectorySiteAutoDiscovery", m.GetBlockActiveDirectorySiteAutoDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockAutomaticLogin", m.GetBlockAutomaticLogin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cacheName", m.GetCacheName())
        if err != nil {
            return err
        }
    }
    if m.GetCredentialBundleIdAccessControlList() != nil {
        err = writer.WriteCollectionOfStringValues("credentialBundleIdAccessControlList", m.GetCredentialBundleIdAccessControlList())
        if err != nil {
            return err
        }
    }
    if m.GetDomainRealms() != nil {
        err = writer.WriteCollectionOfStringValues("domainRealms", m.GetDomainRealms())
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err = writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultRealm", m.GetIsDefaultRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedAppsInBundleIdACLIncluded", m.GetManagedAppsInBundleIdACLIncluded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockModification", m.GetPasswordBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordChangeUrl", m.GetPasswordChangeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordEnableLocalSync", m.GetPasswordEnableLocalSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationNotificationDays", m.GetPasswordExpirationNotificationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumAgeDays", m.GetPasswordMinimumAgeDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequireActiveDirectoryComplexity", m.GetPasswordRequireActiveDirectoryComplexity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordRequirementsDescription", m.GetPasswordRequirementsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("realm", m.GetRealm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireUserPresence", m.GetRequireUserPresence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInHelpText", m.GetSignInHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDirectorySiteCode sets the activeDirectorySiteCode property value. Gets or sets the Active Directory site.
func (m *IosKerberosSingleSignOnExtension) SetActiveDirectorySiteCode(value *string)() {
    m.activeDirectorySiteCode = value
}
// SetBlockActiveDirectorySiteAutoDiscovery sets the blockActiveDirectorySiteAutoDiscovery property value. Enables or disables whether the Kerberos extension can automatically determine its site name.
func (m *IosKerberosSingleSignOnExtension) SetBlockActiveDirectorySiteAutoDiscovery(value *bool)() {
    m.blockActiveDirectorySiteAutoDiscovery = value
}
// SetBlockAutomaticLogin sets the blockAutomaticLogin property value. Enables or disables Keychain usage.
func (m *IosKerberosSingleSignOnExtension) SetBlockAutomaticLogin(value *bool)() {
    m.blockAutomaticLogin = value
}
// SetCacheName sets the cacheName property value. Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
func (m *IosKerberosSingleSignOnExtension) SetCacheName(value *string)() {
    m.cacheName = value
}
// SetCredentialBundleIdAccessControlList sets the credentialBundleIdAccessControlList property value. Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
func (m *IosKerberosSingleSignOnExtension) SetCredentialBundleIdAccessControlList(value []string)() {
    m.credentialBundleIdAccessControlList = value
}
// SetDomainRealms sets the domainRealms property value. Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
func (m *IosKerberosSingleSignOnExtension) SetDomainRealms(value []string)() {
    m.domainRealms = value
}
// SetDomains sets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *IosKerberosSingleSignOnExtension) SetDomains(value []string)() {
    m.domains = value
}
// SetIsDefaultRealm sets the isDefaultRealm property value. When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
func (m *IosKerberosSingleSignOnExtension) SetIsDefaultRealm(value *bool)() {
    m.isDefaultRealm = value
}
// SetManagedAppsInBundleIdACLIncluded sets the managedAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
func (m *IosKerberosSingleSignOnExtension) SetManagedAppsInBundleIdACLIncluded(value *bool)() {
    m.managedAppsInBundleIdACLIncluded = value
}
// SetPasswordBlockModification sets the passwordBlockModification property value. Enables or disables password changes.
func (m *IosKerberosSingleSignOnExtension) SetPasswordBlockModification(value *bool)() {
    m.passwordBlockModification = value
}
// SetPasswordChangeUrl sets the passwordChangeUrl property value. Gets or sets the URL that the user will be sent to when they initiate a password change.
func (m *IosKerberosSingleSignOnExtension) SetPasswordChangeUrl(value *string)() {
    m.passwordChangeUrl = value
}
// SetPasswordEnableLocalSync sets the passwordEnableLocalSync property value. Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
func (m *IosKerberosSingleSignOnExtension) SetPasswordEnableLocalSync(value *bool)() {
    m.passwordEnableLocalSync = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Overrides the default password expiration in days. For most domains, this value is calculated automatically.
func (m *IosKerberosSingleSignOnExtension) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordExpirationNotificationDays sets the passwordExpirationNotificationDays property value. Gets or sets the number of days until the user is notified that their password will expire (default is 15).
func (m *IosKerberosSingleSignOnExtension) SetPasswordExpirationNotificationDays(value *int32)() {
    m.passwordExpirationNotificationDays = value
}
// SetPasswordMinimumAgeDays sets the passwordMinimumAgeDays property value. Gets or sets the minimum number of days until a user can change their password again.
func (m *IosKerberosSingleSignOnExtension) SetPasswordMinimumAgeDays(value *int32)() {
    m.passwordMinimumAgeDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Gets or sets the minimum length of a password.
func (m *IosKerberosSingleSignOnExtension) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Gets or sets the number of previous passwords to block.
func (m *IosKerberosSingleSignOnExtension) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequireActiveDirectoryComplexity sets the passwordRequireActiveDirectoryComplexity property value. Enables or disables whether passwords must meet Active Directory's complexity requirements.
func (m *IosKerberosSingleSignOnExtension) SetPasswordRequireActiveDirectoryComplexity(value *bool)() {
    m.passwordRequireActiveDirectoryComplexity = value
}
// SetPasswordRequirementsDescription sets the passwordRequirementsDescription property value. Gets or sets a description of the password complexity requirements.
func (m *IosKerberosSingleSignOnExtension) SetPasswordRequirementsDescription(value *string)() {
    m.passwordRequirementsDescription = value
}
// SetRealm sets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *IosKerberosSingleSignOnExtension) SetRealm(value *string)() {
    m.realm = value
}
// SetRequireUserPresence sets the requireUserPresence property value. Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
func (m *IosKerberosSingleSignOnExtension) SetRequireUserPresence(value *bool)() {
    m.requireUserPresence = value
}
// SetSignInHelpText sets the signInHelpText property value. Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
func (m *IosKerberosSingleSignOnExtension) SetSignInHelpText(value *string)() {
    m.signInHelpText = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
func (m *IosKerberosSingleSignOnExtension) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
