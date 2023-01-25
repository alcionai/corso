package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementUserRightsSetting represents a user rights setting.
type DeviceManagementUserRightsSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Representing a collection of local users or groups which will be set on device if the state of this setting is Allowed. This collection can contain a maximum of 500 elements.
    localUsersOrGroups []DeviceManagementUserRightsLocalUserOrGroupable
    // The OdataType property
    odataType *string
    // State Management Setting.
    state *StateManagementSetting
}
// NewDeviceManagementUserRightsSetting instantiates a new deviceManagementUserRightsSetting and sets the default values.
func NewDeviceManagementUserRightsSetting()(*DeviceManagementUserRightsSetting) {
    m := &DeviceManagementUserRightsSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementUserRightsSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementUserRightsSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementUserRightsSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementUserRightsSetting) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementUserRightsSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["localUsersOrGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementUserRightsLocalUserOrGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementUserRightsLocalUserOrGroupable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementUserRightsLocalUserOrGroupable)
            }
            m.SetLocalUsersOrGroups(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*StateManagementSetting))
        }
        return nil
    }
    return res
}
// GetLocalUsersOrGroups gets the localUsersOrGroups property value. Representing a collection of local users or groups which will be set on device if the state of this setting is Allowed. This collection can contain a maximum of 500 elements.
func (m *DeviceManagementUserRightsSetting) GetLocalUsersOrGroups()([]DeviceManagementUserRightsLocalUserOrGroupable) {
    return m.localUsersOrGroups
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementUserRightsSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. State Management Setting.
func (m *DeviceManagementUserRightsSetting) GetState()(*StateManagementSetting) {
    return m.state
}
// Serialize serializes information the current object
func (m *DeviceManagementUserRightsSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLocalUsersOrGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalUsersOrGroups()))
        for i, v := range m.GetLocalUsersOrGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("localUsersOrGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementUserRightsSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLocalUsersOrGroups sets the localUsersOrGroups property value. Representing a collection of local users or groups which will be set on device if the state of this setting is Allowed. This collection can contain a maximum of 500 elements.
func (m *DeviceManagementUserRightsSetting) SetLocalUsersOrGroups(value []DeviceManagementUserRightsLocalUserOrGroupable)() {
    m.localUsersOrGroups = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementUserRightsSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. State Management Setting.
func (m *DeviceManagementUserRightsSetting) SetState(value *StateManagementSetting)() {
    m.state = value
}
