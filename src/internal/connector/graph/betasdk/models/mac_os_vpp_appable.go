package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppAppable 
type MacOsVppAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppStoreUrl()(*string)
    GetAssignedLicenses()([]MacOsVppAppAssignedLicenseable)
    GetBundleId()(*string)
    GetLicensingType()(VppLicensingTypeable)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRevokeLicenseActionResults()([]MacOsVppAppRevokeLicensesActionResultable)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    GetVppTokenAccountType()(*VppTokenAccountType)
    GetVppTokenAppleId()(*string)
    GetVppTokenId()(*string)
    GetVppTokenOrganizationName()(*string)
    SetAppStoreUrl(value *string)()
    SetAssignedLicenses(value []MacOsVppAppAssignedLicenseable)()
    SetBundleId(value *string)()
    SetLicensingType(value VppLicensingTypeable)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRevokeLicenseActionResults(value []MacOsVppAppRevokeLicensesActionResultable)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
    SetVppTokenAccountType(value *VppTokenAccountType)()
    SetVppTokenAppleId(value *string)()
    SetVppTokenId(value *string)()
    SetVppTokenOrganizationName(value *string)()
}
