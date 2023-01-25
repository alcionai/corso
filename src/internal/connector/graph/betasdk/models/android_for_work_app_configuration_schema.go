package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkAppConfigurationSchema schema describing an Android for Work application's custom configurations.
type AndroidForWorkAppConfigurationSchema struct {
    Entity
    // UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
    exampleJson []byte
    // Collection of items each representing a named configuration option in the schema
    schemaItems []AndroidForWorkAppConfigurationSchemaItemable
}
// NewAndroidForWorkAppConfigurationSchema instantiates a new androidForWorkAppConfigurationSchema and sets the default values.
func NewAndroidForWorkAppConfigurationSchema()(*AndroidForWorkAppConfigurationSchema) {
    m := &AndroidForWorkAppConfigurationSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkAppConfigurationSchema(), nil
}
// GetExampleJson gets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidForWorkAppConfigurationSchema) GetExampleJson()([]byte) {
    return m.exampleJson
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkAppConfigurationSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exampleJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExampleJson(val)
        }
        return nil
    }
    res["schemaItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkAppConfigurationSchemaItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkAppConfigurationSchemaItemable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidForWorkAppConfigurationSchemaItemable)
            }
            m.SetSchemaItems(res)
        }
        return nil
    }
    return res
}
// GetSchemaItems gets the schemaItems property value. Collection of items each representing a named configuration option in the schema
func (m *AndroidForWorkAppConfigurationSchema) GetSchemaItems()([]AndroidForWorkAppConfigurationSchemaItemable) {
    return m.schemaItems
}
// Serialize serializes information the current object
func (m *AndroidForWorkAppConfigurationSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("exampleJson", m.GetExampleJson())
        if err != nil {
            return err
        }
    }
    if m.GetSchemaItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchemaItems()))
        for i, v := range m.GetSchemaItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("schemaItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExampleJson sets the exampleJson property value. UTF8 encoded byte array containing example JSON string conforming to this schema that demonstrates how to set the configuration for this app
func (m *AndroidForWorkAppConfigurationSchema) SetExampleJson(value []byte)() {
    m.exampleJson = value
}
// SetSchemaItems sets the schemaItems property value. Collection of items each representing a named configuration option in the schema
func (m *AndroidForWorkAppConfigurationSchema) SetSchemaItems(value []AndroidForWorkAppConfigurationSchemaItemable)() {
    m.schemaItems = value
}
