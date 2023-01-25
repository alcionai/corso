package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactDataMatchStoreColumn 
type ExactDataMatchStoreColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ignoredDelimiters property
    ignoredDelimiters []string
    // The isCaseInsensitive property
    isCaseInsensitive *bool
    // The isSearchable property
    isSearchable *bool
    // The name property
    name *string
    // The OdataType property
    odataType *string
}
// NewExactDataMatchStoreColumn instantiates a new exactDataMatchStoreColumn and sets the default values.
func NewExactDataMatchStoreColumn()(*ExactDataMatchStoreColumn) {
    m := &ExactDataMatchStoreColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExactDataMatchStoreColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactDataMatchStoreColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactDataMatchStoreColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactDataMatchStoreColumn) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactDataMatchStoreColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ignoredDelimiters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIgnoredDelimiters(res)
        }
        return nil
    }
    res["isCaseInsensitive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCaseInsensitive(val)
        }
        return nil
    }
    res["isSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
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
// GetIgnoredDelimiters gets the ignoredDelimiters property value. The ignoredDelimiters property
func (m *ExactDataMatchStoreColumn) GetIgnoredDelimiters()([]string) {
    return m.ignoredDelimiters
}
// GetIsCaseInsensitive gets the isCaseInsensitive property value. The isCaseInsensitive property
func (m *ExactDataMatchStoreColumn) GetIsCaseInsensitive()(*bool) {
    return m.isCaseInsensitive
}
// GetIsSearchable gets the isSearchable property value. The isSearchable property
func (m *ExactDataMatchStoreColumn) GetIsSearchable()(*bool) {
    return m.isSearchable
}
// GetName gets the name property value. The name property
func (m *ExactDataMatchStoreColumn) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExactDataMatchStoreColumn) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ExactDataMatchStoreColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIgnoredDelimiters() != nil {
        err := writer.WriteCollectionOfStringValues("ignoredDelimiters", m.GetIgnoredDelimiters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCaseInsensitive", m.GetIsCaseInsensitive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
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
func (m *ExactDataMatchStoreColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIgnoredDelimiters sets the ignoredDelimiters property value. The ignoredDelimiters property
func (m *ExactDataMatchStoreColumn) SetIgnoredDelimiters(value []string)() {
    m.ignoredDelimiters = value
}
// SetIsCaseInsensitive sets the isCaseInsensitive property value. The isCaseInsensitive property
func (m *ExactDataMatchStoreColumn) SetIsCaseInsensitive(value *bool)() {
    m.isCaseInsensitive = value
}
// SetIsSearchable sets the isSearchable property value. The isSearchable property
func (m *ExactDataMatchStoreColumn) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetName sets the name property value. The name property
func (m *ExactDataMatchStoreColumn) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExactDataMatchStoreColumn) SetOdataType(value *string)() {
    m.odataType = value
}
