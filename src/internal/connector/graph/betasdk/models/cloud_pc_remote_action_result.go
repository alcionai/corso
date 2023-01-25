package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcRemoteActionResult 
type CloudPcRemoteActionResult struct {
    // The specified action. Supported values in the Microsoft Endpoint Manager portal are: Reprovision, Resize, Restore. Supported values in enterprise Cloud PC devices are: Reboot, Rename, Reprovision, Troubleshoot.
    actionName *string
    // State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
    actionState *ActionState
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ID of the Cloud PC device on which the remote action is performed. Read-only.
    cloudPcId *string
    // Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID of the Intune managed device on which the remote action is performed. Read-only.
    managedDeviceId *string
    // The OdataType property
    odataType *string
    // Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The details of the Cloud PC status.
    statusDetails CloudPcStatusDetailsable
}
// NewCloudPcRemoteActionResult instantiates a new cloudPcRemoteActionResult and sets the default values.
func NewCloudPcRemoteActionResult()(*CloudPcRemoteActionResult) {
    m := &CloudPcRemoteActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcRemoteActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcRemoteActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcRemoteActionResult(), nil
}
// GetActionName gets the actionName property value. The specified action. Supported values in the Microsoft Endpoint Manager portal are: Reprovision, Resize, Restore. Supported values in enterprise Cloud PC devices are: Reboot, Rename, Reprovision, Troubleshoot.
func (m *CloudPcRemoteActionResult) GetActionName()(*string) {
    return m.actionName
}
// GetActionState gets the actionState property value. State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
func (m *CloudPcRemoteActionResult) GetActionState()(*ActionState) {
    return m.actionState
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRemoteActionResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCloudPcId gets the cloudPcId property value. The ID of the Cloud PC device on which the remote action is performed. Read-only.
func (m *CloudPcRemoteActionResult) GetCloudPcId()(*string) {
    return m.cloudPcId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcRemoteActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
    res["actionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionState(val.(*ActionState))
        }
        return nil
    }
    res["cloudPcId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcId(val)
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
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val.(CloudPcStatusDetailsable))
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcRemoteActionResult) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. The ID of the Intune managed device on which the remote action is performed. Read-only.
func (m *CloudPcRemoteActionResult) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcRemoteActionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDateTime gets the startDateTime property value. Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcRemoteActionResult) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatusDetails gets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPcRemoteActionResult) GetStatusDetails()(CloudPcStatusDetailsable) {
    return m.statusDetails
}
// Serialize serializes information the current object
func (m *CloudPcRemoteActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    if m.GetActionState() != nil {
        cast := (*m.GetActionState()).String()
        err := writer.WriteStringValue("actionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
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
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("statusDetails", m.GetStatusDetails())
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
// SetActionName sets the actionName property value. The specified action. Supported values in the Microsoft Endpoint Manager portal are: Reprovision, Resize, Restore. Supported values in enterprise Cloud PC devices are: Reboot, Rename, Reprovision, Troubleshoot.
func (m *CloudPcRemoteActionResult) SetActionName(value *string)() {
    m.actionName = value
}
// SetActionState sets the actionState property value. State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
func (m *CloudPcRemoteActionResult) SetActionState(value *ActionState)() {
    m.actionState = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRemoteActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCloudPcId sets the cloudPcId property value. The ID of the Cloud PC device on which the remote action is performed. Read-only.
func (m *CloudPcRemoteActionResult) SetCloudPcId(value *string)() {
    m.cloudPcId = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcRemoteActionResult) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The ID of the Intune managed device on which the remote action is performed. Read-only.
func (m *CloudPcRemoteActionResult) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcRemoteActionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDateTime sets the startDateTime property value. Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcRemoteActionResult) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatusDetails sets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPcRemoteActionResult) SetStatusDetails(value CloudPcStatusDetailsable)() {
    m.statusDetails = value
}
