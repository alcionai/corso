package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSKerberosSingleSignOnExtension 
type MacOSKerberosSingleSignOnExtension struct {
    MacOSSingleSignOnExtension
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
    // When set to True, the credential is requested on the next matching Kerberos challenge or network state change. When the credential is expired or missing, a new credential is created. Available for devices running macOS versions 12 and later.
    credentialsCacheMonitored *bool
    // Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
    domainRealms []string
    // Gets or sets a list of hosts or domain names for which the app extension performs SSO.
    domains []string
    // When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
    isDefaultRealm *bool
    // When set to True, the Kerberos extension allows any apps entered with the app bundle ID, managed apps, and standard Kerberos utilities, such as TicketViewer and klist, to access and use the credential. Available for devices running macOS versions 12 and later.
    kerberosAppsInBundleIdACLIncluded *bool
    // When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
    managedAppsInBundleIdACLIncluded *bool
    // Select how other processes use the Kerberos Extension credential.
    modeCredentialUsed *string
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
    // Add creates an ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. This list is used when the servers are not discoverable using DNS. When the servers are discoverable, the list is used for both connectivity checks, and used first for Kerberos traffic. If the servers don’t respond, then the device uses DNS discovery. Delete removes an existing list, and devices use DNS discovery. Available for devices running macOS versions 12 and later.
    preferredKDCs []string
    // Gets or sets the case-sensitive realm name for this profile.
    realm *string
    // Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
    requireUserPresence *bool
    // Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
    signInHelpText *string
    // When set to True, LDAP connections are required to use Transport Layer Security (TLS). Available for devices running macOS versions 11 and later.
    tlsForLDAPRequired *bool
    // This label replaces the user name shown in the Kerberos extension. You can enter a name to match the name of your company or organization. Available for devices running macOS versions 11 and later.
    usernameLabelCustom *string
    // Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
    userPrincipalName *string
    // When set to True, the user isn’t prompted to set up the Kerberos extension until the extension is enabled by the admin, or a Kerberos challenge is received. Available for devices running macOS versions 11 and later.
    userSetupDelayed *bool
}
// NewMacOSKerberosSingleSignOnExtension instantiates a new MacOSKerberosSingleSignOnExtension and sets the default values.
func NewMacOSKerberosSingleSignOnExtension()(*MacOSKerberosSingleSignOnExtension) {
    m := &MacOSKerberosSingleSignOnExtension{
        MacOSSingleSignOnExtension: *NewMacOSSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.macOSKerberosSingleSignOnExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSKerberosSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSKerberosSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSKerberosSingleSignOnExtension(), nil
}
// GetActiveDirectorySiteCode gets the activeDirectorySiteCode property value. Gets or sets the Active Directory site.
func (m *MacOSKerberosSingleSignOnExtension) GetActiveDirectorySiteCode()(*string) {
    return m.activeDirectorySiteCode
}
// GetBlockActiveDirectorySiteAutoDiscovery gets the blockActiveDirectorySiteAutoDiscovery property value. Enables or disables whether the Kerberos extension can automatically determine its site name.
func (m *MacOSKerberosSingleSignOnExtension) GetBlockActiveDirectorySiteAutoDiscovery()(*bool) {
    return m.blockActiveDirectorySiteAutoDiscovery
}
// GetBlockAutomaticLogin gets the blockAutomaticLogin property value. Enables or disables Keychain usage.
func (m *MacOSKerberosSingleSignOnExtension) GetBlockAutomaticLogin()(*bool) {
    return m.blockAutomaticLogin
}
// GetCacheName gets the cacheName property value. Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
func (m *MacOSKerberosSingleSignOnExtension) GetCacheName()(*string) {
    return m.cacheName
}
// GetCredentialBundleIdAccessControlList gets the credentialBundleIdAccessControlList property value. Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
func (m *MacOSKerberosSingleSignOnExtension) GetCredentialBundleIdAccessControlList()([]string) {
    return m.credentialBundleIdAccessControlList
}
// GetCredentialsCacheMonitored gets the credentialsCacheMonitored property value. When set to True, the credential is requested on the next matching Kerberos challenge or network state change. When the credential is expired or missing, a new credential is created. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetCredentialsCacheMonitored()(*bool) {
    return m.credentialsCacheMonitored
}
// GetDomainRealms gets the domainRealms property value. Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
func (m *MacOSKerberosSingleSignOnExtension) GetDomainRealms()([]string) {
    return m.domainRealms
}
// GetDomains gets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *MacOSKerberosSingleSignOnExtension) GetDomains()([]string) {
    return m.domains
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSKerberosSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSSingleSignOnExtension.GetFieldDeserializers()
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
    res["credentialsCacheMonitored"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialsCacheMonitored(val)
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
    res["kerberosAppsInBundleIdACLIncluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosAppsInBundleIdACLIncluded(val)
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
    res["modeCredentialUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModeCredentialUsed(val)
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
    res["preferredKDCs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPreferredKDCs(res)
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
    res["tlsForLDAPRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsForLDAPRequired(val)
        }
        return nil
    }
    res["usernameLabelCustom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameLabelCustom(val)
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
    res["userSetupDelayed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSetupDelayed(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultRealm gets the isDefaultRealm property value. When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
func (m *MacOSKerberosSingleSignOnExtension) GetIsDefaultRealm()(*bool) {
    return m.isDefaultRealm
}
// GetKerberosAppsInBundleIdACLIncluded gets the kerberosAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows any apps entered with the app bundle ID, managed apps, and standard Kerberos utilities, such as TicketViewer and klist, to access and use the credential. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetKerberosAppsInBundleIdACLIncluded()(*bool) {
    return m.kerberosAppsInBundleIdACLIncluded
}
// GetManagedAppsInBundleIdACLIncluded gets the managedAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetManagedAppsInBundleIdACLIncluded()(*bool) {
    return m.managedAppsInBundleIdACLIncluded
}
// GetModeCredentialUsed gets the modeCredentialUsed property value. Select how other processes use the Kerberos Extension credential.
func (m *MacOSKerberosSingleSignOnExtension) GetModeCredentialUsed()(*string) {
    return m.modeCredentialUsed
}
// GetPasswordBlockModification gets the passwordBlockModification property value. Enables or disables password changes.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordBlockModification()(*bool) {
    return m.passwordBlockModification
}
// GetPasswordChangeUrl gets the passwordChangeUrl property value. Gets or sets the URL that the user will be sent to when they initiate a password change.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordChangeUrl()(*string) {
    return m.passwordChangeUrl
}
// GetPasswordEnableLocalSync gets the passwordEnableLocalSync property value. Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordEnableLocalSync()(*bool) {
    return m.passwordEnableLocalSync
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Overrides the default password expiration in days. For most domains, this value is calculated automatically.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordExpirationNotificationDays gets the passwordExpirationNotificationDays property value. Gets or sets the number of days until the user is notified that their password will expire (default is 15).
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordExpirationNotificationDays()(*int32) {
    return m.passwordExpirationNotificationDays
}
// GetPasswordMinimumAgeDays gets the passwordMinimumAgeDays property value. Gets or sets the minimum number of days until a user can change their password again.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordMinimumAgeDays()(*int32) {
    return m.passwordMinimumAgeDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Gets or sets the minimum length of a password.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Gets or sets the number of previous passwords to block.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequireActiveDirectoryComplexity gets the passwordRequireActiveDirectoryComplexity property value. Enables or disables whether passwords must meet Active Directory's complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordRequireActiveDirectoryComplexity()(*bool) {
    return m.passwordRequireActiveDirectoryComplexity
}
// GetPasswordRequirementsDescription gets the passwordRequirementsDescription property value. Gets or sets a description of the password complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) GetPasswordRequirementsDescription()(*string) {
    return m.passwordRequirementsDescription
}
// GetPreferredKDCs gets the preferredKDCs property value. Add creates an ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. This list is used when the servers are not discoverable using DNS. When the servers are discoverable, the list is used for both connectivity checks, and used first for Kerberos traffic. If the servers don’t respond, then the device uses DNS discovery. Delete removes an existing list, and devices use DNS discovery. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetPreferredKDCs()([]string) {
    return m.preferredKDCs
}
// GetRealm gets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *MacOSKerberosSingleSignOnExtension) GetRealm()(*string) {
    return m.realm
}
// GetRequireUserPresence gets the requireUserPresence property value. Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
func (m *MacOSKerberosSingleSignOnExtension) GetRequireUserPresence()(*bool) {
    return m.requireUserPresence
}
// GetSignInHelpText gets the signInHelpText property value. Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetSignInHelpText()(*string) {
    return m.signInHelpText
}
// GetTlsForLDAPRequired gets the tlsForLDAPRequired property value. When set to True, LDAP connections are required to use Transport Layer Security (TLS). Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetTlsForLDAPRequired()(*bool) {
    return m.tlsForLDAPRequired
}
// GetUsernameLabelCustom gets the usernameLabelCustom property value. This label replaces the user name shown in the Kerberos extension. You can enter a name to match the name of your company or organization. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetUsernameLabelCustom()(*string) {
    return m.usernameLabelCustom
}
// GetUserPrincipalName gets the userPrincipalName property value. Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
func (m *MacOSKerberosSingleSignOnExtension) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserSetupDelayed gets the userSetupDelayed property value. When set to True, the user isn’t prompted to set up the Kerberos extension until the extension is enabled by the admin, or a Kerberos challenge is received. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) GetUserSetupDelayed()(*bool) {
    return m.userSetupDelayed
}
// Serialize serializes information the current object
func (m *MacOSKerberosSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSSingleSignOnExtension.Serialize(writer)
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
    {
        err = writer.WriteBoolValue("credentialsCacheMonitored", m.GetCredentialsCacheMonitored())
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
        err = writer.WriteBoolValue("kerberosAppsInBundleIdACLIncluded", m.GetKerberosAppsInBundleIdACLIncluded())
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
        err = writer.WriteStringValue("modeCredentialUsed", m.GetModeCredentialUsed())
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
    if m.GetPreferredKDCs() != nil {
        err = writer.WriteCollectionOfStringValues("preferredKDCs", m.GetPreferredKDCs())
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
        err = writer.WriteBoolValue("tlsForLDAPRequired", m.GetTlsForLDAPRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameLabelCustom", m.GetUsernameLabelCustom())
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
    {
        err = writer.WriteBoolValue("userSetupDelayed", m.GetUserSetupDelayed())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDirectorySiteCode sets the activeDirectorySiteCode property value. Gets or sets the Active Directory site.
func (m *MacOSKerberosSingleSignOnExtension) SetActiveDirectorySiteCode(value *string)() {
    m.activeDirectorySiteCode = value
}
// SetBlockActiveDirectorySiteAutoDiscovery sets the blockActiveDirectorySiteAutoDiscovery property value. Enables or disables whether the Kerberos extension can automatically determine its site name.
func (m *MacOSKerberosSingleSignOnExtension) SetBlockActiveDirectorySiteAutoDiscovery(value *bool)() {
    m.blockActiveDirectorySiteAutoDiscovery = value
}
// SetBlockAutomaticLogin sets the blockAutomaticLogin property value. Enables or disables Keychain usage.
func (m *MacOSKerberosSingleSignOnExtension) SetBlockAutomaticLogin(value *bool)() {
    m.blockAutomaticLogin = value
}
// SetCacheName sets the cacheName property value. Gets or sets the Generic Security Services name of the Kerberos cache to use for this profile.
func (m *MacOSKerberosSingleSignOnExtension) SetCacheName(value *string)() {
    m.cacheName = value
}
// SetCredentialBundleIdAccessControlList sets the credentialBundleIdAccessControlList property value. Gets or sets a list of app Bundle IDs allowed to access the Kerberos Ticket Granting Ticket.
func (m *MacOSKerberosSingleSignOnExtension) SetCredentialBundleIdAccessControlList(value []string)() {
    m.credentialBundleIdAccessControlList = value
}
// SetCredentialsCacheMonitored sets the credentialsCacheMonitored property value. When set to True, the credential is requested on the next matching Kerberos challenge or network state change. When the credential is expired or missing, a new credential is created. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetCredentialsCacheMonitored(value *bool)() {
    m.credentialsCacheMonitored = value
}
// SetDomainRealms sets the domainRealms property value. Gets or sets a list of realms for custom domain-realm mapping. Realms are case sensitive.
func (m *MacOSKerberosSingleSignOnExtension) SetDomainRealms(value []string)() {
    m.domainRealms = value
}
// SetDomains sets the domains property value. Gets or sets a list of hosts or domain names for which the app extension performs SSO.
func (m *MacOSKerberosSingleSignOnExtension) SetDomains(value []string)() {
    m.domains = value
}
// SetIsDefaultRealm sets the isDefaultRealm property value. When true, this profile's realm will be selected as the default. Necessary if multiple Kerberos-type profiles are configured.
func (m *MacOSKerberosSingleSignOnExtension) SetIsDefaultRealm(value *bool)() {
    m.isDefaultRealm = value
}
// SetKerberosAppsInBundleIdACLIncluded sets the kerberosAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows any apps entered with the app bundle ID, managed apps, and standard Kerberos utilities, such as TicketViewer and klist, to access and use the credential. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetKerberosAppsInBundleIdACLIncluded(value *bool)() {
    m.kerberosAppsInBundleIdACLIncluded = value
}
// SetManagedAppsInBundleIdACLIncluded sets the managedAppsInBundleIdACLIncluded property value. When set to True, the Kerberos extension allows managed apps, and any apps entered with the app bundle ID to access the credential. When set to False, the Kerberos extension allows all apps to access the credential. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetManagedAppsInBundleIdACLIncluded(value *bool)() {
    m.managedAppsInBundleIdACLIncluded = value
}
// SetModeCredentialUsed sets the modeCredentialUsed property value. Select how other processes use the Kerberos Extension credential.
func (m *MacOSKerberosSingleSignOnExtension) SetModeCredentialUsed(value *string)() {
    m.modeCredentialUsed = value
}
// SetPasswordBlockModification sets the passwordBlockModification property value. Enables or disables password changes.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordBlockModification(value *bool)() {
    m.passwordBlockModification = value
}
// SetPasswordChangeUrl sets the passwordChangeUrl property value. Gets or sets the URL that the user will be sent to when they initiate a password change.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordChangeUrl(value *string)() {
    m.passwordChangeUrl = value
}
// SetPasswordEnableLocalSync sets the passwordEnableLocalSync property value. Enables or disables password syncing. This won't affect users logged in with a mobile account on macOS.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordEnableLocalSync(value *bool)() {
    m.passwordEnableLocalSync = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Overrides the default password expiration in days. For most domains, this value is calculated automatically.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordExpirationNotificationDays sets the passwordExpirationNotificationDays property value. Gets or sets the number of days until the user is notified that their password will expire (default is 15).
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordExpirationNotificationDays(value *int32)() {
    m.passwordExpirationNotificationDays = value
}
// SetPasswordMinimumAgeDays sets the passwordMinimumAgeDays property value. Gets or sets the minimum number of days until a user can change their password again.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordMinimumAgeDays(value *int32)() {
    m.passwordMinimumAgeDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Gets or sets the minimum length of a password.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Gets or sets the number of previous passwords to block.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequireActiveDirectoryComplexity sets the passwordRequireActiveDirectoryComplexity property value. Enables or disables whether passwords must meet Active Directory's complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordRequireActiveDirectoryComplexity(value *bool)() {
    m.passwordRequireActiveDirectoryComplexity = value
}
// SetPasswordRequirementsDescription sets the passwordRequirementsDescription property value. Gets or sets a description of the password complexity requirements.
func (m *MacOSKerberosSingleSignOnExtension) SetPasswordRequirementsDescription(value *string)() {
    m.passwordRequirementsDescription = value
}
// SetPreferredKDCs sets the preferredKDCs property value. Add creates an ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. This list is used when the servers are not discoverable using DNS. When the servers are discoverable, the list is used for both connectivity checks, and used first for Kerberos traffic. If the servers don’t respond, then the device uses DNS discovery. Delete removes an existing list, and devices use DNS discovery. Available for devices running macOS versions 12 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetPreferredKDCs(value []string)() {
    m.preferredKDCs = value
}
// SetRealm sets the realm property value. Gets or sets the case-sensitive realm name for this profile.
func (m *MacOSKerberosSingleSignOnExtension) SetRealm(value *string)() {
    m.realm = value
}
// SetRequireUserPresence sets the requireUserPresence property value. Gets or sets whether to require authentication via Touch ID, Face ID, or a passcode to access the keychain entry.
func (m *MacOSKerberosSingleSignOnExtension) SetRequireUserPresence(value *bool)() {
    m.requireUserPresence = value
}
// SetSignInHelpText sets the signInHelpText property value. Text displayed to the user at the Kerberos sign in window. Available for devices running iOS and iPadOS versions 14 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetSignInHelpText(value *string)() {
    m.signInHelpText = value
}
// SetTlsForLDAPRequired sets the tlsForLDAPRequired property value. When set to True, LDAP connections are required to use Transport Layer Security (TLS). Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetTlsForLDAPRequired(value *bool)() {
    m.tlsForLDAPRequired = value
}
// SetUsernameLabelCustom sets the usernameLabelCustom property value. This label replaces the user name shown in the Kerberos extension. You can enter a name to match the name of your company or organization. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetUsernameLabelCustom(value *string)() {
    m.usernameLabelCustom = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Gets or sets the principle user name to use for this profile. The realm name does not need to be included.
func (m *MacOSKerberosSingleSignOnExtension) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserSetupDelayed sets the userSetupDelayed property value. When set to True, the user isn’t prompted to set up the Kerberos extension until the extension is enabled by the admin, or a Kerberos challenge is received. Available for devices running macOS versions 11 and later.
func (m *MacOSKerberosSingleSignOnExtension) SetUserSetupDelayed(value *bool)() {
    m.userSetupDelayed = value
}
