package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerSystemUpdateFreezePeriod represents one item in the list of freeze periods for Android Device Owner system updates
type AndroidDeviceOwnerSystemUpdateFreezePeriod struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The day of the end date of the freeze period. Valid values 1 to 31
    endDay *int32
    // The month of the end date of the freeze period. Valid values 1 to 12
    endMonth *int32
    // The OdataType property
    odataType *string
    // The day of the start date of the freeze period. Valid values 1 to 31
    startDay *int32
    // The month of the start date of the freeze period. Valid values 1 to 12
    startMonth *int32
}
// NewAndroidDeviceOwnerSystemUpdateFreezePeriod instantiates a new androidDeviceOwnerSystemUpdateFreezePeriod and sets the default values.
func NewAndroidDeviceOwnerSystemUpdateFreezePeriod()(*AndroidDeviceOwnerSystemUpdateFreezePeriod) {
    m := &AndroidDeviceOwnerSystemUpdateFreezePeriod{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidDeviceOwnerSystemUpdateFreezePeriodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerSystemUpdateFreezePeriodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerSystemUpdateFreezePeriod(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEndDay gets the endDay property value. The day of the end date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetEndDay()(*int32) {
    return m.endDay
}
// GetEndMonth gets the endMonth property value. The month of the end date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetEndMonth()(*int32) {
    return m.endMonth
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDay(val)
        }
        return nil
    }
    res["endMonth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonth(val)
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
    res["startDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDay(val)
        }
        return nil
    }
    res["startMonth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonth(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDay gets the startDay property value. The day of the start date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetStartDay()(*int32) {
    return m.startDay
}
// GetStartMonth gets the startMonth property value. The month of the start date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) GetStartMonth()(*int32) {
    return m.startMonth
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("endDay", m.GetEndDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("endMonth", m.GetEndMonth())
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
        err := writer.WriteInt32Value("startDay", m.GetStartDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("startMonth", m.GetStartMonth())
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
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndDay sets the endDay property value. The day of the end date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetEndDay(value *int32)() {
    m.endDay = value
}
// SetEndMonth sets the endMonth property value. The month of the end date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetEndMonth(value *int32)() {
    m.endMonth = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDay sets the startDay property value. The day of the start date of the freeze period. Valid values 1 to 31
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetStartDay(value *int32)() {
    m.startDay = value
}
// SetStartMonth sets the startMonth property value. The month of the start date of the freeze period. Valid values 1 to 12
func (m *AndroidDeviceOwnerSystemUpdateFreezePeriod) SetStartMonth(value *int32)() {
    m.startMonth = value
}
