package tenantadmin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdleSessionSignOutable 
type IdleSessionSignOutable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEnabled()(*bool)
    GetOdataType()(*string)
    GetSignOutAfterInSeconds()(*int64)
    GetWarnAfterInSeconds()(*int64)
    SetIsEnabled(value *bool)()
    SetOdataType(value *string)()
    SetSignOutAfterInSeconds(value *int64)()
    SetWarnAfterInSeconds(value *int64)()
}
