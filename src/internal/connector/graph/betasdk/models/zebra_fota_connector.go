package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaConnector 
type ZebraFotaConnector struct {
    Entity
    // Complete account enrollment authorization URL. This corresponds to verification_uri_complete in the Zebra API documentations.
    enrollmentAuthorizationUrl *string
    // Tenant enrollment token from Zebra. The token is used to enroll Zebra devices in the FOTA Service via app config.
    enrollmentToken *string
    // Flag indicating if required Firmware Over-the-Air (FOTA) Apps have been approved.
    fotaAppsApproved *bool
    // Date and time when the account was last synched with Zebra
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Represents various states for Zebra FOTA connector.
    state *ZebraFotaConnectorState
}
// NewZebraFotaConnector instantiates a new zebraFotaConnector and sets the default values.
func NewZebraFotaConnector()(*ZebraFotaConnector) {
    m := &ZebraFotaConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateZebraFotaConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaConnector(), nil
}
// GetEnrollmentAuthorizationUrl gets the enrollmentAuthorizationUrl property value. Complete account enrollment authorization URL. This corresponds to verification_uri_complete in the Zebra API documentations.
func (m *ZebraFotaConnector) GetEnrollmentAuthorizationUrl()(*string) {
    return m.enrollmentAuthorizationUrl
}
// GetEnrollmentToken gets the enrollmentToken property value. Tenant enrollment token from Zebra. The token is used to enroll Zebra devices in the FOTA Service via app config.
func (m *ZebraFotaConnector) GetEnrollmentToken()(*string) {
    return m.enrollmentToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enrollmentAuthorizationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentAuthorizationUrl(val)
        }
        return nil
    }
    res["enrollmentToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentToken(val)
        }
        return nil
    }
    res["fotaAppsApproved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFotaAppsApproved(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseZebraFotaConnectorState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ZebraFotaConnectorState))
        }
        return nil
    }
    return res
}
// GetFotaAppsApproved gets the fotaAppsApproved property value. Flag indicating if required Firmware Over-the-Air (FOTA) Apps have been approved.
func (m *ZebraFotaConnector) GetFotaAppsApproved()(*bool) {
    return m.fotaAppsApproved
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Date and time when the account was last synched with Zebra
func (m *ZebraFotaConnector) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetState gets the state property value. Represents various states for Zebra FOTA connector.
func (m *ZebraFotaConnector) GetState()(*ZebraFotaConnectorState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ZebraFotaConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("enrollmentAuthorizationUrl", m.GetEnrollmentAuthorizationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentToken", m.GetEnrollmentToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fotaAppsApproved", m.GetFotaAppsApproved())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
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
// SetEnrollmentAuthorizationUrl sets the enrollmentAuthorizationUrl property value. Complete account enrollment authorization URL. This corresponds to verification_uri_complete in the Zebra API documentations.
func (m *ZebraFotaConnector) SetEnrollmentAuthorizationUrl(value *string)() {
    m.enrollmentAuthorizationUrl = value
}
// SetEnrollmentToken sets the enrollmentToken property value. Tenant enrollment token from Zebra. The token is used to enroll Zebra devices in the FOTA Service via app config.
func (m *ZebraFotaConnector) SetEnrollmentToken(value *string)() {
    m.enrollmentToken = value
}
// SetFotaAppsApproved sets the fotaAppsApproved property value. Flag indicating if required Firmware Over-the-Air (FOTA) Apps have been approved.
func (m *ZebraFotaConnector) SetFotaAppsApproved(value *bool)() {
    m.fotaAppsApproved = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Date and time when the account was last synched with Zebra
func (m *ZebraFotaConnector) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetState sets the state property value. Represents various states for Zebra FOTA connector.
func (m *ZebraFotaConnector) SetState(value *ZebraFotaConnectorState)() {
    m.state = value
}
