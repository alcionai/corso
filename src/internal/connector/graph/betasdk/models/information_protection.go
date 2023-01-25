package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtection 
type InformationProtection struct {
    Entity
    // The bitlocker property
    bitlocker Bitlockerable
    // The dataLossPreventionPolicies property
    dataLossPreventionPolicies []DataLossPreventionPolicyable
    // The policy property
    policy InformationProtectionPolicyable
    // The sensitivityLabels property
    sensitivityLabels []SensitivityLabelable
    // The sensitivityPolicySettings property
    sensitivityPolicySettings SensitivityPolicySettingsable
    // The threatAssessmentRequests property
    threatAssessmentRequests []ThreatAssessmentRequestable
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetBitlocker gets the bitlocker property value. The bitlocker property
func (m *InformationProtection) GetBitlocker()(Bitlockerable) {
    return m.bitlocker
}
// GetDataLossPreventionPolicies gets the dataLossPreventionPolicies property value. The dataLossPreventionPolicies property
func (m *InformationProtection) GetDataLossPreventionPolicies()([]DataLossPreventionPolicyable) {
    return m.dataLossPreventionPolicies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bitlocker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitlockerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitlocker(val.(Bitlockerable))
        }
        return nil
    }
    res["dataLossPreventionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataLossPreventionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataLossPreventionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(DataLossPreventionPolicyable)
            }
            m.SetDataLossPreventionPolicies(res)
        }
        return nil
    }
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(InformationProtectionPolicyable))
        }
        return nil
    }
    res["sensitivityLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                res[i] = v.(SensitivityLabelable)
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    res["sensitivityPolicySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensitivityPolicySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityPolicySettings(val.(SensitivityPolicySettingsable))
        }
        return nil
    }
    res["threatAssessmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThreatAssessmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(ThreatAssessmentRequestable)
            }
            m.SetThreatAssessmentRequests(res)
        }
        return nil
    }
    return res
}
// GetPolicy gets the policy property value. The policy property
func (m *InformationProtection) GetPolicy()(InformationProtectionPolicyable) {
    return m.policy
}
// GetSensitivityLabels gets the sensitivityLabels property value. The sensitivityLabels property
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabelable) {
    return m.sensitivityLabels
}
// GetSensitivityPolicySettings gets the sensitivityPolicySettings property value. The sensitivityPolicySettings property
func (m *InformationProtection) GetSensitivityPolicySettings()(SensitivityPolicySettingsable) {
    return m.sensitivityPolicySettings
}
// GetThreatAssessmentRequests gets the threatAssessmentRequests property value. The threatAssessmentRequests property
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequestable) {
    return m.threatAssessmentRequests
}
// Serialize serializes information the current object
func (m *InformationProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("bitlocker", m.GetBitlocker())
        if err != nil {
            return err
        }
    }
    if m.GetDataLossPreventionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataLossPreventionPolicies()))
        for i, v := range m.GetDataLossPreventionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dataLossPreventionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sensitivityPolicySettings", m.GetSensitivityPolicySettings())
        if err != nil {
            return err
        }
    }
    if m.GetThreatAssessmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBitlocker sets the bitlocker property value. The bitlocker property
func (m *InformationProtection) SetBitlocker(value Bitlockerable)() {
    m.bitlocker = value
}
// SetDataLossPreventionPolicies sets the dataLossPreventionPolicies property value. The dataLossPreventionPolicies property
func (m *InformationProtection) SetDataLossPreventionPolicies(value []DataLossPreventionPolicyable)() {
    m.dataLossPreventionPolicies = value
}
// SetPolicy sets the policy property value. The policy property
func (m *InformationProtection) SetPolicy(value InformationProtectionPolicyable)() {
    m.policy = value
}
// SetSensitivityLabels sets the sensitivityLabels property value. The sensitivityLabels property
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabelable)() {
    m.sensitivityLabels = value
}
// SetSensitivityPolicySettings sets the sensitivityPolicySettings property value. The sensitivityPolicySettings property
func (m *InformationProtection) SetSensitivityPolicySettings(value SensitivityPolicySettingsable)() {
    m.sensitivityPolicySettings = value
}
// SetThreatAssessmentRequests sets the threatAssessmentRequests property value. The threatAssessmentRequests property
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)() {
    m.threatAssessmentRequests = value
}
