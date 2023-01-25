package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsInsight the user experience analytics insight is the recomendation to improve the user experience analytics score.
type UserExperienceAnalyticsInsight struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier of the user experience analytics insight.
    insightId *string
    // The OdataType property
    odataType *string
    // The severity property
    severity *UserExperienceAnalyticsInsightSeverity
    // The unique identifier of the user experience analytics insight.
    userExperienceAnalyticsMetricId *string
    // The value of the user experience analytics insight.
    values []UserExperienceAnalyticsInsightValueable
}
// NewUserExperienceAnalyticsInsight instantiates a new userExperienceAnalyticsInsight and sets the default values.
func NewUserExperienceAnalyticsInsight()(*UserExperienceAnalyticsInsight) {
    m := &UserExperienceAnalyticsInsight{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsInsight(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsInsight) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["insightId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsightId(val)
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
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsInsightSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*UserExperienceAnalyticsInsightSeverity))
        }
        return nil
    }
    res["userExperienceAnalyticsMetricId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsMetricId(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsightValueable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsInsightValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetInsightId gets the insightId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetInsightId()(*string) {
    return m.insightId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsInsight) GetOdataType()(*string) {
    return m.odataType
}
// GetSeverity gets the severity property value. The severity property
func (m *UserExperienceAnalyticsInsight) GetSeverity()(*UserExperienceAnalyticsInsightSeverity) {
    return m.severity
}
// GetUserExperienceAnalyticsMetricId gets the userExperienceAnalyticsMetricId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetUserExperienceAnalyticsMetricId()(*string) {
    return m.userExperienceAnalyticsMetricId
}
// GetValues gets the values property value. The value of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetValues()([]UserExperienceAnalyticsInsightValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("insightId", m.GetInsightId())
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
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userExperienceAnalyticsMetricId", m.GetUserExperienceAnalyticsMetricId())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("values", cast)
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
func (m *UserExperienceAnalyticsInsight) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetInsightId sets the insightId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetInsightId(value *string)() {
    m.insightId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsInsight) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSeverity sets the severity property value. The severity property
func (m *UserExperienceAnalyticsInsight) SetSeverity(value *UserExperienceAnalyticsInsightSeverity)() {
    m.severity = value
}
// SetUserExperienceAnalyticsMetricId sets the userExperienceAnalyticsMetricId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetUserExperienceAnalyticsMetricId(value *string)() {
    m.userExperienceAnalyticsMetricId = value
}
// SetValues sets the values property value. The value of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetValues(value []UserExperienceAnalyticsInsightValueable)() {
    m.values = value
}
