package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationSettingsable 
type OrganizationSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContactInsights()(InsightsSettingsable)
    GetItemInsights()(InsightsSettingsable)
    GetMicrosoftApplicationDataAccess()(MicrosoftApplicationDataAccessSettingsable)
    GetPeopleInsights()(InsightsSettingsable)
    GetProfileCardProperties()([]ProfileCardPropertyable)
    SetContactInsights(value InsightsSettingsable)()
    SetItemInsights(value InsightsSettingsable)()
    SetMicrosoftApplicationDataAccess(value MicrosoftApplicationDataAccessSettingsable)()
    SetPeopleInsights(value InsightsSettingsable)()
    SetProfileCardProperties(value []ProfileCardPropertyable)()
}
