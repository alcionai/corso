package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerSilentCertificateAccess contain the package ID that has the pre-granted access to the certificate.
type AndroidDeviceOwnerSilentCertificateAccess struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Package ID that has the pre-granted access to the certificate.
    packageId *string
}
// NewAndroidDeviceOwnerSilentCertificateAccess instantiates a new androidDeviceOwnerSilentCertificateAccess and sets the default values.
func NewAndroidDeviceOwnerSilentCertificateAccess()(*AndroidDeviceOwnerSilentCertificateAccess) {
    m := &AndroidDeviceOwnerSilentCertificateAccess{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidDeviceOwnerSilentCertificateAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerSilentCertificateAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerSilentCertificateAccess(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerSilentCertificateAccess) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerSilentCertificateAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerSilentCertificateAccess) GetOdataType()(*string) {
    return m.odataType
}
// GetPackageId gets the packageId property value. Package ID that has the pre-granted access to the certificate.
func (m *AndroidDeviceOwnerSilentCertificateAccess) GetPackageId()(*string) {
    return m.packageId
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerSilentCertificateAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("packageId", m.GetPackageId())
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
func (m *AndroidDeviceOwnerSilentCertificateAccess) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerSilentCertificateAccess) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPackageId sets the packageId property value. Package ID that has the pre-granted access to the certificate.
func (m *AndroidDeviceOwnerSilentCertificateAccess) SetPackageId(value *string)() {
    m.packageId = value
}
