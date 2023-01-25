package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Program 
type Program struct {
    Entity
    // Controls associated with the program.
    controls []ProgramControlable
    // The description of the program.
    description *string
    // The name of the program.  Required on create.
    displayName *string
}
// NewProgram instantiates a new Program and sets the default values.
func NewProgram()(*Program) {
    m := &Program{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProgramFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProgramFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProgram(), nil
}
// GetControls gets the controls property value. Controls associated with the program.
func (m *Program) GetControls()([]ProgramControlable) {
    return m.controls
}
// GetDescription gets the description property value. The description of the program.
func (m *Program) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the program.  Required on create.
func (m *Program) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Program) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProgramControlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProgramControlable, len(val))
            for i, v := range val {
                res[i] = v.(ProgramControlable)
            }
            m.SetControls(res)
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
    return res
}
// Serialize serializes information the current object
func (m *Program) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetControls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetControls()))
        for i, v := range m.GetControls() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("controls", cast)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetControls sets the controls property value. Controls associated with the program.
func (m *Program) SetControls(value []ProgramControlable)() {
    m.controls = value
}
// SetDescription sets the description property value. The description of the program.
func (m *Program) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the program.  Required on create.
func (m *Program) SetDisplayName(value *string)() {
    m.displayName = value
}
