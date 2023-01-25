package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Filter 
type Filter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
    categoryFilterGroups []FilterGroupable
    // Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is not satisfied any longer, such object will get de-provisioned'. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
    groups []FilterGroupable
    // *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If an object doesn't satisfy this filter it will not be processed further. Important to understand is that if an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is no longer satisfied, such object will NOT get de-provisioned. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
    inputFilterGroups []FilterGroupable
    // The OdataType property
    odataType *string
}
// NewFilter instantiates a new filter and sets the default values.
func NewFilter()(*Filter) {
    m := &Filter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Filter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCategoryFilterGroups gets the categoryFilterGroups property value. *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) GetCategoryFilterGroups()([]FilterGroupable) {
    return m.categoryFilterGroups
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Filter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["categoryFilterGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterGroupable, len(val))
            for i, v := range val {
                res[i] = v.(FilterGroupable)
            }
            m.SetCategoryFilterGroups(res)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterGroupable, len(val))
            for i, v := range val {
                res[i] = v.(FilterGroupable)
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["inputFilterGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilterGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilterGroupable, len(val))
            for i, v := range val {
                res[i] = v.(FilterGroupable)
            }
            m.SetInputFilterGroups(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is not satisfied any longer, such object will get de-provisioned'. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) GetGroups()([]FilterGroupable) {
    return m.groups
}
// GetInputFilterGroups gets the inputFilterGroups property value. *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If an object doesn't satisfy this filter it will not be processed further. Important to understand is that if an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is no longer satisfied, such object will NOT get de-provisioned. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) GetInputFilterGroups()([]FilterGroupable) {
    return m.inputFilterGroups
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Filter) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Filter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCategoryFilterGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategoryFilterGroups()))
        for i, v := range m.GetCategoryFilterGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("categoryFilterGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInputFilterGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInputFilterGroups()))
        for i, v := range m.GetInputFilterGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("inputFilterGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Filter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCategoryFilterGroups sets the categoryFilterGroups property value. *Experimental* Filter group set used to decide whether given object belongs and should be processed as part of this object mapping. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) SetCategoryFilterGroups(value []FilterGroupable)() {
    m.categoryFilterGroups = value
}
// SetGroups sets the groups property value. Filter group set used to decide whether given object is in scope for provisioning. This is the filter which should be used in most cases. If an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is not satisfied any longer, such object will get de-provisioned'. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) SetGroups(value []FilterGroupable)() {
    m.groups = value
}
// SetInputFilterGroups sets the inputFilterGroups property value. *Experimental* Filter group set used to filter out objects at the early stage of reading them from the directory. If an object doesn't satisfy this filter it will not be processed further. Important to understand is that if an object used to satisfy this filter at a given moment, and then the object or the filter was changed so that filter is no longer satisfied, such object will NOT get de-provisioned. An object is considered in scope if ANY of the groups in the collection is evaluated to true.
func (m *Filter) SetInputFilterGroups(value []FilterGroupable)() {
    m.inputFilterGroups = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Filter) SetOdataType(value *string)() {
    m.odataType = value
}
