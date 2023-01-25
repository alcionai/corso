package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CollapseProperty 
type CollapseProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Defines the collapse group to trim results. The properties in this collection must be sortable/refinable properties. Required.
    fields []string
    // Defines a maximum limit count for this field. This numeric value must be a positive integer. Required.
    limit *int32
    // The OdataType property
    odataType *string
}
// NewCollapseProperty instantiates a new collapseProperty and sets the default values.
func NewCollapseProperty()(*CollapseProperty) {
    m := &CollapseProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCollapsePropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCollapsePropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollapseProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CollapseProperty) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CollapseProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFields(res)
        }
        return nil
    }
    res["limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimit(val)
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
    return res
}
// GetFields gets the fields property value. Defines the collapse group to trim results. The properties in this collection must be sortable/refinable properties. Required.
func (m *CollapseProperty) GetFields()([]string) {
    return m.fields
}
// GetLimit gets the limit property value. Defines a maximum limit count for this field. This numeric value must be a positive integer. Required.
func (m *CollapseProperty) GetLimit()(*int32) {
    return m.limit
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CollapseProperty) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CollapseProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFields() != nil {
        err := writer.WriteCollectionOfStringValues("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("limit", m.GetLimit())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CollapseProperty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFields sets the fields property value. Defines the collapse group to trim results. The properties in this collection must be sortable/refinable properties. Required.
func (m *CollapseProperty) SetFields(value []string)() {
    m.fields = value
}
// SetLimit sets the limit property value. Defines a maximum limit count for this field. This numeric value must be a positive integer. Required.
func (m *CollapseProperty) SetLimit(value *int32)() {
    m.limit = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CollapseProperty) SetOdataType(value *string)() {
    m.odataType = value
}
