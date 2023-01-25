package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkPosition 
type WorkPosition struct {
    ItemFacet
    // Categories that the user has associated with this position.
    categories []string
    // Colleagues that are associated with this position.
    colleagues []RelatedPersonable
    // The detail property
    detail PositionDetailable
    // Denotes whether or not the position is current.
    isCurrent *bool
    // Contains detail of the user's manager in this position.
    manager RelatedPersonable
}
// NewWorkPosition instantiates a new WorkPosition and sets the default values.
func NewWorkPosition()(*WorkPosition) {
    m := &WorkPosition{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.workPosition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkPositionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkPositionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkPosition(), nil
}
// GetCategories gets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) GetCategories()([]string) {
    return m.categories
}
// GetColleagues gets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) GetColleagues()([]RelatedPersonable) {
    return m.colleagues
}
// GetDetail gets the detail property value. The detail property
func (m *WorkPosition) GetDetail()(PositionDetailable) {
    return m.detail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkPosition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["colleagues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPersonable, len(val))
            for i, v := range val {
                res[i] = v.(RelatedPersonable)
            }
            m.SetColleagues(res)
        }
        return nil
    }
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePositionDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PositionDetailable))
        }
        return nil
    }
    res["isCurrent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCurrent(val)
        }
        return nil
    }
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(RelatedPersonable))
        }
        return nil
    }
    return res
}
// GetIsCurrent gets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) GetIsCurrent()(*bool) {
    return m.isCurrent
}
// GetManager gets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) GetManager()(RelatedPersonable) {
    return m.manager
}
// Serialize serializes information the current object
func (m *WorkPosition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetColleagues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColleagues()))
        for i, v := range m.GetColleagues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("colleagues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCurrent", m.GetIsCurrent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategories sets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) SetCategories(value []string)() {
    m.categories = value
}
// SetColleagues sets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) SetColleagues(value []RelatedPersonable)() {
    m.colleagues = value
}
// SetDetail sets the detail property value. The detail property
func (m *WorkPosition) SetDetail(value PositionDetailable)() {
    m.detail = value
}
// SetIsCurrent sets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) SetIsCurrent(value *bool)() {
    m.isCurrent = value
}
// SetManager sets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) SetManager(value RelatedPersonable)() {
    m.manager = value
}
