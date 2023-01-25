package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityActionState 
type SecurityActionState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be extracted from the auth token and not entered manually by the calling application.
    appId *string
    // The OdataType property
    odataType *string
    // Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
    status *OperationStatus
    // Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be extracted from the auth token and not entered manually by the calling application.
    user *string
}
// NewSecurityActionState instantiates a new securityActionState and sets the default values.
func NewSecurityActionState()(*SecurityActionState) {
    m := &SecurityActionState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityActionStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityActionStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityActionState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityActionState) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppId gets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityActionState) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityActionState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*OperationStatus))
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SecurityActionState) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
func (m *SecurityActionState) GetStatus()(*OperationStatus) {
    return m.status
}
// GetUpdatedDateTime gets the updatedDateTime property value. Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *SecurityActionState) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedDateTime
}
// GetUser gets the user property value. The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityActionState) GetUser()(*string) {
    return m.user
}
// Serialize serializes information the current object
func (m *SecurityActionState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *SecurityActionState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppId sets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityActionState) SetAppId(value *string)() {
    m.appId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SecurityActionState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
func (m *SecurityActionState) SetStatus(value *OperationStatus)() {
    m.status = value
}
// SetUpdatedDateTime sets the updatedDateTime property value. Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *SecurityActionState) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedDateTime = value
}
// SetUser sets the user property value. The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityActionState) SetUser(value *string)() {
    m.user = value
}
