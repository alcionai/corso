package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingAbstractImplementationConstraint 
type DeviceManagementSettingAbstractImplementationConstraint struct {
    DeviceManagementConstraint
    // List of value which means not configured for the setting
    allowedAbstractImplementationDefinitionIds []string
}
// NewDeviceManagementSettingAbstractImplementationConstraint instantiates a new DeviceManagementSettingAbstractImplementationConstraint and sets the default values.
func NewDeviceManagementSettingAbstractImplementationConstraint()(*DeviceManagementSettingAbstractImplementationConstraint) {
    m := &DeviceManagementSettingAbstractImplementationConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingAbstractImplementationConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingAbstractImplementationConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingAbstractImplementationConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingAbstractImplementationConstraint(), nil
}
// GetAllowedAbstractImplementationDefinitionIds gets the allowedAbstractImplementationDefinitionIds property value. List of value which means not configured for the setting
func (m *DeviceManagementSettingAbstractImplementationConstraint) GetAllowedAbstractImplementationDefinitionIds()([]string) {
    return m.allowedAbstractImplementationDefinitionIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingAbstractImplementationConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["allowedAbstractImplementationDefinitionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedAbstractImplementationDefinitionIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingAbstractImplementationConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedAbstractImplementationDefinitionIds() != nil {
        err = writer.WriteCollectionOfStringValues("allowedAbstractImplementationDefinitionIds", m.GetAllowedAbstractImplementationDefinitionIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedAbstractImplementationDefinitionIds sets the allowedAbstractImplementationDefinitionIds property value. List of value which means not configured for the setting
func (m *DeviceManagementSettingAbstractImplementationConstraint) SetAllowedAbstractImplementationDefinitionIds(value []string)() {
    m.allowedAbstractImplementationDefinitionIds = value
}
