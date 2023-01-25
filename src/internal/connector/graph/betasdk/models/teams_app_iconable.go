package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppIconable 
type TeamsAppIconable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostedContent()(TeamworkHostedContentable)
    GetWebUrl()(*string)
    SetHostedContent(value TeamworkHostedContentable)()
    SetWebUrl(value *string)()
}
