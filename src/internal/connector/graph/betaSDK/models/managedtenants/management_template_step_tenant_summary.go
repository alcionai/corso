package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagementTemplateStepTenantSummary provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagementTemplateStepTenantSummary struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The assignedTenantsCount property
    assignedTenantsCount *int32
    // The compliantTenantsCount property
    compliantTenantsCount *int32
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The dismissedTenantsCount property
    dismissedTenantsCount *int32
    // The ineligibleTenantsCount property
    ineligibleTenantsCount *int32
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managementTemplateCollectionDisplayName property
    managementTemplateCollectionDisplayName *string
    // The managementTemplateCollectionId property
    managementTemplateCollectionId *string
    // The managementTemplateDisplayName property
    managementTemplateDisplayName *string
    // The managementTemplateId property
    managementTemplateId *string
    // The managementTemplateStepDisplayName property
    managementTemplateStepDisplayName *string
    // The managementTemplateStepId property
    managementTemplateStepId *string
    // The notCompliantTenantsCount property
    notCompliantTenantsCount *int32
}
// NewManagementTemplateStepTenantSummary instantiates a new managementTemplateStepTenantSummary and sets the default values.
func NewManagementTemplateStepTenantSummary()(*ManagementTemplateStepTenantSummary) {
    m := &ManagementTemplateStepTenantSummary{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStepTenantSummary(), nil
}
// GetAssignedTenantsCount gets the assignedTenantsCount property value. The assignedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetAssignedTenantsCount()(*int32) {
    return m.assignedTenantsCount
}
// GetCompliantTenantsCount gets the compliantTenantsCount property value. The compliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetCompliantTenantsCount()(*int32) {
    return m.compliantTenantsCount
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepTenantSummary) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepTenantSummary) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDismissedTenantsCount gets the dismissedTenantsCount property value. The dismissedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetDismissedTenantsCount()(*int32) {
    return m.dismissedTenantsCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepTenantSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTenantsCount(val)
        }
        return nil
    }
    res["compliantTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantTenantsCount(val)
        }
        return nil
    }
    res["createdByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["dismissedTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedTenantsCount(val)
        }
        return nil
    }
    res["ineligibleTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIneligibleTenantsCount(val)
        }
        return nil
    }
    res["lastActionByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["managementTemplateCollectionDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateCollectionDisplayName(val)
        }
        return nil
    }
    res["managementTemplateCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateCollectionId(val)
        }
        return nil
    }
    res["managementTemplateDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateDisplayName(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateStepDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepDisplayName(val)
        }
        return nil
    }
    res["managementTemplateStepId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepId(val)
        }
        return nil
    }
    res["notCompliantTenantsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotCompliantTenantsCount(val)
        }
        return nil
    }
    return res
}
// GetIneligibleTenantsCount gets the ineligibleTenantsCount property value. The ineligibleTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetIneligibleTenantsCount()(*int32) {
    return m.ineligibleTenantsCount
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepTenantSummary) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepTenantSummary) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetManagementTemplateCollectionDisplayName gets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateCollectionDisplayName()(*string) {
    return m.managementTemplateCollectionDisplayName
}
// GetManagementTemplateCollectionId gets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateCollectionId()(*string) {
    return m.managementTemplateCollectionId
}
// GetManagementTemplateDisplayName gets the managementTemplateDisplayName property value. The managementTemplateDisplayName property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateDisplayName()(*string) {
    return m.managementTemplateDisplayName
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateId()(*string) {
    return m.managementTemplateId
}
// GetManagementTemplateStepDisplayName gets the managementTemplateStepDisplayName property value. The managementTemplateStepDisplayName property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateStepDisplayName()(*string) {
    return m.managementTemplateStepDisplayName
}
// GetManagementTemplateStepId gets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ManagementTemplateStepTenantSummary) GetManagementTemplateStepId()(*string) {
    return m.managementTemplateStepId
}
// GetNotCompliantTenantsCount gets the notCompliantTenantsCount property value. The notCompliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) GetNotCompliantTenantsCount()(*int32) {
    return m.notCompliantTenantsCount
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepTenantSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("assignedTenantsCount", m.GetAssignedTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("compliantTenantsCount", m.GetCompliantTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dismissedTenantsCount", m.GetDismissedTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ineligibleTenantsCount", m.GetIneligibleTenantsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateCollectionDisplayName", m.GetManagementTemplateCollectionDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateCollectionId", m.GetManagementTemplateCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateDisplayName", m.GetManagementTemplateDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateStepDisplayName", m.GetManagementTemplateStepDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementTemplateStepId", m.GetManagementTemplateStepId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notCompliantTenantsCount", m.GetNotCompliantTenantsCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTenantsCount sets the assignedTenantsCount property value. The assignedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetAssignedTenantsCount(value *int32)() {
    m.assignedTenantsCount = value
}
// SetCompliantTenantsCount sets the compliantTenantsCount property value. The compliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetCompliantTenantsCount(value *int32)() {
    m.compliantTenantsCount = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepTenantSummary) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepTenantSummary) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDismissedTenantsCount sets the dismissedTenantsCount property value. The dismissedTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetDismissedTenantsCount(value *int32)() {
    m.dismissedTenantsCount = value
}
// SetIneligibleTenantsCount sets the ineligibleTenantsCount property value. The ineligibleTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetIneligibleTenantsCount(value *int32)() {
    m.ineligibleTenantsCount = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepTenantSummary) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepTenantSummary) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetManagementTemplateCollectionDisplayName sets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateCollectionDisplayName(value *string)() {
    m.managementTemplateCollectionDisplayName = value
}
// SetManagementTemplateCollectionId sets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateCollectionId(value *string)() {
    m.managementTemplateCollectionId = value
}
// SetManagementTemplateDisplayName sets the managementTemplateDisplayName property value. The managementTemplateDisplayName property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateDisplayName(value *string)() {
    m.managementTemplateDisplayName = value
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// SetManagementTemplateStepDisplayName sets the managementTemplateStepDisplayName property value. The managementTemplateStepDisplayName property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateStepDisplayName(value *string)() {
    m.managementTemplateStepDisplayName = value
}
// SetManagementTemplateStepId sets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ManagementTemplateStepTenantSummary) SetManagementTemplateStepId(value *string)() {
    m.managementTemplateStepId = value
}
// SetNotCompliantTenantsCount sets the notCompliantTenantsCount property value. The notCompliantTenantsCount property
func (m *ManagementTemplateStepTenantSummary) SetNotCompliantTenantsCount(value *int32)() {
    m.notCompliantTenantsCount = value
}
