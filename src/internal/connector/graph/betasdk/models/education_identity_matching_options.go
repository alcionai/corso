package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationIdentityMatchingOptions 
type EducationIdentityMatchingOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appliesTo property
    appliesTo *EducationUserRole
    // The OdataType property
    odataType *string
    // The name of the source property, which should be a field name in the source data. This property is case-sensitive.
    sourcePropertyName *string
    // The domain to suffix with the source property to match on the target. If provided as null, the source property will be used to match with the target property.
    targetDomain *string
    // The name of the target property, which should be a valid property in Azure AD. This property is case-sensitive.
    targetPropertyName *string
}
// NewEducationIdentityMatchingOptions instantiates a new educationIdentityMatchingOptions and sets the default values.
func NewEducationIdentityMatchingOptions()(*EducationIdentityMatchingOptions) {
    m := &EducationIdentityMatchingOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationIdentityMatchingOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationIdentityMatchingOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationIdentityMatchingOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationIdentityMatchingOptions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppliesTo gets the appliesTo property value. The appliesTo property
func (m *EducationIdentityMatchingOptions) GetAppliesTo()(*EducationUserRole) {
    return m.appliesTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationIdentityMatchingOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val.(*EducationUserRole))
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
    res["sourcePropertyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourcePropertyName(val)
        }
        return nil
    }
    res["targetDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDomain(val)
        }
        return nil
    }
    res["targetPropertyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPropertyName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationIdentityMatchingOptions) GetOdataType()(*string) {
    return m.odataType
}
// GetSourcePropertyName gets the sourcePropertyName property value. The name of the source property, which should be a field name in the source data. This property is case-sensitive.
func (m *EducationIdentityMatchingOptions) GetSourcePropertyName()(*string) {
    return m.sourcePropertyName
}
// GetTargetDomain gets the targetDomain property value. The domain to suffix with the source property to match on the target. If provided as null, the source property will be used to match with the target property.
func (m *EducationIdentityMatchingOptions) GetTargetDomain()(*string) {
    return m.targetDomain
}
// GetTargetPropertyName gets the targetPropertyName property value. The name of the target property, which should be a valid property in Azure AD. This property is case-sensitive.
func (m *EducationIdentityMatchingOptions) GetTargetPropertyName()(*string) {
    return m.targetPropertyName
}
// Serialize serializes information the current object
func (m *EducationIdentityMatchingOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err := writer.WriteStringValue("appliesTo", &cast)
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
        err := writer.WriteStringValue("sourcePropertyName", m.GetSourcePropertyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDomain", m.GetTargetDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetPropertyName", m.GetTargetPropertyName())
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
func (m *EducationIdentityMatchingOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppliesTo sets the appliesTo property value. The appliesTo property
func (m *EducationIdentityMatchingOptions) SetAppliesTo(value *EducationUserRole)() {
    m.appliesTo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationIdentityMatchingOptions) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourcePropertyName sets the sourcePropertyName property value. The name of the source property, which should be a field name in the source data. This property is case-sensitive.
func (m *EducationIdentityMatchingOptions) SetSourcePropertyName(value *string)() {
    m.sourcePropertyName = value
}
// SetTargetDomain sets the targetDomain property value. The domain to suffix with the source property to match on the target. If provided as null, the source property will be used to match with the target property.
func (m *EducationIdentityMatchingOptions) SetTargetDomain(value *string)() {
    m.targetDomain = value
}
// SetTargetPropertyName sets the targetPropertyName property value. The name of the target property, which should be a valid property in Azure AD. This property is case-sensitive.
func (m *EducationIdentityMatchingOptions) SetTargetPropertyName(value *string)() {
    m.targetPropertyName = value
}
