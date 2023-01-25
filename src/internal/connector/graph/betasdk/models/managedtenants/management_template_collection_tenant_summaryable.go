package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagementTemplateCollectionTenantSummaryable 
type ManagementTemplateCollectionTenantSummaryable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompleteStepsCount()(*int32)
    GetCompleteUsersCount()(*int32)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedStepsCount()(*int32)
    GetExcludedUsersCount()(*int32)
    GetExcludedUsersDistinctCount()(*int32)
    GetIncompleteStepsCount()(*int32)
    GetIncompleteUsersCount()(*int32)
    GetIneligibleStepsCount()(*int32)
    GetIsComplete()(*bool)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollectionDisplayName()(*string)
    GetManagementTemplateCollectionId()(*string)
    GetTenantId()(*string)
    SetCompleteStepsCount(value *int32)()
    SetCompleteUsersCount(value *int32)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedStepsCount(value *int32)()
    SetExcludedUsersCount(value *int32)()
    SetExcludedUsersDistinctCount(value *int32)()
    SetIncompleteStepsCount(value *int32)()
    SetIncompleteUsersCount(value *int32)()
    SetIneligibleStepsCount(value *int32)()
    SetIsComplete(value *bool)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollectionDisplayName(value *string)()
    SetManagementTemplateCollectionId(value *string)()
    SetTenantId(value *string)()
}
