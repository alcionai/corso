package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityAction provides operations to manage the collection of site entities.
type SecurityAction struct {
    Entity
    // Reason for invoking this action.
    actionReason *string
    // The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from the auth token and not entered manually by the calling application.
    appId *string
    // Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The azureTenantId should be extracted from the auth token and not entered manually by the calling application.
    azureTenantId *string
    // The clientContext property
    clientContext *string
    // Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Error info when the action fails.
    errorInfo ResultInfoable
    // Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Action name.
    name *string
    // Collection of parameters (key-value pairs) necessary to invoke the action, for example, URL or fileHash to block.). Required.
    parameters []KeyValuePairable
    // Collection of securityActionState to keep the history of an action.
    states []SecurityActionStateable
    // Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
    status *OperationStatus
    // The user principal name of the signed-in user that submitted  (POST) the action. The user should be extracted from the auth token and not entered manually by the calling application.
    user *string
    // Complex Type containing details about the Security product/service vendor, provider, and sub-provider (for example, vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
    vendorInformation SecurityVendorInformationable
}
// NewSecurityAction instantiates a new securityAction and sets the default values.
func NewSecurityAction()(*SecurityAction) {
    m := &SecurityAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityAction(), nil
}
// GetActionReason gets the actionReason property value. Reason for invoking this action.
func (m *SecurityAction) GetActionReason()(*string) {
    return m.actionReason
}
// GetAppId gets the appId property value. The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) GetAppId()(*string) {
    return m.appId
}
// GetAzureTenantId gets the azureTenantId property value. Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The azureTenantId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *SecurityAction) GetClientContext()(*string) {
    return m.clientContext
}
// GetCompletedDateTime gets the completedDateTime property value. Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SecurityAction) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SecurityAction) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetErrorInfo gets the errorInfo property value. Error info when the action fails.
func (m *SecurityAction) GetErrorInfo()(ResultInfoable) {
    return m.errorInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionReason(val)
        }
        return nil
    }
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
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
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
    res["errorInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorInfo(val.(ResultInfoable))
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["states"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityActionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityActionStateable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityActionStateable)
            }
            m.SetStates(res)
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
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    return res
}
// GetLastActionDateTime gets the lastActionDateTime property value. Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SecurityAction) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetName gets the name property value. Action name.
func (m *SecurityAction) GetName()(*string) {
    return m.name
}
// GetParameters gets the parameters property value. Collection of parameters (key-value pairs) necessary to invoke the action, for example, URL or fileHash to block.). Required.
func (m *SecurityAction) GetParameters()([]KeyValuePairable) {
    return m.parameters
}
// GetStates gets the states property value. Collection of securityActionState to keep the history of an action.
func (m *SecurityAction) GetStates()([]SecurityActionStateable) {
    return m.states
}
// GetStatus gets the status property value. Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
func (m *SecurityAction) GetStatus()(*OperationStatus) {
    return m.status
}
// GetUser gets the user property value. The user principal name of the signed-in user that submitted  (POST) the action. The user should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) GetUser()(*string) {
    return m.user
}
// GetVendorInformation gets the vendorInformation property value. Complex Type containing details about the Security product/service vendor, provider, and sub-provider (for example, vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
func (m *SecurityAction) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
}
// Serialize serializes information the current object
func (m *SecurityAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionReason", m.GetActionReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
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
        err = writer.WriteObjectValue("errorInfo", m.GetErrorInfo())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStates()))
        for i, v := range m.GetStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("states", cast)
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
        err = writer.WriteStringValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionReason sets the actionReason property value. Reason for invoking this action.
func (m *SecurityAction) SetActionReason(value *string)() {
    m.actionReason = value
}
// SetAppId sets the appId property value. The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) SetAppId(value *string)() {
    m.appId = value
}
// SetAzureTenantId sets the azureTenantId property value. Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The azureTenantId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *SecurityAction) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetCompletedDateTime sets the completedDateTime property value. Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SecurityAction) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SecurityAction) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetErrorInfo sets the errorInfo property value. Error info when the action fails.
func (m *SecurityAction) SetErrorInfo(value ResultInfoable)() {
    m.errorInfo = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SecurityAction) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetName sets the name property value. Action name.
func (m *SecurityAction) SetName(value *string)() {
    m.name = value
}
// SetParameters sets the parameters property value. Collection of parameters (key-value pairs) necessary to invoke the action, for example, URL or fileHash to block.). Required.
func (m *SecurityAction) SetParameters(value []KeyValuePairable)() {
    m.parameters = value
}
// SetStates sets the states property value. Collection of securityActionState to keep the history of an action.
func (m *SecurityAction) SetStates(value []SecurityActionStateable)() {
    m.states = value
}
// SetStatus sets the status property value. Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
func (m *SecurityAction) SetStatus(value *OperationStatus)() {
    m.status = value
}
// SetUser sets the user property value. The user principal name of the signed-in user that submitted  (POST) the action. The user should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityAction) SetUser(value *string)() {
    m.user = value
}
// SetVendorInformation sets the vendorInformation property value. Complex Type containing details about the Security product/service vendor, provider, and sub-provider (for example, vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
func (m *SecurityAction) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
