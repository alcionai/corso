package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppManagementConfigurationable 
type AppManagementConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyCredentials()([]KeyCredentialConfigurationable)
    GetOdataType()(*string)
    GetPasswordCredentials()([]PasswordCredentialConfigurationable)
    SetKeyCredentials(value []KeyCredentialConfigurationable)()
    SetOdataType(value *string)()
    SetPasswordCredentials(value []PasswordCredentialConfigurationable)()
}
