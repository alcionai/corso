package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyCategoryable 
type GroupPolicyCategoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]GroupPolicyCategoryable)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDefinitions()([]GroupPolicyDefinitionable)
    GetDisplayName()(*string)
    GetIngestionSource()(*IngestionSource)
    GetIsRoot()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParent()(GroupPolicyCategoryable)
    SetChildren(value []GroupPolicyCategoryable)()
    SetDefinitionFile(value GroupPolicyDefinitionFileable)()
    SetDefinitions(value []GroupPolicyDefinitionable)()
    SetDisplayName(value *string)()
    SetIngestionSource(value *IngestionSource)()
    SetIsRoot(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParent(value GroupPolicyCategoryable)()
}
