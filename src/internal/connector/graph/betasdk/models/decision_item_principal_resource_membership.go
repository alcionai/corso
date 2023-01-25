package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DecisionItemPrincipalResourceMembership 
type DecisionItemPrincipalResourceMembership struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The membershipType property
    membershipType *DecisionItemPrincipalResourceMembershipType
    // The OdataType property
    odataType *string
}
// NewDecisionItemPrincipalResourceMembership instantiates a new decisionItemPrincipalResourceMembership and sets the default values.
func NewDecisionItemPrincipalResourceMembership()(*DecisionItemPrincipalResourceMembership) {
    m := &DecisionItemPrincipalResourceMembership{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDecisionItemPrincipalResourceMembershipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDecisionItemPrincipalResourceMembershipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDecisionItemPrincipalResourceMembership(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DecisionItemPrincipalResourceMembership) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DecisionItemPrincipalResourceMembership) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["membershipType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDecisionItemPrincipalResourceMembershipType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipType(val.(*DecisionItemPrincipalResourceMembershipType))
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
// GetMembershipType gets the membershipType property value. The membershipType property
func (m *DecisionItemPrincipalResourceMembership) GetMembershipType()(*DecisionItemPrincipalResourceMembershipType) {
    return m.membershipType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DecisionItemPrincipalResourceMembership) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DecisionItemPrincipalResourceMembership) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMembershipType() != nil {
        cast := (*m.GetMembershipType()).String()
        err := writer.WriteStringValue("membershipType", &cast)
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
func (m *DecisionItemPrincipalResourceMembership) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMembershipType sets the membershipType property value. The membershipType property
func (m *DecisionItemPrincipalResourceMembership) SetMembershipType(value *DecisionItemPrincipalResourceMembershipType)() {
    m.membershipType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DecisionItemPrincipalResourceMembership) SetOdataType(value *string)() {
    m.odataType = value
}
