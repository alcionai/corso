package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPrintUsageSummary 
type UserPrintUsageSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The completedJobCount property
    completedJobCount *int32
    // The incompleteJobCount property
    incompleteJobCount *int32
    // The OdataType property
    odataType *string
    // The user property
    user Identityable
    // The userDisplayName property
    userDisplayName *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewUserPrintUsageSummary instantiates a new userPrintUsageSummary and sets the default values.
func NewUserPrintUsageSummary()(*UserPrintUsageSummary) {
    m := &UserPrintUsageSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserPrintUsageSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserPrintUsageSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPrintUsageSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserPrintUsageSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompletedJobCount gets the completedJobCount property value. The completedJobCount property
func (m *UserPrintUsageSummary) GetCompletedJobCount()(*int32) {
    return m.completedJobCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserPrintUsageSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedJobCount(val)
        }
        return nil
    }
    res["incompleteJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteJobCount(val)
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(Identityable))
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
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
// GetIncompleteJobCount gets the incompleteJobCount property value. The incompleteJobCount property
func (m *UserPrintUsageSummary) GetIncompleteJobCount()(*int32) {
    return m.incompleteJobCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserPrintUsageSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetUser gets the user property value. The user property
func (m *UserPrintUsageSummary) GetUser()(Identityable) {
    return m.user
}
// GetUserDisplayName gets the userDisplayName property value. The userDisplayName property
func (m *UserPrintUsageSummary) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *UserPrintUsageSummary) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *UserPrintUsageSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("completedJobCount", m.GetCompletedJobCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("incompleteJobCount", m.GetIncompleteJobCount())
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
        err := writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
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
func (m *UserPrintUsageSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompletedJobCount sets the completedJobCount property value. The completedJobCount property
func (m *UserPrintUsageSummary) SetCompletedJobCount(value *int32)() {
    m.completedJobCount = value
}
// SetIncompleteJobCount sets the incompleteJobCount property value. The incompleteJobCount property
func (m *UserPrintUsageSummary) SetIncompleteJobCount(value *int32)() {
    m.incompleteJobCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserPrintUsageSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUser sets the user property value. The user property
func (m *UserPrintUsageSummary) SetUser(value Identityable)() {
    m.user = value
}
// SetUserDisplayName sets the userDisplayName property value. The userDisplayName property
func (m *UserPrintUsageSummary) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *UserPrintUsageSummary) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
