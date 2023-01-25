package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MatchingLabel 
type MatchingLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The applicationMode property
    applicationMode *ApplicationMode
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The id property
    id *string
    // The isEndpointProtectionEnabled property
    isEndpointProtectionEnabled *bool
    // The labelActions property
    labelActions []LabelActionBaseable
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The policyTip property
    policyTip *string
    // The priority property
    priority *int32
    // The toolTip property
    toolTip *string
}
// NewMatchingLabel instantiates a new matchingLabel and sets the default values.
func NewMatchingLabel()(*MatchingLabel) {
    m := &MatchingLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMatchingLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMatchingLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMatchingLabel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingLabel) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationMode gets the applicationMode property value. The applicationMode property
func (m *MatchingLabel) GetApplicationMode()(*ApplicationMode) {
    return m.applicationMode
}
// GetDescription gets the description property value. The description property
func (m *MatchingLabel) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *MatchingLabel) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MatchingLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationMode(val.(*ApplicationMode))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
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
    res["isEndpointProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEndpointProtectionEnabled(val)
        }
        return nil
    }
    res["labelActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelActionBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelActionBaseable, len(val))
            for i, v := range val {
                res[i] = v.(LabelActionBaseable)
            }
            m.SetLabelActions(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["policyTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTip(val)
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
    res["toolTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToolTip(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *MatchingLabel) GetId()(*string) {
    return m.id
}
// GetIsEndpointProtectionEnabled gets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *MatchingLabel) GetIsEndpointProtectionEnabled()(*bool) {
    return m.isEndpointProtectionEnabled
}
// GetLabelActions gets the labelActions property value. The labelActions property
func (m *MatchingLabel) GetLabelActions()([]LabelActionBaseable) {
    return m.labelActions
}
// GetName gets the name property value. The name property
func (m *MatchingLabel) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MatchingLabel) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyTip gets the policyTip property value. The policyTip property
func (m *MatchingLabel) GetPolicyTip()(*string) {
    return m.policyTip
}
// GetPriority gets the priority property value. The priority property
func (m *MatchingLabel) GetPriority()(*int32) {
    return m.priority
}
// GetToolTip gets the toolTip property value. The toolTip property
func (m *MatchingLabel) GetToolTip()(*string) {
    return m.toolTip
}
// Serialize serializes information the current object
func (m *MatchingLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApplicationMode() != nil {
        cast := (*m.GetApplicationMode()).String()
        err := writer.WriteStringValue("applicationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEndpointProtectionEnabled", m.GetIsEndpointProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetLabelActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabelActions()))
        for i, v := range m.GetLabelActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("labelActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("policyTip", m.GetPolicyTip())
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
    {
        err := writer.WriteStringValue("toolTip", m.GetToolTip())
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
func (m *MatchingLabel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicationMode sets the applicationMode property value. The applicationMode property
func (m *MatchingLabel) SetApplicationMode(value *ApplicationMode)() {
    m.applicationMode = value
}
// SetDescription sets the description property value. The description property
func (m *MatchingLabel) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *MatchingLabel) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The id property
func (m *MatchingLabel) SetId(value *string)() {
    m.id = value
}
// SetIsEndpointProtectionEnabled sets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *MatchingLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    m.isEndpointProtectionEnabled = value
}
// SetLabelActions sets the labelActions property value. The labelActions property
func (m *MatchingLabel) SetLabelActions(value []LabelActionBaseable)() {
    m.labelActions = value
}
// SetName sets the name property value. The name property
func (m *MatchingLabel) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MatchingLabel) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyTip sets the policyTip property value. The policyTip property
func (m *MatchingLabel) SetPolicyTip(value *string)() {
    m.policyTip = value
}
// SetPriority sets the priority property value. The priority property
func (m *MatchingLabel) SetPriority(value *int32)() {
    m.priority = value
}
// SetToolTip sets the toolTip property value. The toolTip property
func (m *MatchingLabel) SetToolTip(value *string)() {
    m.toolTip = value
}
