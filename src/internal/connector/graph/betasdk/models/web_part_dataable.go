package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebPartDataable 
type WebPartDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudiences()([]string)
    GetDataVersion()(*string)
    GetDescription()(*string)
    GetOdataType()(*string)
    GetProperties()(Jsonable)
    GetServerProcessedContent()(ServerProcessedContentable)
    GetTitle()(*string)
    SetAudiences(value []string)()
    SetDataVersion(value *string)()
    SetDescription(value *string)()
    SetOdataType(value *string)()
    SetProperties(value Jsonable)()
    SetServerProcessedContent(value ServerProcessedContentable)()
    SetTitle(value *string)()
}
