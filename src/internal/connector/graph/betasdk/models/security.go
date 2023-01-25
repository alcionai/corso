package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Security 
type Security struct {
    Entity
    // Notifications for suspicious or potential security issues in a customer’s tenant.
    alerts []Alertable
    // Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
    attackSimulation AttackSimulationRootable
    // The cloudAppSecurityProfiles property
    cloudAppSecurityProfiles []CloudAppSecurityProfileable
    // The domainSecurityProfiles property
    domainSecurityProfiles []DomainSecurityProfileable
    // The fileSecurityProfiles property
    fileSecurityProfiles []FileSecurityProfileable
    // The hostSecurityProfiles property
    hostSecurityProfiles []HostSecurityProfileable
    // The ipSecurityProfiles property
    ipSecurityProfiles []IpSecurityProfileable
    // The providerStatus property
    providerStatus []SecurityProviderStatusable
    // The providerTenantSettings property
    providerTenantSettings []ProviderTenantSettingable
    // The secureScoreControlProfiles property
    secureScoreControlProfiles []SecureScoreControlProfileable
    // Measurements of tenants’ security posture to help protect them from threats.
    secureScores []SecureScoreable
    // The securityActions property
    securityActions []SecurityActionable
    // The subjectRightsRequests property
    subjectRightsRequests []SubjectRightsRequestable
    // The tiIndicators property
    tiIndicators []TiIndicatorable
    // The userSecurityProfiles property
    userSecurityProfiles []UserSecurityProfileable
}
// NewSecurity instantiates a new Security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurity(), nil
}
// GetAlerts gets the alerts property value. Notifications for suspicious or potential security issues in a customer’s tenant.
func (m *Security) GetAlerts()([]Alertable) {
    return m.alerts
}
// GetAttackSimulation gets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) GetAttackSimulation()(AttackSimulationRootable) {
    return m.attackSimulation
}
// GetCloudAppSecurityProfiles gets the cloudAppSecurityProfiles property value. The cloudAppSecurityProfiles property
func (m *Security) GetCloudAppSecurityProfiles()([]CloudAppSecurityProfileable) {
    return m.cloudAppSecurityProfiles
}
// GetDomainSecurityProfiles gets the domainSecurityProfiles property value. The domainSecurityProfiles property
func (m *Security) GetDomainSecurityProfiles()([]DomainSecurityProfileable) {
    return m.domainSecurityProfiles
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alertable, len(val))
            for i, v := range val {
                res[i] = v.(Alertable)
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["attackSimulation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulation(val.(AttackSimulationRootable))
        }
        return nil
    }
    res["cloudAppSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudAppSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudAppSecurityProfileable, len(val))
            for i, v := range val {
                res[i] = v.(CloudAppSecurityProfileable)
            }
            m.SetCloudAppSecurityProfiles(res)
        }
        return nil
    }
    res["domainSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDomainSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DomainSecurityProfileable, len(val))
            for i, v := range val {
                res[i] = v.(DomainSecurityProfileable)
            }
            m.SetDomainSecurityProfiles(res)
        }
        return nil
    }
    res["fileSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileSecurityProfileable, len(val))
            for i, v := range val {
                res[i] = v.(FileSecurityProfileable)
            }
            m.SetFileSecurityProfiles(res)
        }
        return nil
    }
    res["hostSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHostSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HostSecurityProfileable, len(val))
            for i, v := range val {
                res[i] = v.(HostSecurityProfileable)
            }
            m.SetHostSecurityProfiles(res)
        }
        return nil
    }
    res["ipSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpSecurityProfileable, len(val))
            for i, v := range val {
                res[i] = v.(IpSecurityProfileable)
            }
            m.SetIpSecurityProfiles(res)
        }
        return nil
    }
    res["providerStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityProviderStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityProviderStatusable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityProviderStatusable)
            }
            m.SetProviderStatus(res)
        }
        return nil
    }
    res["providerTenantSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProviderTenantSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProviderTenantSettingable, len(val))
            for i, v := range val {
                res[i] = v.(ProviderTenantSettingable)
            }
            m.SetProviderTenantSettings(res)
        }
        return nil
    }
    res["secureScoreControlProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreControlProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlProfileable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreControlProfileable)
            }
            m.SetSecureScoreControlProfiles(res)
        }
        return nil
    }
    res["secureScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreable)
            }
            m.SetSecureScores(res)
        }
        return nil
    }
    res["securityActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityActionable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityActionable)
            }
            m.SetSecurityActions(res)
        }
        return nil
    }
    res["subjectRightsRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectRightsRequestable)
            }
            m.SetSubjectRightsRequests(res)
        }
        return nil
    }
    res["tiIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTiIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TiIndicatorable, len(val))
            for i, v := range val {
                res[i] = v.(TiIndicatorable)
            }
            m.SetTiIndicators(res)
        }
        return nil
    }
    res["userSecurityProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSecurityProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSecurityProfileable, len(val))
            for i, v := range val {
                res[i] = v.(UserSecurityProfileable)
            }
            m.SetUserSecurityProfiles(res)
        }
        return nil
    }
    return res
}
// GetFileSecurityProfiles gets the fileSecurityProfiles property value. The fileSecurityProfiles property
func (m *Security) GetFileSecurityProfiles()([]FileSecurityProfileable) {
    return m.fileSecurityProfiles
}
// GetHostSecurityProfiles gets the hostSecurityProfiles property value. The hostSecurityProfiles property
func (m *Security) GetHostSecurityProfiles()([]HostSecurityProfileable) {
    return m.hostSecurityProfiles
}
// GetIpSecurityProfiles gets the ipSecurityProfiles property value. The ipSecurityProfiles property
func (m *Security) GetIpSecurityProfiles()([]IpSecurityProfileable) {
    return m.ipSecurityProfiles
}
// GetProviderStatus gets the providerStatus property value. The providerStatus property
func (m *Security) GetProviderStatus()([]SecurityProviderStatusable) {
    return m.providerStatus
}
// GetProviderTenantSettings gets the providerTenantSettings property value. The providerTenantSettings property
func (m *Security) GetProviderTenantSettings()([]ProviderTenantSettingable) {
    return m.providerTenantSettings
}
// GetSecureScoreControlProfiles gets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfileable) {
    return m.secureScoreControlProfiles
}
// GetSecureScores gets the secureScores property value. Measurements of tenants’ security posture to help protect them from threats.
func (m *Security) GetSecureScores()([]SecureScoreable) {
    return m.secureScores
}
// GetSecurityActions gets the securityActions property value. The securityActions property
func (m *Security) GetSecurityActions()([]SecurityActionable) {
    return m.securityActions
}
// GetSubjectRightsRequests gets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Security) GetSubjectRightsRequests()([]SubjectRightsRequestable) {
    return m.subjectRightsRequests
}
// GetTiIndicators gets the tiIndicators property value. The tiIndicators property
func (m *Security) GetTiIndicators()([]TiIndicatorable) {
    return m.tiIndicators
}
// GetUserSecurityProfiles gets the userSecurityProfiles property value. The userSecurityProfiles property
func (m *Security) GetUserSecurityProfiles()([]UserSecurityProfileable) {
    return m.userSecurityProfiles
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("attackSimulation", m.GetAttackSimulation())
        if err != nil {
            return err
        }
    }
    if m.GetCloudAppSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudAppSecurityProfiles()))
        for i, v := range m.GetCloudAppSecurityProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudAppSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainSecurityProfiles()))
        for i, v := range m.GetDomainSecurityProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("domainSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFileSecurityProfiles()))
        for i, v := range m.GetFileSecurityProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fileSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHostSecurityProfiles()))
        for i, v := range m.GetHostSecurityProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("hostSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpSecurityProfiles()))
        for i, v := range m.GetIpSecurityProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("ipSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProviderStatus()))
        for i, v := range m.GetProviderStatus() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("providerStatus", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProviderTenantSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProviderTenantSettings()))
        for i, v := range m.GetProviderTenantSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("providerTenantSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScoreControlProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecureScoreControlProfiles()))
        for i, v := range m.GetSecureScoreControlProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecureScores()))
        for i, v := range m.GetSecureScores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecurityActions()))
        for i, v := range m.GetSecurityActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("securityActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectRightsRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjectRightsRequests()))
        for i, v := range m.GetSubjectRightsRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("subjectRightsRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTiIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTiIndicators()))
        for i, v := range m.GetTiIndicators() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tiIndicators", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSecurityProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSecurityProfiles()))
        for i, v := range m.GetUserSecurityProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userSecurityProfiles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. Notifications for suspicious or potential security issues in a customer’s tenant.
func (m *Security) SetAlerts(value []Alertable)() {
    m.alerts = value
}
// SetAttackSimulation sets the attackSimulation property value. Provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *Security) SetAttackSimulation(value AttackSimulationRootable)() {
    m.attackSimulation = value
}
// SetCloudAppSecurityProfiles sets the cloudAppSecurityProfiles property value. The cloudAppSecurityProfiles property
func (m *Security) SetCloudAppSecurityProfiles(value []CloudAppSecurityProfileable)() {
    m.cloudAppSecurityProfiles = value
}
// SetDomainSecurityProfiles sets the domainSecurityProfiles property value. The domainSecurityProfiles property
func (m *Security) SetDomainSecurityProfiles(value []DomainSecurityProfileable)() {
    m.domainSecurityProfiles = value
}
// SetFileSecurityProfiles sets the fileSecurityProfiles property value. The fileSecurityProfiles property
func (m *Security) SetFileSecurityProfiles(value []FileSecurityProfileable)() {
    m.fileSecurityProfiles = value
}
// SetHostSecurityProfiles sets the hostSecurityProfiles property value. The hostSecurityProfiles property
func (m *Security) SetHostSecurityProfiles(value []HostSecurityProfileable)() {
    m.hostSecurityProfiles = value
}
// SetIpSecurityProfiles sets the ipSecurityProfiles property value. The ipSecurityProfiles property
func (m *Security) SetIpSecurityProfiles(value []IpSecurityProfileable)() {
    m.ipSecurityProfiles = value
}
// SetProviderStatus sets the providerStatus property value. The providerStatus property
func (m *Security) SetProviderStatus(value []SecurityProviderStatusable)() {
    m.providerStatus = value
}
// SetProviderTenantSettings sets the providerTenantSettings property value. The providerTenantSettings property
func (m *Security) SetProviderTenantSettings(value []ProviderTenantSettingable)() {
    m.providerTenantSettings = value
}
// SetSecureScoreControlProfiles sets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)() {
    m.secureScoreControlProfiles = value
}
// SetSecureScores sets the secureScores property value. Measurements of tenants’ security posture to help protect them from threats.
func (m *Security) SetSecureScores(value []SecureScoreable)() {
    m.secureScores = value
}
// SetSecurityActions sets the securityActions property value. The securityActions property
func (m *Security) SetSecurityActions(value []SecurityActionable)() {
    m.securityActions = value
}
// SetSubjectRightsRequests sets the subjectRightsRequests property value. The subjectRightsRequests property
func (m *Security) SetSubjectRightsRequests(value []SubjectRightsRequestable)() {
    m.subjectRightsRequests = value
}
// SetTiIndicators sets the tiIndicators property value. The tiIndicators property
func (m *Security) SetTiIndicators(value []TiIndicatorable)() {
    m.tiIndicators = value
}
// SetUserSecurityProfiles sets the userSecurityProfiles property value. The userSecurityProfiles property
func (m *Security) SetUserSecurityProfiles(value []UserSecurityProfileable)() {
    m.userSecurityProfiles = value
}
