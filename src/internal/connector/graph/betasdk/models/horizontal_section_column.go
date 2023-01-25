package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HorizontalSectionColumn provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type HorizontalSectionColumn struct {
    Entity
    // The collection of WebParts in this column.
    webparts []WebPartable
    // Width of the column. A horizontal section is divided into 12 grids. A column should have a value of 1-12 to represent its range spans. For example, there can be two columns both have a width of 6 in a section.
    width *int32
}
// NewHorizontalSectionColumn instantiates a new horizontalSectionColumn and sets the default values.
func NewHorizontalSectionColumn()(*HorizontalSectionColumn) {
    m := &HorizontalSectionColumn{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHorizontalSectionColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHorizontalSectionColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHorizontalSectionColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HorizontalSectionColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["webparts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebPartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebPartable, len(val))
            for i, v := range val {
                res[i] = v.(WebPartable)
            }
            m.SetWebparts(res)
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetWebparts gets the webparts property value. The collection of WebParts in this column.
func (m *HorizontalSectionColumn) GetWebparts()([]WebPartable) {
    return m.webparts
}
// GetWidth gets the width property value. Width of the column. A horizontal section is divided into 12 grids. A column should have a value of 1-12 to represent its range spans. For example, there can be two columns both have a width of 6 in a section.
func (m *HorizontalSectionColumn) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *HorizontalSectionColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetWebparts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebparts()))
        for i, v := range m.GetWebparts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("webparts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWebparts sets the webparts property value. The collection of WebParts in this column.
func (m *HorizontalSectionColumn) SetWebparts(value []WebPartable)() {
    m.webparts = value
}
// SetWidth sets the width property value. Width of the column. A horizontal section is divided into 12 grids. A column should have a value of 1-12 to represent its range spans. For example, there can be two columns both have a width of 6 in a section.
func (m *HorizontalSectionColumn) SetWidth(value *int32)() {
    m.width = value
}
