package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Incidentable 
type Incidentable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlerts()([]Alertable)
    GetAssignedTo()(*string)
    GetClassification()(*AlertClassification)
    GetComments()([]AlertCommentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomTags()([]string)
    GetDetermination()(*AlertDetermination)
    GetDisplayName()(*string)
    GetIncidentWebUrl()(*string)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRedirectIncidentId()(*string)
    GetSeverity()(*AlertSeverity)
    GetStatus()(*IncidentStatus)
    GetTenantId()(*string)
    SetAlerts(value []Alertable)()
    SetAssignedTo(value *string)()
    SetClassification(value *AlertClassification)()
    SetComments(value []AlertCommentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomTags(value []string)()
    SetDetermination(value *AlertDetermination)()
    SetDisplayName(value *string)()
    SetIncidentWebUrl(value *string)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRedirectIncidentId(value *string)()
    SetSeverity(value *AlertSeverity)()
    SetStatus(value *IncidentStatus)()
    SetTenantId(value *string)()
}
