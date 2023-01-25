package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoggedOnUser logged On User
type LoggedOnUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Date time when user logs on
    lastLogOnDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // User id
    userId *string
}
// NewLoggedOnUser instantiates a new loggedOnUser and sets the default values.
func NewLoggedOnUser()(*LoggedOnUser) {
    m := &LoggedOnUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLoggedOnUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoggedOnUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoggedOnUser(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoggedOnUser) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoggedOnUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lastLogOnDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastLogOnDateTime(val)
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
// GetLastLogOnDateTime gets the lastLogOnDateTime property value. Date time when user logs on
func (m *LoggedOnUser) GetLastLogOnDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastLogOnDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LoggedOnUser) GetOdataType()(*string) {
    return m.odataType
}
// GetUserId gets the userId property value. User id
func (m *LoggedOnUser) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *LoggedOnUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastLogOnDateTime", m.GetLastLogOnDateTime())
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
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *LoggedOnUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLastLogOnDateTime sets the lastLogOnDateTime property value. Date time when user logs on
func (m *LoggedOnUser) SetLastLogOnDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastLogOnDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LoggedOnUser) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserId sets the userId property value. User id
func (m *LoggedOnUser) SetUserId(value *string)() {
    m.userId = value
}
