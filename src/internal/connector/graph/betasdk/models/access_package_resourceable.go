package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceable 
type AccessPackageResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageResourceEnvironment()(AccessPackageResourceEnvironmentable)
    GetAccessPackageResourceRoles()([]AccessPackageResourceRoleable)
    GetAccessPackageResourceScopes()([]AccessPackageResourceScopeable)
    GetAddedBy()(*string)
    GetAddedOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAttributes()([]AccessPackageResourceAttributeable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsPendingOnboarding()(*bool)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetResourceType()(*string)
    GetUrl()(*string)
    SetAccessPackageResourceEnvironment(value AccessPackageResourceEnvironmentable)()
    SetAccessPackageResourceRoles(value []AccessPackageResourceRoleable)()
    SetAccessPackageResourceScopes(value []AccessPackageResourceScopeable)()
    SetAddedBy(value *string)()
    SetAddedOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAttributes(value []AccessPackageResourceAttributeable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsPendingOnboarding(value *bool)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetResourceType(value *string)()
    SetUrl(value *string)()
}
