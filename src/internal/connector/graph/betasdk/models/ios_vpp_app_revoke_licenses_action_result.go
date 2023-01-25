package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppRevokeLicensesActionResult defines results for actions on iOS Vpp Apps, contains inherited properties for ActionResult.
type IosVppAppRevokeLicensesActionResult struct {
    // Possible types of reasons for an Apple Volume Purchase Program token action failure.
    actionFailureReason *VppTokenActionFailureReason
    // Action name
    actionName *string
    // The actionState property
    actionState *ActionState
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A count of the number of licenses for which revoke failed.
    failedLicensesCount *int32
    // Time the action state was last updated
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // DeviceId associated with the action.
    managedDeviceId *string
    // The OdataType property
    odataType *string
    // Time the action was initiated
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A count of the number of licenses for which revoke was attempted.
    totalLicensesCount *int32
    // UserId associated with the action.
    userId *string
}
// NewIosVppAppRevokeLicensesActionResult instantiates a new iosVppAppRevokeLicensesActionResult and sets the default values.
func NewIosVppAppRevokeLicensesActionResult()(*IosVppAppRevokeLicensesActionResult) {
    m := &IosVppAppRevokeLicensesActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosVppAppRevokeLicensesActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppRevokeLicensesActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppRevokeLicensesActionResult(), nil
}
// GetActionFailureReason gets the actionFailureReason property value. Possible types of reasons for an Apple Volume Purchase Program token action failure.
func (m *IosVppAppRevokeLicensesActionResult) GetActionFailureReason()(*VppTokenActionFailureReason) {
    return m.actionFailureReason
}
// GetActionName gets the actionName property value. Action name
func (m *IosVppAppRevokeLicensesActionResult) GetActionName()(*string) {
    return m.actionName
}
// GetActionState gets the actionState property value. The actionState property
func (m *IosVppAppRevokeLicensesActionResult) GetActionState()(*ActionState) {
    return m.actionState
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosVppAppRevokeLicensesActionResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFailedLicensesCount gets the failedLicensesCount property value. A count of the number of licenses for which revoke failed.
func (m *IosVppAppRevokeLicensesActionResult) GetFailedLicensesCount()(*int32) {
    return m.failedLicensesCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppRevokeLicensesActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionFailureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenActionFailureReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionFailureReason(val.(*VppTokenActionFailureReason))
        }
        return nil
    }
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
    res["failedLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedLicensesCount(val)
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
    res["totalLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicensesCount(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Time the action state was last updated
func (m *IosVppAppRevokeLicensesActionResult) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. DeviceId associated with the action.
func (m *IosVppAppRevokeLicensesActionResult) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosVppAppRevokeLicensesActionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDateTime gets the startDateTime property value. Time the action was initiated
func (m *IosVppAppRevokeLicensesActionResult) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetTotalLicensesCount gets the totalLicensesCount property value. A count of the number of licenses for which revoke was attempted.
func (m *IosVppAppRevokeLicensesActionResult) GetTotalLicensesCount()(*int32) {
    return m.totalLicensesCount
}
// GetUserId gets the userId property value. UserId associated with the action.
func (m *IosVppAppRevokeLicensesActionResult) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *IosVppAppRevokeLicensesActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionFailureReason() != nil {
        cast := (*m.GetActionFailureReason()).String()
        err := writer.WriteStringValue("actionFailureReason", &cast)
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteInt32Value("failedLicensesCount", m.GetFailedLicensesCount())
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
        err := writer.WriteInt32Value("totalLicensesCount", m.GetTotalLicensesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
// SetActionFailureReason sets the actionFailureReason property value. Possible types of reasons for an Apple Volume Purchase Program token action failure.
func (m *IosVppAppRevokeLicensesActionResult) SetActionFailureReason(value *VppTokenActionFailureReason)() {
    m.actionFailureReason = value
}
// SetActionName sets the actionName property value. Action name
func (m *IosVppAppRevokeLicensesActionResult) SetActionName(value *string)() {
    m.actionName = value
}
// SetActionState sets the actionState property value. The actionState property
func (m *IosVppAppRevokeLicensesActionResult) SetActionState(value *ActionState)() {
    m.actionState = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosVppAppRevokeLicensesActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFailedLicensesCount sets the failedLicensesCount property value. A count of the number of licenses for which revoke failed.
func (m *IosVppAppRevokeLicensesActionResult) SetFailedLicensesCount(value *int32)() {
    m.failedLicensesCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Time the action state was last updated
func (m *IosVppAppRevokeLicensesActionResult) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. DeviceId associated with the action.
func (m *IosVppAppRevokeLicensesActionResult) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosVppAppRevokeLicensesActionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDateTime sets the startDateTime property value. Time the action was initiated
func (m *IosVppAppRevokeLicensesActionResult) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetTotalLicensesCount sets the totalLicensesCount property value. A count of the number of licenses for which revoke was attempted.
func (m *IosVppAppRevokeLicensesActionResult) SetTotalLicensesCount(value *int32)() {
    m.totalLicensesCount = value
}
// SetUserId sets the userId property value. UserId associated with the action.
func (m *IosVppAppRevokeLicensesActionResult) SetUserId(value *string)() {
    m.userId = value
}
