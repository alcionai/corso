package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateConnectorDetails entity used to retrieve information about Intune Certificate Connectors.
type CertificateConnectorDetails struct {
    Entity
    // Connector name (set during enrollment).
    connectorName *string
    // Version of the connector installed.
    connectorVersion *string
    // Date/time when this connector was enrolled.
    enrollmentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date/time when this connector last connected to the service.
    lastCheckinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the machine hosting this connector service.
    machineName *string
}
// NewCertificateConnectorDetails instantiates a new certificateConnectorDetails and sets the default values.
func NewCertificateConnectorDetails()(*CertificateConnectorDetails) {
    m := &CertificateConnectorDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCertificateConnectorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorDetails(), nil
}
// GetConnectorName gets the connectorName property value. Connector name (set during enrollment).
func (m *CertificateConnectorDetails) GetConnectorName()(*string) {
    return m.connectorName
}
// GetConnectorVersion gets the connectorVersion property value. Version of the connector installed.
func (m *CertificateConnectorDetails) GetConnectorVersion()(*string) {
    return m.connectorVersion
}
// GetEnrollmentDateTime gets the enrollmentDateTime property value. Date/time when this connector was enrolled.
func (m *CertificateConnectorDetails) GetEnrollmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrollmentDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorName(val)
        }
        return nil
    }
    res["connectorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorVersion(val)
        }
        return nil
    }
    res["enrollmentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentDateTime(val)
        }
        return nil
    }
    res["lastCheckinDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckinDateTime(val)
        }
        return nil
    }
    res["machineName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineName(val)
        }
        return nil
    }
    return res
}
// GetLastCheckinDateTime gets the lastCheckinDateTime property value. Date/time when this connector last connected to the service.
func (m *CertificateConnectorDetails) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckinDateTime
}
// GetMachineName gets the machineName property value. Name of the machine hosting this connector service.
func (m *CertificateConnectorDetails) GetMachineName()(*string) {
    return m.machineName
}
// Serialize serializes information the current object
func (m *CertificateConnectorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectorName", m.GetConnectorName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectorVersion", m.GetConnectorVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrollmentDateTime", m.GetEnrollmentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckinDateTime", m.GetLastCheckinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("machineName", m.GetMachineName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectorName sets the connectorName property value. Connector name (set during enrollment).
func (m *CertificateConnectorDetails) SetConnectorName(value *string)() {
    m.connectorName = value
}
// SetConnectorVersion sets the connectorVersion property value. Version of the connector installed.
func (m *CertificateConnectorDetails) SetConnectorVersion(value *string)() {
    m.connectorVersion = value
}
// SetEnrollmentDateTime sets the enrollmentDateTime property value. Date/time when this connector was enrolled.
func (m *CertificateConnectorDetails) SetEnrollmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrollmentDateTime = value
}
// SetLastCheckinDateTime sets the lastCheckinDateTime property value. Date/time when this connector last connected to the service.
func (m *CertificateConnectorDetails) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckinDateTime = value
}
// SetMachineName sets the machineName property value. Name of the machine hosting this connector service.
func (m *CertificateConnectorDetails) SetMachineName(value *string)() {
    m.machineName = value
}
