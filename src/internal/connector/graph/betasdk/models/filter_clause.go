package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FilterClause 
type FilterClause struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
    operatorName *string
    // Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
    sourceOperandName *string
    // Values that the source operand will be tested against.
    targetOperand FilterOperandable
}
// NewFilterClause instantiates a new filterClause and sets the default values.
func NewFilterClause()(*FilterClause) {
    m := &FilterClause{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFilterClauseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterClauseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilterClause(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FilterClause) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterClause) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["operatorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatorName(val)
        }
        return nil
    }
    res["sourceOperandName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceOperandName(val)
        }
        return nil
    }
    res["targetOperand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilterOperandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetOperand(val.(FilterOperandable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *FilterClause) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatorName gets the operatorName property value. Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
func (m *FilterClause) GetOperatorName()(*string) {
    return m.operatorName
}
// GetSourceOperandName gets the sourceOperandName property value. Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
func (m *FilterClause) GetSourceOperandName()(*string) {
    return m.sourceOperandName
}
// GetTargetOperand gets the targetOperand property value. Values that the source operand will be tested against.
func (m *FilterClause) GetTargetOperand()(FilterOperandable) {
    return m.targetOperand
}
// Serialize serializes information the current object
func (m *FilterClause) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatorName", m.GetOperatorName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceOperandName", m.GetSourceOperandName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("targetOperand", m.GetTargetOperand())
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
func (m *FilterClause) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FilterClause) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatorName sets the operatorName property value. Name of the operator to be applied to the source and target operands. Must be one of the supported operators. Supported operators can be discovered.
func (m *FilterClause) SetOperatorName(value *string)() {
    m.operatorName = value
}
// SetSourceOperandName sets the sourceOperandName property value. Name of source operand (the operand being tested). The source operand name must match one of the attribute names on the source object.
func (m *FilterClause) SetSourceOperandName(value *string)() {
    m.sourceOperandName = value
}
// SetTargetOperand sets the targetOperand property value. Values that the source operand will be tested against.
func (m *FilterClause) SetTargetOperand(value FilterOperandable)() {
    m.targetOperand = value
}
