package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceNotificationTemplate 
type GovernanceNotificationTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The culture property
    culture *string
    // The id property
    id *string
    // The OdataType property
    odataType *string
    // The source property
    source *string
    // The type property
    type_escaped *string
    // The version property
    version *string
}
// NewGovernanceNotificationTemplate instantiates a new governanceNotificationTemplate and sets the default values.
func NewGovernanceNotificationTemplate()(*GovernanceNotificationTemplate) {
    m := &GovernanceNotificationTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceNotificationTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceNotificationTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceNotificationTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCulture gets the culture property value. The culture property
func (m *GovernanceNotificationTemplate) GetCulture()(*string) {
    return m.culture
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceNotificationTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["culture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCulture(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *GovernanceNotificationTemplate) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernanceNotificationTemplate) GetOdataType()(*string) {
    return m.odataType
}
// GetSource gets the source property value. The source property
func (m *GovernanceNotificationTemplate) GetSource()(*string) {
    return m.source
}
// GetType gets the type property value. The type property
func (m *GovernanceNotificationTemplate) GetType()(*string) {
    return m.type_escaped
}
// GetVersion gets the version property value. The version property
func (m *GovernanceNotificationTemplate) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *GovernanceNotificationTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("culture", m.GetCulture())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *GovernanceNotificationTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCulture sets the culture property value. The culture property
func (m *GovernanceNotificationTemplate) SetCulture(value *string)() {
    m.culture = value
}
// SetId sets the id property value. The id property
func (m *GovernanceNotificationTemplate) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernanceNotificationTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSource sets the source property value. The source property
func (m *GovernanceNotificationTemplate) SetSource(value *string)() {
    m.source = value
}
// SetType sets the type property value. The type property
func (m *GovernanceNotificationTemplate) SetType(value *string)() {
    m.type_escaped = value
}
// SetVersion sets the version property value. The version property
func (m *GovernanceNotificationTemplate) SetVersion(value *string)() {
    m.version = value
}
