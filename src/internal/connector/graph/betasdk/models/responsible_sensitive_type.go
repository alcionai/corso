package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResponsibleSensitiveType 
type ResponsibleSensitiveType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description property
    description *string
    // The id property
    id *string
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The publisherName property
    publisherName *string
    // The rulePackageId property
    rulePackageId *string
    // The rulePackageType property
    rulePackageType *string
}
// NewResponsibleSensitiveType instantiates a new responsibleSensitiveType and sets the default values.
func NewResponsibleSensitiveType()(*ResponsibleSensitiveType) {
    m := &ResponsibleSensitiveType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateResponsibleSensitiveTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResponsibleSensitiveTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResponsibleSensitiveType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResponsibleSensitiveType) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
func (m *ResponsibleSensitiveType) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResponsibleSensitiveType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["rulePackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulePackageId(val)
        }
        return nil
    }
    res["rulePackageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulePackageType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *ResponsibleSensitiveType) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name property
func (m *ResponsibleSensitiveType) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ResponsibleSensitiveType) GetOdataType()(*string) {
    return m.odataType
}
// GetPublisherName gets the publisherName property value. The publisherName property
func (m *ResponsibleSensitiveType) GetPublisherName()(*string) {
    return m.publisherName
}
// GetRulePackageId gets the rulePackageId property value. The rulePackageId property
func (m *ResponsibleSensitiveType) GetRulePackageId()(*string) {
    return m.rulePackageId
}
// GetRulePackageType gets the rulePackageType property value. The rulePackageType property
func (m *ResponsibleSensitiveType) GetRulePackageType()(*string) {
    return m.rulePackageType
}
// Serialize serializes information the current object
func (m *ResponsibleSensitiveType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rulePackageId", m.GetRulePackageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rulePackageType", m.GetRulePackageType())
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
func (m *ResponsibleSensitiveType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *ResponsibleSensitiveType) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The id property
func (m *ResponsibleSensitiveType) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *ResponsibleSensitiveType) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResponsibleSensitiveType) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPublisherName sets the publisherName property value. The publisherName property
func (m *ResponsibleSensitiveType) SetPublisherName(value *string)() {
    m.publisherName = value
}
// SetRulePackageId sets the rulePackageId property value. The rulePackageId property
func (m *ResponsibleSensitiveType) SetRulePackageId(value *string)() {
    m.rulePackageId = value
}
// SetRulePackageType sets the rulePackageType property value. The rulePackageType property
func (m *ResponsibleSensitiveType) SetRulePackageType(value *string)() {
    m.rulePackageType = value
}
