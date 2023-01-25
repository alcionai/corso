package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81SCEPCertificateProfileCollectionResponse 
type Windows81SCEPCertificateProfileCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows81SCEPCertificateProfileable
}
// NewWindows81SCEPCertificateProfileCollectionResponse instantiates a new Windows81SCEPCertificateProfileCollectionResponse and sets the default values.
func NewWindows81SCEPCertificateProfileCollectionResponse()(*Windows81SCEPCertificateProfileCollectionResponse) {
    m := &Windows81SCEPCertificateProfileCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows81SCEPCertificateProfileCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81SCEPCertificateProfileCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81SCEPCertificateProfileCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81SCEPCertificateProfileCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows81SCEPCertificateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows81SCEPCertificateProfileable, len(val))
            for i, v := range val {
                res[i] = v.(Windows81SCEPCertificateProfileable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows81SCEPCertificateProfileCollectionResponse) GetValue()([]Windows81SCEPCertificateProfileable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows81SCEPCertificateProfileCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows81SCEPCertificateProfileCollectionResponse) SetValue(value []Windows81SCEPCertificateProfileable)() {
    m.value = value
}
