package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingApplicability 
type DeviceManagementConfigurationSettingApplicability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // description of the setting
    description *string
    // Describes applicability for the mode the device is in
    deviceMode *DeviceManagementConfigurationDeviceMode
    // The OdataType property
    odataType *string
    // Supported platform types.
    platform *DeviceManagementConfigurationPlatforms
    // Describes which technology this setting can be deployed with
    technologies *DeviceManagementConfigurationTechnologies
}
// NewDeviceManagementConfigurationSettingApplicability instantiates a new deviceManagementConfigurationSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationSettingApplicability()(*DeviceManagementConfigurationSettingApplicability) {
    m := &DeviceManagementConfigurationSettingApplicability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingApplicabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingApplicabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementConfigurationExchangeOnlineSettingApplicability":
                        return NewDeviceManagementConfigurationExchangeOnlineSettingApplicability(), nil
                    case "#microsoft.graph.deviceManagementConfigurationWindowsSettingApplicability":
                        return NewDeviceManagementConfigurationWindowsSettingApplicability(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingApplicability(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingApplicability) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. description of the setting
func (m *DeviceManagementConfigurationSettingApplicability) GetDescription()(*string) {
    return m.description
}
// GetDeviceMode gets the deviceMode property value. Describes applicability for the mode the device is in
func (m *DeviceManagementConfigurationSettingApplicability) GetDeviceMode()(*DeviceManagementConfigurationDeviceMode) {
    return m.deviceMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingApplicability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationDeviceMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMode(val.(*DeviceManagementConfigurationDeviceMode))
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DeviceManagementConfigurationPlatforms))
        }
        return nil
    }
    res["technologies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnologies(val.(*DeviceManagementConfigurationTechnologies))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingApplicability) GetOdataType()(*string) {
    return m.odataType
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *DeviceManagementConfigurationSettingApplicability) GetPlatform()(*DeviceManagementConfigurationPlatforms) {
    return m.platform
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationSettingApplicability) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    return m.technologies
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingApplicability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceMode() != nil {
        cast := (*m.GetDeviceMode()).String()
        err := writer.WriteStringValue("deviceMode", &cast)
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnologies() != nil {
        cast := (*m.GetTechnologies()).String()
        err := writer.WriteStringValue("technologies", &cast)
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
func (m *DeviceManagementConfigurationSettingApplicability) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. description of the setting
func (m *DeviceManagementConfigurationSettingApplicability) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceMode sets the deviceMode property value. Describes applicability for the mode the device is in
func (m *DeviceManagementConfigurationSettingApplicability) SetDeviceMode(value *DeviceManagementConfigurationDeviceMode)() {
    m.deviceMode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingApplicability) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *DeviceManagementConfigurationSettingApplicability) SetPlatform(value *DeviceManagementConfigurationPlatforms)() {
    m.platform = value
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationSettingApplicability) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
