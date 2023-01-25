package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileWiFiConfiguration 
type AndroidWorkProfileWiFiConfiguration struct {
    DeviceConfiguration
    // Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
    connectAutomatically *bool
    // When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
    connectWhenNetworkNameIsHidden *bool
    // Network Name
    networkName *string
    // This is the name of the Wi-Fi network that is broadcast to all devices.
    ssid *string
    // Wi-Fi Security Types for Android.
    wiFiSecurityType *AndroidWiFiSecurityType
}
// NewAndroidWorkProfileWiFiConfiguration instantiates a new AndroidWorkProfileWiFiConfiguration and sets the default values.
func NewAndroidWorkProfileWiFiConfiguration()(*AndroidWorkProfileWiFiConfiguration) {
    m := &AndroidWorkProfileWiFiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileWiFiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidWorkProfileWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidWorkProfileEnterpriseWiFiConfiguration":
                        return NewAndroidWorkProfileEnterpriseWiFiConfiguration(), nil
                }
            }
        }
    }
    return NewAndroidWorkProfileWiFiConfiguration(), nil
}
// GetConnectAutomatically gets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
func (m *AndroidWorkProfileWiFiConfiguration) GetConnectAutomatically()(*bool) {
    return m.connectAutomatically
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *AndroidWorkProfileWiFiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    return m.connectWhenNetworkNameIsHidden
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["connectAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectAutomatically(val)
        }
        return nil
    }
    res["connectWhenNetworkNameIsHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectWhenNetworkNameIsHidden(val)
        }
        return nil
    }
    res["networkName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkName(val)
        }
        return nil
    }
    res["ssid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSsid(val)
        }
        return nil
    }
    res["wiFiSecurityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidWiFiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiSecurityType(val.(*AndroidWiFiSecurityType))
        }
        return nil
    }
    return res
}
// GetNetworkName gets the networkName property value. Network Name
func (m *AndroidWorkProfileWiFiConfiguration) GetNetworkName()(*string) {
    return m.networkName
}
// GetSsid gets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *AndroidWorkProfileWiFiConfiguration) GetSsid()(*string) {
    return m.ssid
}
// GetWiFiSecurityType gets the wiFiSecurityType property value. Wi-Fi Security Types for Android.
func (m *AndroidWorkProfileWiFiConfiguration) GetWiFiSecurityType()(*AndroidWiFiSecurityType) {
    return m.wiFiSecurityType
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("connectAutomatically", m.GetConnectAutomatically())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectWhenNetworkNameIsHidden", m.GetConnectWhenNetworkNameIsHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkName", m.GetNetworkName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ssid", m.GetSsid())
        if err != nil {
            return err
        }
    }
    if m.GetWiFiSecurityType() != nil {
        cast := (*m.GetWiFiSecurityType()).String()
        err = writer.WriteStringValue("wiFiSecurityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectAutomatically sets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
func (m *AndroidWorkProfileWiFiConfiguration) SetConnectAutomatically(value *bool)() {
    m.connectAutomatically = value
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *AndroidWorkProfileWiFiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    m.connectWhenNetworkNameIsHidden = value
}
// SetNetworkName sets the networkName property value. Network Name
func (m *AndroidWorkProfileWiFiConfiguration) SetNetworkName(value *string)() {
    m.networkName = value
}
// SetSsid sets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *AndroidWorkProfileWiFiConfiguration) SetSsid(value *string)() {
    m.ssid = value
}
// SetWiFiSecurityType sets the wiFiSecurityType property value. Wi-Fi Security Types for Android.
func (m *AndroidWorkProfileWiFiConfiguration) SetWiFiSecurityType(value *AndroidWiFiSecurityType)() {
    m.wiFiSecurityType = value
}
