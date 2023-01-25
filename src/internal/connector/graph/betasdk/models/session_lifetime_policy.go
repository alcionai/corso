package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SessionLifetimePolicy 
type SessionLifetimePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The human-readable details of the conditional access session management policy applied to the sign-in.
    detail *string
    // If a conditional access session management policy required the user to authenticate in this sign-in event, this field describes the policy type that required authentication. The possible values are: rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy, signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
    expirationRequirement *ExpirationRequirement
    // The OdataType property
    odataType *string
}
// NewSessionLifetimePolicy instantiates a new sessionLifetimePolicy and sets the default values.
func NewSessionLifetimePolicy()(*SessionLifetimePolicy) {
    m := &SessionLifetimePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSessionLifetimePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSessionLifetimePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSessionLifetimePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SessionLifetimePolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDetail gets the detail property value. The human-readable details of the conditional access session management policy applied to the sign-in.
func (m *SessionLifetimePolicy) GetDetail()(*string) {
    return m.detail
}
// GetExpirationRequirement gets the expirationRequirement property value. If a conditional access session management policy required the user to authenticate in this sign-in event, this field describes the policy type that required authentication. The possible values are: rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy, signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
func (m *SessionLifetimePolicy) GetExpirationRequirement()(*ExpirationRequirement) {
    return m.expirationRequirement
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SessionLifetimePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val)
        }
        return nil
    }
    res["expirationRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExpirationRequirement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationRequirement(val.(*ExpirationRequirement))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SessionLifetimePolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SessionLifetimePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    if m.GetExpirationRequirement() != nil {
        cast := (*m.GetExpirationRequirement()).String()
        err := writer.WriteStringValue("expirationRequirement", &cast)
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
func (m *SessionLifetimePolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDetail sets the detail property value. The human-readable details of the conditional access session management policy applied to the sign-in.
func (m *SessionLifetimePolicy) SetDetail(value *string)() {
    m.detail = value
}
// SetExpirationRequirement sets the expirationRequirement property value. If a conditional access session management policy required the user to authenticate in this sign-in event, this field describes the policy type that required authentication. The possible values are: rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy, signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
func (m *SessionLifetimePolicy) SetExpirationRequirement(value *ExpirationRequirement)() {
    m.expirationRequirement = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SessionLifetimePolicy) SetOdataType(value *string)() {
    m.odataType = value
}
