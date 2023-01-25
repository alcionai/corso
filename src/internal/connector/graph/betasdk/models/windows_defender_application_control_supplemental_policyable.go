package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyable 
type WindowsDefenderApplicationControlSupplementalPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)
    GetContent()([]byte)
    GetContentFileName()(*string)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploySummary()(WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable)
    GetDescription()(*string)
    GetDeviceStatuses()([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetVersion()(*string)
    SetAssignments(value []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)()
    SetContent(value []byte)()
    SetContentFileName(value *string)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploySummary(value WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable)()
    SetDescription(value *string)()
    SetDeviceStatuses(value []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetVersion(value *string)()
}
