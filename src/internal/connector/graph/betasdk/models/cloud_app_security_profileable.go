package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudAppSecurityProfileable 
type CloudAppSecurityProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentPackageUrl()(*string)
    GetDestinationServiceName()(*string)
    GetIsSigned()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManifest()(*string)
    GetName()(*string)
    GetPermissionsRequired()(*ApplicationPermissionsRequired)
    GetPlatform()(*string)
    GetPolicyName()(*string)
    GetPublisher()(*string)
    GetRiskScore()(*string)
    GetTags()([]string)
    GetType()(*string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentPackageUrl(value *string)()
    SetDestinationServiceName(value *string)()
    SetIsSigned(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManifest(value *string)()
    SetName(value *string)()
    SetPermissionsRequired(value *ApplicationPermissionsRequired)()
    SetPlatform(value *string)()
    SetPolicyName(value *string)()
    SetPublisher(value *string)()
    SetRiskScore(value *string)()
    SetTags(value []string)()
    SetType(value *string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
