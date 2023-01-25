package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntent entity that represents an intent to apply settings to a device
type DeviceManagementIntent struct {
    Entity
    // Collection of assignments
    assignments []DeviceManagementIntentAssignmentable
    // Collection of setting categories within the intent
    categories []DeviceManagementIntentSettingCategoryable
    // The user given description
    description *string
    // Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
    deviceSettingStateSummaries []DeviceManagementIntentDeviceSettingStateSummaryable
    // Collection of states of all devices that the intent is applied to
    deviceStates []DeviceManagementIntentDeviceStateable
    // A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
    deviceStateSummary DeviceManagementIntentDeviceStateSummaryable
    // The user given display name
    displayName *string
    // Signifies whether or not the intent is assigned to users
    isAssigned *bool
    // When the intent was last modified
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // Collection of all settings to be applied
    settings []DeviceManagementSettingInstanceable
    // The ID of the template this intent was created from (if any)
    templateId *string
    // Collection of states of all users that the intent is applied to
    userStates []DeviceManagementIntentUserStateable
    // A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
    userStateSummary DeviceManagementIntentUserStateSummaryable
}
// NewDeviceManagementIntent instantiates a new deviceManagementIntent and sets the default values.
func NewDeviceManagementIntent()(*DeviceManagementIntent) {
    m := &DeviceManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntent(), nil
}
// GetAssignments gets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) GetAssignments()([]DeviceManagementIntentAssignmentable) {
    return m.assignments
}
// GetCategories gets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) GetCategories()([]DeviceManagementIntentSettingCategoryable) {
    return m.categories
}
// GetDescription gets the description property value. The user given description
func (m *DeviceManagementIntent) GetDescription()(*string) {
    return m.description
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) GetDeviceSettingStateSummaries()([]DeviceManagementIntentDeviceSettingStateSummaryable) {
    return m.deviceSettingStateSummaries
}
// GetDeviceStates gets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStates()([]DeviceManagementIntentDeviceStateable) {
    return m.deviceStates
}
// GetDeviceStateSummary gets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) GetDeviceStateSummary()(DeviceManagementIntentDeviceStateSummaryable) {
    return m.deviceStateSummary
}
// GetDisplayName gets the displayName property value. The user given display name
func (m *DeviceManagementIntent) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentSettingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentSettingCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentSettingCategoryable)
            }
            m.SetCategories(res)
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
    res["deviceSettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentDeviceSettingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentDeviceSettingStateSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentDeviceSettingStateSummaryable)
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentDeviceStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentDeviceStateable)
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
    res["deviceStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementIntentDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStateSummary(val.(DeviceManagementIntentDeviceStateSummaryable))
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
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
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
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingInstanceable)
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["userStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentUserStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentUserStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentUserStateable)
            }
            m.SetUserStates(res)
        }
        return nil
    }
    res["userStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementIntentUserStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStateSummary(val.(DeviceManagementIntentUserStateSummaryable))
        }
        return nil
    }
    return res
}
// GetIsAssigned gets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) GetIsAssigned()(*bool) {
    return m.isAssigned
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSettings gets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) GetSettings()([]DeviceManagementSettingInstanceable) {
    return m.settings
}
// GetTemplateId gets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) GetTemplateId()(*string) {
    return m.templateId
}
// GetUserStates gets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStates()([]DeviceManagementIntentUserStateable) {
    return m.userStates
}
// GetUserStateSummary gets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) GetUserStateSummary()(DeviceManagementIntentUserStateSummaryable) {
    return m.userStateSummary
}
// Serialize serializes information the current object
func (m *DeviceManagementIntent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStateSummary", m.GetDeviceStateSummary())
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
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
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
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetUserStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserStates()))
        for i, v := range m.GetUserStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStateSummary", m.GetUserStateSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. Collection of assignments
func (m *DeviceManagementIntent) SetAssignments(value []DeviceManagementIntentAssignmentable)() {
    m.assignments = value
}
// SetCategories sets the categories property value. Collection of setting categories within the intent
func (m *DeviceManagementIntent) SetCategories(value []DeviceManagementIntentSettingCategoryable)() {
    m.categories = value
}
// SetDescription sets the description property value. The user given description
func (m *DeviceManagementIntent) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *DeviceManagementIntent) SetDeviceSettingStateSummaries(value []DeviceManagementIntentDeviceSettingStateSummaryable)() {
    m.deviceSettingStateSummaries = value
}
// SetDeviceStates sets the deviceStates property value. Collection of states of all devices that the intent is applied to
func (m *DeviceManagementIntent) SetDeviceStates(value []DeviceManagementIntentDeviceStateable)() {
    m.deviceStates = value
}
// SetDeviceStateSummary sets the deviceStateSummary property value. A summary of device states and counts of devices that belong to corresponding state for all devices that the intent is applied to
func (m *DeviceManagementIntent) SetDeviceStateSummary(value DeviceManagementIntentDeviceStateSummaryable)() {
    m.deviceStateSummary = value
}
// SetDisplayName sets the displayName property value. The user given display name
func (m *DeviceManagementIntent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsAssigned sets the isAssigned property value. Signifies whether or not the intent is assigned to users
func (m *DeviceManagementIntent) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the intent was last modified
func (m *DeviceManagementIntent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementIntent) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSettings sets the settings property value. Collection of all settings to be applied
func (m *DeviceManagementIntent) SetSettings(value []DeviceManagementSettingInstanceable)() {
    m.settings = value
}
// SetTemplateId sets the templateId property value. The ID of the template this intent was created from (if any)
func (m *DeviceManagementIntent) SetTemplateId(value *string)() {
    m.templateId = value
}
// SetUserStates sets the userStates property value. Collection of states of all users that the intent is applied to
func (m *DeviceManagementIntent) SetUserStates(value []DeviceManagementIntentUserStateable)() {
    m.userStates = value
}
// SetUserStateSummary sets the userStateSummary property value. A summary of user states and counts of users that belong to corresponding state for all users that the intent is applied to
func (m *DeviceManagementIntent) SetUserStateSummary(value DeviceManagementIntentUserStateSummaryable)() {
    m.userStateSummary = value
}
