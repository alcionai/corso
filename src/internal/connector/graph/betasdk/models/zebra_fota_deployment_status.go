package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeploymentStatus describes the status for a single FOTA deployment.
type ZebraFotaDeploymentStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A boolean that indicates if a cancellation was requested on the deployment. NOTE: A cancellation request does not guarantee that the deployment was canceled.
    cancelRequested *bool
    // The date and time when this deployment was completed or canceled. The actual date time is determined by the value of state. If the state is canceled, this property holds the cancellation date/time. If the the state is completed, this property holds the completion date/time. If the deployment is not completed before the deployment end date, then completed date/time and end date/time are the same. This is always in the deployment timezone. Note: An installation that is in progress can continue past the deployment end date.
    completeOrCanceledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time when the deployment status was updated from Zebra
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Represents the state of Zebra FOTA deployment.
    state *ZebraFotaDeploymentState
    // An integer that indicates the total number of devices where installation was successful.
    totalAwaitingInstall *int32
    // An integer that indicates the total number of devices where installation was canceled.
    totalCanceled *int32
    // An integer that indicates the total number of devices that have a job in the CREATED state. Typically indicates jobs that did not reach the devices.
    totalCreated *int32
    // An integer that indicates the total number of devices in the deployment.
    totalDevices *int32
    // An integer that indicates the total number of devices where installation was successful.
    totalDownloading *int32
    // An integer that indicates the total number of devices that have failed to download the new OS file.
    totalFailedDownload *int32
    // An integer that indicates the total number of devices that have failed to install the new OS file.
    totalFailedInstall *int32
    // An integer that indicates the total number of devices that received the json and are scheduled.
    totalScheduled *int32
    // An integer that indicates the total number of devices where installation was successful.
    totalSucceededInstall *int32
    // An integer that indicates the total number of devices where no deployment status or end state has not received, even after the scheduled end date was reached.
    totalUnknown *int32
}
// NewZebraFotaDeploymentStatus instantiates a new zebraFotaDeploymentStatus and sets the default values.
func NewZebraFotaDeploymentStatus()(*ZebraFotaDeploymentStatus) {
    m := &ZebraFotaDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateZebraFotaDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaDeploymentStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ZebraFotaDeploymentStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCancelRequested gets the cancelRequested property value. A boolean that indicates if a cancellation was requested on the deployment. NOTE: A cancellation request does not guarantee that the deployment was canceled.
func (m *ZebraFotaDeploymentStatus) GetCancelRequested()(*bool) {
    return m.cancelRequested
}
// GetCompleteOrCanceledDateTime gets the completeOrCanceledDateTime property value. The date and time when this deployment was completed or canceled. The actual date time is determined by the value of state. If the state is canceled, this property holds the cancellation date/time. If the the state is completed, this property holds the completion date/time. If the deployment is not completed before the deployment end date, then completed date/time and end date/time are the same. This is always in the deployment timezone. Note: An installation that is in progress can continue past the deployment end date.
func (m *ZebraFotaDeploymentStatus) GetCompleteOrCanceledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completeOrCanceledDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cancelRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelRequested(val)
        }
        return nil
    }
    res["completeOrCanceledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompleteOrCanceledDateTime(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaDeploymentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ZebraFotaDeploymentState))
        }
        return nil
    }
    res["totalAwaitingInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAwaitingInstall(val)
        }
        return nil
    }
    res["totalCanceled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCanceled(val)
        }
        return nil
    }
    res["totalCreated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCreated(val)
        }
        return nil
    }
    res["totalDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDevices(val)
        }
        return nil
    }
    res["totalDownloading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDownloading(val)
        }
        return nil
    }
    res["totalFailedDownload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFailedDownload(val)
        }
        return nil
    }
    res["totalFailedInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalFailedInstall(val)
        }
        return nil
    }
    res["totalScheduled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalScheduled(val)
        }
        return nil
    }
    res["totalSucceededInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSucceededInstall(val)
        }
        return nil
    }
    res["totalUnknown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUnknown(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Date and time when the deployment status was updated from Zebra
func (m *ZebraFotaDeploymentStatus) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. Represents the state of Zebra FOTA deployment.
func (m *ZebraFotaDeploymentStatus) GetState()(*ZebraFotaDeploymentState) {
    return m.state
}
// GetTotalAwaitingInstall gets the totalAwaitingInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) GetTotalAwaitingInstall()(*int32) {
    return m.totalAwaitingInstall
}
// GetTotalCanceled gets the totalCanceled property value. An integer that indicates the total number of devices where installation was canceled.
func (m *ZebraFotaDeploymentStatus) GetTotalCanceled()(*int32) {
    return m.totalCanceled
}
// GetTotalCreated gets the totalCreated property value. An integer that indicates the total number of devices that have a job in the CREATED state. Typically indicates jobs that did not reach the devices.
func (m *ZebraFotaDeploymentStatus) GetTotalCreated()(*int32) {
    return m.totalCreated
}
// GetTotalDevices gets the totalDevices property value. An integer that indicates the total number of devices in the deployment.
func (m *ZebraFotaDeploymentStatus) GetTotalDevices()(*int32) {
    return m.totalDevices
}
// GetTotalDownloading gets the totalDownloading property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) GetTotalDownloading()(*int32) {
    return m.totalDownloading
}
// GetTotalFailedDownload gets the totalFailedDownload property value. An integer that indicates the total number of devices that have failed to download the new OS file.
func (m *ZebraFotaDeploymentStatus) GetTotalFailedDownload()(*int32) {
    return m.totalFailedDownload
}
// GetTotalFailedInstall gets the totalFailedInstall property value. An integer that indicates the total number of devices that have failed to install the new OS file.
func (m *ZebraFotaDeploymentStatus) GetTotalFailedInstall()(*int32) {
    return m.totalFailedInstall
}
// GetTotalScheduled gets the totalScheduled property value. An integer that indicates the total number of devices that received the json and are scheduled.
func (m *ZebraFotaDeploymentStatus) GetTotalScheduled()(*int32) {
    return m.totalScheduled
}
// GetTotalSucceededInstall gets the totalSucceededInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) GetTotalSucceededInstall()(*int32) {
    return m.totalSucceededInstall
}
// GetTotalUnknown gets the totalUnknown property value. An integer that indicates the total number of devices where no deployment status or end state has not received, even after the scheduled end date was reached.
func (m *ZebraFotaDeploymentStatus) GetTotalUnknown()(*int32) {
    return m.totalUnknown
}
// Serialize serializes information the current object
func (m *ZebraFotaDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("cancelRequested", m.GetCancelRequested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completeOrCanceledDateTime", m.GetCompleteOrCanceledDateTime())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalAwaitingInstall", m.GetTotalAwaitingInstall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCanceled", m.GetTotalCanceled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCreated", m.GetTotalCreated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalDevices", m.GetTotalDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalDownloading", m.GetTotalDownloading())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalFailedDownload", m.GetTotalFailedDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalFailedInstall", m.GetTotalFailedInstall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalScheduled", m.GetTotalScheduled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalSucceededInstall", m.GetTotalSucceededInstall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalUnknown", m.GetTotalUnknown())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ZebraFotaDeploymentStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCancelRequested sets the cancelRequested property value. A boolean that indicates if a cancellation was requested on the deployment. NOTE: A cancellation request does not guarantee that the deployment was canceled.
func (m *ZebraFotaDeploymentStatus) SetCancelRequested(value *bool)() {
    m.cancelRequested = value
}
// SetCompleteOrCanceledDateTime sets the completeOrCanceledDateTime property value. The date and time when this deployment was completed or canceled. The actual date time is determined by the value of state. If the state is canceled, this property holds the cancellation date/time. If the the state is completed, this property holds the completion date/time. If the deployment is not completed before the deployment end date, then completed date/time and end date/time are the same. This is always in the deployment timezone. Note: An installation that is in progress can continue past the deployment end date.
func (m *ZebraFotaDeploymentStatus) SetCompleteOrCanceledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completeOrCanceledDateTime = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Date and time when the deployment status was updated from Zebra
func (m *ZebraFotaDeploymentStatus) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ZebraFotaDeploymentStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. Represents the state of Zebra FOTA deployment.
func (m *ZebraFotaDeploymentStatus) SetState(value *ZebraFotaDeploymentState)() {
    m.state = value
}
// SetTotalAwaitingInstall sets the totalAwaitingInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) SetTotalAwaitingInstall(value *int32)() {
    m.totalAwaitingInstall = value
}
// SetTotalCanceled sets the totalCanceled property value. An integer that indicates the total number of devices where installation was canceled.
func (m *ZebraFotaDeploymentStatus) SetTotalCanceled(value *int32)() {
    m.totalCanceled = value
}
// SetTotalCreated sets the totalCreated property value. An integer that indicates the total number of devices that have a job in the CREATED state. Typically indicates jobs that did not reach the devices.
func (m *ZebraFotaDeploymentStatus) SetTotalCreated(value *int32)() {
    m.totalCreated = value
}
// SetTotalDevices sets the totalDevices property value. An integer that indicates the total number of devices in the deployment.
func (m *ZebraFotaDeploymentStatus) SetTotalDevices(value *int32)() {
    m.totalDevices = value
}
// SetTotalDownloading sets the totalDownloading property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) SetTotalDownloading(value *int32)() {
    m.totalDownloading = value
}
// SetTotalFailedDownload sets the totalFailedDownload property value. An integer that indicates the total number of devices that have failed to download the new OS file.
func (m *ZebraFotaDeploymentStatus) SetTotalFailedDownload(value *int32)() {
    m.totalFailedDownload = value
}
// SetTotalFailedInstall sets the totalFailedInstall property value. An integer that indicates the total number of devices that have failed to install the new OS file.
func (m *ZebraFotaDeploymentStatus) SetTotalFailedInstall(value *int32)() {
    m.totalFailedInstall = value
}
// SetTotalScheduled sets the totalScheduled property value. An integer that indicates the total number of devices that received the json and are scheduled.
func (m *ZebraFotaDeploymentStatus) SetTotalScheduled(value *int32)() {
    m.totalScheduled = value
}
// SetTotalSucceededInstall sets the totalSucceededInstall property value. An integer that indicates the total number of devices where installation was successful.
func (m *ZebraFotaDeploymentStatus) SetTotalSucceededInstall(value *int32)() {
    m.totalSucceededInstall = value
}
// SetTotalUnknown sets the totalUnknown property value. An integer that indicates the total number of devices where no deployment status or end state has not received, even after the scheduled end date was reached.
func (m *ZebraFotaDeploymentStatus) SetTotalUnknown(value *int32)() {
    m.totalUnknown = value
}
