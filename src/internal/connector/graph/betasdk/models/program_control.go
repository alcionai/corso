package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProgramControl provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ProgramControl struct {
    Entity
    // The controlId of the control, in particular the identifier of an access review. Required on create.
    controlId *string
    // The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
    controlTypeId *string
    // The creation date and time of the program control.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the control.
    displayName *string
    // The user who created the program control.
    owner UserIdentityable
    // The program this control is part of.
    program Programable
    // The programId of the program this control is a part of. Required on create.
    programId *string
    // The resource, a group or an app, targeted by this program control's access review.
    resource ProgramResourceable
    // The life cycle status of the control.
    status *string
}
// NewProgramControl instantiates a new programControl and sets the default values.
func NewProgramControl()(*ProgramControl) {
    m := &ProgramControl{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProgramControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProgramControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProgramControl(), nil
}
// GetControlId gets the controlId property value. The controlId of the control, in particular the identifier of an access review. Required on create.
func (m *ProgramControl) GetControlId()(*string) {
    return m.controlId
}
// GetControlTypeId gets the controlTypeId property value. The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
func (m *ProgramControl) GetControlTypeId()(*string) {
    return m.controlTypeId
}
// GetCreatedDateTime gets the createdDateTime property value. The creation date and time of the program control.
func (m *ProgramControl) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the control.
func (m *ProgramControl) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramControl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlId(val)
        }
        return nil
    }
    res["controlTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlTypeId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(UserIdentityable))
        }
        return nil
    }
    res["program"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProgramFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(Programable))
        }
        return nil
    }
    res["programId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgramId(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProgramResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(ProgramResourceable))
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
    return res
}
// GetOwner gets the owner property value. The user who created the program control.
func (m *ProgramControl) GetOwner()(UserIdentityable) {
    return m.owner
}
// GetProgram gets the program property value. The program this control is part of.
func (m *ProgramControl) GetProgram()(Programable) {
    return m.program
}
// GetProgramId gets the programId property value. The programId of the program this control is a part of. Required on create.
func (m *ProgramControl) GetProgramId()(*string) {
    return m.programId
}
// GetResource gets the resource property value. The resource, a group or an app, targeted by this program control's access review.
func (m *ProgramControl) GetResource()(ProgramResourceable) {
    return m.resource
}
// GetStatus gets the status property value. The life cycle status of the control.
func (m *ProgramControl) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *ProgramControl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("controlId", m.GetControlId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("controlTypeId", m.GetControlTypeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("program", m.GetProgram())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("programId", m.GetProgramId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
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
    return nil
}
// SetControlId sets the controlId property value. The controlId of the control, in particular the identifier of an access review. Required on create.
func (m *ProgramControl) SetControlId(value *string)() {
    m.controlId = value
}
// SetControlTypeId sets the controlTypeId property value. The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
func (m *ProgramControl) SetControlTypeId(value *string)() {
    m.controlTypeId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The creation date and time of the program control.
func (m *ProgramControl) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the control.
func (m *ProgramControl) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOwner sets the owner property value. The user who created the program control.
func (m *ProgramControl) SetOwner(value UserIdentityable)() {
    m.owner = value
}
// SetProgram sets the program property value. The program this control is part of.
func (m *ProgramControl) SetProgram(value Programable)() {
    m.program = value
}
// SetProgramId sets the programId property value. The programId of the program this control is a part of. Required on create.
func (m *ProgramControl) SetProgramId(value *string)() {
    m.programId = value
}
// SetResource sets the resource property value. The resource, a group or an app, targeted by this program control's access review.
func (m *ProgramControl) SetResource(value ProgramResourceable)() {
    m.resource = value
}
// SetStatus sets the status property value. The life cycle status of the control.
func (m *ProgramControl) SetStatus(value *string)() {
    m.status = value
}
