package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateProfileInventorySyncStatus a complex type to store the status of a driver and firmware profile inventory sync. The status includes the last successful sync date time and the state of the last sync.
type WindowsDriverUpdateProfileInventorySyncStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Windows DnF update inventory sync state.
    driverInventorySyncState *WindowsDriverUpdateProfileInventorySyncState
    // The last successful sync date and time in UTC.
    lastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewWindowsDriverUpdateProfileInventorySyncStatus instantiates a new windowsDriverUpdateProfileInventorySyncStatus and sets the default values.
func NewWindowsDriverUpdateProfileInventorySyncStatus()(*WindowsDriverUpdateProfileInventorySyncStatus) {
    m := &WindowsDriverUpdateProfileInventorySyncStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsDriverUpdateProfileInventorySyncStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDriverUpdateProfileInventorySyncStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateProfileInventorySyncStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDriverInventorySyncState gets the driverInventorySyncState property value. Windows DnF update inventory sync state.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetDriverInventorySyncState()(*WindowsDriverUpdateProfileInventorySyncState) {
    return m.driverInventorySyncState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["driverInventorySyncState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDriverUpdateProfileInventorySyncState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriverInventorySyncState(val.(*WindowsDriverUpdateProfileInventorySyncState))
        }
        return nil
    }
    res["lastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulSyncDateTime(val)
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
    return res
}
// GetLastSuccessfulSyncDateTime gets the lastSuccessfulSyncDateTime property value. The last successful sync date and time in UTC.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSuccessfulSyncDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateProfileInventorySyncStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDriverInventorySyncState() != nil {
        cast := (*m.GetDriverInventorySyncState()).String()
        err := writer.WriteStringValue("driverInventorySyncState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSuccessfulSyncDateTime", m.GetLastSuccessfulSyncDateTime())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDriverInventorySyncState sets the driverInventorySyncState property value. Windows DnF update inventory sync state.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetDriverInventorySyncState(value *WindowsDriverUpdateProfileInventorySyncState)() {
    m.driverInventorySyncState = value
}
// SetLastSuccessfulSyncDateTime sets the lastSuccessfulSyncDateTime property value. The last successful sync date and time in UTC.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSuccessfulSyncDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetOdataType(value *string)() {
    m.odataType = value
}
