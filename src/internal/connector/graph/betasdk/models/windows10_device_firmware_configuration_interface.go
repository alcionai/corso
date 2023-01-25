package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10DeviceFirmwareConfigurationInterface 
type Windows10DeviceFirmwareConfigurationInterface struct {
    DeviceConfiguration
    // Possible values of a property
    bluetooth *Enablement
    // Possible values of a property
    bootFromBuiltInNetworkAdapters *Enablement
    // Possible values of a property
    bootFromExternalMedia *Enablement
    // Possible values of a property
    cameras *Enablement
    // Defines the permission level granted to users to enable them change Uefi settings
    changeUefiSettingsPermission *ChangeUefiSettingsPermission
    // Possible values of a property
    frontCamera *Enablement
    // Possible values of a property
    infraredCamera *Enablement
    // Possible values of a property
    microphone *Enablement
    // Possible values of a property
    microphonesAndSpeakers *Enablement
    // Possible values of a property
    nearFieldCommunication *Enablement
    // Possible values of a property
    radios *Enablement
    // Possible values of a property
    rearCamera *Enablement
    // Possible values of a property
    sdCard *Enablement
    // Possible values of a property
    simultaneousMultiThreading *Enablement
    // Possible values of a property
    usbTypeAPort *Enablement
    // Possible values of a property
    virtualizationOfCpuAndIO *Enablement
    // Possible values of a property
    wakeOnLAN *Enablement
    // Possible values of a property
    wakeOnPower *Enablement
    // Possible values of a property
    wiFi *Enablement
    // Possible values of a property
    windowsPlatformBinaryTable *Enablement
    // Possible values of a property
    wirelessWideAreaNetwork *Enablement
}
// NewWindows10DeviceFirmwareConfigurationInterface instantiates a new Windows10DeviceFirmwareConfigurationInterface and sets the default values.
func NewWindows10DeviceFirmwareConfigurationInterface()(*Windows10DeviceFirmwareConfigurationInterface) {
    m := &Windows10DeviceFirmwareConfigurationInterface{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10DeviceFirmwareConfigurationInterface";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10DeviceFirmwareConfigurationInterfaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10DeviceFirmwareConfigurationInterface(), nil
}
// GetBluetooth gets the bluetooth property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBluetooth()(*Enablement) {
    return m.bluetooth
}
// GetBootFromBuiltInNetworkAdapters gets the bootFromBuiltInNetworkAdapters property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBootFromBuiltInNetworkAdapters()(*Enablement) {
    return m.bootFromBuiltInNetworkAdapters
}
// GetBootFromExternalMedia gets the bootFromExternalMedia property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetBootFromExternalMedia()(*Enablement) {
    return m.bootFromExternalMedia
}
// GetCameras gets the cameras property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetCameras()(*Enablement) {
    return m.cameras
}
// GetChangeUefiSettingsPermission gets the changeUefiSettingsPermission property value. Defines the permission level granted to users to enable them change Uefi settings
func (m *Windows10DeviceFirmwareConfigurationInterface) GetChangeUefiSettingsPermission()(*ChangeUefiSettingsPermission) {
    return m.changeUefiSettingsPermission
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10DeviceFirmwareConfigurationInterface) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["bluetooth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetooth(val.(*Enablement))
        }
        return nil
    }
    res["bootFromBuiltInNetworkAdapters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootFromBuiltInNetworkAdapters(val.(*Enablement))
        }
        return nil
    }
    res["bootFromExternalMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootFromExternalMedia(val.(*Enablement))
        }
        return nil
    }
    res["cameras"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameras(val.(*Enablement))
        }
        return nil
    }
    res["changeUefiSettingsPermission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChangeUefiSettingsPermission)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeUefiSettingsPermission(val.(*ChangeUefiSettingsPermission))
        }
        return nil
    }
    res["frontCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrontCamera(val.(*Enablement))
        }
        return nil
    }
    res["infraredCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfraredCamera(val.(*Enablement))
        }
        return nil
    }
    res["microphone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophone(val.(*Enablement))
        }
        return nil
    }
    res["microphonesAndSpeakers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophonesAndSpeakers(val.(*Enablement))
        }
        return nil
    }
    res["nearFieldCommunication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNearFieldCommunication(val.(*Enablement))
        }
        return nil
    }
    res["radios"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRadios(val.(*Enablement))
        }
        return nil
    }
    res["rearCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRearCamera(val.(*Enablement))
        }
        return nil
    }
    res["sdCard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSdCard(val.(*Enablement))
        }
        return nil
    }
    res["simultaneousMultiThreading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimultaneousMultiThreading(val.(*Enablement))
        }
        return nil
    }
    res["usbTypeAPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsbTypeAPort(val.(*Enablement))
        }
        return nil
    }
    res["virtualizationOfCpuAndIO"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualizationOfCpuAndIO(val.(*Enablement))
        }
        return nil
    }
    res["wakeOnLAN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWakeOnLAN(val.(*Enablement))
        }
        return nil
    }
    res["wakeOnPower"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWakeOnPower(val.(*Enablement))
        }
        return nil
    }
    res["wiFi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFi(val.(*Enablement))
        }
        return nil
    }
    res["windowsPlatformBinaryTable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsPlatformBinaryTable(val.(*Enablement))
        }
        return nil
    }
    res["wirelessWideAreaNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWirelessWideAreaNetwork(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetFrontCamera gets the frontCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetFrontCamera()(*Enablement) {
    return m.frontCamera
}
// GetInfraredCamera gets the infraredCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetInfraredCamera()(*Enablement) {
    return m.infraredCamera
}
// GetMicrophone gets the microphone property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetMicrophone()(*Enablement) {
    return m.microphone
}
// GetMicrophonesAndSpeakers gets the microphonesAndSpeakers property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetMicrophonesAndSpeakers()(*Enablement) {
    return m.microphonesAndSpeakers
}
// GetNearFieldCommunication gets the nearFieldCommunication property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetNearFieldCommunication()(*Enablement) {
    return m.nearFieldCommunication
}
// GetRadios gets the radios property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetRadios()(*Enablement) {
    return m.radios
}
// GetRearCamera gets the rearCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetRearCamera()(*Enablement) {
    return m.rearCamera
}
// GetSdCard gets the sdCard property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetSdCard()(*Enablement) {
    return m.sdCard
}
// GetSimultaneousMultiThreading gets the simultaneousMultiThreading property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetSimultaneousMultiThreading()(*Enablement) {
    return m.simultaneousMultiThreading
}
// GetUsbTypeAPort gets the usbTypeAPort property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetUsbTypeAPort()(*Enablement) {
    return m.usbTypeAPort
}
// GetVirtualizationOfCpuAndIO gets the virtualizationOfCpuAndIO property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetVirtualizationOfCpuAndIO()(*Enablement) {
    return m.virtualizationOfCpuAndIO
}
// GetWakeOnLAN gets the wakeOnLAN property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWakeOnLAN()(*Enablement) {
    return m.wakeOnLAN
}
// GetWakeOnPower gets the wakeOnPower property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWakeOnPower()(*Enablement) {
    return m.wakeOnPower
}
// GetWiFi gets the wiFi property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWiFi()(*Enablement) {
    return m.wiFi
}
// GetWindowsPlatformBinaryTable gets the windowsPlatformBinaryTable property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWindowsPlatformBinaryTable()(*Enablement) {
    return m.windowsPlatformBinaryTable
}
// GetWirelessWideAreaNetwork gets the wirelessWideAreaNetwork property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) GetWirelessWideAreaNetwork()(*Enablement) {
    return m.wirelessWideAreaNetwork
}
// Serialize serializes information the current object
func (m *Windows10DeviceFirmwareConfigurationInterface) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBluetooth() != nil {
        cast := (*m.GetBluetooth()).String()
        err = writer.WriteStringValue("bluetooth", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBootFromBuiltInNetworkAdapters() != nil {
        cast := (*m.GetBootFromBuiltInNetworkAdapters()).String()
        err = writer.WriteStringValue("bootFromBuiltInNetworkAdapters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBootFromExternalMedia() != nil {
        cast := (*m.GetBootFromExternalMedia()).String()
        err = writer.WriteStringValue("bootFromExternalMedia", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCameras() != nil {
        cast := (*m.GetCameras()).String()
        err = writer.WriteStringValue("cameras", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChangeUefiSettingsPermission() != nil {
        cast := (*m.GetChangeUefiSettingsPermission()).String()
        err = writer.WriteStringValue("changeUefiSettingsPermission", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFrontCamera() != nil {
        cast := (*m.GetFrontCamera()).String()
        err = writer.WriteStringValue("frontCamera", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInfraredCamera() != nil {
        cast := (*m.GetInfraredCamera()).String()
        err = writer.WriteStringValue("infraredCamera", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrophone() != nil {
        cast := (*m.GetMicrophone()).String()
        err = writer.WriteStringValue("microphone", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrophonesAndSpeakers() != nil {
        cast := (*m.GetMicrophonesAndSpeakers()).String()
        err = writer.WriteStringValue("microphonesAndSpeakers", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNearFieldCommunication() != nil {
        cast := (*m.GetNearFieldCommunication()).String()
        err = writer.WriteStringValue("nearFieldCommunication", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRadios() != nil {
        cast := (*m.GetRadios()).String()
        err = writer.WriteStringValue("radios", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRearCamera() != nil {
        cast := (*m.GetRearCamera()).String()
        err = writer.WriteStringValue("rearCamera", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSdCard() != nil {
        cast := (*m.GetSdCard()).String()
        err = writer.WriteStringValue("sdCard", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSimultaneousMultiThreading() != nil {
        cast := (*m.GetSimultaneousMultiThreading()).String()
        err = writer.WriteStringValue("simultaneousMultiThreading", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsbTypeAPort() != nil {
        cast := (*m.GetUsbTypeAPort()).String()
        err = writer.WriteStringValue("usbTypeAPort", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVirtualizationOfCpuAndIO() != nil {
        cast := (*m.GetVirtualizationOfCpuAndIO()).String()
        err = writer.WriteStringValue("virtualizationOfCpuAndIO", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWakeOnLAN() != nil {
        cast := (*m.GetWakeOnLAN()).String()
        err = writer.WriteStringValue("wakeOnLAN", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWakeOnPower() != nil {
        cast := (*m.GetWakeOnPower()).String()
        err = writer.WriteStringValue("wakeOnPower", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWiFi() != nil {
        cast := (*m.GetWiFi()).String()
        err = writer.WriteStringValue("wiFi", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsPlatformBinaryTable() != nil {
        cast := (*m.GetWindowsPlatformBinaryTable()).String()
        err = writer.WriteStringValue("windowsPlatformBinaryTable", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWirelessWideAreaNetwork() != nil {
        cast := (*m.GetWirelessWideAreaNetwork()).String()
        err = writer.WriteStringValue("wirelessWideAreaNetwork", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBluetooth sets the bluetooth property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBluetooth(value *Enablement)() {
    m.bluetooth = value
}
// SetBootFromBuiltInNetworkAdapters sets the bootFromBuiltInNetworkAdapters property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBootFromBuiltInNetworkAdapters(value *Enablement)() {
    m.bootFromBuiltInNetworkAdapters = value
}
// SetBootFromExternalMedia sets the bootFromExternalMedia property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetBootFromExternalMedia(value *Enablement)() {
    m.bootFromExternalMedia = value
}
// SetCameras sets the cameras property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetCameras(value *Enablement)() {
    m.cameras = value
}
// SetChangeUefiSettingsPermission sets the changeUefiSettingsPermission property value. Defines the permission level granted to users to enable them change Uefi settings
func (m *Windows10DeviceFirmwareConfigurationInterface) SetChangeUefiSettingsPermission(value *ChangeUefiSettingsPermission)() {
    m.changeUefiSettingsPermission = value
}
// SetFrontCamera sets the frontCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetFrontCamera(value *Enablement)() {
    m.frontCamera = value
}
// SetInfraredCamera sets the infraredCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetInfraredCamera(value *Enablement)() {
    m.infraredCamera = value
}
// SetMicrophone sets the microphone property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetMicrophone(value *Enablement)() {
    m.microphone = value
}
// SetMicrophonesAndSpeakers sets the microphonesAndSpeakers property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetMicrophonesAndSpeakers(value *Enablement)() {
    m.microphonesAndSpeakers = value
}
// SetNearFieldCommunication sets the nearFieldCommunication property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetNearFieldCommunication(value *Enablement)() {
    m.nearFieldCommunication = value
}
// SetRadios sets the radios property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetRadios(value *Enablement)() {
    m.radios = value
}
// SetRearCamera sets the rearCamera property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetRearCamera(value *Enablement)() {
    m.rearCamera = value
}
// SetSdCard sets the sdCard property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetSdCard(value *Enablement)() {
    m.sdCard = value
}
// SetSimultaneousMultiThreading sets the simultaneousMultiThreading property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetSimultaneousMultiThreading(value *Enablement)() {
    m.simultaneousMultiThreading = value
}
// SetUsbTypeAPort sets the usbTypeAPort property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetUsbTypeAPort(value *Enablement)() {
    m.usbTypeAPort = value
}
// SetVirtualizationOfCpuAndIO sets the virtualizationOfCpuAndIO property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetVirtualizationOfCpuAndIO(value *Enablement)() {
    m.virtualizationOfCpuAndIO = value
}
// SetWakeOnLAN sets the wakeOnLAN property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWakeOnLAN(value *Enablement)() {
    m.wakeOnLAN = value
}
// SetWakeOnPower sets the wakeOnPower property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWakeOnPower(value *Enablement)() {
    m.wakeOnPower = value
}
// SetWiFi sets the wiFi property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWiFi(value *Enablement)() {
    m.wiFi = value
}
// SetWindowsPlatformBinaryTable sets the windowsPlatformBinaryTable property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWindowsPlatformBinaryTable(value *Enablement)() {
    m.windowsPlatformBinaryTable = value
}
// SetWirelessWideAreaNetwork sets the wirelessWideAreaNetwork property value. Possible values of a property
func (m *Windows10DeviceFirmwareConfigurationInterface) SetWirelessWideAreaNetwork(value *Enablement)() {
    m.wirelessWideAreaNetwork = value
}
