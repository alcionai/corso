package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationIdentityDomain 
type EducationIdentityDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appliesTo property
    appliesTo *EducationUserRole
    // Represents the domain for the user account.
    name *string
    // The OdataType property
    odataType *string
}
// NewEducationIdentityDomain instantiates a new educationIdentityDomain and sets the default values.
func NewEducationIdentityDomain()(*EducationIdentityDomain) {
    m := &EducationIdentityDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationIdentityDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationIdentityDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationIdentityDomain(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationIdentityDomain) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppliesTo gets the appliesTo property value. The appliesTo property
func (m *EducationIdentityDomain) GetAppliesTo()(*EducationUserRole) {
    return m.appliesTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationIdentityDomain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetName gets the name property value. Represents the domain for the user account.
func (m *EducationIdentityDomain) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationIdentityDomain) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EducationIdentityDomain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err := writer.WriteStringValue("appliesTo", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationIdentityDomain) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppliesTo sets the appliesTo property value. The appliesTo property
func (m *EducationIdentityDomain) SetAppliesTo(value *EducationUserRole)() {
    m.appliesTo = value
}
// SetName sets the name property value. Represents the domain for the user account.
func (m *EducationIdentityDomain) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationIdentityDomain) SetOdataType(value *string)() {
    m.odataType = value
}
