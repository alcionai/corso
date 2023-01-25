package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationMethodCount 
type UserRegistrationMethodCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of authentication method.
    authenticationMethod *string
    // The OdataType property
    odataType *string
    // Number of users registered.
    userCount *int64
}
// NewUserRegistrationMethodCount instantiates a new userRegistrationMethodCount and sets the default values.
func NewUserRegistrationMethodCount()(*UserRegistrationMethodCount) {
    m := &UserRegistrationMethodCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserRegistrationMethodCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationMethodCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationMethodCount(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationMethodCount) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuthenticationMethod gets the authenticationMethod property value. Name of authentication method.
func (m *UserRegistrationMethodCount) GetAuthenticationMethod()(*string) {
    return m.authenticationMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationMethodCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val)
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
    res["userCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCount(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserRegistrationMethodCount) GetOdataType()(*string) {
    return m.odataType
}
// GetUserCount gets the userCount property value. Number of users registered.
func (m *UserRegistrationMethodCount) GetUserCount()(*int64) {
    return m.userCount
}
// Serialize serializes information the current object
func (m *UserRegistrationMethodCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationMethod", m.GetAuthenticationMethod())
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
        err := writer.WriteInt64Value("userCount", m.GetUserCount())
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
func (m *UserRegistrationMethodCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuthenticationMethod sets the authenticationMethod property value. Name of authentication method.
func (m *UserRegistrationMethodCount) SetAuthenticationMethod(value *string)() {
    m.authenticationMethod = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserRegistrationMethodCount) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserCount sets the userCount property value. Number of users registered.
func (m *UserRegistrationMethodCount) SetUserCount(value *int64)() {
    m.userCount = value
}
