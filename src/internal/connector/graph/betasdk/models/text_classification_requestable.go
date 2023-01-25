package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TextClassificationRequestable 
type TextClassificationRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileExtension()(*string)
    GetMatchTolerancesToInclude()(*MlClassificationMatchTolerance)
    GetScopesToRun()(*SensitiveTypeScope)
    GetSensitiveTypeIds()([]string)
    GetText()(*string)
    SetFileExtension(value *string)()
    SetMatchTolerancesToInclude(value *MlClassificationMatchTolerance)()
    SetScopesToRun(value *SensitiveTypeScope)()
    SetSensitiveTypeIds(value []string)()
    SetText(value *string)()
}
