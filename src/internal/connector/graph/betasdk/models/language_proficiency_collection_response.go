package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LanguageProficiencyCollectionResponse 
type LanguageProficiencyCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []LanguageProficiencyable
}
// NewLanguageProficiencyCollectionResponse instantiates a new LanguageProficiencyCollectionResponse and sets the default values.
func NewLanguageProficiencyCollectionResponse()(*LanguageProficiencyCollectionResponse) {
    m := &LanguageProficiencyCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateLanguageProficiencyCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLanguageProficiencyCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLanguageProficiencyCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LanguageProficiencyCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLanguageProficiencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LanguageProficiencyable, len(val))
            for i, v := range val {
                res[i] = v.(LanguageProficiencyable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *LanguageProficiencyCollectionResponse) GetValue()([]LanguageProficiencyable) {
    return m.value
}
// Serialize serializes information the current object
func (m *LanguageProficiencyCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *LanguageProficiencyCollectionResponse) SetValue(value []LanguageProficiencyable)() {
    m.value = value
}
