package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContentable 
type DetectedSensitiveContentable interface {
    DetectedSensitiveContentBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassificationAttributes()([]ClassificationAttributeable)
    GetClassificationMethod()(*ClassificationMethod)
    GetMatches()([]SensitiveContentLocationable)
    GetScope()(*SensitiveTypeScope)
    GetSensitiveTypeSource()(*SensitiveTypeSource)
    SetClassificationAttributes(value []ClassificationAttributeable)()
    SetClassificationMethod(value *ClassificationMethod)()
    SetMatches(value []SensitiveContentLocationable)()
    SetScope(value *SensitiveTypeScope)()
    SetSensitiveTypeSource(value *SensitiveTypeSource)()
}
