package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationOAuth2ClientCredentialsConnectionSettingsable 
type EducationSynchronizationOAuth2ClientCredentialsConnectionSettingsable interface {
    EducationSynchronizationConnectionSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScope()(*string)
    GetTokenUrl()(*string)
    SetScope(value *string)()
    SetTokenUrl(value *string)()
}
