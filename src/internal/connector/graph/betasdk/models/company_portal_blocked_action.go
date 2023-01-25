package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyPortalBlockedAction blocked actions on the company portal as per platform and device ownership types
type CompanyPortalBlockedAction struct {
    // Action on a device that can be executed in the Company Portal
    action *CompanyPortalAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Owner type of device.
    ownerType *OwnerType
    // Supported platform types.
    platform *DevicePlatformType
}
// NewCompanyPortalBlockedAction instantiates a new companyPortalBlockedAction and sets the default values.
func NewCompanyPortalBlockedAction()(*CompanyPortalBlockedAction) {
    m := &CompanyPortalBlockedAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCompanyPortalBlockedActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyPortalBlockedActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanyPortalBlockedAction(), nil
}
// GetAction gets the action property value. Action on a device that can be executed in the Company Portal
func (m *CompanyPortalBlockedAction) GetAction()(*CompanyPortalAction) {
    return m.action
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyPortalBlockedAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompanyPortalBlockedAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCompanyPortalAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*CompanyPortalAction))
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
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CompanyPortalBlockedAction) GetOdataType()(*string) {
    return m.odataType
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *CompanyPortalBlockedAction) GetOwnerType()(*OwnerType) {
    return m.ownerType
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *CompanyPortalBlockedAction) GetPlatform()(*DevicePlatformType) {
    return m.platform
}
// Serialize serializes information the current object
func (m *CompanyPortalBlockedAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
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
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err := writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
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
// SetAction sets the action property value. Action on a device that can be executed in the Company Portal
func (m *CompanyPortalBlockedAction) SetAction(value *CompanyPortalAction)() {
    m.action = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyPortalBlockedAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CompanyPortalBlockedAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *CompanyPortalBlockedAction) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *CompanyPortalBlockedAction) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
