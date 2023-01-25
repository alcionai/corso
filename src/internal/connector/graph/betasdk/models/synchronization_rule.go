package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationRule 
type SynchronizationRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // true if the synchronization rule can be customized; false if this rule is read-only and should not be changed.
    editable *bool
    // Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported rule identifiers can be found in the synchronization template returned by the API.
    id *string
    // Additional extension properties. Unless instructed explicitly by the support team, metadata values should not be changed.
    metadata []StringKeyStringValuePairable
    // Human-readable name of the synchronization rule. Not nullable.
    name *string
    // Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be synchronized.
    objectMappings []ObjectMappingable
    // The OdataType property
    odataType *string
    // Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be processed first.
    priority *int32
    // Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
    sourceDirectoryName *string
    // Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
    targetDirectoryName *string
}
// NewSynchronizationRule instantiates a new synchronizationRule and sets the default values.
func NewSynchronizationRule()(*SynchronizationRule) {
    m := &SynchronizationRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEditable gets the editable property value. true if the synchronization rule can be customized; false if this rule is read-only and should not be changed.
func (m *SynchronizationRule) GetEditable()(*bool) {
    return m.editable
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["editable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditable(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStringKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyStringValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(StringKeyStringValuePairable)
            }
            m.SetMetadata(res)
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
    res["objectMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateObjectMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ObjectMappingable, len(val))
            for i, v := range val {
                res[i] = v.(ObjectMappingable)
            }
            m.SetObjectMappings(res)
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
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["sourceDirectoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceDirectoryName(val)
        }
        return nil
    }
    res["targetDirectoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDirectoryName(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported rule identifiers can be found in the synchronization template returned by the API.
func (m *SynchronizationRule) GetId()(*string) {
    return m.id
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless instructed explicitly by the support team, metadata values should not be changed.
func (m *SynchronizationRule) GetMetadata()([]StringKeyStringValuePairable) {
    return m.metadata
}
// GetName gets the name property value. Human-readable name of the synchronization rule. Not nullable.
func (m *SynchronizationRule) GetName()(*string) {
    return m.name
}
// GetObjectMappings gets the objectMappings property value. Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be synchronized.
func (m *SynchronizationRule) GetObjectMappings()([]ObjectMappingable) {
    return m.objectMappings
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationRule) GetOdataType()(*string) {
    return m.odataType
}
// GetPriority gets the priority property value. Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be processed first.
func (m *SynchronizationRule) GetPriority()(*int32) {
    return m.priority
}
// GetSourceDirectoryName gets the sourceDirectoryName property value. Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) GetSourceDirectoryName()(*string) {
    return m.sourceDirectoryName
}
// GetTargetDirectoryName gets the targetDirectoryName property value. Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) GetTargetDirectoryName()(*string) {
    return m.targetDirectoryName
}
// Serialize serializes information the current object
func (m *SynchronizationRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("editable", m.GetEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
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
    if m.GetObjectMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetObjectMappings()))
        for i, v := range m.GetObjectMappings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("objectMappings", cast)
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
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceDirectoryName", m.GetSourceDirectoryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDirectoryName", m.GetTargetDirectoryName())
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
func (m *SynchronizationRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEditable sets the editable property value. true if the synchronization rule can be customized; false if this rule is read-only and should not be changed.
func (m *SynchronizationRule) SetEditable(value *bool)() {
    m.editable = value
}
// SetId sets the id property value. Synchronization rule identifier. Must be one of the identifiers recognized by the synchronization engine. Supported rule identifiers can be found in the synchronization template returned by the API.
func (m *SynchronizationRule) SetId(value *string)() {
    m.id = value
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless instructed explicitly by the support team, metadata values should not be changed.
func (m *SynchronizationRule) SetMetadata(value []StringKeyStringValuePairable)() {
    m.metadata = value
}
// SetName sets the name property value. Human-readable name of the synchronization rule. Not nullable.
func (m *SynchronizationRule) SetName(value *string)() {
    m.name = value
}
// SetObjectMappings sets the objectMappings property value. Collection of object mappings supported by the rule. Tells the synchronization engine which objects should be synchronized.
func (m *SynchronizationRule) SetObjectMappings(value []ObjectMappingable)() {
    m.objectMappings = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPriority sets the priority property value. Priority relative to other rules in the synchronizationSchema. Rules with the lowest priority number will be processed first.
func (m *SynchronizationRule) SetPriority(value *int32)() {
    m.priority = value
}
// SetSourceDirectoryName sets the sourceDirectoryName property value. Name of the source directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) SetSourceDirectoryName(value *string)() {
    m.sourceDirectoryName = value
}
// SetTargetDirectoryName sets the targetDirectoryName property value. Name of the target directory. Must match one of the directory definitions in synchronizationSchema.
func (m *SynchronizationRule) SetTargetDirectoryName(value *string)() {
    m.targetDirectoryName = value
}
