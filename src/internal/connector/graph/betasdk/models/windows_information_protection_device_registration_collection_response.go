package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionDeviceRegistrationCollectionResponse 
type WindowsInformationProtectionDeviceRegistrationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []WindowsInformationProtectionDeviceRegistrationable
}
// NewWindowsInformationProtectionDeviceRegistrationCollectionResponse instantiates a new WindowsInformationProtectionDeviceRegistrationCollectionResponse and sets the default values.
func NewWindowsInformationProtectionDeviceRegistrationCollectionResponse()(*WindowsInformationProtectionDeviceRegistrationCollectionResponse) {
    m := &WindowsInformationProtectionDeviceRegistrationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindowsInformationProtectionDeviceRegistrationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionDeviceRegistrationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionDeviceRegistrationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionDeviceRegistrationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionDeviceRegistrationable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionDeviceRegistrationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *WindowsInformationProtectionDeviceRegistrationCollectionResponse) GetValue()([]WindowsInformationProtectionDeviceRegistrationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionDeviceRegistrationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WindowsInformationProtectionDeviceRegistrationCollectionResponse) SetValue(value []WindowsInformationProtectionDeviceRegistrationable)() {
    m.value = value
}
