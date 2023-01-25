package devicemanagement

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// AlertRecord provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AlertRecord struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The impact of the alert event. Consists of a number followed by the aggregation type. For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. 12 affectedCloudPcPercentage means 12% of Cloud PCs are affected.
    alertImpact AlertImpactable
    // The corresponding ID of the alert rule.
    alertRuleId *string
    // The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
    alertRuleTemplate *AlertRuleTemplate
    // The date and time when the alert event was detected. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    detectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display name of the alert record.
    displayName *string
    // The date and time when the alert record was last updated. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the alert event was resolved. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    resolvedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The severity of the alert event. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
    severity *RuleSeverityType
    // The status of the alert record. The possible values are: active, resolved, unknownFutureValue.
    status *AlertStatusType
}
// NewAlertRecord instantiates a new alertRecord and sets the default values.
func NewAlertRecord()(*AlertRecord) {
    m := &AlertRecord{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateAlertRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertRecord(), nil
}
// GetAlertImpact gets the alertImpact property value. The impact of the alert event. Consists of a number followed by the aggregation type. For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. 12 affectedCloudPcPercentage means 12% of Cloud PCs are affected.
func (m *AlertRecord) GetAlertImpact()(AlertImpactable) {
    return m.alertImpact
}
// GetAlertRuleId gets the alertRuleId property value. The corresponding ID of the alert rule.
func (m *AlertRecord) GetAlertRuleId()(*string) {
    return m.alertRuleId
}
// GetAlertRuleTemplate gets the alertRuleTemplate property value. The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *AlertRecord) GetAlertRuleTemplate()(*AlertRuleTemplate) {
    return m.alertRuleTemplate
}
// GetDetectedDateTime gets the detectedDateTime property value. The date and time when the alert event was detected. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AlertRecord) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.detectedDateTime
}
// GetDisplayName gets the displayName property value. The display name of the alert record.
func (m *AlertRecord) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlertImpactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertImpact(val.(AlertImpactable))
        }
        return nil
    }
    res["alertRuleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleId(val)
        }
        return nil
    }
    res["alertRuleTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertRuleTemplate)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleTemplate(val.(*AlertRuleTemplate))
        }
        return nil
    }
    res["detectedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectedDateTime(val)
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
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["resolvedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedDateTime(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleSeverityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*RuleSeverityType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatusType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AlertStatusType))
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date and time when the alert record was last updated. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AlertRecord) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetResolvedDateTime gets the resolvedDateTime property value. The date and time when the alert event was resolved. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AlertRecord) GetResolvedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.resolvedDateTime
}
// GetSeverity gets the severity property value. The severity of the alert event. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *AlertRecord) GetSeverity()(*RuleSeverityType) {
    return m.severity
}
// GetStatus gets the status property value. The status of the alert record. The possible values are: active, resolved, unknownFutureValue.
func (m *AlertRecord) GetStatus()(*AlertStatusType) {
    return m.status
}
// Serialize serializes information the current object
func (m *AlertRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alertImpact", m.GetAlertImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertRuleId", m.GetAlertRuleId())
        if err != nil {
            return err
        }
    }
    if m.GetAlertRuleTemplate() != nil {
        cast := (*m.GetAlertRuleTemplate()).String()
        err = writer.WriteStringValue("alertRuleTemplate", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("detectedDateTime", m.GetDetectedDateTime())
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("resolvedDateTime", m.GetResolvedDateTime())
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
    return nil
}
// SetAlertImpact sets the alertImpact property value. The impact of the alert event. Consists of a number followed by the aggregation type. For example, 6 affectedCloudPcCount means that 6 Cloud PCs are affected. 12 affectedCloudPcPercentage means 12% of Cloud PCs are affected.
func (m *AlertRecord) SetAlertImpact(value AlertImpactable)() {
    m.alertImpact = value
}
// SetAlertRuleId sets the alertRuleId property value. The corresponding ID of the alert rule.
func (m *AlertRecord) SetAlertRuleId(value *string)() {
    m.alertRuleId = value
}
// SetAlertRuleTemplate sets the alertRuleTemplate property value. The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *AlertRecord) SetAlertRuleTemplate(value *AlertRuleTemplate)() {
    m.alertRuleTemplate = value
}
// SetDetectedDateTime sets the detectedDateTime property value. The date and time when the alert event was detected. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AlertRecord) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.detectedDateTime = value
}
// SetDisplayName sets the displayName property value. The display name of the alert record.
func (m *AlertRecord) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date and time when the alert record was last updated. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AlertRecord) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetResolvedDateTime sets the resolvedDateTime property value. The date and time when the alert event was resolved. The Timestamp type represents date and time information using ISO 8601 format. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AlertRecord) SetResolvedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.resolvedDateTime = value
}
// SetSeverity sets the severity property value. The severity of the alert event. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *AlertRecord) SetSeverity(value *RuleSeverityType)() {
    m.severity = value
}
// SetStatus sets the status property value. The status of the alert record. The possible values are: active, resolved, unknownFutureValue.
func (m *AlertRecord) SetStatus(value *AlertStatusType)() {
    m.status = value
}
