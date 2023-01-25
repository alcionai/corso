package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LearningProviderable 
type LearningProviderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetIsEnabled()(*bool)
    GetLearningContents()([]LearningContentable)
    GetLoginWebUrl()(*string)
    GetLongLogoWebUrlForDarkTheme()(*string)
    GetLongLogoWebUrlForLightTheme()(*string)
    GetSquareLogoWebUrlForDarkTheme()(*string)
    GetSquareLogoWebUrlForLightTheme()(*string)
    SetDisplayName(value *string)()
    SetIsEnabled(value *bool)()
    SetLearningContents(value []LearningContentable)()
    SetLoginWebUrl(value *string)()
    SetLongLogoWebUrlForDarkTheme(value *string)()
    SetLongLogoWebUrlForLightTheme(value *string)()
    SetSquareLogoWebUrlForDarkTheme(value *string)()
    SetSquareLogoWebUrlForLightTheme(value *string)()
}
