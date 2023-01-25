package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Taskable 
type Taskable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArguments()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.KeyValuePairable)
    GetCategory()(*LifecycleTaskCategory)
    GetContinueOnError()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExecutionSequence()(*int32)
    GetIsEnabled()(*bool)
    GetTaskDefinitionId()(*string)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    SetArguments(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.KeyValuePairable)()
    SetCategory(value *LifecycleTaskCategory)()
    SetContinueOnError(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExecutionSequence(value *int32)()
    SetIsEnabled(value *bool)()
    SetTaskDefinitionId(value *string)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
}
