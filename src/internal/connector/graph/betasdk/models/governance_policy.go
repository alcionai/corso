package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernancePolicy 
type GovernancePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The decisionMakerCriteria property
    decisionMakerCriteria []GovernanceCriteriaable
    // The notificationPolicy property
    notificationPolicy GovernanceNotificationPolicyable
    // The OdataType property
    odataType *string
}
// NewGovernancePolicy instantiates a new governancePolicy and sets the default values.
func NewGovernancePolicy()(*GovernancePolicy) {
    m := &GovernancePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernancePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernancePolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDecisionMakerCriteria gets the decisionMakerCriteria property value. The decisionMakerCriteria property
func (m *GovernancePolicy) GetDecisionMakerCriteria()([]GovernanceCriteriaable) {
    return m.decisionMakerCriteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decisionMakerCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceCriteriaable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceCriteriaable)
            }
            m.SetDecisionMakerCriteria(res)
        }
        return nil
    }
    res["notificationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceNotificationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationPolicy(val.(GovernanceNotificationPolicyable))
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
// GetNotificationPolicy gets the notificationPolicy property value. The notificationPolicy property
func (m *GovernancePolicy) GetNotificationPolicy()(GovernanceNotificationPolicyable) {
    return m.notificationPolicy
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernancePolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GovernancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDecisionMakerCriteria() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecisionMakerCriteria()))
        for i, v := range m.GetDecisionMakerCriteria() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("decisionMakerCriteria", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notificationPolicy", m.GetNotificationPolicy())
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
func (m *GovernancePolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDecisionMakerCriteria sets the decisionMakerCriteria property value. The decisionMakerCriteria property
func (m *GovernancePolicy) SetDecisionMakerCriteria(value []GovernanceCriteriaable)() {
    m.decisionMakerCriteria = value
}
// SetNotificationPolicy sets the notificationPolicy property value. The notificationPolicy property
func (m *GovernancePolicy) SetNotificationPolicy(value GovernanceNotificationPolicyable)() {
    m.notificationPolicy = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernancePolicy) SetOdataType(value *string)() {
    m.odataType = value
}
