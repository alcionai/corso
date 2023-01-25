package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingDependency dependency information for a setting
type DeviceManagementSettingDependency struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of constraints for the dependency setting value
    constraints []DeviceManagementConstraintable
    // The setting definition ID of the setting depended on
    definitionId *string
    // The OdataType property
    odataType *string
}
// NewDeviceManagementSettingDependency instantiates a new deviceManagementSettingDependency and sets the default values.
func NewDeviceManagementSettingDependency()(*DeviceManagementSettingDependency) {
    m := &DeviceManagementSettingDependency{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingDependencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingDependencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingDependency(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingDependency) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConstraints gets the constraints property value. Collection of constraints for the dependency setting value
func (m *DeviceManagementSettingDependency) GetConstraints()([]DeviceManagementConstraintable) {
    return m.constraints
}
// GetDefinitionId gets the definitionId property value. The setting definition ID of the setting depended on
func (m *DeviceManagementSettingDependency) GetDefinitionId()(*string) {
    return m.definitionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingDependency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["constraints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConstraintable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConstraintable)
            }
            m.SetConstraints(res)
        }
        return nil
    }
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingDependency) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingDependency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConstraints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("constraints", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("definitionId", m.GetDefinitionId())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingDependency) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConstraints sets the constraints property value. Collection of constraints for the dependency setting value
func (m *DeviceManagementSettingDependency) SetConstraints(value []DeviceManagementConstraintable)() {
    m.constraints = value
}
// SetDefinitionId sets the definitionId property value. The setting definition ID of the setting depended on
func (m *DeviceManagementSettingDependency) SetDefinitionId(value *string)() {
    m.definitionId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingDependency) SetOdataType(value *string)() {
    m.odataType = value
}
