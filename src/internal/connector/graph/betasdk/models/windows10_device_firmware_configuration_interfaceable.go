package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10DeviceFirmwareConfigurationInterfaceable 
type Windows10DeviceFirmwareConfigurationInterfaceable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBluetooth()(*Enablement)
    GetBootFromBuiltInNetworkAdapters()(*Enablement)
    GetBootFromExternalMedia()(*Enablement)
    GetCameras()(*Enablement)
    GetChangeUefiSettingsPermission()(*ChangeUefiSettingsPermission)
    GetFrontCamera()(*Enablement)
    GetInfraredCamera()(*Enablement)
    GetMicrophone()(*Enablement)
    GetMicrophonesAndSpeakers()(*Enablement)
    GetNearFieldCommunication()(*Enablement)
    GetRadios()(*Enablement)
    GetRearCamera()(*Enablement)
    GetSdCard()(*Enablement)
    GetSimultaneousMultiThreading()(*Enablement)
    GetUsbTypeAPort()(*Enablement)
    GetVirtualizationOfCpuAndIO()(*Enablement)
    GetWakeOnLAN()(*Enablement)
    GetWakeOnPower()(*Enablement)
    GetWiFi()(*Enablement)
    GetWindowsPlatformBinaryTable()(*Enablement)
    GetWirelessWideAreaNetwork()(*Enablement)
    SetBluetooth(value *Enablement)()
    SetBootFromBuiltInNetworkAdapters(value *Enablement)()
    SetBootFromExternalMedia(value *Enablement)()
    SetCameras(value *Enablement)()
    SetChangeUefiSettingsPermission(value *ChangeUefiSettingsPermission)()
    SetFrontCamera(value *Enablement)()
    SetInfraredCamera(value *Enablement)()
    SetMicrophone(value *Enablement)()
    SetMicrophonesAndSpeakers(value *Enablement)()
    SetNearFieldCommunication(value *Enablement)()
    SetRadios(value *Enablement)()
    SetRearCamera(value *Enablement)()
    SetSdCard(value *Enablement)()
    SetSimultaneousMultiThreading(value *Enablement)()
    SetUsbTypeAPort(value *Enablement)()
    SetVirtualizationOfCpuAndIO(value *Enablement)()
    SetWakeOnLAN(value *Enablement)()
    SetWakeOnPower(value *Enablement)()
    SetWiFi(value *Enablement)()
    SetWindowsPlatformBinaryTable(value *Enablement)()
    SetWirelessWideAreaNetwork(value *Enablement)()
}
