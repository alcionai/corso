package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcRestorePointSetting 
type CloudPcRestorePointSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6, 12, 16, and 24. The default frequency is 12 hours.
    frequencyInHours *int32
    // The OdataType property
    odataType *string
    // If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users cannot use snapshots to restore the Cloud PC.
    userRestoreEnabled *bool
}
// NewCloudPcRestorePointSetting instantiates a new cloudPcRestorePointSetting and sets the default values.
func NewCloudPcRestorePointSetting()(*CloudPcRestorePointSetting) {
    m := &CloudPcRestorePointSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcRestorePointSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcRestorePointSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcRestorePointSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRestorePointSetting) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcRestorePointSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["frequencyInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrequencyInHours(val)
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
    res["userRestoreEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRestoreEnabled(val)
        }
        return nil
    }
    return res
}
// GetFrequencyInHours gets the frequencyInHours property value. The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6, 12, 16, and 24. The default frequency is 12 hours.
func (m *CloudPcRestorePointSetting) GetFrequencyInHours()(*int32) {
    return m.frequencyInHours
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcRestorePointSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetUserRestoreEnabled gets the userRestoreEnabled property value. If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users cannot use snapshots to restore the Cloud PC.
func (m *CloudPcRestorePointSetting) GetUserRestoreEnabled()(*bool) {
    return m.userRestoreEnabled
}
// Serialize serializes information the current object
func (m *CloudPcRestorePointSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("frequencyInHours", m.GetFrequencyInHours())
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
        err := writer.WriteBoolValue("userRestoreEnabled", m.GetUserRestoreEnabled())
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
func (m *CloudPcRestorePointSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFrequencyInHours sets the frequencyInHours property value. The time interval in hours to take snapshots (restore points) of a Cloud PC automatically. Possible values are 4, 6, 12, 16, and 24. The default frequency is 12 hours.
func (m *CloudPcRestorePointSetting) SetFrequencyInHours(value *int32)() {
    m.frequencyInHours = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcRestorePointSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserRestoreEnabled sets the userRestoreEnabled property value. If true, the user has the ability to use snapshots to restore Cloud PCs. If false, non-admin users cannot use snapshots to restore the Cloud PC.
func (m *CloudPcRestorePointSetting) SetUserRestoreEnabled(value *bool)() {
    m.userRestoreEnabled = value
}
