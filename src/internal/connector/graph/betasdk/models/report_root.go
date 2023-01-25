package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportRoot 
type ReportRoot struct {
    Entity
    // Represents a detailed summary of an application sign-in.
    applicationSignInDetailedSummary []ApplicationSignInDetailedSummaryable
    // Container for navigation properties for Azure AD authentication methods resources.
    authenticationMethods AuthenticationMethodsRootable
    // Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
    credentialUserRegistrationDetails []CredentialUserRegistrationDetailsable
    // The dailyPrintUsage property
    dailyPrintUsage []PrintUsageable
    // The dailyPrintUsageByPrinter property
    dailyPrintUsageByPrinter []PrintUsageByPrinterable
    // The dailyPrintUsageByUser property
    dailyPrintUsageByUser []PrintUsageByUserable
    // The dailyPrintUsageSummariesByPrinter property
    dailyPrintUsageSummariesByPrinter []PrintUsageByPrinterable
    // The dailyPrintUsageSummariesByUser property
    dailyPrintUsageSummariesByUser []PrintUsageByUserable
    // The monthlyPrintUsageByPrinter property
    monthlyPrintUsageByPrinter []PrintUsageByPrinterable
    // The monthlyPrintUsageByUser property
    monthlyPrintUsageByUser []PrintUsageByUserable
    // The monthlyPrintUsageSummariesByPrinter property
    monthlyPrintUsageSummariesByPrinter []PrintUsageByPrinterable
    // The monthlyPrintUsageSummariesByUser property
    monthlyPrintUsageSummariesByUser []PrintUsageByUserable
    // Provides the ability to launch a realistic simulated phishing attack that organizations can learn from.
    security SecurityReportsRootable
    // Represents the self-service password reset (SSPR) usage for a given tenant.
    userCredentialUsageDetails []UserCredentialUsageDetailsable
}
// NewReportRoot instantiates a new ReportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateReportRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportRoot(), nil
}
// GetApplicationSignInDetailedSummary gets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummaryable) {
    return m.applicationSignInDetailedSummary
}
// GetAuthenticationMethods gets the authenticationMethods property value. Container for navigation properties for Azure AD authentication methods resources.
func (m *ReportRoot) GetAuthenticationMethods()(AuthenticationMethodsRootable) {
    return m.authenticationMethods
}
// GetCredentialUserRegistrationDetails gets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetailsable) {
    return m.credentialUserRegistrationDetails
}
// GetDailyPrintUsage gets the dailyPrintUsage property value. The dailyPrintUsage property
func (m *ReportRoot) GetDailyPrintUsage()([]PrintUsageable) {
    return m.dailyPrintUsage
}
// GetDailyPrintUsageByPrinter gets the dailyPrintUsageByPrinter property value. The dailyPrintUsageByPrinter property
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    return m.dailyPrintUsageByPrinter
}
// GetDailyPrintUsageByUser gets the dailyPrintUsageByUser property value. The dailyPrintUsageByUser property
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUserable) {
    return m.dailyPrintUsageByUser
}
// GetDailyPrintUsageSummariesByPrinter gets the dailyPrintUsageSummariesByPrinter property value. The dailyPrintUsageSummariesByPrinter property
func (m *ReportRoot) GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable) {
    return m.dailyPrintUsageSummariesByPrinter
}
// GetDailyPrintUsageSummariesByUser gets the dailyPrintUsageSummariesByUser property value. The dailyPrintUsageSummariesByUser property
func (m *ReportRoot) GetDailyPrintUsageSummariesByUser()([]PrintUsageByUserable) {
    return m.dailyPrintUsageSummariesByUser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationSignInDetailedSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicationSignInDetailedSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicationSignInDetailedSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ApplicationSignInDetailedSummaryable)
            }
            m.SetApplicationSignInDetailedSummary(res)
        }
        return nil
    }
    res["authenticationMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethods(val.(AuthenticationMethodsRootable))
        }
        return nil
    }
    res["credentialUserRegistrationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCredentialUserRegistrationDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CredentialUserRegistrationDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(CredentialUserRegistrationDetailsable)
            }
            m.SetCredentialUserRegistrationDetails(res)
        }
        return nil
    }
    res["dailyPrintUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageable)
            }
            m.SetDailyPrintUsage(res)
        }
        return nil
    }
    res["dailyPrintUsageByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByPrinterable)
            }
            m.SetDailyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByUserable)
            }
            m.SetDailyPrintUsageByUser(res)
        }
        return nil
    }
    res["dailyPrintUsageSummariesByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByPrinterable)
            }
            m.SetDailyPrintUsageSummariesByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageSummariesByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByUserable)
            }
            m.SetDailyPrintUsageSummariesByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByPrinterable)
            }
            m.SetMonthlyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByUserable)
            }
            m.SetMonthlyPrintUsageByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageSummariesByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByPrinterable)
            }
            m.SetMonthlyPrintUsageSummariesByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageSummariesByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                res[i] = v.(PrintUsageByUserable)
            }
            m.SetMonthlyPrintUsageSummariesByUser(res)
        }
        return nil
    }
    res["security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityReportsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurity(val.(SecurityReportsRootable))
        }
        return nil
    }
    res["userCredentialUsageDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserCredentialUsageDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserCredentialUsageDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(UserCredentialUsageDetailsable)
            }
            m.SetUserCredentialUsageDetails(res)
        }
        return nil
    }
    return res
}
// GetMonthlyPrintUsageByPrinter gets the monthlyPrintUsageByPrinter property value. The monthlyPrintUsageByPrinter property
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    return m.monthlyPrintUsageByPrinter
}
// GetMonthlyPrintUsageByUser gets the monthlyPrintUsageByUser property value. The monthlyPrintUsageByUser property
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUserable) {
    return m.monthlyPrintUsageByUser
}
// GetMonthlyPrintUsageSummariesByPrinter gets the monthlyPrintUsageSummariesByPrinter property value. The monthlyPrintUsageSummariesByPrinter property
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable) {
    return m.monthlyPrintUsageSummariesByPrinter
}
// GetMonthlyPrintUsageSummariesByUser gets the monthlyPrintUsageSummariesByUser property value. The monthlyPrintUsageSummariesByUser property
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUserable) {
    return m.monthlyPrintUsageSummariesByUser
}
// GetSecurity gets the security property value. Provides the ability to launch a realistic simulated phishing attack that organizations can learn from.
func (m *ReportRoot) GetSecurity()(SecurityReportsRootable) {
    return m.security
}
// GetUserCredentialUsageDetails gets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) GetUserCredentialUsageDetails()([]UserCredentialUsageDetailsable) {
    return m.userCredentialUsageDetails
}
// Serialize serializes information the current object
func (m *ReportRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicationSignInDetailedSummary() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApplicationSignInDetailedSummary()))
        for i, v := range m.GetApplicationSignInDetailedSummary() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("applicationSignInDetailedSummary", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationMethods", m.GetAuthenticationMethods())
        if err != nil {
            return err
        }
    }
    if m.GetCredentialUserRegistrationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentialUserRegistrationDetails()))
        for i, v := range m.GetCredentialUserRegistrationDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsage() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsage()))
        for i, v := range m.GetDailyPrintUsage() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsage", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageByPrinter()))
        for i, v := range m.GetDailyPrintUsageByPrinter() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageByUser()))
        for i, v := range m.GetDailyPrintUsageByUser() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageSummariesByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageSummariesByPrinter()))
        for i, v := range m.GetDailyPrintUsageSummariesByPrinter() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageSummariesByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageSummariesByUser()))
        for i, v := range m.GetDailyPrintUsageSummariesByUser() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageByPrinter() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageByUser()))
        for i, v := range m.GetMonthlyPrintUsageByUser() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageSummariesByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageSummariesByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageSummariesByPrinter() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageSummariesByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageSummariesByUser()))
        for i, v := range m.GetMonthlyPrintUsageSummariesByUser() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("security", m.GetSecurity())
        if err != nil {
            return err
        }
    }
    if m.GetUserCredentialUsageDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserCredentialUsageDetails()))
        for i, v := range m.GetUserCredentialUsageDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userCredentialUsageDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationSignInDetailedSummary sets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummaryable)() {
    m.applicationSignInDetailedSummary = value
}
// SetAuthenticationMethods sets the authenticationMethods property value. Container for navigation properties for Azure AD authentication methods resources.
func (m *ReportRoot) SetAuthenticationMethods(value AuthenticationMethodsRootable)() {
    m.authenticationMethods = value
}
// SetCredentialUserRegistrationDetails sets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetailsable)() {
    m.credentialUserRegistrationDetails = value
}
// SetDailyPrintUsage sets the dailyPrintUsage property value. The dailyPrintUsage property
func (m *ReportRoot) SetDailyPrintUsage(value []PrintUsageable)() {
    m.dailyPrintUsage = value
}
// SetDailyPrintUsageByPrinter sets the dailyPrintUsageByPrinter property value. The dailyPrintUsageByPrinter property
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    m.dailyPrintUsageByPrinter = value
}
// SetDailyPrintUsageByUser sets the dailyPrintUsageByUser property value. The dailyPrintUsageByUser property
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUserable)() {
    m.dailyPrintUsageByUser = value
}
// SetDailyPrintUsageSummariesByPrinter sets the dailyPrintUsageSummariesByPrinter property value. The dailyPrintUsageSummariesByPrinter property
func (m *ReportRoot) SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)() {
    m.dailyPrintUsageSummariesByPrinter = value
}
// SetDailyPrintUsageSummariesByUser sets the dailyPrintUsageSummariesByUser property value. The dailyPrintUsageSummariesByUser property
func (m *ReportRoot) SetDailyPrintUsageSummariesByUser(value []PrintUsageByUserable)() {
    m.dailyPrintUsageSummariesByUser = value
}
// SetMonthlyPrintUsageByPrinter sets the monthlyPrintUsageByPrinter property value. The monthlyPrintUsageByPrinter property
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    m.monthlyPrintUsageByPrinter = value
}
// SetMonthlyPrintUsageByUser sets the monthlyPrintUsageByUser property value. The monthlyPrintUsageByUser property
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)() {
    m.monthlyPrintUsageByUser = value
}
// SetMonthlyPrintUsageSummariesByPrinter sets the monthlyPrintUsageSummariesByPrinter property value. The monthlyPrintUsageSummariesByPrinter property
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)() {
    m.monthlyPrintUsageSummariesByPrinter = value
}
// SetMonthlyPrintUsageSummariesByUser sets the monthlyPrintUsageSummariesByUser property value. The monthlyPrintUsageSummariesByUser property
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUserable)() {
    m.monthlyPrintUsageSummariesByUser = value
}
// SetSecurity sets the security property value. Provides the ability to launch a realistic simulated phishing attack that organizations can learn from.
func (m *ReportRoot) SetSecurity(value SecurityReportsRootable)() {
    m.security = value
}
// SetUserCredentialUsageDetails sets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) SetUserCredentialUsageDetails(value []UserCredentialUsageDetailsable)() {
    m.userCredentialUsageDetails = value
}
