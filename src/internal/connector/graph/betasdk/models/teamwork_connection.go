package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkConnection 
type TeamworkConnection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are: unknown, connected, disconnected, unknownFutureValue.
    connectionStatus *TeamworkConnectionStatus
    // Time at which the state was last changed. For example, indicates connected since when the state is connected and disconnected since when the state is disconnected.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewTeamworkConnection instantiates a new teamworkConnection and sets the default values.
func NewTeamworkConnection()(*TeamworkConnection) {
    m := &TeamworkConnection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkConnection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkConnection) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConnectionStatus gets the connectionStatus property value. Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are: unknown, connected, disconnected, unknownFutureValue.
func (m *TeamworkConnection) GetConnectionStatus()(*TeamworkConnectionStatus) {
    return m.connectionStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionStatus(val.(*TeamworkConnectionStatus))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time at which the state was last changed. For example, indicates connected since when the state is connected and disconnected since when the state is disconnected.
func (m *TeamworkConnection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkConnection) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConnectionStatus() != nil {
        cast := (*m.GetConnectionStatus()).String()
        err := writer.WriteStringValue("connectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
func (m *TeamworkConnection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConnectionStatus sets the connectionStatus property value. Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are: unknown, connected, disconnected, unknownFutureValue.
func (m *TeamworkConnection) SetConnectionStatus(value *TeamworkConnectionStatus)() {
    m.connectionStatus = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time at which the state was last changed. For example, indicates connected since when the state is connected and disconnected since when the state is disconnected.
func (m *TeamworkConnection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkConnection) SetOdataType(value *string)() {
    m.odataType = value
}
