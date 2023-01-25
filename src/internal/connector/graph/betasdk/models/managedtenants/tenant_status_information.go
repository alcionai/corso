package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantStatusInformation 
type TenantStatusInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
    delegatedPrivilegeStatus *DelegatedPrivilegeStatus
    // The date and time the delegated admin privileges status was updated. Optional. Read-only.
    lastDelegatedPrivilegeRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The identifier for the account that offboarded the managed tenant. Optional. Read-only.
    offboardedByUserId *string
    // The date and time when the managed tenant was offboarded. Optional. Read-only.
    offboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identifier for the account that onboarded the managed tenant. Optional. Read-only.
    onboardedByUserId *string
    // The date and time when the managed tenant was onboarded. Optional. Read-only.
    onboardedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
    onboardingStatus *TenantOnboardingStatus
    // Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
    tenantOnboardingEligibilityReason *TenantOnboardingEligibilityReason
    // The collection of workload statues for the managed tenant. Optional. Read-only.
    workloadStatuses []WorkloadStatusable
}
// NewTenantStatusInformation instantiates a new tenantStatusInformation and sets the default values.
func NewTenantStatusInformation()(*TenantStatusInformation) {
    m := &TenantStatusInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantStatusInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantStatusInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantStatusInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDelegatedPrivilegeStatus gets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
func (m *TenantStatusInformation) GetDelegatedPrivilegeStatus()(*DelegatedPrivilegeStatus) {
    return m.delegatedPrivilegeStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantStatusInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["delegatedPrivilegeStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedPrivilegeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegatedPrivilegeStatus(val.(*DelegatedPrivilegeStatus))
        }
        return nil
    }
    res["lastDelegatedPrivilegeRefreshDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDelegatedPrivilegeRefreshDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["offboardedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedByUserId(val)
        }
        return nil
    }
    res["offboardedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffboardedDateTime(val)
        }
        return nil
    }
    res["onboardedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedByUserId(val)
        }
        return nil
    }
    res["onboardedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardedDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*TenantOnboardingStatus))
        }
        return nil
    }
    res["tenantOnboardingEligibilityReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantOnboardingEligibilityReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantOnboardingEligibilityReason(val.(*TenantOnboardingEligibilityReason))
        }
        return nil
    }
    res["workloadStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkloadStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadStatusable, len(val))
            for i, v := range val {
                res[i] = v.(WorkloadStatusable)
            }
            m.SetWorkloadStatuses(res)
        }
        return nil
    }
    return res
}
// GetLastDelegatedPrivilegeRefreshDateTime gets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastDelegatedPrivilegeRefreshDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantStatusInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetOffboardedByUserId gets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedByUserId()(*string) {
    return m.offboardedByUserId
}
// GetOffboardedDateTime gets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.offboardedDateTime
}
// GetOnboardedByUserId gets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedByUserId()(*string) {
    return m.onboardedByUserId
}
// GetOnboardedDateTime gets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onboardedDateTime
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetOnboardingStatus()(*TenantOnboardingStatus) {
    return m.onboardingStatus
}
// GetTenantOnboardingEligibilityReason gets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) GetTenantOnboardingEligibilityReason()(*TenantOnboardingEligibilityReason) {
    return m.tenantOnboardingEligibilityReason
}
// GetWorkloadStatuses gets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) GetWorkloadStatuses()([]WorkloadStatusable) {
    return m.workloadStatuses
}
// Serialize serializes information the current object
func (m *TenantStatusInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDelegatedPrivilegeStatus() != nil {
        cast := (*m.GetDelegatedPrivilegeStatus()).String()
        err := writer.WriteStringValue("delegatedPrivilegeStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastDelegatedPrivilegeRefreshDateTime", m.GetLastDelegatedPrivilegeRefreshDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("offboardedByUserId", m.GetOffboardedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("offboardedDateTime", m.GetOffboardedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onboardedByUserId", m.GetOnboardedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("onboardedDateTime", m.GetOnboardedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err := writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantOnboardingEligibilityReason() != nil {
        cast := (*m.GetTenantOnboardingEligibilityReason()).String()
        err := writer.WriteStringValue("tenantOnboardingEligibilityReason", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkloadStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkloadStatuses()))
        for i, v := range m.GetWorkloadStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("workloadStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantStatusInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDelegatedPrivilegeStatus sets the delegatedPrivilegeStatus property value. The status of the delegated admin privilege relationship between the managing entity and the managed tenant. Possible values are: none, delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges, delegatedAndGranularDelegetedAdminPrivileges. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: granularDelegatedAdminPrivileges , delegatedAndGranularDelegetedAdminPrivileges. Optional. Read-only.
func (m *TenantStatusInformation) SetDelegatedPrivilegeStatus(value *DelegatedPrivilegeStatus)() {
    m.delegatedPrivilegeStatus = value
}
// SetLastDelegatedPrivilegeRefreshDateTime sets the lastDelegatedPrivilegeRefreshDateTime property value. The date and time the delegated admin privileges status was updated. Optional. Read-only.
func (m *TenantStatusInformation) SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDelegatedPrivilegeRefreshDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantStatusInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffboardedByUserId sets the offboardedByUserId property value. The identifier for the account that offboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedByUserId(value *string)() {
    m.offboardedByUserId = value
}
// SetOffboardedDateTime sets the offboardedDateTime property value. The date and time when the managed tenant was offboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.offboardedDateTime = value
}
// SetOnboardedByUserId sets the onboardedByUserId property value. The identifier for the account that onboarded the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedByUserId(value *string)() {
    m.onboardedByUserId = value
}
// SetOnboardedDateTime sets the onboardedDateTime property value. The date and time when the managed tenant was onboarded. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onboardedDateTime = value
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboarding status for the managed tenant.. Possible values are: ineligible, inProcess, active, inactive, unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetOnboardingStatus(value *TenantOnboardingStatus)() {
    m.onboardingStatus = value
}
// SetTenantOnboardingEligibilityReason sets the tenantOnboardingEligibilityReason property value. Organization's onboarding eligibility reason in Microsoft 365 Lighthouse.. Possible values are: none, contractType, delegatedAdminPrivileges,usersCount,license and unknownFutureValue. Optional. Read-only.
func (m *TenantStatusInformation) SetTenantOnboardingEligibilityReason(value *TenantOnboardingEligibilityReason)() {
    m.tenantOnboardingEligibilityReason = value
}
// SetWorkloadStatuses sets the workloadStatuses property value. The collection of workload statues for the managed tenant. Optional. Read-only.
func (m *TenantStatusInformation) SetWorkloadStatuses(value []WorkloadStatusable)() {
    m.workloadStatuses = value
}
