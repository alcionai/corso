package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFormatProtection 
type WorkbookFormatProtection struct {
    Entity
    // The formulaHidden property
    formulaHidden *bool
    // The locked property
    locked *bool
}
// NewWorkbookFormatProtection instantiates a new WorkbookFormatProtection and sets the default values.
func NewWorkbookFormatProtection()(*WorkbookFormatProtection) {
    m := &WorkbookFormatProtection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFormatProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFormatProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFormatProtection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFormatProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["formulaHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulaHidden(val)
        }
        return nil
    }
    res["locked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocked(val)
        }
        return nil
    }
    return res
}
// GetFormulaHidden gets the formulaHidden property value. The formulaHidden property
func (m *WorkbookFormatProtection) GetFormulaHidden()(*bool) {
    return m.formulaHidden
}
// GetLocked gets the locked property value. The locked property
func (m *WorkbookFormatProtection) GetLocked()(*bool) {
    return m.locked
}
// Serialize serializes information the current object
func (m *WorkbookFormatProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("formulaHidden", m.GetFormulaHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locked", m.GetLocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormulaHidden sets the formulaHidden property value. The formulaHidden property
func (m *WorkbookFormatProtection) SetFormulaHidden(value *bool)() {
    m.formulaHidden = value
}
// SetLocked sets the locked property value. The locked property
func (m *WorkbookFormatProtection) SetLocked(value *bool)() {
    m.locked = value
}
