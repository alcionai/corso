package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HorizontalSection provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type HorizontalSection struct {
    Entity
    // The set of vertical columns in this section.
    columns []HorizontalSectionColumnable
    // Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural, soft, strong, unknownFutureValue.
    emphasis *SectionEmphasisType
    // Layout type of the section. The possible values are: none, oneColumn, twoColumns, threeColumns, oneThirdLeftColumn, oneThirdRightColumn, fullWidth, unknownFutureValue.
    layout *HorizontalSectionLayoutType
}
// NewHorizontalSection instantiates a new horizontalSection and sets the default values.
func NewHorizontalSection()(*HorizontalSection) {
    m := &HorizontalSection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHorizontalSectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHorizontalSectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHorizontalSection(), nil
}
// GetColumns gets the columns property value. The set of vertical columns in this section.
func (m *HorizontalSection) GetColumns()([]HorizontalSectionColumnable) {
    return m.columns
}
// GetEmphasis gets the emphasis property value. Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural, soft, strong, unknownFutureValue.
func (m *HorizontalSection) GetEmphasis()(*SectionEmphasisType) {
    return m.emphasis
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HorizontalSection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["columns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHorizontalSectionColumnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HorizontalSectionColumnable, len(val))
            for i, v := range val {
                res[i] = v.(HorizontalSectionColumnable)
            }
            m.SetColumns(res)
        }
        return nil
    }
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
    res["layout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHorizontalSectionLayoutType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLayout(val.(*HorizontalSectionLayoutType))
        }
        return nil
    }
    return res
}
// GetLayout gets the layout property value. Layout type of the section. The possible values are: none, oneColumn, twoColumns, threeColumns, oneThirdLeftColumn, oneThirdRightColumn, fullWidth, unknownFutureValue.
func (m *HorizontalSection) GetLayout()(*HorizontalSectionLayoutType) {
    return m.layout
}
// Serialize serializes information the current object
func (m *HorizontalSection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColumns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmphasis() != nil {
        cast := (*m.GetEmphasis()).String()
        err = writer.WriteStringValue("emphasis", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLayout() != nil {
        cast := (*m.GetLayout()).String()
        err = writer.WriteStringValue("layout", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColumns sets the columns property value. The set of vertical columns in this section.
func (m *HorizontalSection) SetColumns(value []HorizontalSectionColumnable)() {
    m.columns = value
}
// SetEmphasis sets the emphasis property value. Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural, soft, strong, unknownFutureValue.
func (m *HorizontalSection) SetEmphasis(value *SectionEmphasisType)() {
    m.emphasis = value
}
// SetLayout sets the layout property value. Layout type of the section. The possible values are: none, oneColumn, twoColumns, threeColumns, oneThirdLeftColumn, oneThirdRightColumn, fullWidth, unknownFutureValue.
func (m *HorizontalSection) SetLayout(value *HorizontalSectionLayoutType)() {
    m.layout = value
}
