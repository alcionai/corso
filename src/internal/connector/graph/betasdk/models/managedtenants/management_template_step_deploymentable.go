package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagementTemplateStepDeploymentable 
type ManagementTemplateStepDeploymentable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetError()(GraphAPIErrorDetailsable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*ManagementTemplateDeploymentStatus)
    GetTemplateStepVersion()(ManagementTemplateStepVersionable)
    GetTenantId()(*string)
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetError(value GraphAPIErrorDetailsable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *ManagementTemplateDeploymentStatus)()
    SetTemplateStepVersion(value ManagementTemplateStepVersionable)()
    SetTenantId(value *string)()
}
