package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAccount 
type UserAccount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The lastSeenDateTime property
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The riskScore property
    riskScore *string
    // The service property
    service *string
    // The signinName property
    signinName *string
    // The status property
    status *AccountStatus
}
// NewUserAccount instantiates a new userAccount and sets the default values.
func NewUserAccount()(*UserAccount) {
    m := &UserAccount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserAccount(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAccount) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *UserAccount) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
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
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["signinName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigninName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccountStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AccountStatus))
        }
        return nil
    }
    return res
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *UserAccount) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserAccount) GetOdataType()(*string) {
    return m.odataType
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *UserAccount) GetRiskScore()(*string) {
    return m.riskScore
}
// GetService gets the service property value. The service property
func (m *UserAccount) GetService()(*string) {
    return m.service
}
// GetSigninName gets the signinName property value. The signinName property
func (m *UserAccount) GetSigninName()(*string) {
    return m.signinName
}
// GetStatus gets the status property value. The status property
func (m *UserAccount) GetStatus()(*AccountStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *UserAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
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
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signinName", m.GetSigninName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAccount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *UserAccount) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *UserAccount) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserAccount) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *UserAccount) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetService sets the service property value. The service property
func (m *UserAccount) SetService(value *string)() {
    m.service = value
}
// SetSigninName sets the signinName property value. The signinName property
func (m *UserAccount) SetSigninName(value *string)() {
    m.signinName = value
}
// SetStatus sets the status property value. The status property
func (m *UserAccount) SetStatus(value *AccountStatus)() {
    m.status = value
}
