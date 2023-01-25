package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerticalSection 
type VerticalSection struct {
    Entity
    // Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural, soft, strong, unknownFutureValue.
    emphasis *SectionEmphasisType
    // The set of web parts in this section.
    webparts []WebPartable
}
// NewVerticalSection instantiates a new verticalSection and sets the default values.
func NewVerticalSection()(*VerticalSection) {
    m := &VerticalSection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVerticalSectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerticalSectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerticalSection(), nil
}
// GetEmphasis gets the emphasis property value. Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural, soft, strong, unknownFutureValue.
func (m *VerticalSection) GetEmphasis()(*SectionEmphasisType) {
    return m.emphasis
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerticalSection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emphasis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSectionEmphasisType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmphasis(val.(*SectionEmphasisType))
        }
        return nil
    }
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
    return res
}
// GetWebparts gets the webparts property value. The set of web parts in this section.
func (m *VerticalSection) GetWebparts()([]WebPartable) {
    return m.webparts
}
// Serialize serializes information the current object
func (m *VerticalSection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmphasis() != nil {
        cast := (*m.GetEmphasis()).String()
        err = writer.WriteStringValue("emphasis", &cast)
        if err != nil {
            return err
        }
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
    return nil
}
// SetEmphasis sets the emphasis property value. Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural, soft, strong, unknownFutureValue.
func (m *VerticalSection) SetEmphasis(value *SectionEmphasisType)() {
    m.emphasis = value
}
// SetWebparts sets the webparts property value. The set of web parts in this section.
func (m *VerticalSection) SetWebparts(value []WebPartable)() {
    m.webparts = value
}
