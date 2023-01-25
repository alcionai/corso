package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingInstance base type for a setting instance
type DeviceManagementSettingInstance struct {
    Entity
    // The ID of the setting definition for this instance
    definitionId *string
    // JSON representation of the value
    valueJson *string
}
// NewDeviceManagementSettingInstance instantiates a new deviceManagementSettingInstance and sets the default values.
func NewDeviceManagementSettingInstance()(*DeviceManagementSettingInstance) {
    m := &DeviceManagementSettingInstance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementAbstractComplexSettingInstance":
                        return NewDeviceManagementAbstractComplexSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementBooleanSettingInstance":
                        return NewDeviceManagementBooleanSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementCollectionSettingInstance":
                        return NewDeviceManagementCollectionSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementComplexSettingInstance":
                        return NewDeviceManagementComplexSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementIntegerSettingInstance":
                        return NewDeviceManagementIntegerSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementStringSettingInstance":
                        return NewDeviceManagementStringSettingInstance(), nil
                }
            }
        }
    }
    return NewDeviceManagementSettingInstance(), nil
}
// GetDefinitionId gets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingInstance) GetDefinitionId()(*string) {
    return m.definitionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionId(val)
        }
        return nil
    }
    res["valueJson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueJson(val)
        }
        return nil
    }
    return res
}
// GetValueJson gets the valueJson property value. JSON representation of the value
func (m *DeviceManagementSettingInstance) GetValueJson()(*string) {
    return m.valueJson
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("definitionId", m.GetDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueJson", m.GetValueJson())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinitionId sets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingInstance) SetDefinitionId(value *string)() {
    m.definitionId = value
}
// SetValueJson sets the valueJson property value. JSON representation of the value
func (m *DeviceManagementSettingInstance) SetValueJson(value *string)() {
    m.valueJson = value
}
