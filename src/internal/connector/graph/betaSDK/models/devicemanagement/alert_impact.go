package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertImpact 
type AlertImpact struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
    aggregationType *AggregationType
    // The OdataType property
    odataType *string
    // The number value of the impact. For the aggregation types of count and affectedCloudPcCount, the value indicates the number of affected instances. For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. For the aggregation types of percentage and affectedCloudPcPercentage, the value indicates the percent of affected instances. For example, 12 affectedCloudPcPercentage means that 12% of Cloud PCs are affected.
    value *int32
}
// NewAlertImpact instantiates a new alertImpact and sets the default values.
func NewAlertImpact()(*AlertImpact) {
    m := &AlertImpact{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlertImpactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertImpactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertImpact(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertImpact) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAggregationType gets the aggregationType property value. The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *AlertImpact) GetAggregationType()(*AggregationType) {
    return m.aggregationType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertImpact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAggregationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregationType(val.(*AggregationType))
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlertImpact) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The number value of the impact. For the aggregation types of count and affectedCloudPcCount, the value indicates the number of affected instances. For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. For the aggregation types of percentage and affectedCloudPcPercentage, the value indicates the percent of affected instances. For example, 12 affectedCloudPcPercentage means that 12% of Cloud PCs are affected.
func (m *AlertImpact) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *AlertImpact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregationType() != nil {
        cast := (*m.GetAggregationType()).String()
        err := writer.WriteStringValue("aggregationType", &cast)
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
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *AlertImpact) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAggregationType sets the aggregationType property value. The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *AlertImpact) SetAggregationType(value *AggregationType)() {
    m.aggregationType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertImpact) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The number value of the impact. For the aggregation types of count and affectedCloudPcCount, the value indicates the number of affected instances. For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. For the aggregation types of percentage and affectedCloudPcPercentage, the value indicates the percent of affected instances. For example, 12 affectedCloudPcPercentage means that 12% of Cloud PCs are affected.
func (m *AlertImpact) SetValue(value *int32)() {
    m.value = value
}
