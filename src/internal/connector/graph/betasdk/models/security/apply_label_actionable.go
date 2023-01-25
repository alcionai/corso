package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplyLabelActionable 
type ApplyLabelActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]InformationProtectionActionable)
    GetActionSource()(*ActionSource)
    GetResponsibleSensitiveTypeIds()([]string)
    GetSensitivityLabelId()(*string)
    SetActions(value []InformationProtectionActionable)()
    SetActionSource(value *ActionSource)()
    SetResponsibleSensitiveTypeIds(value []string)()
    SetSensitivityLabelId(value *string)()
}
