package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPCConnectivityIssue the user experience analyte connectivity issue entity.
type CloudPCConnectivityIssue struct {
    Entity
    // The Intune DeviceId of the device the connection is associated with.
    deviceId *string
    // The error code of the connectivity issue.
    errorCode *string
    // The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC) time.
    errorDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The detailed description of what went wrong.
    errorDescription *string
    // The recommended action to fix the corresponding error.
    recommendedAction *string
    // The unique id of user who initialize the connection.
    userId *string
}
// NewCloudPCConnectivityIssue instantiates a new cloudPCConnectivityIssue and sets the default values.
func NewCloudPCConnectivityIssue()(*CloudPCConnectivityIssue) {
    m := &CloudPCConnectivityIssue{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPCConnectivityIssueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCConnectivityIssueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPCConnectivityIssue(), nil
}
// GetDeviceId gets the deviceId property value. The Intune DeviceId of the device the connection is associated with.
func (m *CloudPCConnectivityIssue) GetDeviceId()(*string) {
    return m.deviceId
}
// GetErrorCode gets the errorCode property value. The error code of the connectivity issue.
func (m *CloudPCConnectivityIssue) GetErrorCode()(*string) {
    return m.errorCode
}
// GetErrorDateTime gets the errorDateTime property value. The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC) time.
func (m *CloudPCConnectivityIssue) GetErrorDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.errorDateTime
}
// GetErrorDescription gets the errorDescription property value. The detailed description of what went wrong.
func (m *CloudPCConnectivityIssue) GetErrorDescription()(*string) {
    return m.errorDescription
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPCConnectivityIssue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDateTime(val)
        }
        return nil
    }
    res["errorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
        }
        return nil
    }
    res["recommendedAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedAction(val)
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
// GetRecommendedAction gets the recommendedAction property value. The recommended action to fix the corresponding error.
func (m *CloudPCConnectivityIssue) GetRecommendedAction()(*string) {
    return m.recommendedAction
}
// GetUserId gets the userId property value. The unique id of user who initialize the connection.
func (m *CloudPCConnectivityIssue) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *CloudPCConnectivityIssue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("errorDateTime", m.GetErrorDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorDescription", m.GetErrorDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The Intune DeviceId of the device the connection is associated with.
func (m *CloudPCConnectivityIssue) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetErrorCode sets the errorCode property value. The error code of the connectivity issue.
func (m *CloudPCConnectivityIssue) SetErrorCode(value *string)() {
    m.errorCode = value
}
// SetErrorDateTime sets the errorDateTime property value. The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC) time.
func (m *CloudPCConnectivityIssue) SetErrorDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.errorDateTime = value
}
// SetErrorDescription sets the errorDescription property value. The detailed description of what went wrong.
func (m *CloudPCConnectivityIssue) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
// SetRecommendedAction sets the recommendedAction property value. The recommended action to fix the corresponding error.
func (m *CloudPCConnectivityIssue) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
// SetUserId sets the userId property value. The unique id of user who initialize the connection.
func (m *CloudPCConnectivityIssue) SetUserId(value *string)() {
    m.userId = value
}
