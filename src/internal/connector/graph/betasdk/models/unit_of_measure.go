package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnitOfMeasure provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UnitOfMeasure struct {
    Entity
    // The code property
    code *string
    // The displayName property
    displayName *string
    // The internationalStandardCode property
    internationalStandardCode *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUnitOfMeasure instantiates a new unitOfMeasure and sets the default values.
func NewUnitOfMeasure()(*UnitOfMeasure) {
    m := &UnitOfMeasure{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnitOfMeasureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnitOfMeasureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnitOfMeasure(), nil
}
// GetCode gets the code property value. The code property
func (m *UnitOfMeasure) GetCode()(*string) {
    return m.code
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *UnitOfMeasure) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnitOfMeasure) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
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
    res["internationalStandardCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternationalStandardCode(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetInternationalStandardCode gets the internationalStandardCode property value. The internationalStandardCode property
func (m *UnitOfMeasure) GetInternationalStandardCode()(*string) {
    return m.internationalStandardCode
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UnitOfMeasure) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *UnitOfMeasure) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
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
        err = writer.WriteStringValue("internationalStandardCode", m.GetInternationalStandardCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCode sets the code property value. The code property
func (m *UnitOfMeasure) SetCode(value *string)() {
    m.code = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *UnitOfMeasure) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInternationalStandardCode sets the internationalStandardCode property value. The internationalStandardCode property
func (m *UnitOfMeasure) SetInternationalStandardCode(value *string)() {
    m.internationalStandardCode = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UnitOfMeasure) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
