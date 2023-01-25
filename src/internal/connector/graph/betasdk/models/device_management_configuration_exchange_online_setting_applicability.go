package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationExchangeOnlineSettingApplicability 
type DeviceManagementConfigurationExchangeOnlineSettingApplicability struct {
    DeviceManagementConfigurationSettingApplicability
}
// NewDeviceManagementConfigurationExchangeOnlineSettingApplicability instantiates a new DeviceManagementConfigurationExchangeOnlineSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationExchangeOnlineSettingApplicability()(*DeviceManagementConfigurationExchangeOnlineSettingApplicability) {
    m := &DeviceManagementConfigurationExchangeOnlineSettingApplicability{
        DeviceManagementConfigurationSettingApplicability: *NewDeviceManagementConfigurationSettingApplicability(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationExchangeOnlineSettingApplicability";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationExchangeOnlineSettingApplicabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationExchangeOnlineSettingApplicabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationExchangeOnlineSettingApplicability(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationExchangeOnlineSettingApplicability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingApplicability.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationExchangeOnlineSettingApplicability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingApplicability.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
