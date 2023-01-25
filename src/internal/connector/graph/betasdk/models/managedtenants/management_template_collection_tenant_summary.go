package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagementTemplateCollectionTenantSummary provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagementTemplateCollectionTenantSummary struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The completeStepsCount property
    completeStepsCount *int32
    // The completeUsersCount property
    completeUsersCount *int32
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The dismissedStepsCount property
    dismissedStepsCount *int32
    // The excludedUsersCount property
    excludedUsersCount *int32
    // The excludedUsersDistinctCount property
    excludedUsersDistinctCount *int32
    // The incompleteStepsCount property
    incompleteStepsCount *int32
    // The incompleteUsersCount property
    incompleteUsersCount *int32
    // The ineligibleStepsCount property
    ineligibleStepsCount *int32
    // The isComplete property
    isComplete *bool
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managementTemplateCollectionDisplayName property
    managementTemplateCollectionDisplayName *string
    // The managementTemplateCollectionId property
    managementTemplateCollectionId *string
    // The tenantId property
    tenantId *string
}
// NewManagementTemplateCollectionTenantSummary instantiates a new managementTemplateCollectionTenantSummary and sets the default values.
func NewManagementTemplateCollectionTenantSummary()(*ManagementTemplateCollectionTenantSummary) {
    m := &ManagementTemplateCollectionTenantSummary{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateCollectionTenantSummary(), nil
}
// GetCompleteStepsCount gets the completeStepsCount property value. The completeStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetCompleteStepsCount()(*int32) {
    return m.completeStepsCount
}
// GetCompleteUsersCount gets the completeUsersCount property value. The completeUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) GetCompleteUsersCount()(*int32) {
    return m.completeUsersCount
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateCollectionTenantSummary) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateCollectionTenantSummary) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDismissedStepsCount gets the dismissedStepsCount property value. The dismissedStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetDismissedStepsCount()(*int32) {
    return m.dismissedStepsCount
}
// GetExcludedUsersCount gets the excludedUsersCount property value. The excludedUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) GetExcludedUsersCount()(*int32) {
    return m.excludedUsersCount
}
// GetExcludedUsersDistinctCount gets the excludedUsersDistinctCount property value. The excludedUsersDistinctCount property
func (m *ManagementTemplateCollectionTenantSummary) GetExcludedUsersDistinctCount()(*int32) {
    return m.excludedUsersDistinctCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateCollectionTenantSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completeStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleteStepsCount(val)
        }
        return nil
    }
    res["completeUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleteUsersCount(val)
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
    res["dismissedStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedStepsCount(val)
        }
        return nil
    }
    res["excludedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedUsersCount(val)
        }
        return nil
    }
    res["excludedUsersDistinctCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedUsersDistinctCount(val)
        }
        return nil
    }
    res["incompleteStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteStepsCount(val)
        }
        return nil
    }
    res["incompleteUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteUsersCount(val)
        }
        return nil
    }
    res["ineligibleStepsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIneligibleStepsCount(val)
        }
        return nil
    }
    res["isComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsComplete(val)
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
    return res
}
// GetIncompleteStepsCount gets the incompleteStepsCount property value. The incompleteStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetIncompleteStepsCount()(*int32) {
    return m.incompleteStepsCount
}
// GetIncompleteUsersCount gets the incompleteUsersCount property value. The incompleteUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) GetIncompleteUsersCount()(*int32) {
    return m.incompleteUsersCount
}
// GetIneligibleStepsCount gets the ineligibleStepsCount property value. The ineligibleStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) GetIneligibleStepsCount()(*int32) {
    return m.ineligibleStepsCount
}
// GetIsComplete gets the isComplete property value. The isComplete property
func (m *ManagementTemplateCollectionTenantSummary) GetIsComplete()(*bool) {
    return m.isComplete
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateCollectionTenantSummary) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateCollectionTenantSummary) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetManagementTemplateCollectionDisplayName gets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateCollectionTenantSummary) GetManagementTemplateCollectionDisplayName()(*string) {
    return m.managementTemplateCollectionDisplayName
}
// GetManagementTemplateCollectionId gets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateCollectionTenantSummary) GetManagementTemplateCollectionId()(*string) {
    return m.managementTemplateCollectionId
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagementTemplateCollectionTenantSummary) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagementTemplateCollectionTenantSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("completeStepsCount", m.GetCompleteStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("completeUsersCount", m.GetCompleteUsersCount())
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
        err = writer.WriteInt32Value("dismissedStepsCount", m.GetDismissedStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("excludedUsersCount", m.GetExcludedUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("excludedUsersDistinctCount", m.GetExcludedUsersDistinctCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("incompleteStepsCount", m.GetIncompleteStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("incompleteUsersCount", m.GetIncompleteUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ineligibleStepsCount", m.GetIneligibleStepsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isComplete", m.GetIsComplete())
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
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompleteStepsCount sets the completeStepsCount property value. The completeStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetCompleteStepsCount(value *int32)() {
    m.completeStepsCount = value
}
// SetCompleteUsersCount sets the completeUsersCount property value. The completeUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) SetCompleteUsersCount(value *int32)() {
    m.completeUsersCount = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateCollectionTenantSummary) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateCollectionTenantSummary) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDismissedStepsCount sets the dismissedStepsCount property value. The dismissedStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetDismissedStepsCount(value *int32)() {
    m.dismissedStepsCount = value
}
// SetExcludedUsersCount sets the excludedUsersCount property value. The excludedUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) SetExcludedUsersCount(value *int32)() {
    m.excludedUsersCount = value
}
// SetExcludedUsersDistinctCount sets the excludedUsersDistinctCount property value. The excludedUsersDistinctCount property
func (m *ManagementTemplateCollectionTenantSummary) SetExcludedUsersDistinctCount(value *int32)() {
    m.excludedUsersDistinctCount = value
}
// SetIncompleteStepsCount sets the incompleteStepsCount property value. The incompleteStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetIncompleteStepsCount(value *int32)() {
    m.incompleteStepsCount = value
}
// SetIncompleteUsersCount sets the incompleteUsersCount property value. The incompleteUsersCount property
func (m *ManagementTemplateCollectionTenantSummary) SetIncompleteUsersCount(value *int32)() {
    m.incompleteUsersCount = value
}
// SetIneligibleStepsCount sets the ineligibleStepsCount property value. The ineligibleStepsCount property
func (m *ManagementTemplateCollectionTenantSummary) SetIneligibleStepsCount(value *int32)() {
    m.ineligibleStepsCount = value
}
// SetIsComplete sets the isComplete property value. The isComplete property
func (m *ManagementTemplateCollectionTenantSummary) SetIsComplete(value *bool)() {
    m.isComplete = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateCollectionTenantSummary) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateCollectionTenantSummary) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetManagementTemplateCollectionDisplayName sets the managementTemplateCollectionDisplayName property value. The managementTemplateCollectionDisplayName property
func (m *ManagementTemplateCollectionTenantSummary) SetManagementTemplateCollectionDisplayName(value *string)() {
    m.managementTemplateCollectionDisplayName = value
}
// SetManagementTemplateCollectionId sets the managementTemplateCollectionId property value. The managementTemplateCollectionId property
func (m *ManagementTemplateCollectionTenantSummary) SetManagementTemplateCollectionId(value *string)() {
    m.managementTemplateCollectionId = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagementTemplateCollectionTenantSummary) SetTenantId(value *string)() {
    m.tenantId = value
}
