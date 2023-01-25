package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CallRecordable 
type CallRecordable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJoinWebUrl()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModalities()([]Modality)
    GetOrganizer()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)
    GetParticipants()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)
    GetSessions()([]Sessionable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetType()(*CallType)
    GetVersion()(*int64)
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJoinWebUrl(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModalities(value []Modality)()
    SetOrganizer(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)()
    SetParticipants(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)()
    SetSessions(value []Sessionable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetType(value *CallType)()
    SetVersion(value *int64)()
}
