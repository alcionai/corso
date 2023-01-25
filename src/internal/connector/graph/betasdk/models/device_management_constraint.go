package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConstraint base entity for a constraint
type DeviceManagementConstraint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
}
// NewDeviceManagementConstraint instantiates a new deviceManagementConstraint and sets the default values.
func NewDeviceManagementConstraint()(*DeviceManagementConstraint) {
    m := &DeviceManagementConstraint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementEnumConstraint":
                        return NewDeviceManagementEnumConstraint(), nil
                    case "#microsoft.graph.deviceManagementIntentSettingSecretConstraint":
                        return NewDeviceManagementIntentSettingSecretConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingAbstractImplementationConstraint":
                        return NewDeviceManagementSettingAbstractImplementationConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingAppConstraint":
                        return NewDeviceManagementSettingAppConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingBooleanConstraint":
                        return NewDeviceManagementSettingBooleanConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingCollectionConstraint":
                        return NewDeviceManagementSettingCollectionConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingEnrollmentTypeConstraint":
                        return NewDeviceManagementSettingEnrollmentTypeConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingFileConstraint":
                        return NewDeviceManagementSettingFileConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingIntegerConstraint":
                        return NewDeviceManagementSettingIntegerConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingProfileConstraint":
                        return NewDeviceManagementSettingProfileConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingRegexConstraint":
                        return NewDeviceManagementSettingRegexConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingRequiredConstraint":
                        return NewDeviceManagementSettingRequiredConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingSddlConstraint":
                        return NewDeviceManagementSettingSddlConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingStringLengthConstraint":
                        return NewDeviceManagementSettingStringLengthConstraint(), nil
                    case "#microsoft.graph.deviceManagementSettingXmlConstraint":
                        return NewDeviceManagementSettingXmlConstraint(), nil
                }
            }
        }
    }
    return NewDeviceManagementConstraint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConstraint) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *DeviceManagementConstraint) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceManagementConstraint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConstraint) SetOdataType(value *string)() {
    m.odataType = value
}
