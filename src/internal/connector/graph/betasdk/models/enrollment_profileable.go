package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentProfileable 
type EnrollmentProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationEndpointUrl()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnableAuthenticationViaCompanyPortal()(*bool)
    GetRequireCompanyPortalOnSetupAssistantEnrolledDevices()(*bool)
    GetRequiresUserAuthentication()(*bool)
    SetConfigurationEndpointUrl(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnableAuthenticationViaCompanyPortal(value *bool)()
    SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(value *bool)()
    SetRequiresUserAuthentication(value *bool)()
}
