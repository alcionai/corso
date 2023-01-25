package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelServerable 
type MicrosoftTunnelServerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgentImageDigest()(*string)
    GetDisplayName()(*string)
    GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetServerImageDigest()(*string)
    GetTunnelServerHealthStatus()(*MicrosoftTunnelServerHealthStatus)
    SetAgentImageDigest(value *string)()
    SetDisplayName(value *string)()
    SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetServerImageDigest(value *string)()
    SetTunnelServerHealthStatus(value *MicrosoftTunnelServerHealthStatus)()
}
