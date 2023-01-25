package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalLockConfigurationable 
type ServicePrincipalLockConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllProperties()(*bool)
    GetCredentialsWithUsageSign()(*bool)
    GetCredentialsWithUsageVerify()(*bool)
    GetIsEnabled()(*bool)
    GetOdataType()(*string)
    GetTokenEncryptionKeyId()(*bool)
    SetAllProperties(value *bool)()
    SetCredentialsWithUsageSign(value *bool)()
    SetCredentialsWithUsageVerify(value *bool)()
    SetIsEnabled(value *bool)()
    SetOdataType(value *string)()
    SetTokenEncryptionKeyId(value *bool)()
}
