package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionDurationInDaysable 
type RetentionDurationInDaysable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RetentionDurationable
    GetDays()(*int32)
    SetDays(value *int32)()
}
