package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FilterGroup 
type FilterGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Filter clauses (conditions) of this group. All clauses in a group must be satisfied in order for the filter group to evaluate to true.
    clauses []FilterClauseable
    // Human-readable name of the filter group.
    name *string
    // The OdataType property
    odataType *string
}
// NewFilterGroup instantiates a new filterGroup and sets the default values.
func NewFilterGroup()(*FilterGroup) {
    m := &FilterGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFilterGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilterGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FilterGroup) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClauses gets the clauses property value. Filter clauses (conditions) of this group. All clauses in a group must be satisfied in order for the filter group to evaluate to true.
func (m *FilterGroup) GetClauses()([]FilterClauseable) {
    return m.clauses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clauses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterClauseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterClauseable, len(val))
            for i, v := range val {
                res[i] = v.(FilterClauseable)
            }
            m.SetClauses(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. Human-readable name of the filter group.
func (m *FilterGroup) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *FilterGroup) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FilterGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClauses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClauses()))
        for i, v := range m.GetClauses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("clauses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *FilterGroup) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClauses sets the clauses property value. Filter clauses (conditions) of this group. All clauses in a group must be satisfied in order for the filter group to evaluate to true.
func (m *FilterGroup) SetClauses(value []FilterClauseable)() {
    m.clauses = value
}
// SetName sets the name property value. Human-readable name of the filter group.
func (m *FilterGroup) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FilterGroup) SetOdataType(value *string)() {
    m.odataType = value
}
