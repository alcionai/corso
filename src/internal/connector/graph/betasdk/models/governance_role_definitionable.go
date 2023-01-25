package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleDefinitionable 
type GovernanceRoleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetResource()(GovernanceResourceable)
    GetResourceId()(*string)
    GetRoleSetting()(GovernanceRoleSettingable)
    GetTemplateId()(*string)
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetResource(value GovernanceResourceable)()
    SetResourceId(value *string)()
    SetRoleSetting(value GovernanceRoleSettingable)()
    SetTemplateId(value *string)()
}
