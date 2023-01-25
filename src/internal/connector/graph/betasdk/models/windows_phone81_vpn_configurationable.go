package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81VpnConfigurationable 
type WindowsPhone81VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows81VpnConfigurationable
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetBypassVpnOnCompanyWifi()(*bool)
    GetBypassVpnOnHomeWifi()(*bool)
    GetDnsSuffixSearchList()([]string)
    GetIdentityCertificate()(WindowsPhone81CertificateProfileBaseable)
    GetRememberUserCredentials()(*bool)
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetBypassVpnOnCompanyWifi(value *bool)()
    SetBypassVpnOnHomeWifi(value *bool)()
    SetDnsSuffixSearchList(value []string)()
    SetIdentityCertificate(value WindowsPhone81CertificateProfileBaseable)()
    SetRememberUserCredentials(value *bool)()
}
