package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileSecurityProfile provides operations to manage the collection of site entities.
type FileSecurityProfile struct {
    Entity
    // The activityGroupNames property
    activityGroupNames []string
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The azureTenantId property
    azureTenantId *string
    // The certificateThumbprint property
    certificateThumbprint *string
    // The extensions property
    extensions []string
    // The fileType property
    fileType *string
    // The firstSeenDateTime property
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The hashes property
    hashes []FileHashable
    // The lastSeenDateTime property
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The malwareStates property
    malwareStates []MalwareStateable
    // The names property
    names []string
    // The riskScore property
    riskScore *string
    // The size property
    size *int64
    // The tags property
    tags []string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
    // The vulnerabilityStates property
    vulnerabilityStates []VulnerabilityStateable
}
// NewFileSecurityProfile instantiates a new fileSecurityProfile and sets the default values.
func NewFileSecurityProfile()(*FileSecurityProfile) {
    m := &FileSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFileSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileSecurityProfile(), nil
}
// GetActivityGroupNames gets the activityGroupNames property value. The activityGroupNames property
func (m *FileSecurityProfile) GetActivityGroupNames()([]string) {
    return m.activityGroupNames
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *FileSecurityProfile) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *FileSecurityProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetCertificateThumbprint gets the certificateThumbprint property value. The certificateThumbprint property
func (m *FileSecurityProfile) GetCertificateThumbprint()(*string) {
    return m.certificateThumbprint
}
// GetExtensions gets the extensions property value. The extensions property
func (m *FileSecurityProfile) GetExtensions()([]string) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["certificateThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateThumbprint(val)
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["fileType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileType(val)
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
    res["hashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileHashable, len(val))
            for i, v := range val {
                res[i] = v.(FileHashable)
            }
            m.SetHashes(res)
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
    res["malwareStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMalwareStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MalwareStateable, len(val))
            for i, v := range val {
                res[i] = v.(MalwareStateable)
            }
            m.SetMalwareStates(res)
        }
        return nil
    }
    res["names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNames(res)
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
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
    res["vulnerabilityStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVulnerabilityStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VulnerabilityStateable, len(val))
            for i, v := range val {
                res[i] = v.(VulnerabilityStateable)
            }
            m.SetVulnerabilityStates(res)
        }
        return nil
    }
    return res
}
// GetFileType gets the fileType property value. The fileType property
func (m *FileSecurityProfile) GetFileType()(*string) {
    return m.fileType
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *FileSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHashes gets the hashes property value. The hashes property
func (m *FileSecurityProfile) GetHashes()([]FileHashable) {
    return m.hashes
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *FileSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetMalwareStates gets the malwareStates property value. The malwareStates property
func (m *FileSecurityProfile) GetMalwareStates()([]MalwareStateable) {
    return m.malwareStates
}
// GetNames gets the names property value. The names property
func (m *FileSecurityProfile) GetNames()([]string) {
    return m.names
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *FileSecurityProfile) GetRiskScore()(*string) {
    return m.riskScore
}
// GetSize gets the size property value. The size property
func (m *FileSecurityProfile) GetSize()(*int64) {
    return m.size
}
// GetTags gets the tags property value. The tags property
func (m *FileSecurityProfile) GetTags()([]string) {
    return m.tags
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *FileSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// GetVulnerabilityStates gets the vulnerabilityStates property value. The vulnerabilityStates property
func (m *FileSecurityProfile) GetVulnerabilityStates()([]VulnerabilityStateable) {
    return m.vulnerabilityStates
}
// Serialize serializes information the current object
func (m *FileSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("certificateThumbprint", m.GetCertificateThumbprint())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        err = writer.WriteCollectionOfStringValues("extensions", m.GetExtensions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileType", m.GetFileType())
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
    if m.GetHashes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHashes()))
        for i, v := range m.GetHashes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("hashes", cast)
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
    if m.GetMalwareStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMalwareStates()))
        for i, v := range m.GetMalwareStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("malwareStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        err = writer.WriteCollectionOfStringValues("names", m.GetNames())
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
    {
        err = writer.WriteInt64Value("size", m.GetSize())
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
    if m.GetVulnerabilityStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVulnerabilityStates()))
        for i, v := range m.GetVulnerabilityStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("vulnerabilityStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityGroupNames sets the activityGroupNames property value. The activityGroupNames property
func (m *FileSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *FileSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *FileSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetCertificateThumbprint sets the certificateThumbprint property value. The certificateThumbprint property
func (m *FileSecurityProfile) SetCertificateThumbprint(value *string)() {
    m.certificateThumbprint = value
}
// SetExtensions sets the extensions property value. The extensions property
func (m *FileSecurityProfile) SetExtensions(value []string)() {
    m.extensions = value
}
// SetFileType sets the fileType property value. The fileType property
func (m *FileSecurityProfile) SetFileType(value *string)() {
    m.fileType = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *FileSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHashes sets the hashes property value. The hashes property
func (m *FileSecurityProfile) SetHashes(value []FileHashable)() {
    m.hashes = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *FileSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetMalwareStates sets the malwareStates property value. The malwareStates property
func (m *FileSecurityProfile) SetMalwareStates(value []MalwareStateable)() {
    m.malwareStates = value
}
// SetNames sets the names property value. The names property
func (m *FileSecurityProfile) SetNames(value []string)() {
    m.names = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *FileSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetSize sets the size property value. The size property
func (m *FileSecurityProfile) SetSize(value *int64)() {
    m.size = value
}
// SetTags sets the tags property value. The tags property
func (m *FileSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *FileSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
// SetVulnerabilityStates sets the vulnerabilityStates property value. The vulnerabilityStates property
func (m *FileSecurityProfile) SetVulnerabilityStates(value []VulnerabilityStateable)() {
    m.vulnerabilityStates = value
}
