package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10ImportedPFXCertificateProfileCollectionResponse 
type Windows10ImportedPFXCertificateProfileCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows10ImportedPFXCertificateProfileable
}
// NewWindows10ImportedPFXCertificateProfileCollectionResponse instantiates a new Windows10ImportedPFXCertificateProfileCollectionResponse and sets the default values.
func NewWindows10ImportedPFXCertificateProfileCollectionResponse()(*Windows10ImportedPFXCertificateProfileCollectionResponse) {
    m := &Windows10ImportedPFXCertificateProfileCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows10ImportedPFXCertificateProfileCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10ImportedPFXCertificateProfileCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10ImportedPFXCertificateProfileCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10ImportedPFXCertificateProfileCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10ImportedPFXCertificateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10ImportedPFXCertificateProfileable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10ImportedPFXCertificateProfileable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows10ImportedPFXCertificateProfileCollectionResponse) GetValue()([]Windows10ImportedPFXCertificateProfileable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows10ImportedPFXCertificateProfileCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10ImportedPFXCertificateProfileCollectionResponse) SetValue(value []Windows10ImportedPFXCertificateProfileable)() {
    m.value = value
}
