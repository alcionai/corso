package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerEnrollmentProfileable 
type AndroidDeviceOwnerEnrollmentProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountId()(*string)
    GetConfigureWifi()(*bool)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnrolledDeviceCount()(*int32)
    GetEnrollmentMode()(*AndroidDeviceOwnerEnrollmentMode)
    GetEnrollmentTokenType()(*AndroidDeviceOwnerEnrollmentTokenType)
    GetEnrollmentTokenUsageCount()(*int32)
    GetIsTeamsDeviceProfile()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetQrCodeContent()(*string)
    GetQrCodeImage()(MimeContentable)
    GetRoleScopeTagIds()([]string)
    GetTokenCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTokenValue()(*string)
    GetWifiHidden()(*bool)
    GetWifiPassword()(*string)
    GetWifiSecurityType()(*AospWifiSecurityType)
    GetWifiSsid()(*string)
    SetAccountId(value *string)()
    SetConfigureWifi(value *bool)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnrolledDeviceCount(value *int32)()
    SetEnrollmentMode(value *AndroidDeviceOwnerEnrollmentMode)()
    SetEnrollmentTokenType(value *AndroidDeviceOwnerEnrollmentTokenType)()
    SetEnrollmentTokenUsageCount(value *int32)()
    SetIsTeamsDeviceProfile(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetQrCodeContent(value *string)()
    SetQrCodeImage(value MimeContentable)()
    SetRoleScopeTagIds(value []string)()
    SetTokenCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTokenValue(value *string)()
    SetWifiHidden(value *bool)()
    SetWifiPassword(value *string)()
    SetWifiSecurityType(value *AospWifiSecurityType)()
    SetWifiSsid(value *string)()
}
