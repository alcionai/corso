package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSettingsable 
type UserSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContactMergeSuggestions()(ContactMergeSuggestionsable)
    GetContributionToContentDiscoveryAsOrganizationDisabled()(*bool)
    GetContributionToContentDiscoveryDisabled()(*bool)
    GetItemInsights()(UserInsightsSettingsable)
    GetRegionalAndLanguageSettings()(RegionalAndLanguageSettingsable)
    GetShiftPreferences()(ShiftPreferencesable)
    SetContactMergeSuggestions(value ContactMergeSuggestionsable)()
    SetContributionToContentDiscoveryAsOrganizationDisabled(value *bool)()
    SetContributionToContentDiscoveryDisabled(value *bool)()
    SetItemInsights(value UserInsightsSettingsable)()
    SetRegionalAndLanguageSettings(value RegionalAndLanguageSettingsable)()
    SetShiftPreferences(value ShiftPreferencesable)()
}
