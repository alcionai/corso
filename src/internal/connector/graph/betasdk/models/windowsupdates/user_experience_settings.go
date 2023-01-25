package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceSettings 
type UserExperienceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the number of days after an update is installed, during which the user of the device can control when the device restarts.
    daysUntilForcedReboot *int32
    // The OdataType property
    odataType *string
}
// NewUserExperienceSettings instantiates a new userExperienceSettings and sets the default values.
func NewUserExperienceSettings()(*UserExperienceSettings) {
    m := &UserExperienceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDaysUntilForcedReboot gets the daysUntilForcedReboot property value. Specifies the number of days after an update is installed, during which the user of the device can control when the device restarts.
func (m *UserExperienceSettings) GetDaysUntilForcedReboot()(*int32) {
    return m.daysUntilForcedReboot
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *UserExperienceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDaysUntilForcedReboot sets the daysUntilForcedReboot property value. Specifies the number of days after an update is installed, during which the user of the device can control when the device restarts.
func (m *UserExperienceSettings) SetDaysUntilForcedReboot(value *int32)() {
    m.daysUntilForcedReboot = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceSettings) SetOdataType(value *string)() {
    m.odataType = value
}
