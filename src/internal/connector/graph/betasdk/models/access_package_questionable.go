package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageQuestionable 
type AccessPackageQuestionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetIsAnswerEditable()(*bool)
    GetIsRequired()(*bool)
    GetOdataType()(*string)
    GetSequence()(*int32)
    GetText()(AccessPackageLocalizedContentable)
    SetId(value *string)()
    SetIsAnswerEditable(value *bool)()
    SetIsRequired(value *bool)()
    SetOdataType(value *string)()
    SetSequence(value *int32)()
    SetText(value AccessPackageLocalizedContentable)()
}
