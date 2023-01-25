package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceNotificationPolicy 
type GovernanceNotificationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enabledTemplateTypes property
    enabledTemplateTypes []string
    // The notificationTemplates property
    notificationTemplates []GovernanceNotificationTemplateable
    // The OdataType property
    odataType *string
}
// NewGovernanceNotificationPolicy instantiates a new governanceNotificationPolicy and sets the default values.
func NewGovernanceNotificationPolicy()(*GovernanceNotificationPolicy) {
    m := &GovernanceNotificationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceNotificationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceNotificationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceNotificationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnabledTemplateTypes gets the enabledTemplateTypes property value. The enabledTemplateTypes property
func (m *GovernanceNotificationPolicy) GetEnabledTemplateTypes()([]string) {
    return m.enabledTemplateTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceNotificationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabledTemplateTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnabledTemplateTypes(res)
        }
        return nil
    }
    res["notificationTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceNotificationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceNotificationTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceNotificationTemplateable)
            }
            m.SetNotificationTemplates(res)
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
// GetNotificationTemplates gets the notificationTemplates property value. The notificationTemplates property
func (m *GovernanceNotificationPolicy) GetNotificationTemplates()([]GovernanceNotificationTemplateable) {
    return m.notificationTemplates
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernanceNotificationPolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GovernanceNotificationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnabledTemplateTypes() != nil {
        err := writer.WriteCollectionOfStringValues("enabledTemplateTypes", m.GetEnabledTemplateTypes())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationTemplates()))
        for i, v := range m.GetNotificationTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("notificationTemplates", cast)
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
func (m *GovernanceNotificationPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnabledTemplateTypes sets the enabledTemplateTypes property value. The enabledTemplateTypes property
func (m *GovernanceNotificationPolicy) SetEnabledTemplateTypes(value []string)() {
    m.enabledTemplateTypes = value
}
// SetNotificationTemplates sets the notificationTemplates property value. The notificationTemplates property
func (m *GovernanceNotificationPolicy) SetNotificationTemplates(value []GovernanceNotificationTemplateable)() {
    m.notificationTemplates = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernanceNotificationPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
