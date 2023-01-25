package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFeatureUpdateProfileable 
type WindowsFeatureUpdateProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]WindowsFeatureUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeployableContentDisplayName()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFeatureUpdateVersion()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetRolloutSettings()(WindowsUpdateRolloutSettingsable)
    SetAssignments(value []WindowsFeatureUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeployableContentDisplayName(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFeatureUpdateVersion(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetRolloutSettings(value WindowsUpdateRolloutSettingsable)()
}
