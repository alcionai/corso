package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppSupersedenceCollectionResponse 
type MobileAppSupersedenceCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []MobileAppSupersedenceable
}
// NewMobileAppSupersedenceCollectionResponse instantiates a new MobileAppSupersedenceCollectionResponse and sets the default values.
func NewMobileAppSupersedenceCollectionResponse()(*MobileAppSupersedenceCollectionResponse) {
    m := &MobileAppSupersedenceCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMobileAppSupersedenceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppSupersedenceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppSupersedenceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppSupersedenceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppSupersedenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppSupersedenceable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppSupersedenceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MobileAppSupersedenceCollectionResponse) GetValue()([]MobileAppSupersedenceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MobileAppSupersedenceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MobileAppSupersedenceCollectionResponse) SetValue(value []MobileAppSupersedenceable)() {
    m.value = value
}
