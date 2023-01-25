package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable 
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeployedDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    SetDeployedDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
}
