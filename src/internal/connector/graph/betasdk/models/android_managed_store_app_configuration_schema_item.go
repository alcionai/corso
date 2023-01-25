package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppConfigurationSchemaItem single configuration item inside an Android application's custom configuration schema.
type AndroidManagedStoreAppConfigurationSchemaItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Data type for a configuration item inside an Android application's custom configuration schema
    dataType *AndroidManagedStoreAppConfigurationSchemaItemDataType
    // Default value for boolean type items, if specified by the app developer
    defaultBoolValue *bool
    // Default value for integer type items, if specified by the app developer
    defaultIntValue *int32
    // Default value for string array type items, if specified by the app developer
    defaultStringArrayValue []string
    // Default value for string type items, if specified by the app developer
    defaultStringValue *string
    // Description of what the item controls within the application
    description *string
    // Human readable name
    displayName *string
    // Unique index the application uses to maintain nested schema items
    index *int32
    // The OdataType property
    odataType *string
    // Index of parent schema item to track nested schema items
    parentIndex *int32
    // Unique key the application uses to identify the item
    schemaItemKey *string
    // List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
    selections []KeyValuePairable
}
// NewAndroidManagedStoreAppConfigurationSchemaItem instantiates a new androidManagedStoreAppConfigurationSchemaItem and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchemaItem()(*AndroidManagedStoreAppConfigurationSchemaItem) {
    m := &AndroidManagedStoreAppConfigurationSchemaItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppConfigurationSchemaItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppConfigurationSchemaItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDataType gets the dataType property value. Data type for a configuration item inside an Android application's custom configuration schema
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDataType()(*AndroidManagedStoreAppConfigurationSchemaItemDataType) {
    return m.dataType
}
// GetDefaultBoolValue gets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultBoolValue()(*bool) {
    return m.defaultBoolValue
}
// GetDefaultIntValue gets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultIntValue()(*int32) {
    return m.defaultIntValue
}
// GetDefaultStringArrayValue gets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultStringArrayValue()([]string) {
    return m.defaultStringArrayValue
}
// GetDefaultStringValue gets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDefaultStringValue()(*string) {
    return m.defaultStringValue
}
// GetDescription gets the description property value. Description of what the item controls within the application
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Human readable name
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAppConfigurationSchemaItemDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*AndroidManagedStoreAppConfigurationSchemaItemDataType))
        }
        return nil
    }
    res["defaultBoolValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultBoolValue(val)
        }
        return nil
    }
    res["defaultIntValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultIntValue(val)
        }
        return nil
    }
    res["defaultStringArrayValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefaultStringArrayValue(res)
        }
        return nil
    }
    res["defaultStringValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultStringValue(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
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
    res["parentIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentIndex(val)
        }
        return nil
    }
    res["schemaItemKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaItemKey(val)
        }
        return nil
    }
    res["selections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetSelections(res)
        }
        return nil
    }
    return res
}
// GetIndex gets the index property value. Unique index the application uses to maintain nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetIndex()(*int32) {
    return m.index
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetOdataType()(*string) {
    return m.odataType
}
// GetParentIndex gets the parentIndex property value. Index of parent schema item to track nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetParentIndex()(*int32) {
    return m.parentIndex
}
// GetSchemaItemKey gets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetSchemaItemKey()(*string) {
    return m.schemaItemKey
}
// GetSelections gets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidManagedStoreAppConfigurationSchemaItem) GetSelections()([]KeyValuePairable) {
    return m.selections
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppConfigurationSchemaItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := (*m.GetDataType()).String()
        err := writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("defaultBoolValue", m.GetDefaultBoolValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("defaultIntValue", m.GetDefaultIntValue())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultStringArrayValue() != nil {
        err := writer.WriteCollectionOfStringValues("defaultStringArrayValue", m.GetDefaultStringArrayValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultStringValue", m.GetDefaultStringValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
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
        err := writer.WriteInt32Value("parentIndex", m.GetParentIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("schemaItemKey", m.GetSchemaItemKey())
        if err != nil {
            return err
        }
    }
    if m.GetSelections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSelections()))
        for i, v := range m.GetSelections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("selections", cast)
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
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDataType sets the dataType property value. Data type for a configuration item inside an Android application's custom configuration schema
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDataType(value *AndroidManagedStoreAppConfigurationSchemaItemDataType)() {
    m.dataType = value
}
// SetDefaultBoolValue sets the defaultBoolValue property value. Default value for boolean type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultBoolValue(value *bool)() {
    m.defaultBoolValue = value
}
// SetDefaultIntValue sets the defaultIntValue property value. Default value for integer type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultIntValue(value *int32)() {
    m.defaultIntValue = value
}
// SetDefaultStringArrayValue sets the defaultStringArrayValue property value. Default value for string array type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultStringArrayValue(value []string)() {
    m.defaultStringArrayValue = value
}
// SetDefaultStringValue sets the defaultStringValue property value. Default value for string type items, if specified by the app developer
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDefaultStringValue(value *string)() {
    m.defaultStringValue = value
}
// SetDescription sets the description property value. Description of what the item controls within the application
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Human readable name
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIndex sets the index property value. Unique index the application uses to maintain nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetIndex(value *int32)() {
    m.index = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParentIndex sets the parentIndex property value. Index of parent schema item to track nested schema items
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetParentIndex(value *int32)() {
    m.parentIndex = value
}
// SetSchemaItemKey sets the schemaItemKey property value. Unique key the application uses to identify the item
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetSchemaItemKey(value *string)() {
    m.schemaItemKey = value
}
// SetSelections sets the selections property value. List of human readable name/value pairs for the valid values that can be set for this item (Choice and Multiselect items only)
func (m *AndroidManagedStoreAppConfigurationSchemaItem) SetSelections(value []KeyValuePairable)() {
    m.selections = value
}
