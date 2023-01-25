package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegionalAndLanguageSettingsable 
type RegionalAndLanguageSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthoringLanguages()([]LocaleInfoable)
    GetDefaultDisplayLanguage()(LocaleInfoable)
    GetDefaultRegionalFormat()(LocaleInfoable)
    GetDefaultSpeechInputLanguage()(LocaleInfoable)
    GetDefaultTranslationLanguage()(LocaleInfoable)
    GetRegionalFormatOverrides()(RegionalFormatOverridesable)
    GetTranslationPreferences()(TranslationPreferencesable)
    SetAuthoringLanguages(value []LocaleInfoable)()
    SetDefaultDisplayLanguage(value LocaleInfoable)()
    SetDefaultRegionalFormat(value LocaleInfoable)()
    SetDefaultSpeechInputLanguage(value LocaleInfoable)()
    SetDefaultTranslationLanguage(value LocaleInfoable)()
    SetRegionalFormatOverrides(value RegionalFormatOverridesable)()
    SetTranslationPreferences(value TranslationPreferencesable)()
}
