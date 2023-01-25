package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentNotificationConfigurationable 
type DeviceEnrollmentNotificationConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrandingOptions()(*EnrollmentNotificationBrandingOptions)
    GetDefaultLocale()(*string)
    GetNotificationMessageTemplateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetNotificationTemplates()([]string)
    GetPlatformType()(*EnrollmentRestrictionPlatformType)
    GetTemplateType()(*EnrollmentNotificationTemplateType)
    SetBrandingOptions(value *EnrollmentNotificationBrandingOptions)()
    SetDefaultLocale(value *string)()
    SetNotificationMessageTemplateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetNotificationTemplates(value []string)()
    SetPlatformType(value *EnrollmentRestrictionPlatformType)()
    SetTemplateType(value *EnrollmentNotificationTemplateType)()
}
