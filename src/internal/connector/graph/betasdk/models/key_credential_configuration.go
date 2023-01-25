package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyCredentialConfiguration 
type KeyCredentialConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The maxLifetime property
    maxLifetime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
    // Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
    restrictForAppsCreatedAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
    restrictionType *AppKeyCredentialRestrictionType
}
// NewKeyCredentialConfiguration instantiates a new keyCredentialConfiguration and sets the default values.
func NewKeyCredentialConfiguration()(*KeyCredentialConfiguration) {
    m := &KeyCredentialConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKeyCredentialConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyCredentialConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyCredentialConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredentialConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyCredentialConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxLifetime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLifetime(val)
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
    res["restrictForAppsCreatedAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictForAppsCreatedAfterDateTime(val)
        }
        return nil
    }
    res["restrictionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppKeyCredentialRestrictionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictionType(val.(*AppKeyCredentialRestrictionType))
        }
        return nil
    }
    return res
}
// GetMaxLifetime gets the maxLifetime property value. The maxLifetime property
func (m *KeyCredentialConfiguration) GetMaxLifetime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.maxLifetime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *KeyCredentialConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetRestrictForAppsCreatedAfterDateTime gets the restrictForAppsCreatedAfterDateTime property value. Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
func (m *KeyCredentialConfiguration) GetRestrictForAppsCreatedAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.restrictForAppsCreatedAfterDateTime
}
// GetRestrictionType gets the restrictionType property value. The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *KeyCredentialConfiguration) GetRestrictionType()(*AppKeyCredentialRestrictionType) {
    return m.restrictionType
}
// Serialize serializes information the current object
func (m *KeyCredentialConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("maxLifetime", m.GetMaxLifetime())
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
        err := writer.WriteTimeValue("restrictForAppsCreatedAfterDateTime", m.GetRestrictForAppsCreatedAfterDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictionType() != nil {
        cast := (*m.GetRestrictionType()).String()
        err := writer.WriteStringValue("restrictionType", &cast)
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
func (m *KeyCredentialConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMaxLifetime sets the maxLifetime property value. The maxLifetime property
func (m *KeyCredentialConfiguration) SetMaxLifetime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.maxLifetime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *KeyCredentialConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRestrictForAppsCreatedAfterDateTime sets the restrictForAppsCreatedAfterDateTime property value. Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications, the enforcement date would be back dated. To apply to all applications regardless of their creation date, this property would be null. Nullable.
func (m *KeyCredentialConfiguration) SetRestrictForAppsCreatedAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restrictForAppsCreatedAfterDateTime = value
}
// SetRestrictionType sets the restrictionType property value. The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of restrictionType can be used only once per policy.
func (m *KeyCredentialConfiguration) SetRestrictionType(value *AppKeyCredentialRestrictionType)() {
    m.restrictionType = value
}
