package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceConfiguration 
type TeamworkDeviceConfiguration struct {
    Entity
    // The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
    cameraConfiguration TeamworkCameraConfigurationable
    // Identity of the user who created the device configuration document.
    createdBy IdentitySetable
    // The UTC date and time when the device configuration document was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display configuration.
    displayConfiguration TeamworkDisplayConfigurationable
    // The hardware configuration. Applicable only for Teams Rooms-enabled devices.
    hardwareConfiguration TeamworkHardwareConfigurationable
    // Identity of the user who last modified the device configuration.
    lastModifiedBy IdentitySetable
    // The UTC date and time when the device configuration was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The microphone configuration. Applicable only for Teams Rooms-enabled devices.
    microphoneConfiguration TeamworkMicrophoneConfigurationable
    // Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
    softwareVersions TeamworkDeviceSoftwareVersionsable
    // The speaker configuration. Applicable only for Teams Rooms-enabled devices.
    speakerConfiguration TeamworkSpeakerConfigurationable
    // The system configuration. Not applicable for Teams Rooms-enabled devices.
    systemConfiguration TeamworkSystemConfigurationable
    // The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
    teamsClientConfiguration TeamworkTeamsClientConfigurationable
}
// NewTeamworkDeviceConfiguration instantiates a new teamworkDeviceConfiguration and sets the default values.
func NewTeamworkDeviceConfiguration()(*TeamworkDeviceConfiguration) {
    m := &TeamworkDeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDeviceConfiguration(), nil
}
// GetCameraConfiguration gets the cameraConfiguration property value. The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetCameraConfiguration()(TeamworkCameraConfigurationable) {
    return m.cameraConfiguration
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device configuration document.
func (m *TeamworkDeviceConfiguration) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device configuration document was created.
func (m *TeamworkDeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayConfiguration gets the displayConfiguration property value. The display configuration.
func (m *TeamworkDeviceConfiguration) GetDisplayConfiguration()(TeamworkDisplayConfigurationable) {
    return m.displayConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cameraConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkCameraConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraConfiguration(val.(TeamworkCameraConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDisplayConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayConfiguration(val.(TeamworkDisplayConfigurationable))
        }
        return nil
    }
    res["hardwareConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHardwareConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareConfiguration(val.(TeamworkHardwareConfigurationable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["microphoneConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkMicrophoneConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneConfiguration(val.(TeamworkMicrophoneConfigurationable))
        }
        return nil
    }
    res["softwareVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDeviceSoftwareVersionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareVersions(val.(TeamworkDeviceSoftwareVersionsable))
        }
        return nil
    }
    res["speakerConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSpeakerConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerConfiguration(val.(TeamworkSpeakerConfigurationable))
        }
        return nil
    }
    res["systemConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkSystemConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemConfiguration(val.(TeamworkSystemConfigurationable))
        }
        return nil
    }
    res["teamsClientConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkTeamsClientConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientConfiguration(val.(TeamworkTeamsClientConfigurationable))
        }
        return nil
    }
    return res
}
// GetHardwareConfiguration gets the hardwareConfiguration property value. The hardware configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetHardwareConfiguration()(TeamworkHardwareConfigurationable) {
    return m.hardwareConfiguration
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device configuration.
func (m *TeamworkDeviceConfiguration) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device configuration was last modified.
func (m *TeamworkDeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMicrophoneConfiguration gets the microphoneConfiguration property value. The microphone configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetMicrophoneConfiguration()(TeamworkMicrophoneConfigurationable) {
    return m.microphoneConfiguration
}
// GetSoftwareVersions gets the softwareVersions property value. Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
func (m *TeamworkDeviceConfiguration) GetSoftwareVersions()(TeamworkDeviceSoftwareVersionsable) {
    return m.softwareVersions
}
// GetSpeakerConfiguration gets the speakerConfiguration property value. The speaker configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetSpeakerConfiguration()(TeamworkSpeakerConfigurationable) {
    return m.speakerConfiguration
}
// GetSystemConfiguration gets the systemConfiguration property value. The system configuration. Not applicable for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetSystemConfiguration()(TeamworkSystemConfigurationable) {
    return m.systemConfiguration
}
// GetTeamsClientConfiguration gets the teamsClientConfiguration property value. The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) GetTeamsClientConfiguration()(TeamworkTeamsClientConfigurationable) {
    return m.teamsClientConfiguration
}
// Serialize serializes information the current object
func (m *TeamworkDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cameraConfiguration", m.GetCameraConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("displayConfiguration", m.GetDisplayConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hardwareConfiguration", m.GetHardwareConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("microphoneConfiguration", m.GetMicrophoneConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareVersions", m.GetSoftwareVersions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("speakerConfiguration", m.GetSpeakerConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("systemConfiguration", m.GetSystemConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsClientConfiguration", m.GetTeamsClientConfiguration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCameraConfiguration sets the cameraConfiguration property value. The camera configuration. Applicable only for Microsoft Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetCameraConfiguration(value TeamworkCameraConfigurationable)() {
    m.cameraConfiguration = value
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device configuration document.
func (m *TeamworkDeviceConfiguration) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device configuration document was created.
func (m *TeamworkDeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayConfiguration sets the displayConfiguration property value. The display configuration.
func (m *TeamworkDeviceConfiguration) SetDisplayConfiguration(value TeamworkDisplayConfigurationable)() {
    m.displayConfiguration = value
}
// SetHardwareConfiguration sets the hardwareConfiguration property value. The hardware configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetHardwareConfiguration(value TeamworkHardwareConfigurationable)() {
    m.hardwareConfiguration = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device configuration.
func (m *TeamworkDeviceConfiguration) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device configuration was last modified.
func (m *TeamworkDeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMicrophoneConfiguration sets the microphoneConfiguration property value. The microphone configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetMicrophoneConfiguration(value TeamworkMicrophoneConfigurationable)() {
    m.microphoneConfiguration = value
}
// SetSoftwareVersions sets the softwareVersions property value. Information related to software versions for the device, such as firmware, operating system, Teams client, and admin agent.
func (m *TeamworkDeviceConfiguration) SetSoftwareVersions(value TeamworkDeviceSoftwareVersionsable)() {
    m.softwareVersions = value
}
// SetSpeakerConfiguration sets the speakerConfiguration property value. The speaker configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetSpeakerConfiguration(value TeamworkSpeakerConfigurationable)() {
    m.speakerConfiguration = value
}
// SetSystemConfiguration sets the systemConfiguration property value. The system configuration. Not applicable for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetSystemConfiguration(value TeamworkSystemConfigurationable)() {
    m.systemConfiguration = value
}
// SetTeamsClientConfiguration sets the teamsClientConfiguration property value. The Teams client configuration. Applicable only for Teams Rooms-enabled devices.
func (m *TeamworkDeviceConfiguration) SetTeamsClientConfiguration(value TeamworkTeamsClientConfigurationable)() {
    m.teamsClientConfiguration = value
}
