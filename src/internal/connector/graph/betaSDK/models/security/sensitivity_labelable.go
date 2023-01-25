package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// SensitivityLabelable 
type SensitivityLabelable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetContentFormats()([]string)
    GetDescription()(*string)
    GetHasProtection()(*bool)
    GetIsActive()(*bool)
    GetIsAppliable()(*bool)
    GetName()(*string)
    GetParent()(SensitivityLabelable)
    GetSensitivity()(*int32)
    GetTooltip()(*string)
    SetColor(value *string)()
    SetContentFormats(value []string)()
    SetDescription(value *string)()
    SetHasProtection(value *bool)()
    SetIsActive(value *bool)()
    SetIsAppliable(value *bool)()
    SetName(value *string)()
    SetParent(value SensitivityLabelable)()
    SetSensitivity(value *int32)()
    SetTooltip(value *string)()
}
