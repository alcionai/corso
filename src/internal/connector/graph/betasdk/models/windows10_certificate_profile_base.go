package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10CertificateProfileBase 
type Windows10CertificateProfileBase struct {
    WindowsCertificateProfileBase
}
// NewWindows10CertificateProfileBase instantiates a new Windows10CertificateProfileBase and sets the default values.
func NewWindows10CertificateProfileBase()(*Windows10CertificateProfileBase) {
    m := &Windows10CertificateProfileBase{
        WindowsCertificateProfileBase: *NewWindowsCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10CertificateProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10CertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10CertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windows10PkcsCertificateProfile":
                        return NewWindows10PkcsCertificateProfile(), nil
                }
            }
        }
    }
    return NewWindows10CertificateProfileBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10CertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsCertificateProfileBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *Windows10CertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsCertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
