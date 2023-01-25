package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileSecurityProfileable 
type FileSecurityProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityGroupNames()([]string)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCertificateThumbprint()(*string)
    GetExtensions()([]string)
    GetFileType()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHashes()([]FileHashable)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareStates()([]MalwareStateable)
    GetNames()([]string)
    GetRiskScore()(*string)
    GetSize()(*int64)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    GetVulnerabilityStates()([]VulnerabilityStateable)
    SetActivityGroupNames(value []string)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCertificateThumbprint(value *string)()
    SetExtensions(value []string)()
    SetFileType(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHashes(value []FileHashable)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareStates(value []MalwareStateable)()
    SetNames(value []string)()
    SetRiskScore(value *string)()
    SetSize(value *int64)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
    SetVulnerabilityStates(value []VulnerabilityStateable)()
}
