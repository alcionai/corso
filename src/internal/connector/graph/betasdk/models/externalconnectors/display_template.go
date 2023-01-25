package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// DisplayTemplate 
type DisplayTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The text identifier for the display template; for example, contosoTickets. Maximum 16 characters. Only alphanumeric characters allowed.
    id *string
    // The layout property
    layout ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Jsonable
    // The OdataType property
    odataType *string
    // Defines the priority of a display template. A display template with priority 1 is evaluated before a template with priority 4. Gaps in priority values are supported. Must be positive value.
    priority *int32
    // Specifies additional rules for selecting this display template based on the item schema. Optional.
    rules []PropertyRuleable
}
// NewDisplayTemplate instantiates a new displayTemplate and sets the default values.
func NewDisplayTemplate()(*DisplayTemplate) {
    m := &DisplayTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDisplayTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDisplayTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDisplayTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DisplayTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DisplayTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["layout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLayout(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Jsonable))
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
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePropertyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PropertyRuleable, len(val))
            for i, v := range val {
                res[i] = v.(PropertyRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The text identifier for the display template; for example, contosoTickets. Maximum 16 characters. Only alphanumeric characters allowed.
func (m *DisplayTemplate) GetId()(*string) {
    return m.id
}
// GetLayout gets the layout property value. The layout property
func (m *DisplayTemplate) GetLayout()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Jsonable) {
    return m.layout
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DisplayTemplate) GetOdataType()(*string) {
    return m.odataType
}
// GetPriority gets the priority property value. Defines the priority of a display template. A display template with priority 1 is evaluated before a template with priority 4. Gaps in priority values are supported. Must be positive value.
func (m *DisplayTemplate) GetPriority()(*int32) {
    return m.priority
}
// GetRules gets the rules property value. Specifies additional rules for selecting this display template based on the item schema. Optional.
func (m *DisplayTemplate) GetRules()([]PropertyRuleable) {
    return m.rules
}
// Serialize serializes information the current object
func (m *DisplayTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("layout", m.GetLayout())
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
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
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
func (m *DisplayTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetId sets the id property value. The text identifier for the display template; for example, contosoTickets. Maximum 16 characters. Only alphanumeric characters allowed.
func (m *DisplayTemplate) SetId(value *string)() {
    m.id = value
}
// SetLayout sets the layout property value. The layout property
func (m *DisplayTemplate) SetLayout(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Jsonable)() {
    m.layout = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DisplayTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPriority sets the priority property value. Defines the priority of a display template. A display template with priority 1 is evaluated before a template with priority 4. Gaps in priority values are supported. Must be positive value.
func (m *DisplayTemplate) SetPriority(value *int32)() {
    m.priority = value
}
// SetRules sets the rules property value. Specifies additional rules for selecting this display template based on the item schema. Optional.
func (m *DisplayTemplate) SetRules(value []PropertyRuleable)() {
    m.rules = value
}
