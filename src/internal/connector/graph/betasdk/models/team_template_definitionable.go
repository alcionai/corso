package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplateDefinitionable 
type TeamTemplateDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudience()(*TeamTemplateAudience)
    GetCategories()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIconUrl()(*string)
    GetLanguageTag()(*string)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParentTemplateId()(*string)
    GetPublisherName()(*string)
    GetShortDescription()(*string)
    GetTeamDefinition()(Teamable)
    SetAudience(value *TeamTemplateAudience)()
    SetCategories(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIconUrl(value *string)()
    SetLanguageTag(value *string)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParentTemplateId(value *string)()
    SetPublisherName(value *string)()
    SetShortDescription(value *string)()
    SetTeamDefinition(value Teamable)()
}
