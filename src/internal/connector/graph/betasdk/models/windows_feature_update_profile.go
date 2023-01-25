package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFeatureUpdateProfile windows Feature Update Profile
type WindowsFeatureUpdateProfile struct {
    Entity
    // The list of group assignments of the profile.
    assignments []WindowsFeatureUpdateProfileAssignmentable
    // The date time that the profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Friendly display name of the quality update profile deployable content
    deployableContentDisplayName *string
    // The description of the profile which is specified by the user.
    description *string
    // The display name of the profile.
    displayName *string
    // The last supported date for a feature update
    endOfSupportDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
    featureUpdateVersion *string
    // The date time that the profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Feature Update entity.
    roleScopeTagIds []string
    // The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
    rolloutSettings WindowsUpdateRolloutSettingsable
}
// NewWindowsFeatureUpdateProfile instantiates a new windowsFeatureUpdateProfile and sets the default values.
func NewWindowsFeatureUpdateProfile()(*WindowsFeatureUpdateProfile) {
    m := &WindowsFeatureUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsFeatureUpdateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFeatureUpdateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFeatureUpdateProfile(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsFeatureUpdateProfile) GetAssignments()([]WindowsFeatureUpdateProfileAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsFeatureUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeployableContentDisplayName gets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsFeatureUpdateProfile) GetDeployableContentDisplayName()(*string) {
    return m.deployableContentDisplayName
}
// GetDescription gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsFeatureUpdateProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the profile.
func (m *WindowsFeatureUpdateProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndOfSupportDate gets the endOfSupportDate property value. The last supported date for a feature update
func (m *WindowsFeatureUpdateProfile) GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endOfSupportDate
}
// GetFeatureUpdateVersion gets the featureUpdateVersion property value. The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
func (m *WindowsFeatureUpdateProfile) GetFeatureUpdateVersion()(*string) {
    return m.featureUpdateVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsFeatureUpdateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsFeatureUpdateProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFeatureUpdateProfileAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsFeatureUpdateProfileAssignmentable)
            }
            m.SetAssignments(res)
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
    res["deployableContentDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployableContentDisplayName(val)
        }
        return nil
    }
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endOfSupportDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndOfSupportDate(val)
        }
        return nil
    }
    res["featureUpdateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdateVersion(val)
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
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["rolloutSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsUpdateRolloutSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRolloutSettings(val.(WindowsUpdateRolloutSettingsable))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsFeatureUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Feature Update entity.
func (m *WindowsFeatureUpdateProfile) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetRolloutSettings gets the rolloutSettings property value. The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
func (m *WindowsFeatureUpdateProfile) GetRolloutSettings()(WindowsUpdateRolloutSettingsable) {
    return m.rolloutSettings
}
// Serialize serializes information the current object
func (m *WindowsFeatureUpdateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
        err = writer.WriteStringValue("deployableContentDisplayName", m.GetDeployableContentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endOfSupportDate", m.GetEndOfSupportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("featureUpdateVersion", m.GetFeatureUpdateVersion())
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
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rolloutSettings", m.GetRolloutSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments of the profile.
func (m *WindowsFeatureUpdateProfile) SetAssignments(value []WindowsFeatureUpdateProfileAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsFeatureUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeployableContentDisplayName sets the deployableContentDisplayName property value. Friendly display name of the quality update profile deployable content
func (m *WindowsFeatureUpdateProfile) SetDeployableContentDisplayName(value *string)() {
    m.deployableContentDisplayName = value
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsFeatureUpdateProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the profile.
func (m *WindowsFeatureUpdateProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndOfSupportDate sets the endOfSupportDate property value. The last supported date for a feature update
func (m *WindowsFeatureUpdateProfile) SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endOfSupportDate = value
}
// SetFeatureUpdateVersion sets the featureUpdateVersion property value. The feature update version that will be deployed to the devices targeted by this profile. The version could be any supported version for example 1709, 1803 or 1809 and so on.
func (m *WindowsFeatureUpdateProfile) SetFeatureUpdateVersion(value *string)() {
    m.featureUpdateVersion = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsFeatureUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Feature Update entity.
func (m *WindowsFeatureUpdateProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetRolloutSettings sets the rolloutSettings property value. The windows update rollout settings, including offer start date time, offer end date time, and days between each set of offers.
func (m *WindowsFeatureUpdateProfile) SetRolloutSettings(value WindowsUpdateRolloutSettingsable)() {
    m.rolloutSettings = value
}
