package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KerberosSingleSignOnExtensionable 
type KerberosSingleSignOnExtensionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SingleSignOnExtensionable
    GetActiveDirectorySiteCode()(*string)
    GetBlockActiveDirectorySiteAutoDiscovery()(*bool)
    GetBlockAutomaticLogin()(*bool)
    GetCacheName()(*string)
    GetCredentialBundleIdAccessControlList()([]string)
    GetDomainRealms()([]string)
    GetDomains()([]string)
    GetIsDefaultRealm()(*bool)
    GetPasswordBlockModification()(*bool)
    GetPasswordChangeUrl()(*string)
    GetPasswordEnableLocalSync()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordExpirationNotificationDays()(*int32)
    GetPasswordMinimumAgeDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequireActiveDirectoryComplexity()(*bool)
    GetPasswordRequirementsDescription()(*string)
    GetRealm()(*string)
    GetRequireUserPresence()(*bool)
    GetUserPrincipalName()(*string)
    SetActiveDirectorySiteCode(value *string)()
    SetBlockActiveDirectorySiteAutoDiscovery(value *bool)()
    SetBlockAutomaticLogin(value *bool)()
    SetCacheName(value *string)()
    SetCredentialBundleIdAccessControlList(value []string)()
    SetDomainRealms(value []string)()
    SetDomains(value []string)()
    SetIsDefaultRealm(value *bool)()
    SetPasswordBlockModification(value *bool)()
    SetPasswordChangeUrl(value *string)()
    SetPasswordEnableLocalSync(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordExpirationNotificationDays(value *int32)()
    SetPasswordMinimumAgeDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequireActiveDirectoryComplexity(value *bool)()
    SetPasswordRequirementsDescription(value *string)()
    SetRealm(value *string)()
    SetRequireUserPresence(value *bool)()
    SetUserPrincipalName(value *string)()
}
