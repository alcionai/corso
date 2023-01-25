package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignedUserLicense 
type IosVppAppAssignedUserLicense struct {
    IosVppAppAssignedLicense
}
// NewIosVppAppAssignedUserLicense instantiates a new IosVppAppAssignedUserLicense and sets the default values.
func NewIosVppAppAssignedUserLicense()(*IosVppAppAssignedUserLicense) {
    m := &IosVppAppAssignedUserLicense{
        IosVppAppAssignedLicense: *NewIosVppAppAssignedLicense(),
    }
    return m
}
// CreateIosVppAppAssignedUserLicenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignedUserLicenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppAssignedUserLicense(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppAssignedUserLicense) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosVppAppAssignedLicense.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IosVppAppAssignedUserLicense) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosVppAppAssignedLicense.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
