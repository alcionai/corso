package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddWatermarkActionable 
type AddWatermarkActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFontColor()(*string)
    GetFontName()(*string)
    GetFontSize()(*int32)
    GetLayout()(*WatermarkLayout)
    GetText()(*string)
    GetUiElementName()(*string)
    SetFontColor(value *string)()
    SetFontName(value *string)()
    SetFontSize(value *int32)()
    SetLayout(value *WatermarkLayout)()
    SetText(value *string)()
    SetUiElementName(value *string)()
}
