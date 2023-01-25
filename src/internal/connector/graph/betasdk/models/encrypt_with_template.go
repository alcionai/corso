package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptWithTemplate 
type EncryptWithTemplate struct {
    EncryptContent
    // The availableForEncryption property
    availableForEncryption *bool
    // The templateId property
    templateId *string
}
// NewEncryptWithTemplate instantiates a new EncryptWithTemplate and sets the default values.
func NewEncryptWithTemplate()(*EncryptWithTemplate) {
    m := &EncryptWithTemplate{
        EncryptContent: *NewEncryptContent(),
    }
    odataTypeValue := "#microsoft.graph.encryptWithTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEncryptWithTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEncryptWithTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptWithTemplate(), nil
}
// GetAvailableForEncryption gets the availableForEncryption property value. The availableForEncryption property
func (m *EncryptWithTemplate) GetAvailableForEncryption()(*bool) {
    return m.availableForEncryption
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptWithTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EncryptContent.GetFieldDeserializers()
    res["availableForEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableForEncryption(val)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. The templateId property
func (m *EncryptWithTemplate) GetTemplateId()(*string) {
    return m.templateId
}
// Serialize serializes information the current object
func (m *EncryptWithTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EncryptContent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("availableForEncryption", m.GetAvailableForEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailableForEncryption sets the availableForEncryption property value. The availableForEncryption property
func (m *EncryptWithTemplate) SetAvailableForEncryption(value *bool)() {
    m.availableForEncryption = value
}
// SetTemplateId sets the templateId property value. The templateId property
func (m *EncryptWithTemplate) SetTemplateId(value *string)() {
    m.templateId = value
}
