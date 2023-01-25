package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkMobileAppConfigurationable 
type AndroidForWorkMobileAppConfigurationable interface {
    ManagedDeviceMobileAppConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectedAppsEnabled()(*bool)
    GetPackageId()(*string)
    GetPayloadJson()(*string)
    GetPermissionActions()([]AndroidPermissionActionable)
    GetProfileApplicability()(*AndroidProfileApplicability)
    SetConnectedAppsEnabled(value *bool)()
    SetPackageId(value *string)()
    SetPayloadJson(value *string)()
    SetPermissionActions(value []AndroidPermissionActionable)()
    SetProfileApplicability(value *AndroidProfileApplicability)()
}
