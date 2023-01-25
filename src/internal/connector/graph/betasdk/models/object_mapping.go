package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectMapping 
type ObjectMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Attribute mappings define which attributes to map from the source object into the target object and how they should flow. A number of functions are available to support the transformation of the original source values.
    attributeMappings []AttributeMappingable
    // When true, this object mapping will be processed during synchronization. When false, this object mapping will be skipped.
    enabled *bool
    // The flowTypes property
    flowTypes *ObjectFlowTypes
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntryable
    // Human-friendly name of the object mapping.
    name *string
    // The OdataType property
    odataType *string
    // Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want to only provision users that are located in the US.
    scope Filterable
    // Name of the object in the source directory. Must match the object name from the source directory definition.
    sourceObjectName *string
    // Name of the object in target directory. Must match the object name from the target directory definition.
    targetObjectName *string
}
// NewObjectMapping instantiates a new objectMapping and sets the default values.
func NewObjectMapping()(*ObjectMapping) {
    m := &ObjectMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateObjectMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectMapping(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectMapping) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttributeMappings gets the attributeMappings property value. Attribute mappings define which attributes to map from the source object into the target object and how they should flow. A number of functions are available to support the transformation of the original source values.
func (m *ObjectMapping) GetAttributeMappings()([]AttributeMappingable) {
    return m.attributeMappings
}
// GetEnabled gets the enabled property value. When true, this object mapping will be processed during synchronization. When false, this object mapping will be skipped.
func (m *ObjectMapping) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributeMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeMappingable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeMappingable)
            }
            m.SetAttributeMappings(res)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["flowTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectFlowTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlowTypes(val.(*ObjectFlowTypes))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetadataEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetadataEntryable, len(val))
            for i, v := range val {
                res[i] = v.(MetadataEntryable)
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
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(Filterable))
        }
        return nil
    }
    res["sourceObjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceObjectName(val)
        }
        return nil
    }
    res["targetObjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetObjectName(val)
        }
        return nil
    }
    return res
}
// GetFlowTypes gets the flowTypes property value. The flowTypes property
func (m *ObjectMapping) GetFlowTypes()(*ObjectFlowTypes) {
    return m.flowTypes
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *ObjectMapping) GetMetadata()([]MetadataEntryable) {
    return m.metadata
}
// GetName gets the name property value. Human-friendly name of the object mapping.
func (m *ObjectMapping) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ObjectMapping) GetOdataType()(*string) {
    return m.odataType
}
// GetScope gets the scope property value. Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want to only provision users that are located in the US.
func (m *ObjectMapping) GetScope()(Filterable) {
    return m.scope
}
// GetSourceObjectName gets the sourceObjectName property value. Name of the object in the source directory. Must match the object name from the source directory definition.
func (m *ObjectMapping) GetSourceObjectName()(*string) {
    return m.sourceObjectName
}
// GetTargetObjectName gets the targetObjectName property value. Name of the object in target directory. Must match the object name from the target directory definition.
func (m *ObjectMapping) GetTargetObjectName()(*string) {
    return m.targetObjectName
}
// Serialize serializes information the current object
func (m *ObjectMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttributeMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributeMappings()))
        for i, v := range m.GetAttributeMappings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attributeMappings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetFlowTypes() != nil {
        cast := (*m.GetFlowTypes()).String()
        err := writer.WriteStringValue("flowTypes", &cast)
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceObjectName", m.GetSourceObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetObjectName", m.GetTargetObjectName())
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
func (m *ObjectMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttributeMappings sets the attributeMappings property value. Attribute mappings define which attributes to map from the source object into the target object and how they should flow. A number of functions are available to support the transformation of the original source values.
func (m *ObjectMapping) SetAttributeMappings(value []AttributeMappingable)() {
    m.attributeMappings = value
}
// SetEnabled sets the enabled property value. When true, this object mapping will be processed during synchronization. When false, this object mapping will be skipped.
func (m *ObjectMapping) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetFlowTypes sets the flowTypes property value. The flowTypes property
func (m *ObjectMapping) SetFlowTypes(value *ObjectFlowTypes)() {
    m.flowTypes = value
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *ObjectMapping) SetMetadata(value []MetadataEntryable)() {
    m.metadata = value
}
// SetName sets the name property value. Human-friendly name of the object mapping.
func (m *ObjectMapping) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ObjectMapping) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScope sets the scope property value. Defines a filter to be used when deciding whether a given object should be provisioned. For example, you might want to only provision users that are located in the US.
func (m *ObjectMapping) SetScope(value Filterable)() {
    m.scope = value
}
// SetSourceObjectName sets the sourceObjectName property value. Name of the object in the source directory. Must match the object name from the source directory definition.
func (m *ObjectMapping) SetSourceObjectName(value *string)() {
    m.sourceObjectName = value
}
// SetTargetObjectName sets the targetObjectName property value. Name of the object in target directory. Must match the object name from the target directory definition.
func (m *ObjectMapping) SetTargetObjectName(value *string)() {
    m.targetObjectName = value
}
