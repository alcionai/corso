package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpeditedWindowsQualityUpdateSettings a complex type to store the expedited quality update settings such as release date and days until forced reboot.
type ExpeditedWindowsQualityUpdateSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of days after installation that forced reboot will happen.
    daysUntilForcedReboot *int32
    // The OdataType property
    odataType *string
    // The release date to identify a quality update.
    qualityUpdateRelease *string
}
// NewExpeditedWindowsQualityUpdateSettings instantiates a new expeditedWindowsQualityUpdateSettings and sets the default values.
func NewExpeditedWindowsQualityUpdateSettings()(*ExpeditedWindowsQualityUpdateSettings) {
    m := &ExpeditedWindowsQualityUpdateSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExpeditedWindowsQualityUpdateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpeditedWindowsQualityUpdateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExpeditedWindowsQualityUpdateSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpeditedWindowsQualityUpdateSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDaysUntilForcedReboot gets the daysUntilForcedReboot property value. The number of days after installation that forced reboot will happen.
func (m *ExpeditedWindowsQualityUpdateSettings) GetDaysUntilForcedReboot()(*int32) {
    return m.daysUntilForcedReboot
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpeditedWindowsQualityUpdateSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["daysUntilForcedReboot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaysUntilForcedReboot(val)
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
    res["qualityUpdateRelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateRelease(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExpeditedWindowsQualityUpdateSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetQualityUpdateRelease gets the qualityUpdateRelease property value. The release date to identify a quality update.
func (m *ExpeditedWindowsQualityUpdateSettings) GetQualityUpdateRelease()(*string) {
    return m.qualityUpdateRelease
}
// Serialize serializes information the current object
func (m *ExpeditedWindowsQualityUpdateSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("daysUntilForcedReboot", m.GetDaysUntilForcedReboot())
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
        err := writer.WriteStringValue("qualityUpdateRelease", m.GetQualityUpdateRelease())
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
func (m *ExpeditedWindowsQualityUpdateSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDaysUntilForcedReboot sets the daysUntilForcedReboot property value. The number of days after installation that forced reboot will happen.
func (m *ExpeditedWindowsQualityUpdateSettings) SetDaysUntilForcedReboot(value *int32)() {
    m.daysUntilForcedReboot = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExpeditedWindowsQualityUpdateSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQualityUpdateRelease sets the qualityUpdateRelease property value. The release date to identify a quality update.
func (m *ExpeditedWindowsQualityUpdateSettings) SetQualityUpdateRelease(value *string)() {
    m.qualityUpdateRelease = value
}
