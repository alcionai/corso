package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RuleThresholdable 
type RuleThresholdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregation()(*AggregationType)
    GetOdataType()(*string)
    GetOperator()(*OperatorType)
    GetTarget()(*int32)
    SetAggregation(value *AggregationType)()
    SetOdataType(value *string)()
    SetOperator(value *OperatorType)()
    SetTarget(value *int32)()
}
