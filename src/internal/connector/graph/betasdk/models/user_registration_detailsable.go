package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationDetailsable 
type UserRegistrationDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultMfaMethod()(*DefaultMfaMethodType)
    GetIsAdmin()(*bool)
    GetIsMfaCapable()(*bool)
    GetIsMfaRegistered()(*bool)
    GetIsPasswordlessCapable()(*bool)
    GetIsSsprCapable()(*bool)
    GetIsSsprEnabled()(*bool)
    GetIsSsprRegistered()(*bool)
    GetMethodsRegistered()([]string)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    GetUserType()(*SignInUserType)
    SetDefaultMfaMethod(value *DefaultMfaMethodType)()
    SetIsAdmin(value *bool)()
    SetIsMfaCapable(value *bool)()
    SetIsMfaRegistered(value *bool)()
    SetIsPasswordlessCapable(value *bool)()
    SetIsSsprCapable(value *bool)()
    SetIsSsprEnabled(value *bool)()
    SetIsSsprRegistered(value *bool)()
    SetMethodsRegistered(value []string)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *SignInUserType)()
}
