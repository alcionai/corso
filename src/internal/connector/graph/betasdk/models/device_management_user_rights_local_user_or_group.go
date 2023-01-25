package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementUserRightsLocalUserOrGroup represents information for a local user or group used for user rights setting.
type DeviceManagementUserRightsLocalUserOrGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Admin’s description of this local user or group.
    description *string
    // The name of this local user or group.
    name *string
    // The OdataType property
    odataType *string
    // The security identifier of this local user or group (e.g. S-1-5-32-544).
    securityIdentifier *string
}
// NewDeviceManagementUserRightsLocalUserOrGroup instantiates a new deviceManagementUserRightsLocalUserOrGroup and sets the default values.
func NewDeviceManagementUserRightsLocalUserOrGroup()(*DeviceManagementUserRightsLocalUserOrGroup) {
    m := &DeviceManagementUserRightsLocalUserOrGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementUserRightsLocalUserOrGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementUserRightsLocalUserOrGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementUserRightsLocalUserOrGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementUserRightsLocalUserOrGroup) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. Admin’s description of this local user or group.
func (m *DeviceManagementUserRightsLocalUserOrGroup) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementUserRightsLocalUserOrGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["securityIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityIdentifier(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of this local user or group.
func (m *DeviceManagementUserRightsLocalUserOrGroup) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementUserRightsLocalUserOrGroup) GetOdataType()(*string) {
    return m.odataType
}
// GetSecurityIdentifier gets the securityIdentifier property value. The security identifier of this local user or group (e.g. S-1-5-32-544).
func (m *DeviceManagementUserRightsLocalUserOrGroup) GetSecurityIdentifier()(*string) {
    return m.securityIdentifier
}
// Serialize serializes information the current object
func (m *DeviceManagementUserRightsLocalUserOrGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("securityIdentifier", m.GetSecurityIdentifier())
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
func (m *DeviceManagementUserRightsLocalUserOrGroup) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. Admin’s description of this local user or group.
func (m *DeviceManagementUserRightsLocalUserOrGroup) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of this local user or group.
func (m *DeviceManagementUserRightsLocalUserOrGroup) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementUserRightsLocalUserOrGroup) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSecurityIdentifier sets the securityIdentifier property value. The security identifier of this local user or group (e.g. S-1-5-32-544).
func (m *DeviceManagementUserRightsLocalUserOrGroup) SetSecurityIdentifier(value *string)() {
    m.securityIdentifier = value
}
