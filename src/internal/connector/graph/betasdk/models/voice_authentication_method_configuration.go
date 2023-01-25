package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VoiceAuthenticationMethodConfiguration 
type VoiceAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // A collection of groups that are enabled to use the authentication method. Expanded by default.
    includeTargets []VoiceAuthenticationMethodTargetable
    // true if users can register office phones, otherwise, false.
    isOfficePhoneAllowed *bool
}
// NewVoiceAuthenticationMethodConfiguration instantiates a new VoiceAuthenticationMethodConfiguration and sets the default values.
func NewVoiceAuthenticationMethodConfiguration()(*VoiceAuthenticationMethodConfiguration) {
    m := &VoiceAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.voiceAuthenticationMethodConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateVoiceAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVoiceAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVoiceAuthenticationMethodConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VoiceAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVoiceAuthenticationMethodTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VoiceAuthenticationMethodTargetable, len(val))
            for i, v := range val {
                res[i] = v.(VoiceAuthenticationMethodTargetable)
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    res["isOfficePhoneAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOfficePhoneAllowed(val)
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of groups that are enabled to use the authentication method. Expanded by default.
func (m *VoiceAuthenticationMethodConfiguration) GetIncludeTargets()([]VoiceAuthenticationMethodTargetable) {
    return m.includeTargets
}
// GetIsOfficePhoneAllowed gets the isOfficePhoneAllowed property value. true if users can register office phones, otherwise, false.
func (m *VoiceAuthenticationMethodConfiguration) GetIsOfficePhoneAllowed()(*bool) {
    return m.isOfficePhoneAllowed
}
// Serialize serializes information the current object
func (m *VoiceAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOfficePhoneAllowed", m.GetIsOfficePhoneAllowed())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIncludeTargets sets the includeTargets property value. A collection of groups that are enabled to use the authentication method. Expanded by default.
func (m *VoiceAuthenticationMethodConfiguration) SetIncludeTargets(value []VoiceAuthenticationMethodTargetable)() {
    m.includeTargets = value
}
// SetIsOfficePhoneAllowed sets the isOfficePhoneAllowed property value. true if users can register office phones, otherwise, false.
func (m *VoiceAuthenticationMethodConfiguration) SetIsOfficePhoneAllowed(value *bool)() {
    m.isOfficePhoneAllowed = value
}
