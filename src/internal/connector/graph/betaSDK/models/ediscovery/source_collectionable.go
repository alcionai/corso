package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// SourceCollectionable 
type SourceCollectionable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalSources()([]DataSourceable)
    GetAddToReviewSetOperation()(AddToReviewSetOperationable)
    GetContentQuery()(*string)
    GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodianSources()([]DataSourceable)
    GetDataSourceScopes()(*DataSourceScopes)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable)
    GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoncustodialSources()([]NoncustodialDataSourceable)
    SetAdditionalSources(value []DataSourceable)()
    SetAddToReviewSetOperation(value AddToReviewSetOperationable)()
    SetContentQuery(value *string)()
    SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodianSources(value []DataSourceable)()
    SetDataSourceScopes(value *DataSourceScopes)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)()
    SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoncustodialSources(value []NoncustodialDataSourceable)()
}
