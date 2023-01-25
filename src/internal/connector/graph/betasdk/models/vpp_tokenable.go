package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppTokenable 
type VppTokenable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleId()(*string)
    GetAutomaticallyUpdateApps()(*bool)
    GetClaimTokenManagementFromExternalMdm()(*bool)
    GetCountryOrRegion()(*string)
    GetDataSharingConsentGranted()(*bool)
    GetDisplayName()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncStatus()(*VppTokenSyncStatus)
    GetLocationName()(*string)
    GetOrganizationName()(*string)
    GetRoleScopeTagIds()([]string)
    GetState()(*VppTokenState)
    GetToken()(*string)
    GetTokenActionResults()([]VppTokenActionResultable)
    GetVppTokenAccountType()(*VppTokenAccountType)
    SetAppleId(value *string)()
    SetAutomaticallyUpdateApps(value *bool)()
    SetClaimTokenManagementFromExternalMdm(value *bool)()
    SetCountryOrRegion(value *string)()
    SetDataSharingConsentGranted(value *bool)()
    SetDisplayName(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncStatus(value *VppTokenSyncStatus)()
    SetLocationName(value *string)()
    SetOrganizationName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetState(value *VppTokenState)()
    SetToken(value *string)()
    SetTokenActionResults(value []VppTokenActionResultable)()
    SetVppTokenAccountType(value *VppTokenAccountType)()
}
