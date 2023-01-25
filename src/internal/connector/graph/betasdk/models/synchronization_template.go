package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationTemplate provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SynchronizationTemplate struct {
    Entity
    // Identifier of the application this template belongs to.
    applicationId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // true if this template is recommended to be the default for the application.
    default_escaped *bool
    // Description of the template.
    description *string
    // true if this template should appear in the collection of templates available for the application instance (service principal).
    discoverable *bool
    // One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
    factoryTag *string
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntryable
    // Default synchronization schema for the jobs based on this template.
    schema SynchronizationSchemaable
}
// NewSynchronizationTemplate instantiates a new synchronizationTemplate and sets the default values.
func NewSynchronizationTemplate()(*SynchronizationTemplate) {
    m := &SynchronizationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationTemplate(), nil
}
// GetApplicationId gets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) GetApplicationId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.applicationId
}
// GetDefault gets the default property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) GetDefault()(*bool) {
    return m.default_escaped
}
// GetDescription gets the description property value. Description of the template.
func (m *SynchronizationTemplate) GetDescription()(*string) {
    return m.description
}
// GetDiscoverable gets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) GetDiscoverable()(*bool) {
    return m.discoverable
}
// GetFactoryTag gets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) GetFactoryTag()(*string) {
    return m.factoryTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault(val)
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
    res["discoverable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverable(val)
        }
        return nil
    }
    res["factoryTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactoryTag(val)
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
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(SynchronizationSchemaable))
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) GetMetadata()([]MetadataEntryable) {
    return m.metadata
}
// GetSchema gets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) GetSchema()(SynchronizationSchemaable) {
    return m.schema
}
// Serialize serializes information the current object
func (m *SynchronizationTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteUUIDValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("default", m.GetDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("discoverable", m.GetDiscoverable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("factoryTag", m.GetFactoryTag())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationId sets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) SetApplicationId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.applicationId = value
}
// SetDefault sets the default property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) SetDefault(value *bool)() {
    m.default_escaped = value
}
// SetDescription sets the description property value. Description of the template.
func (m *SynchronizationTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDiscoverable sets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) SetDiscoverable(value *bool)() {
    m.discoverable = value
}
// SetFactoryTag sets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) SetFactoryTag(value *string)() {
    m.factoryTag = value
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) SetMetadata(value []MetadataEntryable)() {
    m.metadata = value
}
// SetSchema sets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) SetSchema(value SynchronizationSchemaable)() {
    m.schema = value
}
