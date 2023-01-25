package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkFeaturesConfiguration 
type TeamworkFeaturesConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Email address to send logs and feedback.
    emailToSendLogsAndFeedback *string
    // True if auto screen shared is enabled.
    isAutoScreenShareEnabled *bool
    // True if Bluetooth beaconing is enabled.
    isBluetoothBeaconingEnabled *bool
    // True if hiding meeting names is enabled.
    isHideMeetingNamesEnabled *bool
    // True if sending logs and feedback is enabled.
    isSendLogsAndFeedbackEnabled *bool
    // The OdataType property
    odataType *string
}
// NewTeamworkFeaturesConfiguration instantiates a new teamworkFeaturesConfiguration and sets the default values.
func NewTeamworkFeaturesConfiguration()(*TeamworkFeaturesConfiguration) {
    m := &TeamworkFeaturesConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkFeaturesConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkFeaturesConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkFeaturesConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkFeaturesConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEmailToSendLogsAndFeedback gets the emailToSendLogsAndFeedback property value. Email address to send logs and feedback.
func (m *TeamworkFeaturesConfiguration) GetEmailToSendLogsAndFeedback()(*string) {
    return m.emailToSendLogsAndFeedback
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkFeaturesConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailToSendLogsAndFeedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailToSendLogsAndFeedback(val)
        }
        return nil
    }
    res["isAutoScreenShareEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutoScreenShareEnabled(val)
        }
        return nil
    }
    res["isBluetoothBeaconingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBluetoothBeaconingEnabled(val)
        }
        return nil
    }
    res["isHideMeetingNamesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHideMeetingNamesEnabled(val)
        }
        return nil
    }
    res["isSendLogsAndFeedbackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSendLogsAndFeedbackEnabled(val)
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
// GetIsAutoScreenShareEnabled gets the isAutoScreenShareEnabled property value. True if auto screen shared is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsAutoScreenShareEnabled()(*bool) {
    return m.isAutoScreenShareEnabled
}
// GetIsBluetoothBeaconingEnabled gets the isBluetoothBeaconingEnabled property value. True if Bluetooth beaconing is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsBluetoothBeaconingEnabled()(*bool) {
    return m.isBluetoothBeaconingEnabled
}
// GetIsHideMeetingNamesEnabled gets the isHideMeetingNamesEnabled property value. True if hiding meeting names is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsHideMeetingNamesEnabled()(*bool) {
    return m.isHideMeetingNamesEnabled
}
// GetIsSendLogsAndFeedbackEnabled gets the isSendLogsAndFeedbackEnabled property value. True if sending logs and feedback is enabled.
func (m *TeamworkFeaturesConfiguration) GetIsSendLogsAndFeedbackEnabled()(*bool) {
    return m.isSendLogsAndFeedbackEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkFeaturesConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkFeaturesConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emailToSendLogsAndFeedback", m.GetEmailToSendLogsAndFeedback())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutoScreenShareEnabled", m.GetIsAutoScreenShareEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBluetoothBeaconingEnabled", m.GetIsBluetoothBeaconingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHideMeetingNamesEnabled", m.GetIsHideMeetingNamesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSendLogsAndFeedbackEnabled", m.GetIsSendLogsAndFeedbackEnabled())
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
func (m *TeamworkFeaturesConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEmailToSendLogsAndFeedback sets the emailToSendLogsAndFeedback property value. Email address to send logs and feedback.
func (m *TeamworkFeaturesConfiguration) SetEmailToSendLogsAndFeedback(value *string)() {
    m.emailToSendLogsAndFeedback = value
}
// SetIsAutoScreenShareEnabled sets the isAutoScreenShareEnabled property value. True if auto screen shared is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsAutoScreenShareEnabled(value *bool)() {
    m.isAutoScreenShareEnabled = value
}
// SetIsBluetoothBeaconingEnabled sets the isBluetoothBeaconingEnabled property value. True if Bluetooth beaconing is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsBluetoothBeaconingEnabled(value *bool)() {
    m.isBluetoothBeaconingEnabled = value
}
// SetIsHideMeetingNamesEnabled sets the isHideMeetingNamesEnabled property value. True if hiding meeting names is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsHideMeetingNamesEnabled(value *bool)() {
    m.isHideMeetingNamesEnabled = value
}
// SetIsSendLogsAndFeedbackEnabled sets the isSendLogsAndFeedbackEnabled property value. True if sending logs and feedback is enabled.
func (m *TeamworkFeaturesConfiguration) SetIsSendLogsAndFeedbackEnabled(value *bool)() {
    m.isSendLogsAndFeedbackEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkFeaturesConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
