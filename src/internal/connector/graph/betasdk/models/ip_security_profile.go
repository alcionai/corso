package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpSecurityProfile provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type IpSecurityProfile struct {
    Entity
    // The activityGroupNames property
    activityGroupNames []string
    // The address property
    address *string
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The azureTenantId property
    azureTenantId *string
    // The countHits property
    countHits *int32
    // The countHosts property
    countHosts *int32
    // The firstSeenDateTime property
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ipCategories property
    ipCategories []IpCategoryable
    // The ipReferenceData property
    ipReferenceData []IpReferenceDataable
    // The lastSeenDateTime property
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The riskScore property
    riskScore *string
    // The tags property
    tags []string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
}
// NewIpSecurityProfile instantiates a new ipSecurityProfile and sets the default values.
func NewIpSecurityProfile()(*IpSecurityProfile) {
    m := &IpSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIpSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpSecurityProfile(), nil
}
// GetActivityGroupNames gets the activityGroupNames property value. The activityGroupNames property
func (m *IpSecurityProfile) GetActivityGroupNames()([]string) {
    return m.activityGroupNames
}
// GetAddress gets the address property value. The address property
func (m *IpSecurityProfile) GetAddress()(*string) {
    return m.address
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *IpSecurityProfile) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *IpSecurityProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetCountHits gets the countHits property value. The countHits property
func (m *IpSecurityProfile) GetCountHits()(*int32) {
    return m.countHits
}
// GetCountHosts gets the countHosts property value. The countHosts property
func (m *IpSecurityProfile) GetCountHosts()(*int32) {
    return m.countHosts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["countHosts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountHosts(val)
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
    res["ipCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(IpCategoryable)
            }
            m.SetIpCategories(res)
        }
        return nil
    }
    res["ipReferenceData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpReferenceDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpReferenceDataable, len(val))
            for i, v := range val {
                res[i] = v.(IpReferenceDataable)
            }
            m.SetIpReferenceData(res)
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
func (m *IpSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetIpCategories gets the ipCategories property value. The ipCategories property
func (m *IpSecurityProfile) GetIpCategories()([]IpCategoryable) {
    return m.ipCategories
}
// GetIpReferenceData gets the ipReferenceData property value. The ipReferenceData property
func (m *IpSecurityProfile) GetIpReferenceData()([]IpReferenceDataable) {
    return m.ipReferenceData
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *IpSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *IpSecurityProfile) GetRiskScore()(*string) {
    return m.riskScore
}
// GetTags gets the tags property value. The tags property
func (m *IpSecurityProfile) GetTags()([]string) {
    return m.tags
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *IpSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// Serialize serializes information the current object
func (m *IpSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("address", m.GetAddress())
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
        err = writer.WriteInt32Value("countHosts", m.GetCountHosts())
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
    if m.GetIpCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpCategories()))
        for i, v := range m.GetIpCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("ipCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIpReferenceData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIpReferenceData()))
        for i, v := range m.GetIpReferenceData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("ipReferenceData", cast)
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
func (m *IpSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// SetAddress sets the address property value. The address property
func (m *IpSecurityProfile) SetAddress(value *string)() {
    m.address = value
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *IpSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *IpSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetCountHits sets the countHits property value. The countHits property
func (m *IpSecurityProfile) SetCountHits(value *int32)() {
    m.countHits = value
}
// SetCountHosts sets the countHosts property value. The countHosts property
func (m *IpSecurityProfile) SetCountHosts(value *int32)() {
    m.countHosts = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *IpSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetIpCategories sets the ipCategories property value. The ipCategories property
func (m *IpSecurityProfile) SetIpCategories(value []IpCategoryable)() {
    m.ipCategories = value
}
// SetIpReferenceData sets the ipReferenceData property value. The ipReferenceData property
func (m *IpSecurityProfile) SetIpReferenceData(value []IpReferenceDataable)() {
    m.ipReferenceData = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *IpSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *IpSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetTags sets the tags property value. The tags property
func (m *IpSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *IpSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
