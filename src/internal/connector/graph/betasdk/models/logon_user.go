package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LogonUser 
type LogonUser struct {
    // Domain of user account used to logon.
    accountDomain *string
    // Account name of user account used to logon.
    accountName *string
    // User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
    accountType *UserAccountSecurityType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // User logon ID.
    logonId *string
    // Collection of the logon types observed for the logged on user from when first to last seen. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
    logonTypes []LogonType
    // The OdataType property
    odataType *string
}
// NewLogonUser instantiates a new logonUser and sets the default values.
func NewLogonUser()(*LogonUser) {
    m := &LogonUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLogonUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLogonUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLogonUser(), nil
}
// GetAccountDomain gets the accountDomain property value. Domain of user account used to logon.
func (m *LogonUser) GetAccountDomain()(*string) {
    return m.accountDomain
}
// GetAccountName gets the accountName property value. Account name of user account used to logon.
func (m *LogonUser) GetAccountName()(*string) {
    return m.accountName
}
// GetAccountType gets the accountType property value. User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *LogonUser) GetAccountType()(*UserAccountSecurityType) {
    return m.accountType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LogonUser) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LogonUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountDomain(val)
        }
        return nil
    }
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["accountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAccountSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountType(val.(*UserAccountSecurityType))
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
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
    res["logonId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonId(val)
        }
        return nil
    }
    res["logonTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseLogonType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LogonType, len(val))
            for i, v := range val {
                res[i] = *(v.(*LogonType))
            }
            m.SetLogonTypes(res)
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
// GetFirstSeenDateTime gets the firstSeenDateTime property value. DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetLogonId gets the logonId property value. User logon ID.
func (m *LogonUser) GetLogonId()(*string) {
    return m.logonId
}
// GetLogonTypes gets the logonTypes property value. Collection of the logon types observed for the logged on user from when first to last seen. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *LogonUser) GetLogonTypes()([]LogonType) {
    return m.logonTypes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LogonUser) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LogonUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountDomain", m.GetAccountDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    if m.GetAccountType() != nil {
        cast := (*m.GetAccountType()).String()
        err := writer.WriteStringValue("accountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
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
        err := writer.WriteStringValue("logonId", m.GetLogonId())
        if err != nil {
            return err
        }
    }
    if m.GetLogonTypes() != nil {
        err := writer.WriteCollectionOfStringValues("logonTypes", SerializeLogonType(m.GetLogonTypes()))
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
// SetAccountDomain sets the accountDomain property value. Domain of user account used to logon.
func (m *LogonUser) SetAccountDomain(value *string)() {
    m.accountDomain = value
}
// SetAccountName sets the accountName property value. Account name of user account used to logon.
func (m *LogonUser) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAccountType sets the accountType property value. User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *LogonUser) SetAccountType(value *UserAccountSecurityType)() {
    m.accountType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LogonUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetLogonId sets the logonId property value. User logon ID.
func (m *LogonUser) SetLogonId(value *string)() {
    m.logonId = value
}
// SetLogonTypes sets the logonTypes property value. Collection of the logon types observed for the logged on user from when first to last seen. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *LogonUser) SetLogonTypes(value []LogonType)() {
    m.logonTypes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LogonUser) SetOdataType(value *string)() {
    m.odataType = value
}
