package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApplicationSegmentable 
type WebApplicationSegmentable interface {
    ApplicationSegmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateUrl()(*string)
    GetCorsConfigurations()([]CorsConfiguration_v2able)
    GetExternalUrl()(*string)
    GetInternalUrl()(*string)
    SetAlternateUrl(value *string)()
    SetCorsConfigurations(value []CorsConfiguration_v2able)()
    SetExternalUrl(value *string)()
    SetInternalUrl(value *string)()
}
