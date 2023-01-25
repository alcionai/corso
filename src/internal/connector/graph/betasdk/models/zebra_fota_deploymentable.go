package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeploymentable 
type ZebraFotaDeploymentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentAssignments()([]AndroidFotaDeploymentAssignmentable)
    GetDeploymentSettings()(ZebraFotaDeploymentSettingsable)
    GetDeploymentStatus()(ZebraFotaDeploymentStatusable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    SetDeploymentAssignments(value []AndroidFotaDeploymentAssignmentable)()
    SetDeploymentSettings(value ZebraFotaDeploymentSettingsable)()
    SetDeploymentStatus(value ZebraFotaDeploymentStatusable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
}
