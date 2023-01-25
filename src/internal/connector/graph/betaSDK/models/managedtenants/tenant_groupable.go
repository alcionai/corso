package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TenantGroupable 
type TenantGroupable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllTenantsIncluded()(*bool)
    GetDisplayName()(*string)
    GetManagementActions()([]ManagementActionInfoable)
    GetManagementIntents()([]ManagementIntentInfoable)
    GetTenantIds()([]string)
    SetAllTenantsIncluded(value *bool)()
    SetDisplayName(value *string)()
    SetManagementActions(value []ManagementActionInfoable)()
    SetManagementIntents(value []ManagementIntentInfoable)()
    SetTenantIds(value []string)()
}
