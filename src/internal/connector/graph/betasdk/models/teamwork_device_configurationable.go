package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceConfigurationable 
type TeamworkDeviceConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCameraConfiguration()(TeamworkCameraConfigurationable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayConfiguration()(TeamworkDisplayConfigurationable)
    GetHardwareConfiguration()(TeamworkHardwareConfigurationable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrophoneConfiguration()(TeamworkMicrophoneConfigurationable)
    GetSoftwareVersions()(TeamworkDeviceSoftwareVersionsable)
    GetSpeakerConfiguration()(TeamworkSpeakerConfigurationable)
    GetSystemConfiguration()(TeamworkSystemConfigurationable)
    GetTeamsClientConfiguration()(TeamworkTeamsClientConfigurationable)
    SetCameraConfiguration(value TeamworkCameraConfigurationable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayConfiguration(value TeamworkDisplayConfigurationable)()
    SetHardwareConfiguration(value TeamworkHardwareConfigurationable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrophoneConfiguration(value TeamworkMicrophoneConfigurationable)()
    SetSoftwareVersions(value TeamworkDeviceSoftwareVersionsable)()
    SetSpeakerConfiguration(value TeamworkSpeakerConfigurationable)()
    SetSystemConfiguration(value TeamworkSystemConfigurationable)()
    SetTeamsClientConfiguration(value TeamworkTeamsClientConfigurationable)()
}
