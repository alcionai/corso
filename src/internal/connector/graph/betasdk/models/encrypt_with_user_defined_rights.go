package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptWithUserDefinedRights 
type EncryptWithUserDefinedRights struct {
    EncryptContent
    // The allowAdHocPermissions property
    allowAdHocPermissions *bool
    // The allowMailForwarding property
    allowMailForwarding *bool
    // The decryptionRightsManagementTemplateId property
    decryptionRightsManagementTemplateId *string
}
// NewEncryptWithUserDefinedRights instantiates a new EncryptWithUserDefinedRights and sets the default values.
func NewEncryptWithUserDefinedRights()(*EncryptWithUserDefinedRights) {
    m := &EncryptWithUserDefinedRights{
        EncryptContent: *NewEncryptContent(),
    }
    odataTypeValue := "#microsoft.graph.encryptWithUserDefinedRights";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEncryptWithUserDefinedRightsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEncryptWithUserDefinedRightsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptWithUserDefinedRights(), nil
}
// GetAllowAdHocPermissions gets the allowAdHocPermissions property value. The allowAdHocPermissions property
func (m *EncryptWithUserDefinedRights) GetAllowAdHocPermissions()(*bool) {
    return m.allowAdHocPermissions
}
// GetAllowMailForwarding gets the allowMailForwarding property value. The allowMailForwarding property
func (m *EncryptWithUserDefinedRights) GetAllowMailForwarding()(*bool) {
    return m.allowMailForwarding
}
// GetDecryptionRightsManagementTemplateId gets the decryptionRightsManagementTemplateId property value. The decryptionRightsManagementTemplateId property
func (m *EncryptWithUserDefinedRights) GetDecryptionRightsManagementTemplateId()(*string) {
    return m.decryptionRightsManagementTemplateId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptWithUserDefinedRights) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EncryptContent.GetFieldDeserializers()
    res["allowAdHocPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAdHocPermissions(val)
        }
        return nil
    }
    res["allowMailForwarding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMailForwarding(val)
        }
        return nil
    }
    res["decryptionRightsManagementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecryptionRightsManagementTemplateId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EncryptWithUserDefinedRights) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EncryptContent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAdHocPermissions", m.GetAllowAdHocPermissions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowMailForwarding", m.GetAllowMailForwarding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("decryptionRightsManagementTemplateId", m.GetDecryptionRightsManagementTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAdHocPermissions sets the allowAdHocPermissions property value. The allowAdHocPermissions property
func (m *EncryptWithUserDefinedRights) SetAllowAdHocPermissions(value *bool)() {
    m.allowAdHocPermissions = value
}
// SetAllowMailForwarding sets the allowMailForwarding property value. The allowMailForwarding property
func (m *EncryptWithUserDefinedRights) SetAllowMailForwarding(value *bool)() {
    m.allowMailForwarding = value
}
// SetDecryptionRightsManagementTemplateId sets the decryptionRightsManagementTemplateId property value. The decryptionRightsManagementTemplateId property
func (m *EncryptWithUserDefinedRights) SetDecryptionRightsManagementTemplateId(value *string)() {
    m.decryptionRightsManagementTemplateId = value
}
