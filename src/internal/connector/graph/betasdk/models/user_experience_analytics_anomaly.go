package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomaly the user experience analytics anomaly entity contains anomaly details.
type UserExperienceAnalyticsAnomaly struct {
    Entity
    // Indicates the first occurrence date and time for the anomaly.
    anomalyFirstOccurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the anomaly.
    anomalyId *string
    // Indicates the latest occurrence date and time for the anomaly.
    anomalyLatestOccurrenceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the anomaly.
    anomalyName *string
    // Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
    anomalyType *UserExperienceAnalyticsAnomalyType
    // The name of the application or module that caused the anomaly.
    assetName *string
    // The publisher of the application or module that caused the anomaly.
    assetPublisher *string
    // The version of the application or module that caused the anomaly.
    assetVersion *string
    // The unique identifier of the anomaly detection model.
    detectionModelId *string
    // The number of devices impacted by the anomaly. Valid values -2147483648 to 2147483647
    deviceImpactedCount *int32
    // The unique identifier of the anomaly detection model.
    issueId *string
    // Indicates the severity of the anomaly. Eg: anomaly severity can be high, medium, low, informational or other.
    severity *UserExperienceAnalyticsAnomalySeverity
    // Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
    state *UserExperienceAnalyticsAnomalyState
}
// NewUserExperienceAnalyticsAnomaly instantiates a new userExperienceAnalyticsAnomaly and sets the default values.
func NewUserExperienceAnalyticsAnomaly()(*UserExperienceAnalyticsAnomaly) {
    m := &UserExperienceAnalyticsAnomaly{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomaly(), nil
}
// GetAnomalyFirstOccurrenceDateTime gets the anomalyFirstOccurrenceDateTime property value. Indicates the first occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.anomalyFirstOccurrenceDateTime
}
// GetAnomalyId gets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyId()(*string) {
    return m.anomalyId
}
// GetAnomalyLatestOccurrenceDateTime gets the anomalyLatestOccurrenceDateTime property value. Indicates the latest occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.anomalyLatestOccurrenceDateTime
}
// GetAnomalyName gets the anomalyName property value. The name of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyName()(*string) {
    return m.anomalyName
}
// GetAnomalyType gets the anomalyType property value. Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
func (m *UserExperienceAnalyticsAnomaly) GetAnomalyType()(*UserExperienceAnalyticsAnomalyType) {
    return m.anomalyType
}
// GetAssetName gets the assetName property value. The name of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAssetName()(*string) {
    return m.assetName
}
// GetAssetPublisher gets the assetPublisher property value. The publisher of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAssetPublisher()(*string) {
    return m.assetPublisher
}
// GetAssetVersion gets the assetVersion property value. The version of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) GetAssetVersion()(*string) {
    return m.assetVersion
}
// GetDetectionModelId gets the detectionModelId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) GetDetectionModelId()(*string) {
    return m.detectionModelId
}
// GetDeviceImpactedCount gets the deviceImpactedCount property value. The number of devices impacted by the anomaly. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomaly) GetDeviceImpactedCount()(*int32) {
    return m.deviceImpactedCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAnomaly) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["anomalyFirstOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyFirstOccurrenceDateTime(val)
        }
        return nil
    }
    res["anomalyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyId(val)
        }
        return nil
    }
    res["anomalyLatestOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyLatestOccurrenceDateTime(val)
        }
        return nil
    }
    res["anomalyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyName(val)
        }
        return nil
    }
    res["anomalyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyType(val.(*UserExperienceAnalyticsAnomalyType))
        }
        return nil
    }
    res["assetName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetName(val)
        }
        return nil
    }
    res["assetPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetPublisher(val)
        }
        return nil
    }
    res["assetVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetVersion(val)
        }
        return nil
    }
    res["detectionModelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionModelId(val)
        }
        return nil
    }
    res["deviceImpactedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceImpactedCount(val)
        }
        return nil
    }
    res["issueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueId(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalySeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*UserExperienceAnalyticsAnomalySeverity))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsAnomalyState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*UserExperienceAnalyticsAnomalyState))
        }
        return nil
    }
    return res
}
// GetIssueId gets the issueId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) GetIssueId()(*string) {
    return m.issueId
}
// GetSeverity gets the severity property value. Indicates the severity of the anomaly. Eg: anomaly severity can be high, medium, low, informational or other.
func (m *UserExperienceAnalyticsAnomaly) GetSeverity()(*UserExperienceAnalyticsAnomalySeverity) {
    return m.severity
}
// GetState gets the state property value. Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
func (m *UserExperienceAnalyticsAnomaly) GetState()(*UserExperienceAnalyticsAnomalyState) {
    return m.state
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomaly) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("anomalyFirstOccurrenceDateTime", m.GetAnomalyFirstOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("anomalyId", m.GetAnomalyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("anomalyLatestOccurrenceDateTime", m.GetAnomalyLatestOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("anomalyName", m.GetAnomalyName())
        if err != nil {
            return err
        }
    }
    if m.GetAnomalyType() != nil {
        cast := (*m.GetAnomalyType()).String()
        err = writer.WriteStringValue("anomalyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assetName", m.GetAssetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assetPublisher", m.GetAssetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assetVersion", m.GetAssetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("detectionModelId", m.GetDetectionModelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceImpactedCount", m.GetDeviceImpactedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issueId", m.GetIssueId())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnomalyFirstOccurrenceDateTime sets the anomalyFirstOccurrenceDateTime property value. Indicates the first occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.anomalyFirstOccurrenceDateTime = value
}
// SetAnomalyId sets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyId(value *string)() {
    m.anomalyId = value
}
// SetAnomalyLatestOccurrenceDateTime sets the anomalyLatestOccurrenceDateTime property value. Indicates the latest occurrence date and time for the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.anomalyLatestOccurrenceDateTime = value
}
// SetAnomalyName sets the anomalyName property value. The name of the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyName(value *string)() {
    m.anomalyName = value
}
// SetAnomalyType sets the anomalyType property value. Indicates the category of the anomaly. Eg: anomaly type can be device, application, stop error, driver or other.
func (m *UserExperienceAnalyticsAnomaly) SetAnomalyType(value *UserExperienceAnalyticsAnomalyType)() {
    m.anomalyType = value
}
// SetAssetName sets the assetName property value. The name of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAssetName(value *string)() {
    m.assetName = value
}
// SetAssetPublisher sets the assetPublisher property value. The publisher of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAssetPublisher(value *string)() {
    m.assetPublisher = value
}
// SetAssetVersion sets the assetVersion property value. The version of the application or module that caused the anomaly.
func (m *UserExperienceAnalyticsAnomaly) SetAssetVersion(value *string)() {
    m.assetVersion = value
}
// SetDetectionModelId sets the detectionModelId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) SetDetectionModelId(value *string)() {
    m.detectionModelId = value
}
// SetDeviceImpactedCount sets the deviceImpactedCount property value. The number of devices impacted by the anomaly. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAnomaly) SetDeviceImpactedCount(value *int32)() {
    m.deviceImpactedCount = value
}
// SetIssueId sets the issueId property value. The unique identifier of the anomaly detection model.
func (m *UserExperienceAnalyticsAnomaly) SetIssueId(value *string)() {
    m.issueId = value
}
// SetSeverity sets the severity property value. Indicates the severity of the anomaly. Eg: anomaly severity can be high, medium, low, informational or other.
func (m *UserExperienceAnalyticsAnomaly) SetSeverity(value *UserExperienceAnalyticsAnomalySeverity)() {
    m.severity = value
}
// SetState sets the state property value. Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
func (m *UserExperienceAnalyticsAnomaly) SetState(value *UserExperienceAnalyticsAnomalyState)() {
    m.state = value
}
