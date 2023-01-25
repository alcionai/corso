package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinitionable 
type GroupPolicyDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(GroupPolicyCategoryable)
    GetCategoryPath()(*string)
    GetClassType()(*GroupPolicyDefinitionClassType)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDisplayName()(*string)
    GetExplainText()(*string)
    GetGroupPolicyCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetHasRelatedDefinitions()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMinDeviceCspVersion()(*string)
    GetMinUserCspVersion()(*string)
    GetNextVersionDefinition()(GroupPolicyDefinitionable)
    GetPolicyType()(*GroupPolicyType)
    GetPresentations()([]GroupPolicyPresentationable)
    GetPreviousVersionDefinition()(GroupPolicyDefinitionable)
    GetSupportedOn()(*string)
    GetVersion()(*string)
    SetCategory(value GroupPolicyCategoryable)()
    SetCategoryPath(value *string)()
    SetClassType(value *GroupPolicyDefinitionClassType)()
    SetDefinitionFile(value GroupPolicyDefinitionFileable)()
    SetDisplayName(value *string)()
    SetExplainText(value *string)()
    SetGroupPolicyCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetHasRelatedDefinitions(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMinDeviceCspVersion(value *string)()
    SetMinUserCspVersion(value *string)()
    SetNextVersionDefinition(value GroupPolicyDefinitionable)()
    SetPolicyType(value *GroupPolicyType)()
    SetPresentations(value []GroupPolicyPresentationable)()
    SetPreviousVersionDefinition(value GroupPolicyDefinitionable)()
    SetSupportedOn(value *string)()
    SetVersion(value *string)()
}
