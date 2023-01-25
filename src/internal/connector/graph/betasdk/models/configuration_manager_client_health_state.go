package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerClientHealthState configuration manager client health state
type ConfigurationManagerClientHealthState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Error code for failed state.
    errorCode *int32
    // Datetime for last sync with configuration manager management point.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Configuration manager client state
    state *ConfigurationManagerClientState
}
// NewConfigurationManagerClientHealthState instantiates a new configurationManagerClientHealthState and sets the default values.
func NewConfigurationManagerClientHealthState()(*ConfigurationManagerClientHealthState) {
    m := &ConfigurationManagerClientHealthState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfigurationManagerClientHealthStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerClientHealthStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerClientHealthState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientHealthState) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetErrorCode gets the errorCode property value. Error code for failed state.
func (m *ConfigurationManagerClientHealthState) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerClientHealthState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerClientState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ConfigurationManagerClientState))
        }
        return nil
    }
    return res
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Datetime for last sync with configuration manager management point.
func (m *ConfigurationManagerClientHealthState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConfigurationManagerClientHealthState) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. Configuration manager client state
func (m *ConfigurationManagerClientHealthState) GetState()(*ConfigurationManagerClientState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ConfigurationManagerClientHealthState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ConfigurationManagerClientHealthState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetErrorCode sets the errorCode property value. Error code for failed state.
func (m *ConfigurationManagerClientHealthState) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Datetime for last sync with configuration manager management point.
func (m *ConfigurationManagerClientHealthState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConfigurationManagerClientHealthState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. Configuration manager client state
func (m *ConfigurationManagerClientHealthState) SetState(value *ConfigurationManagerClientState)() {
    m.state = value
}
