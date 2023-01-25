package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RankedEmailAddress 
type RankedEmailAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The email address.
    address *string
    // The OdataType property
    odataType *string
    // The rank of the email address. A rank is used as a sort key, in relation to the other returned results. A higher rank value corresponds to a more relevant result. Relevance is determined by communication, collaboration, and business relationship signals.
    rank *float64
}
// NewRankedEmailAddress instantiates a new rankedEmailAddress and sets the default values.
func NewRankedEmailAddress()(*RankedEmailAddress) {
    m := &RankedEmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRankedEmailAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRankedEmailAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRankedEmailAddress(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RankedEmailAddress) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddress gets the address property value. The email address.
func (m *RankedEmailAddress) GetAddress()(*string) {
    return m.address
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RankedEmailAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RankedEmailAddress) GetOdataType()(*string) {
    return m.odataType
}
// GetRank gets the rank property value. The rank of the email address. A rank is used as a sort key, in relation to the other returned results. A higher rank value corresponds to a more relevant result. Relevance is determined by communication, collaboration, and business relationship signals.
func (m *RankedEmailAddress) GetRank()(*float64) {
    return m.rank
}
// Serialize serializes information the current object
func (m *RankedEmailAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
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
        err := writer.WriteFloat64Value("rank", m.GetRank())
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
func (m *RankedEmailAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddress sets the address property value. The email address.
func (m *RankedEmailAddress) SetAddress(value *string)() {
    m.address = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RankedEmailAddress) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRank sets the rank property value. The rank of the email address. A rank is used as a sort key, in relation to the other returned results. A higher rank value corresponds to a more relevant result. Relevance is determined by communication, collaboration, and business relationship signals.
func (m *RankedEmailAddress) SetRank(value *float64)() {
    m.rank = value
}
