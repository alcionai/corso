package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RuleThreshold 
type RuleThreshold struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates the built-in aggregation methods. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
    aggregation *AggregationType
    // The OdataType property
    odataType *string
    // Indicates the built-in operator. The possible values are: greaterOrEqual, equal, greater, less, lessOrEqual, notEqual, unknownFutureValue.
    operator *OperatorType
    // The target threshold value.
    target *int32
}
// NewRuleThreshold instantiates a new ruleThreshold and sets the default values.
func NewRuleThreshold()(*RuleThreshold) {
    m := &RuleThreshold{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRuleThresholdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRuleThresholdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleThreshold(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RuleThreshold) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAggregation gets the aggregation property value. Indicates the built-in aggregation methods. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *RuleThreshold) GetAggregation()(*AggregationType) {
    return m.aggregation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RuleThreshold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAggregationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregation(val.(*AggregationType))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperatorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*OperatorType))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RuleThreshold) GetOdataType()(*string) {
    return m.odataType
}
// GetOperator gets the operator property value. Indicates the built-in operator. The possible values are: greaterOrEqual, equal, greater, less, lessOrEqual, notEqual, unknownFutureValue.
func (m *RuleThreshold) GetOperator()(*OperatorType) {
    return m.operator
}
// GetTarget gets the target property value. The target threshold value.
func (m *RuleThreshold) GetTarget()(*int32) {
    return m.target
}
// Serialize serializes information the current object
func (m *RuleThreshold) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregation() != nil {
        cast := (*m.GetAggregation()).String()
        err := writer.WriteStringValue("aggregation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RuleThreshold) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAggregation sets the aggregation property value. Indicates the built-in aggregation methods. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *RuleThreshold) SetAggregation(value *AggregationType)() {
    m.aggregation = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RuleThreshold) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperator sets the operator property value. Indicates the built-in operator. The possible values are: greaterOrEqual, equal, greater, less, lessOrEqual, notEqual, unknownFutureValue.
func (m *RuleThreshold) SetOperator(value *OperatorType)() {
    m.operator = value
}
// SetTarget sets the target property value. The target threshold value.
func (m *RuleThreshold) SetTarget(value *int32)() {
    m.target = value
}
