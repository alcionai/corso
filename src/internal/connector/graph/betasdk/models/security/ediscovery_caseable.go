package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// EdiscoveryCaseable 
type EdiscoveryCaseable interface {
    Case_escapedable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClosedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodians()([]EdiscoveryCustodianable)
    GetExternalId()(*string)
    GetLegalHolds()([]EdiscoveryHoldPolicyable)
    GetNoncustodialDataSources()([]EdiscoveryNoncustodialDataSourceable)
    GetOperations()([]CaseOperationable)
    GetReviewSets()([]EdiscoveryReviewSetable)
    GetSearches()([]EdiscoverySearchable)
    GetSettings()(EdiscoveryCaseSettingsable)
    GetTags()([]EdiscoveryReviewTagable)
    SetClosedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodians(value []EdiscoveryCustodianable)()
    SetExternalId(value *string)()
    SetLegalHolds(value []EdiscoveryHoldPolicyable)()
    SetNoncustodialDataSources(value []EdiscoveryNoncustodialDataSourceable)()
    SetOperations(value []CaseOperationable)()
    SetReviewSets(value []EdiscoveryReviewSetable)()
    SetSearches(value []EdiscoverySearchable)()
    SetSettings(value EdiscoveryCaseSettingsable)()
    SetTags(value []EdiscoveryReviewTagable)()
}
