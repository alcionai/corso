package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TenantDetailedInformationable 
type TenantDetailedInformationable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryCode()(*string)
    GetCountryName()(*string)
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    GetIndustryName()(*string)
    GetRegion()(*string)
    GetSegmentName()(*string)
    GetTenantId()(*string)
    GetVerticalName()(*string)
    SetCity(value *string)()
    SetCountryCode(value *string)()
    SetCountryName(value *string)()
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
    SetIndustryName(value *string)()
    SetRegion(value *string)()
    SetSegmentName(value *string)()
    SetTenantId(value *string)()
    SetVerticalName(value *string)()
}
