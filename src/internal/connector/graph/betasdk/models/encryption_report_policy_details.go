package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptionReportPolicyDetails policy Details for Encryption Report
type EncryptionReportPolicyDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Policy Id for Encryption Report
    policyId *string
    // Policy Name for Encryption Report
    policyName *string
}
// NewEncryptionReportPolicyDetails instantiates a new encryptionReportPolicyDetails and sets the default values.
func NewEncryptionReportPolicyDetails()(*EncryptionReportPolicyDetails) {
    m := &EncryptionReportPolicyDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEncryptionReportPolicyDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEncryptionReportPolicyDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptionReportPolicyDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptionReportPolicyDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptionReportPolicyDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EncryptionReportPolicyDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyId gets the policyId property value. Policy Id for Encryption Report
func (m *EncryptionReportPolicyDetails) GetPolicyId()(*string) {
    return m.policyId
}
// GetPolicyName gets the policyName property value. Policy Name for Encryption Report
func (m *EncryptionReportPolicyDetails) GetPolicyName()(*string) {
    return m.policyName
}
// Serialize serializes information the current object
func (m *EncryptionReportPolicyDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
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
func (m *EncryptionReportPolicyDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EncryptionReportPolicyDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyId sets the policyId property value. Policy Id for Encryption Report
func (m *EncryptionReportPolicyDetails) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetPolicyName sets the policyName property value. Policy Name for Encryption Report
func (m *EncryptionReportPolicyDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
