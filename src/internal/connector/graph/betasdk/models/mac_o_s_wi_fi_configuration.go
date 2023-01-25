package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSWiFiConfiguration 
type MacOSWiFiConfiguration struct {
    DeviceConfiguration
    // Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
    connectAutomatically *bool
    // Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
    connectWhenNetworkNameIsHidden *bool
    // Network Name
    networkName *string
    // This is the pre-shared key for WPA Personal Wi-Fi network.
    preSharedKey *string
    // URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is typically the location of PAC (Proxy Auto Configuration) file.
    proxyAutomaticConfigurationUrl *string
    // IP Address or DNS hostname of the proxy server when manual configuration is selected.
    proxyManualAddress *string
    // Port of the proxy server when manual configuration is selected.
    proxyManualPort *int32
    // Wi-Fi Proxy Settings.
    proxySettings *WiFiProxySetting
    // This is the name of the Wi-Fi network that is broadcast to all devices.
    ssid *string
    // Wi-Fi Security Types.
    wiFiSecurityType *WiFiSecurityType
}
// NewMacOSWiFiConfiguration instantiates a new MacOSWiFiConfiguration and sets the default values.
func NewMacOSWiFiConfiguration()(*MacOSWiFiConfiguration) {
    m := &MacOSWiFiConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSWiFiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSWiFiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSWiFiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.macOSEnterpriseWiFiConfiguration":
                        return NewMacOSEnterpriseWiFiConfiguration(), nil
                }
            }
        }
    }
    return NewMacOSWiFiConfiguration(), nil
}
// GetConnectAutomatically gets the connectAutomatically property value. Connect automatically when this network is in range. Setting this to true will skip the user prompt and automatically connect the device to Wi-Fi network.
func (m *MacOSWiFiConfiguration) GetConnectAutomatically()(*bool) {
    return m.connectAutomatically
}
// GetConnectWhenNetworkNameIsHidden gets the connectWhenNetworkNameIsHidden property value. Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *MacOSWiFiConfiguration) GetConnectWhenNetworkNameIsHidden()(*bool) {
    return m.connectWhenNetworkNameIsHidden
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSWiFiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["preSharedKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreSharedKey(val)
        }
        return nil
    }
    res["proxyAutomaticConfigurationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyAutomaticConfigurationUrl(val)
        }
        return nil
    }
    res["proxyManualAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyManualAddress(val)
        }
        return nil
    }
    res["proxyManualPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxyManualPort(val)
        }
        return nil
    }
    res["proxySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWiFiProxySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxySettings(val.(*WiFiProxySetting))
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
        val, err := n.GetEnumValue(ParseWiFiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiSecurityType(val.(*WiFiSecurityType))
        }
        return nil
    }
    return res
}
// GetNetworkName gets the networkName property value. Network Name
func (m *MacOSWiFiConfiguration) GetNetworkName()(*string) {
    return m.networkName
}
// GetPreSharedKey gets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *MacOSWiFiConfiguration) GetPreSharedKey()(*string) {
    return m.preSharedKey
}
// GetProxyAutomaticConfigurationUrl gets the proxyAutomaticConfigurationUrl property value. URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is typically the location of PAC (Proxy Auto Configuration) file.
func (m *MacOSWiFiConfiguration) GetProxyAutomaticConfigurationUrl()(*string) {
    return m.proxyAutomaticConfigurationUrl
}
// GetProxyManualAddress gets the proxyManualAddress property value. IP Address or DNS hostname of the proxy server when manual configuration is selected.
func (m *MacOSWiFiConfiguration) GetProxyManualAddress()(*string) {
    return m.proxyManualAddress
}
// GetProxyManualPort gets the proxyManualPort property value. Port of the proxy server when manual configuration is selected.
func (m *MacOSWiFiConfiguration) GetProxyManualPort()(*int32) {
    return m.proxyManualPort
}
// GetProxySettings gets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *MacOSWiFiConfiguration) GetProxySettings()(*WiFiProxySetting) {
    return m.proxySettings
}
// GetSsid gets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *MacOSWiFiConfiguration) GetSsid()(*string) {
    return m.ssid
}
// GetWiFiSecurityType gets the wiFiSecurityType property value. Wi-Fi Security Types.
func (m *MacOSWiFiConfiguration) GetWiFiSecurityType()(*WiFiSecurityType) {
    return m.wiFiSecurityType
}
// Serialize serializes information the current object
func (m *MacOSWiFiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("preSharedKey", m.GetPreSharedKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("proxyAutomaticConfigurationUrl", m.GetProxyAutomaticConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("proxyManualAddress", m.GetProxyManualAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("proxyManualPort", m.GetProxyManualPort())
        if err != nil {
            return err
        }
    }
    if m.GetProxySettings() != nil {
        cast := (*m.GetProxySettings()).String()
        err = writer.WriteStringValue("proxySettings", &cast)
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
func (m *MacOSWiFiConfiguration) SetConnectAutomatically(value *bool)() {
    m.connectAutomatically = value
}
// SetConnectWhenNetworkNameIsHidden sets the connectWhenNetworkNameIsHidden property value. Connect when the network is not broadcasting its name (SSID). When set to true, this profile forces the device to connect to a network that doesn't broadcast its SSID to all devices.
func (m *MacOSWiFiConfiguration) SetConnectWhenNetworkNameIsHidden(value *bool)() {
    m.connectWhenNetworkNameIsHidden = value
}
// SetNetworkName sets the networkName property value. Network Name
func (m *MacOSWiFiConfiguration) SetNetworkName(value *string)() {
    m.networkName = value
}
// SetPreSharedKey sets the preSharedKey property value. This is the pre-shared key for WPA Personal Wi-Fi network.
func (m *MacOSWiFiConfiguration) SetPreSharedKey(value *string)() {
    m.preSharedKey = value
}
// SetProxyAutomaticConfigurationUrl sets the proxyAutomaticConfigurationUrl property value. URL of the proxy server automatic configuration script when automatic configuration is selected. This URL is typically the location of PAC (Proxy Auto Configuration) file.
func (m *MacOSWiFiConfiguration) SetProxyAutomaticConfigurationUrl(value *string)() {
    m.proxyAutomaticConfigurationUrl = value
}
// SetProxyManualAddress sets the proxyManualAddress property value. IP Address or DNS hostname of the proxy server when manual configuration is selected.
func (m *MacOSWiFiConfiguration) SetProxyManualAddress(value *string)() {
    m.proxyManualAddress = value
}
// SetProxyManualPort sets the proxyManualPort property value. Port of the proxy server when manual configuration is selected.
func (m *MacOSWiFiConfiguration) SetProxyManualPort(value *int32)() {
    m.proxyManualPort = value
}
// SetProxySettings sets the proxySettings property value. Wi-Fi Proxy Settings.
func (m *MacOSWiFiConfiguration) SetProxySettings(value *WiFiProxySetting)() {
    m.proxySettings = value
}
// SetSsid sets the ssid property value. This is the name of the Wi-Fi network that is broadcast to all devices.
func (m *MacOSWiFiConfiguration) SetSsid(value *string)() {
    m.ssid = value
}
// SetWiFiSecurityType sets the wiFiSecurityType property value. Wi-Fi Security Types.
func (m *MacOSWiFiConfiguration) SetWiFiSecurityType(value *WiFiSecurityType)() {
    m.wiFiSecurityType = value
}
