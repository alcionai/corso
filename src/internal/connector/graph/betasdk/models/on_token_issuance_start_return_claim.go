package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnTokenIssuanceStartReturnClaim 
type OnTokenIssuanceStartReturnClaim struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The claimIdInApiResponse property
    claimIdInApiResponse *string
    // The OdataType property
    odataType *string
}
// NewOnTokenIssuanceStartReturnClaim instantiates a new onTokenIssuanceStartReturnClaim and sets the default values.
func NewOnTokenIssuanceStartReturnClaim()(*OnTokenIssuanceStartReturnClaim) {
    m := &OnTokenIssuanceStartReturnClaim{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnTokenIssuanceStartReturnClaimFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnTokenIssuanceStartReturnClaimFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartReturnClaim(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnTokenIssuanceStartReturnClaim) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClaimIdInApiResponse gets the claimIdInApiResponse property value. The claimIdInApiResponse property
func (m *OnTokenIssuanceStartReturnClaim) GetClaimIdInApiResponse()(*string) {
    return m.claimIdInApiResponse
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnTokenIssuanceStartReturnClaim) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["claimIdInApiResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaimIdInApiResponse(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartReturnClaim) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OnTokenIssuanceStartReturnClaim) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("claimIdInApiResponse", m.GetClaimIdInApiResponse())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnTokenIssuanceStartReturnClaim) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClaimIdInApiResponse sets the claimIdInApiResponse property value. The claimIdInApiResponse property
func (m *OnTokenIssuanceStartReturnClaim) SetClaimIdInApiResponse(value *string)() {
    m.claimIdInApiResponse = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartReturnClaim) SetOdataType(value *string)() {
    m.odataType = value
}
