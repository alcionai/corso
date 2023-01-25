package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Securityable 
type Securityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlerts()([]Alertable)
    GetAttackSimulation()(AttackSimulationRootable)
    GetCloudAppSecurityProfiles()([]CloudAppSecurityProfileable)
    GetDomainSecurityProfiles()([]DomainSecurityProfileable)
    GetFileSecurityProfiles()([]FileSecurityProfileable)
    GetHostSecurityProfiles()([]HostSecurityProfileable)
    GetIpSecurityProfiles()([]IpSecurityProfileable)
    GetProviderStatus()([]SecurityProviderStatusable)
    GetProviderTenantSettings()([]ProviderTenantSettingable)
    GetSecureScoreControlProfiles()([]SecureScoreControlProfileable)
    GetSecureScores()([]SecureScoreable)
    GetSecurityActions()([]SecurityActionable)
    GetSubjectRightsRequests()([]SubjectRightsRequestable)
    GetTiIndicators()([]TiIndicatorable)
    GetUserSecurityProfiles()([]UserSecurityProfileable)
    SetAlerts(value []Alertable)()
    SetAttackSimulation(value AttackSimulationRootable)()
    SetCloudAppSecurityProfiles(value []CloudAppSecurityProfileable)()
    SetDomainSecurityProfiles(value []DomainSecurityProfileable)()
    SetFileSecurityProfiles(value []FileSecurityProfileable)()
    SetHostSecurityProfiles(value []HostSecurityProfileable)()
    SetIpSecurityProfiles(value []IpSecurityProfileable)()
    SetProviderStatus(value []SecurityProviderStatusable)()
    SetProviderTenantSettings(value []ProviderTenantSettingable)()
    SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)()
    SetSecureScores(value []SecureScoreable)()
    SetSecurityActions(value []SecurityActionable)()
    SetSubjectRightsRequests(value []SubjectRightsRequestable)()
    SetTiIndicators(value []TiIndicatorable)()
    SetUserSecurityProfiles(value []UserSecurityProfileable)()
}
