package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagedTenantAlertRule provides operations to call the add method.
type ManagedTenantAlertRule struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The alertDisplayName property
    alertDisplayName *string
    // The alerts property
    alerts []ManagedTenantAlertable
    // The alertTTL property
    alertTTL *int32
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastRunDateTime property
    lastRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The notificationFinalDestinations property
    notificationFinalDestinations *NotificationDestination
    // The ruleDefinition property
    ruleDefinition ManagedTenantAlertRuleDefinitionable
    // The severity property
    severity *AlertSeverity
    // The targets property
    targets []NotificationTargetable
    // The tenantIds property
    tenantIds []TenantInfoable
}
// NewManagedTenantAlertRule instantiates a new managedTenantAlertRule and sets the default values.
func NewManagedTenantAlertRule()(*ManagedTenantAlertRule) {
    m := &ManagedTenantAlertRule{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagedTenantAlertRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantAlertRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantAlertRule(), nil
}
// GetAlertDisplayName gets the alertDisplayName property value. The alertDisplayName property
func (m *ManagedTenantAlertRule) GetAlertDisplayName()(*string) {
    return m.alertDisplayName
}
// GetAlerts gets the alerts property value. The alerts property
func (m *ManagedTenantAlertRule) GetAlerts()([]ManagedTenantAlertable) {
    return m.alerts
}
// GetAlertTTL gets the alertTTL property value. The alertTTL property
func (m *ManagedTenantAlertRule) GetAlertTTL()(*int32) {
    return m.alertTTL
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlertRule) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlertRule) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description property
func (m *ManagedTenantAlertRule) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ManagedTenantAlertRule) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantAlertRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertDisplayName(val)
        }
        return nil
    }
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertable)
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["alertTTL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertTTL(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["lastRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRunDateTime(val)
        }
        return nil
    }
    res["notificationFinalDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationDestination)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationFinalDestinations(val.(*NotificationDestination))
        }
        return nil
    }
    res["ruleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleDefinition(val.(ManagedTenantAlertRuleDefinitionable))
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
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationTargetable, len(val))
            for i, v := range val {
                res[i] = v.(NotificationTargetable)
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["tenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(TenantInfoable)
            }
            m.SetTenantIds(res)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlertRule) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlertRule) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetLastRunDateTime gets the lastRunDateTime property value. The lastRunDateTime property
func (m *ManagedTenantAlertRule) GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRunDateTime
}
// GetNotificationFinalDestinations gets the notificationFinalDestinations property value. The notificationFinalDestinations property
func (m *ManagedTenantAlertRule) GetNotificationFinalDestinations()(*NotificationDestination) {
    return m.notificationFinalDestinations
}
// GetRuleDefinition gets the ruleDefinition property value. The ruleDefinition property
func (m *ManagedTenantAlertRule) GetRuleDefinition()(ManagedTenantAlertRuleDefinitionable) {
    return m.ruleDefinition
}
// GetSeverity gets the severity property value. The severity property
func (m *ManagedTenantAlertRule) GetSeverity()(*AlertSeverity) {
    return m.severity
}
// GetTargets gets the targets property value. The targets property
func (m *ManagedTenantAlertRule) GetTargets()([]NotificationTargetable) {
    return m.targets
}
// GetTenantIds gets the tenantIds property value. The tenantIds property
func (m *ManagedTenantAlertRule) GetTenantIds()([]TenantInfoable) {
    return m.tenantIds
}
// Serialize serializes information the current object
func (m *ManagedTenantAlertRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alertDisplayName", m.GetAlertDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("alertTTL", m.GetAlertTTL())
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
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteTimeValue("lastRunDateTime", m.GetLastRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationFinalDestinations() != nil {
        cast := (*m.GetNotificationFinalDestinations()).String()
        err = writer.WriteStringValue("notificationFinalDestinations", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ruleDefinition", m.GetRuleDefinition())
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
    if m.GetTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantIds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantIds()))
        for i, v := range m.GetTenantIds() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantIds", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertDisplayName sets the alertDisplayName property value. The alertDisplayName property
func (m *ManagedTenantAlertRule) SetAlertDisplayName(value *string)() {
    m.alertDisplayName = value
}
// SetAlerts sets the alerts property value. The alerts property
func (m *ManagedTenantAlertRule) SetAlerts(value []ManagedTenantAlertable)() {
    m.alerts = value
}
// SetAlertTTL sets the alertTTL property value. The alertTTL property
func (m *ManagedTenantAlertRule) SetAlertTTL(value *int32)() {
    m.alertTTL = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagedTenantAlertRule) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagedTenantAlertRule) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *ManagedTenantAlertRule) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ManagedTenantAlertRule) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagedTenantAlertRule) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagedTenantAlertRule) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetLastRunDateTime sets the lastRunDateTime property value. The lastRunDateTime property
func (m *ManagedTenantAlertRule) SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRunDateTime = value
}
// SetNotificationFinalDestinations sets the notificationFinalDestinations property value. The notificationFinalDestinations property
func (m *ManagedTenantAlertRule) SetNotificationFinalDestinations(value *NotificationDestination)() {
    m.notificationFinalDestinations = value
}
// SetRuleDefinition sets the ruleDefinition property value. The ruleDefinition property
func (m *ManagedTenantAlertRule) SetRuleDefinition(value ManagedTenantAlertRuleDefinitionable)() {
    m.ruleDefinition = value
}
// SetSeverity sets the severity property value. The severity property
func (m *ManagedTenantAlertRule) SetSeverity(value *AlertSeverity)() {
    m.severity = value
}
// SetTargets sets the targets property value. The targets property
func (m *ManagedTenantAlertRule) SetTargets(value []NotificationTargetable)() {
    m.targets = value
}
// SetTenantIds sets the tenantIds property value. The tenantIds property
func (m *ManagedTenantAlertRule) SetTenantIds(value []TenantInfoable)() {
    m.tenantIds = value
}
