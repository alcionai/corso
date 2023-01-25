package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EvaluateLabelJobResult 
type EvaluateLabelJobResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The responsiblePolicy property
    responsiblePolicy ResponsiblePolicyable
    // The responsibleSensitiveTypes property
    responsibleSensitiveTypes []ResponsibleSensitiveTypeable
    // The sensitivityLabel property
    sensitivityLabel MatchingLabelable
}
// NewEvaluateLabelJobResult instantiates a new evaluateLabelJobResult and sets the default values.
func NewEvaluateLabelJobResult()(*EvaluateLabelJobResult) {
    m := &EvaluateLabelJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateLabelJobResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateLabelJobResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateLabelJobResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["responsiblePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResponsiblePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiblePolicy(val.(ResponsiblePolicyable))
        }
        return nil
    }
    res["responsibleSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResponsibleSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResponsibleSensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(ResponsibleSensitiveTypeable)
            }
            m.SetResponsibleSensitiveTypes(res)
        }
        return nil
    }
    res["sensitivityLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMatchingLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabel(val.(MatchingLabelable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EvaluateLabelJobResult) GetOdataType()(*string) {
    return m.odataType
}
// GetResponsiblePolicy gets the responsiblePolicy property value. The responsiblePolicy property
func (m *EvaluateLabelJobResult) GetResponsiblePolicy()(ResponsiblePolicyable) {
    return m.responsiblePolicy
}
// GetResponsibleSensitiveTypes gets the responsibleSensitiveTypes property value. The responsibleSensitiveTypes property
func (m *EvaluateLabelJobResult) GetResponsibleSensitiveTypes()([]ResponsibleSensitiveTypeable) {
    return m.responsibleSensitiveTypes
}
// GetSensitivityLabel gets the sensitivityLabel property value. The sensitivityLabel property
func (m *EvaluateLabelJobResult) GetSensitivityLabel()(MatchingLabelable) {
    return m.sensitivityLabel
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("responsiblePolicy", m.GetResponsiblePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetResponsibleSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResponsibleSensitiveTypes()))
        for i, v := range m.GetResponsibleSensitiveTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("responsibleSensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sensitivityLabel", m.GetSensitivityLabel())
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
func (m *EvaluateLabelJobResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EvaluateLabelJobResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResponsiblePolicy sets the responsiblePolicy property value. The responsiblePolicy property
func (m *EvaluateLabelJobResult) SetResponsiblePolicy(value ResponsiblePolicyable)() {
    m.responsiblePolicy = value
}
// SetResponsibleSensitiveTypes sets the responsibleSensitiveTypes property value. The responsibleSensitiveTypes property
func (m *EvaluateLabelJobResult) SetResponsibleSensitiveTypes(value []ResponsibleSensitiveTypeable)() {
    m.responsibleSensitiveTypes = value
}
// SetSensitivityLabel sets the sensitivityLabel property value. The sensitivityLabel property
func (m *EvaluateLabelJobResult) SetSensitivityLabel(value MatchingLabelable)() {
    m.sensitivityLabel = value
}
