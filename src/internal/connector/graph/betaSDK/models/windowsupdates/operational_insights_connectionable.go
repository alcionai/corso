package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationalInsightsConnectionable 
type OperationalInsightsConnectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceConnectionable
    GetAzureResourceGroupName()(*string)
    GetAzureSubscriptionId()(*string)
    GetWorkspaceName()(*string)
    SetAzureResourceGroupName(value *string)()
    SetAzureSubscriptionId(value *string)()
    SetWorkspaceName(value *string)()
}
