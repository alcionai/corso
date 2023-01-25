package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEnrollmentCompanyCode a class to hold specialty enrollment data used for enrolling via Google's Android Management API, such as Token, Url, and QR code content
type AndroidEnrollmentCompanyCode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Enrollment Token used by the User to enroll their device.
    enrollmentToken *string
    // The OdataType property
    odataType *string
    // String used to generate a QR code for the token.
    qrCodeContent *string
    // Generated QR code for the token.
    qrCodeImage MimeContentable
}
// NewAndroidEnrollmentCompanyCode instantiates a new androidEnrollmentCompanyCode and sets the default values.
func NewAndroidEnrollmentCompanyCode()(*AndroidEnrollmentCompanyCode) {
    m := &AndroidEnrollmentCompanyCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidEnrollmentCompanyCodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidEnrollmentCompanyCodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidEnrollmentCompanyCode(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidEnrollmentCompanyCode) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnrollmentToken gets the enrollmentToken property value. Enrollment Token used by the User to enroll their device.
func (m *AndroidEnrollmentCompanyCode) GetEnrollmentToken()(*string) {
    return m.enrollmentToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidEnrollmentCompanyCode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentToken(val)
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
    res["qrCodeContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeContent(val)
        }
        return nil
    }
    res["qrCodeImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeImage(val.(MimeContentable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidEnrollmentCompanyCode) GetOdataType()(*string) {
    return m.odataType
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidEnrollmentCompanyCode) GetQrCodeContent()(*string) {
    return m.qrCodeContent
}
// GetQrCodeImage gets the qrCodeImage property value. Generated QR code for the token.
func (m *AndroidEnrollmentCompanyCode) GetQrCodeImage()(MimeContentable) {
    return m.qrCodeImage
}
// Serialize serializes information the current object
func (m *AndroidEnrollmentCompanyCode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("enrollmentToken", m.GetEnrollmentToken())
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
        err := writer.WriteStringValue("qrCodeContent", m.GetQrCodeContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("qrCodeImage", m.GetQrCodeImage())
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
func (m *AndroidEnrollmentCompanyCode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnrollmentToken sets the enrollmentToken property value. Enrollment Token used by the User to enroll their device.
func (m *AndroidEnrollmentCompanyCode) SetEnrollmentToken(value *string)() {
    m.enrollmentToken = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidEnrollmentCompanyCode) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidEnrollmentCompanyCode) SetQrCodeContent(value *string)() {
    m.qrCodeContent = value
}
// SetQrCodeImage sets the qrCodeImage property value. Generated QR code for the token.
func (m *AndroidEnrollmentCompanyCode) SetQrCodeImage(value MimeContentable)() {
    m.qrCodeImage = value
}
