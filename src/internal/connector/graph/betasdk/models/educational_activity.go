package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationalActivity 
type EducationalActivity struct {
    ItemFacet
    // The month and year the user graduated or completed the activity.
    completionMonthYear *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The month and year the user completed the educational activity referenced.
    endMonthYear *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The institution property
    institution InstitutionDataable
    // The program property
    program EducationalActivityDetailable
    // The month and year the user commenced the activity referenced.
    startMonthYear *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewEducationalActivity instantiates a new EducationalActivity and sets the default values.
func NewEducationalActivity()(*EducationalActivity) {
    m := &EducationalActivity{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.educationalActivity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationalActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationalActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationalActivity(), nil
}
// GetCompletionMonthYear gets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) GetCompletionMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.completionMonthYear
}
// GetEndMonthYear gets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endMonthYear
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationalActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["completionMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionMonthYear(val)
        }
        return nil
    }
    res["endMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonthYear(val)
        }
        return nil
    }
    res["institution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInstitutionDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstitution(val.(InstitutionDataable))
        }
        return nil
    }
    res["program"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationalActivityDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(EducationalActivityDetailable))
        }
        return nil
    }
    res["startMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonthYear(val)
        }
        return nil
    }
    return res
}
// GetInstitution gets the institution property value. The institution property
func (m *EducationalActivity) GetInstitution()(InstitutionDataable) {
    return m.institution
}
// GetProgram gets the program property value. The program property
func (m *EducationalActivity) GetProgram()(EducationalActivityDetailable) {
    return m.program
}
// GetStartMonthYear gets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startMonthYear
}
// Serialize serializes information the current object
func (m *EducationalActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("completionMonthYear", m.GetCompletionMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endMonthYear", m.GetEndMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("institution", m.GetInstitution())
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
        err = writer.WriteDateOnlyValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletionMonthYear sets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) SetCompletionMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.completionMonthYear = value
}
// SetEndMonthYear sets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endMonthYear = value
}
// SetInstitution sets the institution property value. The institution property
func (m *EducationalActivity) SetInstitution(value InstitutionDataable)() {
    m.institution = value
}
// SetProgram sets the program property value. The program property
func (m *EducationalActivity) SetProgram(value EducationalActivityDetailable)() {
    m.program = value
}
// SetStartMonthYear sets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startMonthYear = value
}
