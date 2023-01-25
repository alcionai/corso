package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HybridAgentUpdaterConfiguration 
type HybridAgentUpdaterConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the agent is available.
    allowUpdateConfigurationOverride *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    deferUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The updateWindow property
    updateWindow UpdateWindowable
}
// NewHybridAgentUpdaterConfiguration instantiates a new hybridAgentUpdaterConfiguration and sets the default values.
func NewHybridAgentUpdaterConfiguration()(*HybridAgentUpdaterConfiguration) {
    m := &HybridAgentUpdaterConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateHybridAgentUpdaterConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHybridAgentUpdaterConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHybridAgentUpdaterConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HybridAgentUpdaterConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowUpdateConfigurationOverride gets the allowUpdateConfigurationOverride property value. Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the agent is available.
func (m *HybridAgentUpdaterConfiguration) GetAllowUpdateConfigurationOverride()(*bool) {
    return m.allowUpdateConfigurationOverride
}
// GetDeferUpdateDateTime gets the deferUpdateDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *HybridAgentUpdaterConfiguration) GetDeferUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deferUpdateDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HybridAgentUpdaterConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowUpdateConfigurationOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUpdateConfigurationOverride(val)
        }
        return nil
    }
    res["deferUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeferUpdateDateTime(val)
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
    res["updateWindow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdateWindowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindow(val.(UpdateWindowable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *HybridAgentUpdaterConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetUpdateWindow gets the updateWindow property value. The updateWindow property
func (m *HybridAgentUpdaterConfiguration) GetUpdateWindow()(UpdateWindowable) {
    return m.updateWindow
}
// Serialize serializes information the current object
func (m *HybridAgentUpdaterConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowUpdateConfigurationOverride", m.GetAllowUpdateConfigurationOverride())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("deferUpdateDateTime", m.GetDeferUpdateDateTime())
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
        err := writer.WriteObjectValue("updateWindow", m.GetUpdateWindow())
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
func (m *HybridAgentUpdaterConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowUpdateConfigurationOverride sets the allowUpdateConfigurationOverride property value. Indicates if updater configuration will be skipped and the agent will receive an update when the next version of the agent is available.
func (m *HybridAgentUpdaterConfiguration) SetAllowUpdateConfigurationOverride(value *bool)() {
    m.allowUpdateConfigurationOverride = value
}
// SetDeferUpdateDateTime sets the deferUpdateDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *HybridAgentUpdaterConfiguration) SetDeferUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deferUpdateDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HybridAgentUpdaterConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUpdateWindow sets the updateWindow property value. The updateWindow property
func (m *HybridAgentUpdaterConfiguration) SetUpdateWindow(value UpdateWindowable)() {
    m.updateWindow = value
}
