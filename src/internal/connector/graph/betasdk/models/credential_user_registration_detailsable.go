package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationDetailsable 
type CredentialUserRegistrationDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthMethods()([]RegistrationAuthMethod)
    GetIsCapable()(*bool)
    GetIsEnabled()(*bool)
    GetIsMfaRegistered()(*bool)
    GetIsRegistered()(*bool)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetAuthMethods(value []RegistrationAuthMethod)()
    SetIsCapable(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsMfaRegistered(value *bool)()
    SetIsRegistered(value *bool)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
