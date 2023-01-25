package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationFeatureCount 
type UserRegistrationFeatureCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The feature property
    feature *AuthenticationMethodFeature
    // The OdataType property
    odataType *string
    // Number of users.
    userCount *int64
}
// NewUserRegistrationFeatureCount instantiates a new userRegistrationFeatureCount and sets the default values.
func NewUserRegistrationFeatureCount()(*UserRegistrationFeatureCount) {
    m := &UserRegistrationFeatureCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserRegistrationFeatureCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationFeatureCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationFeatureCount(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationFeatureCount) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFeature gets the feature property value. The feature property
func (m *UserRegistrationFeatureCount) GetFeature()(*AuthenticationMethodFeature) {
    return m.feature
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationFeatureCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["feature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodFeature)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val.(*AuthenticationMethodFeature))
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
func (m *UserRegistrationFeatureCount) GetOdataType()(*string) {
    return m.odataType
}
// GetUserCount gets the userCount property value. Number of users.
func (m *UserRegistrationFeatureCount) GetUserCount()(*int64) {
    return m.userCount
}
// Serialize serializes information the current object
func (m *UserRegistrationFeatureCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFeature() != nil {
        cast := (*m.GetFeature()).String()
        err := writer.WriteStringValue("feature", &cast)
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
func (m *UserRegistrationFeatureCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFeature sets the feature property value. The feature property
func (m *UserRegistrationFeatureCount) SetFeature(value *AuthenticationMethodFeature)() {
    m.feature = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserRegistrationFeatureCount) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserCount sets the userCount property value. Number of users.
func (m *UserRegistrationFeatureCount) SetUserCount(value *int64)() {
    m.userCount = value
}
