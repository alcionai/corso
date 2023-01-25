package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualAppointmentable 
type VirtualAppointmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppointmentClientJoinWebUrl()(*string)
    GetAppointmentClients()([]VirtualAppointmentUserable)
    GetExternalAppointmentId()(*string)
    GetExternalAppointmentUrl()(*string)
    GetSettings()(VirtualAppointmentSettingsable)
    SetAppointmentClientJoinWebUrl(value *string)()
    SetAppointmentClients(value []VirtualAppointmentUserable)()
    SetExternalAppointmentId(value *string)()
    SetExternalAppointmentUrl(value *string)()
    SetSettings(value VirtualAppointmentSettingsable)()
}
