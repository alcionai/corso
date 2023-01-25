package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcServicePlan 
type CloudPcServicePlan struct {
    Entity
    // The name for the service plan. Read-only.
    displayName *string
    // The size of the RAM in GB. Read-only.
    ramInGB *int32
    // The size of the OS Disk in GB. Read-only.
    storageInGB *int32
    // The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
    type_escaped *CloudPcServicePlanType
    // The size of the user profile disk in GB. Read-only.
    userProfileInGB *int32
    // The number of vCPUs. Read-only.
    vCpuCount *int32
}
// NewCloudPcServicePlan instantiates a new CloudPcServicePlan and sets the default values.
func NewCloudPcServicePlan()(*CloudPcServicePlan) {
    m := &CloudPcServicePlan{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcServicePlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcServicePlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcServicePlan(), nil
}
// GetDisplayName gets the displayName property value. The name for the service plan. Read-only.
func (m *CloudPcServicePlan) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcServicePlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["ramInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamInGB(val)
        }
        return nil
    }
    res["storageInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageInGB(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcServicePlanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*CloudPcServicePlanType))
        }
        return nil
    }
    res["userProfileInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserProfileInGB(val)
        }
        return nil
    }
    res["vCpuCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVCpuCount(val)
        }
        return nil
    }
    return res
}
// GetRamInGB gets the ramInGB property value. The size of the RAM in GB. Read-only.
func (m *CloudPcServicePlan) GetRamInGB()(*int32) {
    return m.ramInGB
}
// GetStorageInGB gets the storageInGB property value. The size of the OS Disk in GB. Read-only.
func (m *CloudPcServicePlan) GetStorageInGB()(*int32) {
    return m.storageInGB
}
// GetType gets the type property value. The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
func (m *CloudPcServicePlan) GetType()(*CloudPcServicePlanType) {
    return m.type_escaped
}
// GetUserProfileInGB gets the userProfileInGB property value. The size of the user profile disk in GB. Read-only.
func (m *CloudPcServicePlan) GetUserProfileInGB()(*int32) {
    return m.userProfileInGB
}
// GetVCpuCount gets the vCpuCount property value. The number of vCPUs. Read-only.
func (m *CloudPcServicePlan) GetVCpuCount()(*int32) {
    return m.vCpuCount
}
// Serialize serializes information the current object
func (m *CloudPcServicePlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ramInGB", m.GetRamInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("storageInGB", m.GetStorageInGB())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("userProfileInGB", m.GetUserProfileInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("vCpuCount", m.GetVCpuCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name for the service plan. Read-only.
func (m *CloudPcServicePlan) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetRamInGB sets the ramInGB property value. The size of the RAM in GB. Read-only.
func (m *CloudPcServicePlan) SetRamInGB(value *int32)() {
    m.ramInGB = value
}
// SetStorageInGB sets the storageInGB property value. The size of the OS Disk in GB. Read-only.
func (m *CloudPcServicePlan) SetStorageInGB(value *int32)() {
    m.storageInGB = value
}
// SetType sets the type property value. The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
func (m *CloudPcServicePlan) SetType(value *CloudPcServicePlanType)() {
    m.type_escaped = value
}
// SetUserProfileInGB sets the userProfileInGB property value. The size of the user profile disk in GB. Read-only.
func (m *CloudPcServicePlan) SetUserProfileInGB(value *int32)() {
    m.userProfileInGB = value
}
// SetVCpuCount sets the vCpuCount property value. The number of vCPUs. Read-only.
func (m *CloudPcServicePlan) SetVCpuCount(value *int32)() {
    m.vCpuCount = value
}
