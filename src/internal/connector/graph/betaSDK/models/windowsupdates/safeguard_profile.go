package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SafeguardProfile 
type SafeguardProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the category of safeguards. The possible values are: likelyIssues, unknownFutureValue.
    category *SafeguardCategory
    // The OdataType property
    odataType *string
}
// NewSafeguardProfile instantiates a new safeguardProfile and sets the default values.
func NewSafeguardProfile()(*SafeguardProfile) {
    m := &SafeguardProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSafeguardProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSafeguardProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSafeguardProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SafeguardProfile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCategory gets the category property value. Specifies the category of safeguards. The possible values are: likelyIssues, unknownFutureValue.
func (m *SafeguardProfile) GetCategory()(*SafeguardCategory) {
    return m.category
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SafeguardProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSafeguardCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*SafeguardCategory))
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
func (m *SafeguardProfile) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SafeguardProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
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
func (m *SafeguardProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCategory sets the category property value. Specifies the category of safeguards. The possible values are: likelyIssues, unknownFutureValue.
func (m *SafeguardProfile) SetCategory(value *SafeguardCategory)() {
    m.category = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SafeguardProfile) SetOdataType(value *string)() {
    m.odataType = value
}
