package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmbeddedSIMActivationCode the embedded SIM activation code as provided by the mobile operator.
type EmbeddedSIMActivationCode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
    integratedCircuitCardIdentifier *string
    // The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
    matchingIdentifier *string
    // The OdataType property
    odataType *string
    // The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
    smdpPlusServerAddress *string
}
// NewEmbeddedSIMActivationCode instantiates a new embeddedSIMActivationCode and sets the default values.
func NewEmbeddedSIMActivationCode()(*EmbeddedSIMActivationCode) {
    m := &EmbeddedSIMActivationCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEmbeddedSIMActivationCodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMActivationCodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCode(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmbeddedSIMActivationCode) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["integratedCircuitCardIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegratedCircuitCardIdentifier(val)
        }
        return nil
    }
    res["matchingIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingIdentifier(val)
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
    res["smdpPlusServerAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmdpPlusServerAddress(val)
        }
        return nil
    }
    return res
}
// GetIntegratedCircuitCardIdentifier gets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) GetIntegratedCircuitCardIdentifier()(*string) {
    return m.integratedCircuitCardIdentifier
}
// GetMatchingIdentifier gets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) GetMatchingIdentifier()(*string) {
    return m.matchingIdentifier
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EmbeddedSIMActivationCode) GetOdataType()(*string) {
    return m.odataType
}
// GetSmdpPlusServerAddress gets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) GetSmdpPlusServerAddress()(*string) {
    return m.smdpPlusServerAddress
}
// Serialize serializes information the current object
func (m *EmbeddedSIMActivationCode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("integratedCircuitCardIdentifier", m.GetIntegratedCircuitCardIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("matchingIdentifier", m.GetMatchingIdentifier())
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
        err := writer.WriteStringValue("smdpPlusServerAddress", m.GetSmdpPlusServerAddress())
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
func (m *EmbeddedSIMActivationCode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIntegratedCircuitCardIdentifier sets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) SetIntegratedCircuitCardIdentifier(value *string)() {
    m.integratedCircuitCardIdentifier = value
}
// SetMatchingIdentifier sets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) SetMatchingIdentifier(value *string)() {
    m.matchingIdentifier = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EmbeddedSIMActivationCode) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSmdpPlusServerAddress sets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) SetSmdpPlusServerAddress(value *string)() {
    m.smdpPlusServerAddress = value
}
