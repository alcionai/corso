package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelServerLogCollectionResponse entity that stores the server log collection status.
type MicrosoftTunnelServerLogCollectionResponse struct {
    Entity
    // The end time of the logs collected
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time when the log collection is expired
    expiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time when the log collection was requested
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // ID of the server the log collection is requested upon
    serverId *string
    // The size of the logs in bytes
    sizeInBytes *int64
    // The start time of the logs collected
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Enum type that represent the status of log collection
    status *MicrosoftTunnelLogCollectionStatus
}
// NewMicrosoftTunnelServerLogCollectionResponse instantiates a new microsoftTunnelServerLogCollectionResponse and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponse()(*MicrosoftTunnelServerLogCollectionResponse) {
    m := &MicrosoftTunnelServerLogCollectionResponse{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelServerLogCollectionResponse(), nil
}
// GetEndDateTime gets the endDateTime property value. The end time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExpiryDateTime gets the expiryDateTime property value. The time when the log collection is expired
func (m *MicrosoftTunnelServerLogCollectionResponse) GetExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiryDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelServerLogCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["expiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDateTime(val)
        }
        return nil
    }
    res["requestDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
        }
        return nil
    }
    res["serverId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerId(val)
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftTunnelLogCollectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MicrosoftTunnelLogCollectionStatus))
        }
        return nil
    }
    return res
}
// GetRequestDateTime gets the requestDateTime property value. The time when the log collection was requested
func (m *MicrosoftTunnelServerLogCollectionResponse) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestDateTime
}
// GetServerId gets the serverId property value. ID of the server the log collection is requested upon
func (m *MicrosoftTunnelServerLogCollectionResponse) GetServerId()(*string) {
    return m.serverId
}
// GetSizeInBytes gets the sizeInBytes property value. The size of the logs in bytes
func (m *MicrosoftTunnelServerLogCollectionResponse) GetSizeInBytes()(*int64) {
    return m.sizeInBytes
}
// GetStartDateTime gets the startDateTime property value. The start time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. Enum type that represent the status of log collection
func (m *MicrosoftTunnelServerLogCollectionResponse) GetStatus()(*MicrosoftTunnelLogCollectionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelServerLogCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiryDateTime", m.GetExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverId", m.GetServerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
    return nil
}
// SetEndDateTime sets the endDateTime property value. The end time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExpiryDateTime sets the expiryDateTime property value. The time when the log collection is expired
func (m *MicrosoftTunnelServerLogCollectionResponse) SetExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryDateTime = value
}
// SetRequestDateTime sets the requestDateTime property value. The time when the log collection was requested
func (m *MicrosoftTunnelServerLogCollectionResponse) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
// SetServerId sets the serverId property value. ID of the server the log collection is requested upon
func (m *MicrosoftTunnelServerLogCollectionResponse) SetServerId(value *string)() {
    m.serverId = value
}
// SetSizeInBytes sets the sizeInBytes property value. The size of the logs in bytes
func (m *MicrosoftTunnelServerLogCollectionResponse) SetSizeInBytes(value *int64)() {
    m.sizeInBytes = value
}
// SetStartDateTime sets the startDateTime property value. The start time of the logs collected
func (m *MicrosoftTunnelServerLogCollectionResponse) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. Enum type that represent the status of log collection
func (m *MicrosoftTunnelServerLogCollectionResponse) SetStatus(value *MicrosoftTunnelLogCollectionStatus)() {
    m.status = value
}
