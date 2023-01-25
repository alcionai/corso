package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleScopeTagInfo a class containing the properties of Role Scope Tag Object.
type RoleScopeTagInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Scope Tag Display name.
    displayName *string
    // The OdataType property
    odataType *string
    // Scope Tag Id.
    roleScopeTagId *string
}
// NewRoleScopeTagInfo instantiates a new roleScopeTagInfo and sets the default values.
func NewRoleScopeTagInfo()(*RoleScopeTagInfo) {
    m := &RoleScopeTagInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleScopeTagInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleScopeTagInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTagInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleScopeTagInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Scope Tag Display name.
func (m *RoleScopeTagInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTagInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["roleScopeTagId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleScopeTagId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RoleScopeTagInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetRoleScopeTagId gets the roleScopeTagId property value. Scope Tag Id.
func (m *RoleScopeTagInfo) GetRoleScopeTagId()(*string) {
    return m.roleScopeTagId
}
// Serialize serializes information the current object
func (m *RoleScopeTagInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("roleScopeTagId", m.GetRoleScopeTagId())
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
func (m *RoleScopeTagInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Scope Tag Display name.
func (m *RoleScopeTagInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RoleScopeTagInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoleScopeTagId sets the roleScopeTagId property value. Scope Tag Id.
func (m *RoleScopeTagInfo) SetRoleScopeTagId(value *string)() {
    m.roleScopeTagId = value
}
