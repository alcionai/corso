package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationServiceable 
type DataClassificationServiceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassifyFileJobs()([]JobResponseBaseable)
    GetClassifyTextJobs()([]JobResponseBaseable)
    GetEvaluateDlpPoliciesJobs()([]JobResponseBaseable)
    GetEvaluateLabelJobs()([]JobResponseBaseable)
    GetExactMatchDataStores()([]ExactMatchDataStoreable)
    GetExactMatchUploadAgents()([]ExactMatchUploadAgentable)
    GetJobs()([]JobResponseBaseable)
    GetSensitiveTypes()([]SensitiveTypeable)
    GetSensitivityLabels()([]SensitivityLabelable)
    SetClassifyFileJobs(value []JobResponseBaseable)()
    SetClassifyTextJobs(value []JobResponseBaseable)()
    SetEvaluateDlpPoliciesJobs(value []JobResponseBaseable)()
    SetEvaluateLabelJobs(value []JobResponseBaseable)()
    SetExactMatchDataStores(value []ExactMatchDataStoreable)()
    SetExactMatchUploadAgents(value []ExactMatchUploadAgentable)()
    SetJobs(value []JobResponseBaseable)()
    SetSensitiveTypes(value []SensitiveTypeable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
}
