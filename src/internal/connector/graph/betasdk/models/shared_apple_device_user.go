package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedAppleDeviceUser 
type SharedAppleDeviceUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Data quota
    dataQuota *int64
    // Data to sync
    dataToSync *bool
    // Data quota
    dataUsed *int64
    // The OdataType property
    odataType *string
    // User name
    userPrincipalName *string
}
// NewSharedAppleDeviceUser instantiates a new sharedAppleDeviceUser and sets the default values.
func NewSharedAppleDeviceUser()(*SharedAppleDeviceUser) {
    m := &SharedAppleDeviceUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSharedAppleDeviceUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedAppleDeviceUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedAppleDeviceUser(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharedAppleDeviceUser) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDataQuota gets the dataQuota property value. Data quota
func (m *SharedAppleDeviceUser) GetDataQuota()(*int64) {
    return m.dataQuota
}
// GetDataToSync gets the dataToSync property value. Data to sync
func (m *SharedAppleDeviceUser) GetDataToSync()(*bool) {
    return m.dataToSync
}
// GetDataUsed gets the dataUsed property value. Data quota
func (m *SharedAppleDeviceUser) GetDataUsed()(*int64) {
    return m.dataUsed
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedAppleDeviceUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataQuota"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataQuota(val)
        }
        return nil
    }
    res["dataToSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataToSync(val)
        }
        return nil
    }
    res["dataUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataUsed(val)
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
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SharedAppleDeviceUser) GetOdataType()(*string) {
    return m.odataType
}
// GetUserPrincipalName gets the userPrincipalName property value. User name
func (m *SharedAppleDeviceUser) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *SharedAppleDeviceUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("dataQuota", m.GetDataQuota())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dataToSync", m.GetDataToSync())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("dataUsed", m.GetDataUsed())
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
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *SharedAppleDeviceUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDataQuota sets the dataQuota property value. Data quota
func (m *SharedAppleDeviceUser) SetDataQuota(value *int64)() {
    m.dataQuota = value
}
// SetDataToSync sets the dataToSync property value. Data to sync
func (m *SharedAppleDeviceUser) SetDataToSync(value *bool)() {
    m.dataToSync = value
}
// SetDataUsed sets the dataUsed property value. Data quota
func (m *SharedAppleDeviceUser) SetDataUsed(value *int64)() {
    m.dataUsed = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharedAppleDeviceUser) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User name
func (m *SharedAppleDeviceUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
