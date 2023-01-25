package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcExternalPartnerSetting 
type CloudPcExternalPartnerSetting struct {
    Entity
    // Enable or disable the connection to an external partner. If true, an external partner API will accept incoming calls from external partners. Required. Supports $filter (eq).
    enableConnection *bool
    // Last data sync time for this external partner. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The external partner ID.
    partnerId *string
    // The status property
    status *CloudPcExternalPartnerStatus
    // Status details message.
    statusDetails *string
}
// NewCloudPcExternalPartnerSetting instantiates a new CloudPcExternalPartnerSetting and sets the default values.
func NewCloudPcExternalPartnerSetting()(*CloudPcExternalPartnerSetting) {
    m := &CloudPcExternalPartnerSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcExternalPartnerSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcExternalPartnerSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcExternalPartnerSetting(), nil
}
// GetEnableConnection gets the enableConnection property value. Enable or disable the connection to an external partner. If true, an external partner API will accept incoming calls from external partners. Required. Supports $filter (eq).
func (m *CloudPcExternalPartnerSetting) GetEnableConnection()(*bool) {
    return m.enableConnection
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcExternalPartnerSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enableConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableConnection(val)
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
    res["partnerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcExternalPartnerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcExternalPartnerStatus))
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val)
        }
        return nil
    }
    return res
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last data sync time for this external partner. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
func (m *CloudPcExternalPartnerSetting) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetPartnerId gets the partnerId property value. The external partner ID.
func (m *CloudPcExternalPartnerSetting) GetPartnerId()(*string) {
    return m.partnerId
}
// GetStatus gets the status property value. The status property
func (m *CloudPcExternalPartnerSetting) GetStatus()(*CloudPcExternalPartnerStatus) {
    return m.status
}
// GetStatusDetails gets the statusDetails property value. Status details message.
func (m *CloudPcExternalPartnerSetting) GetStatusDetails()(*string) {
    return m.statusDetails
}
// Serialize serializes information the current object
func (m *CloudPcExternalPartnerSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enableConnection", m.GetEnableConnection())
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
    {
        err = writer.WriteStringValue("partnerId", m.GetPartnerId())
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
        err = writer.WriteStringValue("statusDetails", m.GetStatusDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnableConnection sets the enableConnection property value. Enable or disable the connection to an external partner. If true, an external partner API will accept incoming calls from external partners. Required. Supports $filter (eq).
func (m *CloudPcExternalPartnerSetting) SetEnableConnection(value *bool)() {
    m.enableConnection = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last data sync time for this external partner. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
func (m *CloudPcExternalPartnerSetting) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetPartnerId sets the partnerId property value. The external partner ID.
func (m *CloudPcExternalPartnerSetting) SetPartnerId(value *string)() {
    m.partnerId = value
}
// SetStatus sets the status property value. The status property
func (m *CloudPcExternalPartnerSetting) SetStatus(value *CloudPcExternalPartnerStatus)() {
    m.status = value
}
// SetStatusDetails sets the statusDetails property value. Status details message.
func (m *CloudPcExternalPartnerSetting) SetStatusDetails(value *string)() {
    m.statusDetails = value
}
