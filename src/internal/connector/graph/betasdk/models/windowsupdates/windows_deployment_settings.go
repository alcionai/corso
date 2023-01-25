package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDeploymentSettings 
type WindowsDeploymentSettings struct {
    DeploymentSettings
    // Settings governing the user's update experience on a device.
    userExperience UserExperienceSettingsable
}
// NewWindowsDeploymentSettings instantiates a new WindowsDeploymentSettings and sets the default values.
func NewWindowsDeploymentSettings()(*WindowsDeploymentSettings) {
    m := &WindowsDeploymentSettings{
        DeploymentSettings: *NewDeploymentSettings(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.windowsDeploymentSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsDeploymentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDeploymentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDeploymentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDeploymentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeploymentSettings.GetFieldDeserializers()
    res["userExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperience(val.(UserExperienceSettingsable))
        }
        return nil
    }
    return res
}
// GetUserExperience gets the userExperience property value. Settings governing the user's update experience on a device.
func (m *WindowsDeploymentSettings) GetUserExperience()(UserExperienceSettingsable) {
    return m.userExperience
}
// Serialize serializes information the current object
func (m *WindowsDeploymentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeploymentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("userExperience", m.GetUserExperience())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserExperience sets the userExperience property value. Settings governing the user's update experience on a device.
func (m *WindowsDeploymentSettings) SetUserExperience(value UserExperienceSettingsable)() {
    m.userExperience = value
}
