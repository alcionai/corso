package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagementTemplateable 
type ManagementTemplateable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*ManagementCategory)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInformationLinks()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ActionUrlable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollections()([]ManagementTemplateCollectionable)
    GetManagementTemplateSteps()([]ManagementTemplateStepable)
    GetParameters()([]TemplateParameterable)
    GetPriority()(*int32)
    GetProvider()(*ManagementProvider)
    GetUserImpact()(*string)
    GetVersion()(*int32)
    GetWorkloadActions()([]WorkloadActionable)
    SetCategory(value *ManagementCategory)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInformationLinks(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ActionUrlable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollections(value []ManagementTemplateCollectionable)()
    SetManagementTemplateSteps(value []ManagementTemplateStepable)()
    SetParameters(value []TemplateParameterable)()
    SetPriority(value *int32)()
    SetProvider(value *ManagementProvider)()
    SetUserImpact(value *string)()
    SetVersion(value *int32)()
    SetWorkloadActions(value []WorkloadActionable)()
}
