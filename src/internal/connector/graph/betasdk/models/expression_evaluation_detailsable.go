package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpressionEvaluationDetailsable 
type ExpressionEvaluationDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpression()(*string)
    GetExpressionEvaluationDetails()([]ExpressionEvaluationDetailsable)
    GetExpressionResult()(*bool)
    GetOdataType()(*string)
    GetPropertyToEvaluate()(PropertyToEvaluateable)
    SetExpression(value *string)()
    SetExpressionEvaluationDetails(value []ExpressionEvaluationDetailsable)()
    SetExpressionResult(value *bool)()
    SetOdataType(value *string)()
    SetPropertyToEvaluate(value PropertyToEvaluateable)()
}
