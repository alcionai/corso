package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalCreationConditionSetable 
type ServicePrincipalCreationConditionSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationIds()([]string)
    GetApplicationPublisherIds()([]string)
    GetApplicationsFromVerifiedPublisherOnly()(*bool)
    GetApplicationTenantIds()([]string)
    GetCertifiedApplicationsOnly()(*bool)
    SetApplicationIds(value []string)()
    SetApplicationPublisherIds(value []string)()
    SetApplicationsFromVerifiedPublisherOnly(value *bool)()
    SetApplicationTenantIds(value []string)()
    SetCertifiedApplicationsOnly(value *bool)()
}
