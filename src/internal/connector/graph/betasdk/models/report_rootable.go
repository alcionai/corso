package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportRootable 
type ReportRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummaryable)
    GetAuthenticationMethods()(AuthenticationMethodsRootable)
    GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetailsable)
    GetDailyPrintUsage()([]PrintUsageable)
    GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageByUser()([]PrintUsageByUserable)
    GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageSummariesByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUserable)
    GetSecurity()(SecurityReportsRootable)
    GetUserCredentialUsageDetails()([]UserCredentialUsageDetailsable)
    SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummaryable)()
    SetAuthenticationMethods(value AuthenticationMethodsRootable)()
    SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetailsable)()
    SetDailyPrintUsage(value []PrintUsageable)()
    SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageByUser(value []PrintUsageByUserable)()
    SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageSummariesByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUserable)()
    SetSecurity(value SecurityReportsRootable)()
    SetUserCredentialUsageDetails(value []UserCredentialUsageDetailsable)()
}
