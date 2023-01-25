package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainSecurityProfile provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DomainSecurityProfile struct {
    Entity
    // The activityGroupNames property
    activityGroupNames []string
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The azureTenantId property
    azureTenantId *string
    // The countHits property
    countHits *int32
    // The countInOrg property
    countInOrg *int32
    // The domainCategories property
    domainCategories []ReputationCategoryable
    // The domainRegisteredDateTime property
    domainRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The firstSeenDateTime property
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastSeenDateTime property
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name property
    name *string
    // The registrant property
    registrant DomainRegistrantable
    // The riskScore property
    riskScore *string
    // The tags property
    tags []string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
}
// NewDomainSecurityProfile instantiates a new domainSecurityProfile and sets the default values.
func NewDomainSecurityProfile()(*DomainSecurityProfile) {
    m := &DomainSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDomainSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainSecurityProfile(), nil
}
// GetActivityGroupNames gets the activityGroupNames property value. The activityGroupNames property
func (m *DomainSecurityProfile) GetActivityGroupNames()([]string) {
    return m.activityGroupNames
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *DomainSecurityProfile) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *DomainSecurityProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetCountHits gets the countHits property value. The countHits property
func (m *DomainSecurityProfile) GetCountHits()(*int32) {
    return m.countHits
}
// GetCountInOrg gets the countInOrg property value. The countInOrg property
func (m *DomainSecurityProfile) GetCountInOrg()(*int32) {
    return m.countInOrg
}
// GetDomainCategories gets the domainCategories property value. The domainCategories property
func (m *DomainSecurityProfile) GetDomainCategories()([]ReputationCategoryable) {
    return m.domainCategories
}
// GetDomainRegisteredDateTime gets the domainRegisteredDateTime property value. The domainRegisteredDateTime property
func (m *DomainSecurityProfile) GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.domainRegisteredDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityGroupNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetActivityGroupNames(res)
        }
        return nil
    }
    res["azureSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
        }
        return nil
    }
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["countHits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountHits(val)
        }
        return nil
    }
    res["countInOrg"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountInOrg(val)
        }
        return nil
    }
    res["domainCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateReputationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReputationCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(ReputationCategoryable)
            }
            m.SetDomainCategories(res)
        }
        return nil
    }
    res["domainRegisteredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainRegisteredDateTime(val)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["registrant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDomainRegistrantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrant(val.(DomainRegistrantable))
        }
        return nil
    }
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DomainSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *DomainSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetName gets the name property value. The name property
func (m *DomainSecurityProfile) GetName()(*string) {
    return m.name
}
// GetRegistrant gets the registrant property value. The registrant property
func (m *DomainSecurityProfile) GetRegistrant()(DomainRegistrantable) {
    return m.registrant
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *DomainSecurityProfile) GetRiskScore()(*string) {
    return m.riskScore
}
// GetTags gets the tags property value. The tags property
func (m *DomainSecurityProfile) GetTags()([]string) {
    return m.tags
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *DomainSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// Serialize serializes information the current object
func (m *DomainSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityGroupNames() != nil {
        err = writer.WriteCollectionOfStringValues("activityGroupNames", m.GetActivityGroupNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("countHits", m.GetCountHits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("countInOrg", m.GetCountInOrg())
        if err != nil {
            return err
        }
    }
    if m.GetDomainCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainCategories()))
        for i, v := range m.GetDomainCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("domainCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("domainRegisteredDateTime", m.GetDomainRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registrant", m.GetRegistrant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityGroupNames sets the activityGroupNames property value. The activityGroupNames property
func (m *DomainSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *DomainSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *DomainSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetCountHits sets the countHits property value. The countHits property
func (m *DomainSecurityProfile) SetCountHits(value *int32)() {
    m.countHits = value
}
// SetCountInOrg sets the countInOrg property value. The countInOrg property
func (m *DomainSecurityProfile) SetCountInOrg(value *int32)() {
    m.countInOrg = value
}
// SetDomainCategories sets the domainCategories property value. The domainCategories property
func (m *DomainSecurityProfile) SetDomainCategories(value []ReputationCategoryable)() {
    m.domainCategories = value
}
// SetDomainRegisteredDateTime sets the domainRegisteredDateTime property value. The domainRegisteredDateTime property
func (m *DomainSecurityProfile) SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.domainRegisteredDateTime = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DomainSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *DomainSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetName sets the name property value. The name property
func (m *DomainSecurityProfile) SetName(value *string)() {
    m.name = value
}
// SetRegistrant sets the registrant property value. The registrant property
func (m *DomainSecurityProfile) SetRegistrant(value DomainRegistrantable)() {
    m.registrant = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *DomainSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetTags sets the tags property value. The tags property
func (m *DomainSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *DomainSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
