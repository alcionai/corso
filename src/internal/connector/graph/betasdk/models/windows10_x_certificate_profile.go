package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XCertificateProfile 
type Windows10XCertificateProfile struct {
    DeviceManagementResourceAccessProfileBase
}
// NewWindows10XCertificateProfile instantiates a new Windows10XCertificateProfile and sets the default values.
func NewWindows10XCertificateProfile()(*Windows10XCertificateProfile) {
    m := &Windows10XCertificateProfile{
        DeviceManagementResourceAccessProfileBase: *NewDeviceManagementResourceAccessProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10XCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10XCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windows10XSCEPCertificateProfile":
                        return NewWindows10XSCEPCertificateProfile(), nil
                }
            }
        }
    }
    return NewWindows10XCertificateProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementResourceAccessProfileBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Windows10XCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementResourceAccessProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
