package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActivityOLDable 
type ItemActivityOLDable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(ItemActionSetable)
    GetActor()(IdentitySetable)
    GetDriveItem()(DriveItemable)
    GetListItem()(ListItemable)
    GetTimes()(ItemActivityTimeSetable)
    SetAction(value ItemActionSetable)()
    SetActor(value IdentitySetable)()
    SetDriveItem(value DriveItemable)()
    SetListItem(value ListItemable)()
    SetTimes(value ItemActivityTimeSetable)()
}
