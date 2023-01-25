package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// AuditEventable 
type AuditEventable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityId()(*string)
    GetCategory()(*string)
    GetHttpVerb()(*string)
    GetInitiatedByAppId()(*string)
    GetInitiatedByUpn()(*string)
    GetInitiatedByUserId()(*string)
    GetIpAddress()(*string)
    GetRequestBody()(*string)
    GetRequestUrl()(*string)
    GetTenantIds()(*string)
    GetTenantNames()(*string)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityId(value *string)()
    SetCategory(value *string)()
    SetHttpVerb(value *string)()
    SetInitiatedByAppId(value *string)()
    SetInitiatedByUpn(value *string)()
    SetInitiatedByUserId(value *string)()
    SetIpAddress(value *string)()
    SetRequestBody(value *string)()
    SetRequestUrl(value *string)()
    SetTenantIds(value *string)()
    SetTenantNames(value *string)()
}
