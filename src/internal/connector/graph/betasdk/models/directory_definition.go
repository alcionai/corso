package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryDefinition provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DirectoryDefinition struct {
    Entity
    // The discoverabilities property
    discoverabilities *DirectoryDefinitionDiscoverabilities
    // Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    discoveryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the directory. Must be unique within the synchronization schema. Not nullable.
    name *string
    // Collection of objects supported by the directory.
    objects []ObjectDefinitionable
    // The readOnly property
    readOnly *bool
    // Read only value that indicates version discovered. null if discovery has not yet occurred.
    version *string
}
// NewDirectoryDefinition instantiates a new directoryDefinition and sets the default values.
func NewDirectoryDefinition()(*DirectoryDefinition) {
    m := &DirectoryDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryDefinition(), nil
}
// GetDiscoverabilities gets the discoverabilities property value. The discoverabilities property
func (m *DirectoryDefinition) GetDiscoverabilities()(*DirectoryDefinitionDiscoverabilities) {
    return m.discoverabilities
}
// GetDiscoveryDateTime gets the discoveryDateTime property value. Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DirectoryDefinition) GetDiscoveryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.discoveryDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["discoverabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryDefinitionDiscoverabilities)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverabilities(val.(*DirectoryDefinitionDiscoverabilities))
        }
        return nil
    }
    res["discoveryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoveryDateTime(val)
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
    res["objects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateObjectDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ObjectDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ObjectDefinitionable)
            }
            m.SetObjects(res)
        }
        return nil
    }
    res["readOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnly(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the directory. Must be unique within the synchronization schema. Not nullable.
func (m *DirectoryDefinition) GetName()(*string) {
    return m.name
}
// GetObjects gets the objects property value. Collection of objects supported by the directory.
func (m *DirectoryDefinition) GetObjects()([]ObjectDefinitionable) {
    return m.objects
}
// GetReadOnly gets the readOnly property value. The readOnly property
func (m *DirectoryDefinition) GetReadOnly()(*bool) {
    return m.readOnly
}
// GetVersion gets the version property value. Read only value that indicates version discovered. null if discovery has not yet occurred.
func (m *DirectoryDefinition) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DirectoryDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDiscoverabilities() != nil {
        cast := (*m.GetDiscoverabilities()).String()
        err = writer.WriteStringValue("discoverabilities", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("discoveryDateTime", m.GetDiscoveryDateTime())
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
    if m.GetObjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetObjects()))
        for i, v := range m.GetObjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("objects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("readOnly", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDiscoverabilities sets the discoverabilities property value. The discoverabilities property
func (m *DirectoryDefinition) SetDiscoverabilities(value *DirectoryDefinitionDiscoverabilities)() {
    m.discoverabilities = value
}
// SetDiscoveryDateTime sets the discoveryDateTime property value. Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DirectoryDefinition) SetDiscoveryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.discoveryDateTime = value
}
// SetName sets the name property value. Name of the directory. Must be unique within the synchronization schema. Not nullable.
func (m *DirectoryDefinition) SetName(value *string)() {
    m.name = value
}
// SetObjects sets the objects property value. Collection of objects supported by the directory.
func (m *DirectoryDefinition) SetObjects(value []ObjectDefinitionable)() {
    m.objects = value
}
// SetReadOnly sets the readOnly property value. The readOnly property
func (m *DirectoryDefinition) SetReadOnly(value *bool)() {
    m.readOnly = value
}
// SetVersion sets the version property value. Read only value that indicates version discovered. null if discovery has not yet occurred.
func (m *DirectoryDefinition) SetVersion(value *string)() {
    m.version = value
}
