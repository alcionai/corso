package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelSummaryable 
type ChannelSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGuestsCount()(*int32)
    GetHasMembersFromOtherTenants()(*bool)
    GetMembersCount()(*int32)
    GetOdataType()(*string)
    GetOwnersCount()(*int32)
    SetGuestsCount(value *int32)()
    SetHasMembersFromOtherTenants(value *bool)()
    SetMembersCount(value *int32)()
    SetOdataType(value *string)()
    SetOwnersCount(value *int32)()
}
