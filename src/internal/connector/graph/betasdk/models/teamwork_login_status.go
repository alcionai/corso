package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkLoginStatus 
type TeamworkLoginStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Information about the Exchange connection.
    exchangeConnection TeamworkConnectionable
    // The OdataType property
    odataType *string
    // Information about the Skype for Business connection.
    skypeConnection TeamworkConnectionable
    // Information about the Teams connection.
    teamsConnection TeamworkConnectionable
}
// NewTeamworkLoginStatus instantiates a new teamworkLoginStatus and sets the default values.
func NewTeamworkLoginStatus()(*TeamworkLoginStatus) {
    m := &TeamworkLoginStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkLoginStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkLoginStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkLoginStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkLoginStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExchangeConnection gets the exchangeConnection property value. Information about the Exchange connection.
func (m *TeamworkLoginStatus) GetExchangeConnection()(TeamworkConnectionable) {
    return m.exchangeConnection
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkLoginStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exchangeConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeConnection(val.(TeamworkConnectionable))
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
    res["skypeConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeConnection(val.(TeamworkConnectionable))
        }
        return nil
    }
    res["teamsConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsConnection(val.(TeamworkConnectionable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkLoginStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetSkypeConnection gets the skypeConnection property value. Information about the Skype for Business connection.
func (m *TeamworkLoginStatus) GetSkypeConnection()(TeamworkConnectionable) {
    return m.skypeConnection
}
// GetTeamsConnection gets the teamsConnection property value. Information about the Teams connection.
func (m *TeamworkLoginStatus) GetTeamsConnection()(TeamworkConnectionable) {
    return m.teamsConnection
}
// Serialize serializes information the current object
func (m *TeamworkLoginStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("exchangeConnection", m.GetExchangeConnection())
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
        err := writer.WriteObjectValue("skypeConnection", m.GetSkypeConnection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("teamsConnection", m.GetTeamsConnection())
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
func (m *TeamworkLoginStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExchangeConnection sets the exchangeConnection property value. Information about the Exchange connection.
func (m *TeamworkLoginStatus) SetExchangeConnection(value TeamworkConnectionable)() {
    m.exchangeConnection = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkLoginStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSkypeConnection sets the skypeConnection property value. Information about the Skype for Business connection.
func (m *TeamworkLoginStatus) SetSkypeConnection(value TeamworkConnectionable)() {
    m.skypeConnection = value
}
// SetTeamsConnection sets the teamsConnection property value. Information about the Teams connection.
func (m *TeamworkLoginStatus) SetTeamsConnection(value TeamworkConnectionable)() {
    m.teamsConnection = value
}
