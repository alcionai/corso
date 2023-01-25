package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadDetail 
type PayloadDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Payload coachmark details.
    coachmarks []PayloadCoachmarkable
    // Payload content details.
    content *string
    // The OdataType property
    odataType *string
    // The phishing URL used to target a user.
    phishingUrl *string
}
// NewPayloadDetail instantiates a new payloadDetail and sets the default values.
func NewPayloadDetail()(*PayloadDetail) {
    m := &PayloadDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePayloadDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.emailPayloadDetail":
                        return NewEmailPayloadDetail(), nil
                }
            }
        }
    }
    return NewPayloadDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCoachmarks gets the coachmarks property value. Payload coachmark details.
func (m *PayloadDetail) GetCoachmarks()([]PayloadCoachmarkable) {
    return m.coachmarks
}
// GetContent gets the content property value. Payload content details.
func (m *PayloadDetail) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["coachmarks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePayloadCoachmarkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PayloadCoachmarkable, len(val))
            for i, v := range val {
                res[i] = v.(PayloadCoachmarkable)
            }
            m.SetCoachmarks(res)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
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
    res["phishingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhishingUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PayloadDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetPhishingUrl gets the phishingUrl property value. The phishing URL used to target a user.
func (m *PayloadDetail) GetPhishingUrl()(*string) {
    return m.phishingUrl
}
// Serialize serializes information the current object
func (m *PayloadDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCoachmarks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCoachmarks()))
        for i, v := range m.GetCoachmarks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("coachmarks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content", m.GetContent())
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
        err := writer.WriteStringValue("phishingUrl", m.GetPhishingUrl())
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
func (m *PayloadDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCoachmarks sets the coachmarks property value. Payload coachmark details.
func (m *PayloadDetail) SetCoachmarks(value []PayloadCoachmarkable)() {
    m.coachmarks = value
}
// SetContent sets the content property value. Payload content details.
func (m *PayloadDetail) SetContent(value *string)() {
    m.content = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PayloadDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPhishingUrl sets the phishingUrl property value. The phishing URL used to target a user.
func (m *PayloadDetail) SetPhishingUrl(value *string)() {
    m.phishingUrl = value
}
