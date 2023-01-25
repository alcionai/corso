package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagedTenantAlert provides operations to call the add method.
type ManagedTenantAlert struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The alertData property
    alertData AlertDataable
    // The alertDataReferenceStrings property
    alertDataReferenceStrings []AlertDataReferenceStringable
    // The alertLogs property
    alertLogs []ManagedTenantAlertLogable
    // The alertRule property
    alertRule ManagedTenantAlertRuleable
    // The alertRuleDisplayName property
    alertRuleDisplayName *string
    // The apiNotifications property
    apiNotifications []ManagedTenantApiNotificationable
    // The assignedToUserId property
    assignedToUserId *string
    // The correlationCount property
    correlationCount *int32
    // The correlationId property
    correlationId *string
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The emailNotifications property
    emailNotifications []ManagedTenantEmailNotificationable
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The message property
    message *string
    // The severity property
    severity *AlertSeverity
    // The status property
    status *AlertStatus
    // The tenantId property
    tenantId *string
    // The title property
    title *string
}
// NewManagedTenantAlert instantiates a new managedTenantAlert and sets the default values.
func NewManagedTenantAlert()(*ManagedTenantAlert) {
    m := &ManagedTenantAlert{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagedTenantAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantAlert(), nil
}
// GetAlertData gets the alertData property value. The alertData property
func (m *ManagedTenantAlert) GetAlertData()(AlertDataable) {
    return m.alertData
}
// GetAlertDataReferenceStrings gets the alertDataReferenceStrings property value. The alertDataReferenceStrings property
func (m *ManagedTenantAlert) GetAlertDataReferenceStrings()([]AlertDataReferenceStringable) {
    return m.alertDataReferenceStrings
}
// GetAlertLogs gets the alertLogs property value. The alertLogs property
func (m *ManagedTenantAlert) GetAlertLogs()([]ManagedTenantAlertLogable) {
    return m.alertLogs
}
// GetAlertRule gets the alertRule property value. The alertRule property
func (m *ManagedTenantAlert) GetAlertRule()(ManagedTenantAlertRuleable) {
    return m.alertRule
}
// GetAlertRuleDisplayName gets the alertRuleDisplayName property value. The alertRuleDisplayName property
func (m *ManagedTenantAlert) GetAlertRuleDisplayName()(*string) {
    return m.alertRuleDisplayName
}
// GetApiNotifications gets the apiNotifications property value. The apiNotifications property
func (m *ManagedTenantAlert) GetApiNotifications()([]ManagedTenantApiNotificationable) {
    return m.apiNotifications
}
// GetAssignedToUserId gets the assignedToUserId property value. The assignedToUserId property
func (m *ManagedTenantAlert) GetAssignedToUserId()(*string) {
    return m.assignedToUserId
}
// GetCorrelationCount gets the correlationCount property value. The correlationCount property
func (m *ManagedTenantAlert) GetCorrelationCount()(*int32) {
    return m.correlationCount
}
// GetCorrelationId gets the correlationId property value. The correlationId property
func (m *ManagedTenantAlert) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlert) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlert) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetEmailNotifications gets the emailNotifications property value. The emailNotifications property
func (m *ManagedTenantAlert) GetEmailNotifications()([]ManagedTenantEmailNotificationable) {
    return m.emailNotifications
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlertDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertData(val.(AlertDataable))
        }
        return nil
    }
    res["alertDataReferenceStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertDataReferenceStringFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertDataReferenceStringable, len(val))
            for i, v := range val {
                res[i] = v.(AlertDataReferenceStringable)
            }
            m.SetAlertDataReferenceStrings(res)
        }
        return nil
    }
    res["alertLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertLogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertLogable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertLogable)
            }
            m.SetAlertLogs(res)
        }
        return nil
    }
    res["alertRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedTenantAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRule(val.(ManagedTenantAlertRuleable))
        }
        return nil
    }
    res["alertRuleDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleDisplayName(val)
        }
        return nil
    }
    res["apiNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantApiNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantApiNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantApiNotificationable)
            }
            m.SetApiNotifications(res)
        }
        return nil
    }
    res["assignedToUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedToUserId(val)
        }
        return nil
    }
    res["correlationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationCount(val)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
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
    res["emailNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantEmailNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantEmailNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantEmailNotificationable)
            }
            m.SetEmailNotifications(res)
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
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*AlertSeverity))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AlertStatus))
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlert) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlert) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetMessage gets the message property value. The message property
func (m *ManagedTenantAlert) GetMessage()(*string) {
    return m.message
}
// GetSeverity gets the severity property value. The severity property
func (m *ManagedTenantAlert) GetSeverity()(*AlertSeverity) {
    return m.severity
}
// GetStatus gets the status property value. The status property
func (m *ManagedTenantAlert) GetStatus()(*AlertStatus) {
    return m.status
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantAlert) GetTenantId()(*string) {
    return m.tenantId
}
// GetTitle gets the title property value. The title property
func (m *ManagedTenantAlert) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *ManagedTenantAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alertData", m.GetAlertData())
        if err != nil {
            return err
        }
    }
    if m.GetAlertDataReferenceStrings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertDataReferenceStrings()))
        for i, v := range m.GetAlertDataReferenceStrings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alertDataReferenceStrings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlertLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertLogs()))
        for i, v := range m.GetAlertLogs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alertLogs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("alertRule", m.GetAlertRule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertRuleDisplayName", m.GetAlertRuleDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetApiNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApiNotifications()))
        for i, v := range m.GetApiNotifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apiNotifications", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignedToUserId", m.GetAssignedToUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("correlationCount", m.GetCorrelationCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
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
    if m.GetEmailNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailNotifications()))
        for i, v := range m.GetEmailNotifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailNotifications", cast)
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
        err = writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
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
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertData sets the alertData property value. The alertData property
func (m *ManagedTenantAlert) SetAlertData(value AlertDataable)() {
    m.alertData = value
}
// SetAlertDataReferenceStrings sets the alertDataReferenceStrings property value. The alertDataReferenceStrings property
func (m *ManagedTenantAlert) SetAlertDataReferenceStrings(value []AlertDataReferenceStringable)() {
    m.alertDataReferenceStrings = value
}
// SetAlertLogs sets the alertLogs property value. The alertLogs property
func (m *ManagedTenantAlert) SetAlertLogs(value []ManagedTenantAlertLogable)() {
    m.alertLogs = value
}
// SetAlertRule sets the alertRule property value. The alertRule property
func (m *ManagedTenantAlert) SetAlertRule(value ManagedTenantAlertRuleable)() {
    m.alertRule = value
}
// SetAlertRuleDisplayName sets the alertRuleDisplayName property value. The alertRuleDisplayName property
func (m *ManagedTenantAlert) SetAlertRuleDisplayName(value *string)() {
    m.alertRuleDisplayName = value
}
// SetApiNotifications sets the apiNotifications property value. The apiNotifications property
func (m *ManagedTenantAlert) SetApiNotifications(value []ManagedTenantApiNotificationable)() {
    m.apiNotifications = value
}
// SetAssignedToUserId sets the assignedToUserId property value. The assignedToUserId property
func (m *ManagedTenantAlert) SetAssignedToUserId(value *string)() {
    m.assignedToUserId = value
}
// SetCorrelationCount sets the correlationCount property value. The correlationCount property
func (m *ManagedTenantAlert) SetCorrelationCount(value *int32)() {
    m.correlationCount = value
}
// SetCorrelationId sets the correlationId property value. The correlationId property
func (m *ManagedTenantAlert) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlert) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlert) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetEmailNotifications sets the emailNotifications property value. The emailNotifications property
func (m *ManagedTenantAlert) SetEmailNotifications(value []ManagedTenantEmailNotificationable)() {
    m.emailNotifications = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlert) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlert) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetMessage sets the message property value. The message property
func (m *ManagedTenantAlert) SetMessage(value *string)() {
    m.message = value
}
// SetSeverity sets the severity property value. The severity property
func (m *ManagedTenantAlert) SetSeverity(value *AlertSeverity)() {
    m.severity = value
}
// SetStatus sets the status property value. The status property
func (m *ManagedTenantAlert) SetStatus(value *AlertStatus)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantAlert) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetTitle sets the title property value. The title property
func (m *ManagedTenantAlert) SetTitle(value *string)() {
    m.title = value
}
