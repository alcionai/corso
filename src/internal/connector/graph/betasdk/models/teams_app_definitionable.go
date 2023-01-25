package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppDefinitionable 
type TeamsAppDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedInstallationScopes()(*TeamsAppInstallationScopes)
    GetAzureADAppId()(*string)
    GetBot()(TeamworkBotable)
    GetColorIcon()(TeamsAppIconable)
    GetCreatedBy()(IdentitySetable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOutlineIcon()(TeamsAppIconable)
    GetPublishingState()(*TeamsAppPublishingState)
    GetShortdescription()(*string)
    GetTeamsAppId()(*string)
    GetVersion()(*string)
    SetAllowedInstallationScopes(value *TeamsAppInstallationScopes)()
    SetAzureADAppId(value *string)()
    SetBot(value TeamworkBotable)()
    SetColorIcon(value TeamsAppIconable)()
    SetCreatedBy(value IdentitySetable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOutlineIcon(value TeamsAppIconable)()
    SetPublishingState(value *TeamsAppPublishingState)()
    SetShortdescription(value *string)()
    SetTeamsAppId(value *string)()
    SetVersion(value *string)()
}
