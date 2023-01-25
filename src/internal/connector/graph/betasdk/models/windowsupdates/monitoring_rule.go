package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MonitoringRule 
type MonitoringRule struct {
    // The action triggered when the threshold for the given signal is met. Possible values are: alertError, pauseDeployment, unknownFutureValue.
    action *MonitoringAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The signal to monitor. Possible values are: rollback, unknownFutureValue.
    signal *MonitoringSignal
    // The threshold for a signal at which to trigger action. An integer from 1 to 100 (inclusive).
    threshold *int32
}
// NewMonitoringRule instantiates a new monitoringRule and sets the default values.
func NewMonitoringRule()(*MonitoringRule) {
    m := &MonitoringRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMonitoringRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonitoringRule(), nil
}
// GetAction gets the action property value. The action triggered when the threshold for the given signal is met. Possible values are: alertError, pauseDeployment, unknownFutureValue.
func (m *MonitoringRule) GetAction()(*MonitoringAction) {
    return m.action
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MonitoringRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MonitoringRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMonitoringAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*MonitoringAction))
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
    res["signal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMonitoringSignal)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignal(val.(*MonitoringSignal))
        }
        return nil
    }
    res["threshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreshold(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MonitoringRule) GetOdataType()(*string) {
    return m.odataType
}
// GetSignal gets the signal property value. The signal to monitor. Possible values are: rollback, unknownFutureValue.
func (m *MonitoringRule) GetSignal()(*MonitoringSignal) {
    return m.signal
}
// GetThreshold gets the threshold property value. The threshold for a signal at which to trigger action. An integer from 1 to 100 (inclusive).
func (m *MonitoringRule) GetThreshold()(*int32) {
    return m.threshold
}
// Serialize serializes information the current object
func (m *MonitoringRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
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
    if m.GetSignal() != nil {
        cast := (*m.GetSignal()).String()
        err := writer.WriteStringValue("signal", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("threshold", m.GetThreshold())
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
// SetAction sets the action property value. The action triggered when the threshold for the given signal is met. Possible values are: alertError, pauseDeployment, unknownFutureValue.
func (m *MonitoringRule) SetAction(value *MonitoringAction)() {
    m.action = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MonitoringRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MonitoringRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSignal sets the signal property value. The signal to monitor. Possible values are: rollback, unknownFutureValue.
func (m *MonitoringRule) SetSignal(value *MonitoringSignal)() {
    m.signal = value
}
// SetThreshold sets the threshold property value. The threshold for a signal at which to trigger action. An integer from 1 to 100 (inclusive).
func (m *MonitoringRule) SetThreshold(value *int32)() {
    m.threshold = value
}
