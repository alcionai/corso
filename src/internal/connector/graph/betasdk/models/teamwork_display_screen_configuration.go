package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDisplayScreenConfiguration 
type TeamworkDisplayScreenConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
    backlightBrightness *int32
    // Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
    backlightTimeout *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
    isHighContrastEnabled *bool
    // True if screensaver is enabled. Not applicable for Teams Rooms devices.
    isScreensaverEnabled *bool
    // The OdataType property
    odataType *string
    // Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
    screensaverTimeout *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewTeamworkDisplayScreenConfiguration instantiates a new teamworkDisplayScreenConfiguration and sets the default values.
func NewTeamworkDisplayScreenConfiguration()(*TeamworkDisplayScreenConfiguration) {
    m := &TeamworkDisplayScreenConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkDisplayScreenConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDisplayScreenConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDisplayScreenConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayScreenConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBacklightBrightness gets the backlightBrightness property value. The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetBacklightBrightness()(*int32) {
    return m.backlightBrightness
}
// GetBacklightTimeout gets the backlightTimeout property value. Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetBacklightTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.backlightTimeout
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDisplayScreenConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["backlightBrightness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBacklightBrightness(val)
        }
        return nil
    }
    res["backlightTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBacklightTimeout(val)
        }
        return nil
    }
    res["isHighContrastEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHighContrastEnabled(val)
        }
        return nil
    }
    res["isScreensaverEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScreensaverEnabled(val)
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
    res["screensaverTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreensaverTimeout(val)
        }
        return nil
    }
    return res
}
// GetIsHighContrastEnabled gets the isHighContrastEnabled property value. True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetIsHighContrastEnabled()(*bool) {
    return m.isHighContrastEnabled
}
// GetIsScreensaverEnabled gets the isScreensaverEnabled property value. True if screensaver is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetIsScreensaverEnabled()(*bool) {
    return m.isScreensaverEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayScreenConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetScreensaverTimeout gets the screensaverTimeout property value. Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) GetScreensaverTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.screensaverTimeout
}
// Serialize serializes information the current object
func (m *TeamworkDisplayScreenConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("backlightBrightness", m.GetBacklightBrightness())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("backlightTimeout", m.GetBacklightTimeout())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHighContrastEnabled", m.GetIsHighContrastEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScreensaverEnabled", m.GetIsScreensaverEnabled())
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
        err := writer.WriteISODurationValue("screensaverTimeout", m.GetScreensaverTimeout())
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
func (m *TeamworkDisplayScreenConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBacklightBrightness sets the backlightBrightness property value. The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetBacklightBrightness(value *int32)() {
    m.backlightBrightness = value
}
// SetBacklightTimeout sets the backlightTimeout property value. Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetBacklightTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.backlightTimeout = value
}
// SetIsHighContrastEnabled sets the isHighContrastEnabled property value. True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetIsHighContrastEnabled(value *bool)() {
    m.isHighContrastEnabled = value
}
// SetIsScreensaverEnabled sets the isScreensaverEnabled property value. True if screensaver is enabled. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetIsScreensaverEnabled(value *bool)() {
    m.isScreensaverEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayScreenConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScreensaverTimeout sets the screensaverTimeout property value. Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayScreenConfiguration) SetScreensaverTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.screensaverTimeout = value
}
