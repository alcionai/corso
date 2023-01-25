package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosSingleSignOnSettingsable 
type IosSingleSignOnSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedAppsList()([]AppListItemable)
    GetAllowedUrls()([]string)
    GetDisplayName()(*string)
    GetKerberosPrincipalName()(*string)
    GetKerberosRealm()(*string)
    GetOdataType()(*string)
    SetAllowedAppsList(value []AppListItemable)()
    SetAllowedUrls(value []string)()
    SetDisplayName(value *string)()
    SetKerberosPrincipalName(value *string)()
    SetKerberosRealm(value *string)()
    SetOdataType(value *string)()
}
