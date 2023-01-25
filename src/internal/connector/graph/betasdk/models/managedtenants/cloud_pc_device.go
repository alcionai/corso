package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CloudPcDevice provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcDevice struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed. Required. Read-only.
    cloudPcStatus *string
    // The specification of the cloud PC device. Required. Read-only.
    deviceSpecification *string
    // The display name  of the cloud PC device. Required. Read-only.
    displayName *string
    // Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managed device identifier of the cloud PC device. Optional. Read-only.
    managedDeviceId *string
    // The managed device display name of the cloud PC device. Optional. Read-only.
    managedDeviceName *string
    // The provisioning policy identifier for the cloud PC device. Required. Read-only.
    provisioningPolicyId *string
    // The service plan name of the cloud PC device. Required. Read-only.
    servicePlanName *string
    // The service plan type of the cloud PC device. Required. Read-only.
    servicePlanType *string
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string
    // The user principal name (UPN) of the user assigned to the cloud PC device. Required. Read-only.
    userPrincipalName *string
}
// NewCloudPcDevice instantiates a new cloudPcDevice and sets the default values.
func NewCloudPcDevice()(*CloudPcDevice) {
    m := &CloudPcDevice{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateCloudPcDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDevice(), nil
}
// GetCloudPcStatus gets the cloudPcStatus property value. The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed. Required. Read-only.
func (m *CloudPcDevice) GetCloudPcStatus()(*string) {
    return m.cloudPcStatus
}
// GetDeviceSpecification gets the deviceSpecification property value. The specification of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) GetDeviceSpecification()(*string) {
    return m.deviceSpecification
}
// GetDisplayName gets the displayName property value. The display name  of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcStatus(val)
        }
        return nil
    }
    res["deviceSpecification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceSpecification(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["provisioningPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    res["servicePlanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanType(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
func (m *CloudPcDevice) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. The managed device identifier of the cloud PC device. Optional. Read-only.
func (m *CloudPcDevice) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManagedDeviceName gets the managedDeviceName property value. The managed device display name of the cloud PC device. Optional. Read-only.
func (m *CloudPcDevice) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetProvisioningPolicyId gets the provisioningPolicyId property value. The provisioning policy identifier for the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) GetProvisioningPolicyId()(*string) {
    return m.provisioningPolicyId
}
// GetServicePlanName gets the servicePlanName property value. The service plan name of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) GetServicePlanName()(*string) {
    return m.servicePlanName
}
// GetServicePlanType gets the servicePlanType property value. The service plan type of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) GetServicePlanType()(*string) {
    return m.servicePlanType
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CloudPcDevice) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CloudPcDevice) GetTenantId()(*string) {
    return m.tenantId
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *CloudPcDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cloudPcStatus", m.GetCloudPcStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceSpecification", m.GetDeviceSpecification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisioningPolicyId", m.GetProvisioningPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanType", m.GetServicePlanType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudPcStatus sets the cloudPcStatus property value. The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed. Required. Read-only.
func (m *CloudPcDevice) SetCloudPcStatus(value *string)() {
    m.cloudPcStatus = value
}
// SetDeviceSpecification sets the deviceSpecification property value. The specification of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) SetDeviceSpecification(value *string)() {
    m.deviceSpecification = value
}
// SetDisplayName sets the displayName property value. The display name  of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
func (m *CloudPcDevice) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The managed device identifier of the cloud PC device. Optional. Read-only.
func (m *CloudPcDevice) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. The managed device display name of the cloud PC device. Optional. Read-only.
func (m *CloudPcDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetProvisioningPolicyId sets the provisioningPolicyId property value. The provisioning policy identifier for the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) SetProvisioningPolicyId(value *string)() {
    m.provisioningPolicyId = value
}
// SetServicePlanName sets the servicePlanName property value. The service plan name of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
// SetServicePlanType sets the servicePlanType property value. The service plan type of the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) SetServicePlanType(value *string)() {
    m.servicePlanType = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CloudPcDevice) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CloudPcDevice) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the cloud PC device. Required. Read-only.
func (m *CloudPcDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
