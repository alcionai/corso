package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplate 
type TeamTemplate struct {
    Entity
    // The definitions property
    definitions []TeamTemplateDefinitionable
}
// NewTeamTemplate instantiates a new TeamTemplate and sets the default values.
func NewTeamTemplate()(*TeamTemplate) {
    m := &TeamTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplate(), nil
}
// GetDefinitions gets the definitions property value. The definitions property
func (m *TeamTemplate) GetDefinitions()([]TeamTemplateDefinitionable) {
    return m.definitions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamTemplateDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamTemplateDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(TeamTemplateDefinitionable)
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinitions sets the definitions property value. The definitions property
func (m *TeamTemplate) SetDefinitions(value []TeamTemplateDefinitionable)() {
    m.definitions = value
}
