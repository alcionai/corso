package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionable 
type InformationProtectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBitlocker()(Bitlockerable)
    GetDataLossPreventionPolicies()([]DataLossPreventionPolicyable)
    GetPolicy()(InformationProtectionPolicyable)
    GetSensitivityLabels()([]SensitivityLabelable)
    GetSensitivityPolicySettings()(SensitivityPolicySettingsable)
    GetThreatAssessmentRequests()([]ThreatAssessmentRequestable)
    SetBitlocker(value Bitlockerable)()
    SetDataLossPreventionPolicies(value []DataLossPreventionPolicyable)()
    SetPolicy(value InformationProtectionPolicyable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
    SetSensitivityPolicySettings(value SensitivityPolicySettingsable)()
    SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)()
}
