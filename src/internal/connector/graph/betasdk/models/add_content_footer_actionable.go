package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddContentFooterActionable 
type AddContentFooterActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlignment()(*ContentAlignment)
    GetFontColor()(*string)
    GetFontName()(*string)
    GetFontSize()(*int32)
    GetMargin()(*int32)
    GetText()(*string)
    GetUiElementName()(*string)
    SetAlignment(value *ContentAlignment)()
    SetFontColor(value *string)()
    SetFontName(value *string)()
    SetFontSize(value *int32)()
    SetMargin(value *int32)()
    SetText(value *string)()
    SetUiElementName(value *string)()
}
