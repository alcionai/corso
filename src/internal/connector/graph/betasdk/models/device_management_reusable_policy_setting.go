package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementReusablePolicySetting graph model for a reusable setting
type DeviceManagementReusablePolicySetting struct {
    Entity
    // reusable setting creation date and time. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // reusable setting description supplied by user.
    description *string
    // reusable setting display name supplied by user.
    displayName *string
    // date and time when reusable setting was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // configuration policies referencing the current reusable setting. This property is read-only.
    referencingConfigurationPolicies []DeviceManagementConfigurationPolicyable
    // count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
    referencingConfigurationPolicyCount *int32
    // setting definition id associated with this reusable setting.
    settingDefinitionId *string
    // reusable setting configuration instance
    settingInstance DeviceManagementConfigurationSettingInstanceable
    // version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
    version *int32
}
// NewDeviceManagementReusablePolicySetting instantiates a new deviceManagementReusablePolicySetting and sets the default values.
func NewDeviceManagementReusablePolicySetting()(*DeviceManagementReusablePolicySetting) {
    m := &DeviceManagementReusablePolicySetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementReusablePolicySetting(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementReusablePolicySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["referencingConfigurationPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationPolicyable)
            }
            m.SetReferencingConfigurationPolicies(res)
        }
        return nil
    }
    res["referencingConfigurationPolicyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencingConfigurationPolicyCount(val)
        }
        return nil
    }
    res["settingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    res["settingInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstance(val.(DeviceManagementConfigurationSettingInstanceable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetReferencingConfigurationPolicies gets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicyable) {
    return m.referencingConfigurationPolicies
}
// GetReferencingConfigurationPolicyCount gets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicyCount()(*int32) {
    return m.referencingConfigurationPolicyCount
}
// GetSettingDefinitionId gets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) GetSettingDefinitionId()(*string) {
    return m.settingDefinitionId
}
// GetSettingInstance gets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable) {
    return m.settingInstance
}
// GetVersion gets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceManagementReusablePolicySetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetReferencingConfigurationPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferencingConfigurationPolicies()))
        for i, v := range m.GetReferencingConfigurationPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("referencingConfigurationPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settingInstance", m.GetSettingInstance())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetReferencingConfigurationPolicies sets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)() {
    m.referencingConfigurationPolicies = value
}
// SetReferencingConfigurationPolicyCount sets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicyCount(value *int32)() {
    m.referencingConfigurationPolicyCount = value
}
// SetSettingDefinitionId sets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
// SetSettingInstance sets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)() {
    m.settingInstance = value
}
// SetVersion sets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetVersion(value *int32)() {
    m.version = value
}
