package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceScopeable 
type AccessPackageResourceScopeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageResource()(AccessPackageResourceable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsRootScope()(*bool)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetRoleOriginId()(*string)
    GetUrl()(*string)
    SetAccessPackageResource(value AccessPackageResourceable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsRootScope(value *bool)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetRoleOriginId(value *string)()
    SetUrl(value *string)()
}
