package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomSecurityAttributeDefinition provides operations to manage the collection of site entities.
type CustomSecurityAttributeDefinition struct {
    Entity
    // Values that are predefined for this custom security attribute.This navigation property is not returned by default and must be specified in an $expand query. For example, /directory/customSecurityAttributeDefinitions?$expand=allowedValues.
    allowedValues []AllowedValueable
    // Name of the attribute set. Case insensitive.
    attributeSet *string
    // Description of the custom security attribute. Can be up to 128 characters long and include Unicode characters. Can be changed later.
    description *string
    // Indicates whether multiple values can be assigned to the custom security attribute. Cannot be changed later. If type is set to Boolean, isCollection cannot be set to true.
    isCollection *bool
    // Indicates whether custom security attribute values will be indexed for searching on objects that are assigned attribute values. Cannot be changed later.
    isSearchable *bool
    // Name of the custom security attribute. Must be unique within an attribute set. Can be up to 32 characters long and include Unicode characters. Cannot contain spaces or special characters. Cannot be changed later. Case insensitive.
    name *string
    // Specifies whether the custom security attribute is active or deactivated. Acceptable values are Available and Deprecated. Can be changed later.
    status *string
    // Data type for the custom security attribute values. Supported types are Boolean, Integer, and String. Cannot be changed later.
    type_escaped *string
    // Indicates whether only predefined values can be assigned to the custom security attribute. If set to false, free-form values are allowed. Can later be changed from true to false, but cannot be changed from false to true. If type is set to Boolean, usePreDefinedValuesOnly cannot be set to true.
    usePreDefinedValuesOnly *bool
}
// NewCustomSecurityAttributeDefinition instantiates a new customSecurityAttributeDefinition and sets the default values.
func NewCustomSecurityAttributeDefinition()(*CustomSecurityAttributeDefinition) {
    m := &CustomSecurityAttributeDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomSecurityAttributeDefinition(), nil
}
// GetAllowedValues gets the allowedValues property value. Values that are predefined for this custom security attribute.This navigation property is not returned by default and must be specified in an $expand query. For example, /directory/customSecurityAttributeDefinitions?$expand=allowedValues.
func (m *CustomSecurityAttributeDefinition) GetAllowedValues()([]AllowedValueable) {
    return m.allowedValues
}
// GetAttributeSet gets the attributeSet property value. Name of the attribute set. Case insensitive.
func (m *CustomSecurityAttributeDefinition) GetAttributeSet()(*string) {
    return m.attributeSet
}
// GetDescription gets the description property value. Description of the custom security attribute. Can be up to 128 characters long and include Unicode characters. Can be changed later.
func (m *CustomSecurityAttributeDefinition) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomSecurityAttributeDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAllowedValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AllowedValueable, len(val))
            for i, v := range val {
                res[i] = v.(AllowedValueable)
            }
            m.SetAllowedValues(res)
        }
        return nil
    }
    res["attributeSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeSet(val)
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
    res["isCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCollection(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["usePreDefinedValuesOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsePreDefinedValuesOnly(val)
        }
        return nil
    }
    return res
}
// GetIsCollection gets the isCollection property value. Indicates whether multiple values can be assigned to the custom security attribute. Cannot be changed later. If type is set to Boolean, isCollection cannot be set to true.
func (m *CustomSecurityAttributeDefinition) GetIsCollection()(*bool) {
    return m.isCollection
}
// GetIsSearchable gets the isSearchable property value. Indicates whether custom security attribute values will be indexed for searching on objects that are assigned attribute values. Cannot be changed later.
func (m *CustomSecurityAttributeDefinition) GetIsSearchable()(*bool) {
    return m.isSearchable
}
// GetName gets the name property value. Name of the custom security attribute. Must be unique within an attribute set. Can be up to 32 characters long and include Unicode characters. Cannot contain spaces or special characters. Cannot be changed later. Case insensitive.
func (m *CustomSecurityAttributeDefinition) GetName()(*string) {
    return m.name
}
// GetStatus gets the status property value. Specifies whether the custom security attribute is active or deactivated. Acceptable values are Available and Deprecated. Can be changed later.
func (m *CustomSecurityAttributeDefinition) GetStatus()(*string) {
    return m.status
}
// GetType gets the type property value. Data type for the custom security attribute values. Supported types are Boolean, Integer, and String. Cannot be changed later.
func (m *CustomSecurityAttributeDefinition) GetType()(*string) {
    return m.type_escaped
}
// GetUsePreDefinedValuesOnly gets the usePreDefinedValuesOnly property value. Indicates whether only predefined values can be assigned to the custom security attribute. If set to false, free-form values are allowed. Can later be changed from true to false, but cannot be changed from false to true. If type is set to Boolean, usePreDefinedValuesOnly cannot be set to true.
func (m *CustomSecurityAttributeDefinition) GetUsePreDefinedValuesOnly()(*bool) {
    return m.usePreDefinedValuesOnly
}
// Serialize serializes information the current object
func (m *CustomSecurityAttributeDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedValues()))
        for i, v := range m.GetAllowedValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("allowedValues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("attributeSet", m.GetAttributeSet())
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
        err = writer.WriteBoolValue("isCollection", m.GetIsCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usePreDefinedValuesOnly", m.GetUsePreDefinedValuesOnly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedValues sets the allowedValues property value. Values that are predefined for this custom security attribute.This navigation property is not returned by default and must be specified in an $expand query. For example, /directory/customSecurityAttributeDefinitions?$expand=allowedValues.
func (m *CustomSecurityAttributeDefinition) SetAllowedValues(value []AllowedValueable)() {
    m.allowedValues = value
}
// SetAttributeSet sets the attributeSet property value. Name of the attribute set. Case insensitive.
func (m *CustomSecurityAttributeDefinition) SetAttributeSet(value *string)() {
    m.attributeSet = value
}
// SetDescription sets the description property value. Description of the custom security attribute. Can be up to 128 characters long and include Unicode characters. Can be changed later.
func (m *CustomSecurityAttributeDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetIsCollection sets the isCollection property value. Indicates whether multiple values can be assigned to the custom security attribute. Cannot be changed later. If type is set to Boolean, isCollection cannot be set to true.
func (m *CustomSecurityAttributeDefinition) SetIsCollection(value *bool)() {
    m.isCollection = value
}
// SetIsSearchable sets the isSearchable property value. Indicates whether custom security attribute values will be indexed for searching on objects that are assigned attribute values. Cannot be changed later.
func (m *CustomSecurityAttributeDefinition) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetName sets the name property value. Name of the custom security attribute. Must be unique within an attribute set. Can be up to 32 characters long and include Unicode characters. Cannot contain spaces or special characters. Cannot be changed later. Case insensitive.
func (m *CustomSecurityAttributeDefinition) SetName(value *string)() {
    m.name = value
}
// SetStatus sets the status property value. Specifies whether the custom security attribute is active or deactivated. Acceptable values are Available and Deprecated. Can be changed later.
func (m *CustomSecurityAttributeDefinition) SetStatus(value *string)() {
    m.status = value
}
// SetType sets the type property value. Data type for the custom security attribute values. Supported types are Boolean, Integer, and String. Cannot be changed later.
func (m *CustomSecurityAttributeDefinition) SetType(value *string)() {
    m.type_escaped = value
}
// SetUsePreDefinedValuesOnly sets the usePreDefinedValuesOnly property value. Indicates whether only predefined values can be assigned to the custom security attribute. If set to false, free-form values are allowed. Can later be changed from true to false, but cannot be changed from false to true. If type is set to Boolean, usePreDefinedValuesOnly cannot be set to true.
func (m *CustomSecurityAttributeDefinition) SetUsePreDefinedValuesOnly(value *bool)() {
    m.usePreDefinedValuesOnly = value
}
