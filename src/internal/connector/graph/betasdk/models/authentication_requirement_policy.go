package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationRequirementPolicy 
type AuthenticationRequirementPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Provides additional detail on the feature identified in requirementProvider.
    detail *string
    // The OdataType property
    odataType *string
    // Identifies what Azure AD feature requires MFA in this policy. Possible values are: user, request, servicePrincipal, v1ConditionalAccess, multiConditionalAccess, tenantSessionRiskPolicy, accountCompromisePolicies, v1ConditionalAccessDependency, v1ConditionalAccessPolicyIdRequested, mfaRegistrationRequiredByIdentityProtectionPolicy, baselineProtection, mfaRegistrationRequiredByBaselineProtection, mfaRegistrationRequiredByMultiConditionalAccess, enforcedForCspAdmins, securityDefaults, mfaRegistrationRequiredBySecurityDefaults, proofUpCodeRequest, crossTenantOutboundRule, gpsLocationCondition, riskBasedPolicy, unknownFutureValue.
    requirementProvider *RequirementProvider
}
// NewAuthenticationRequirementPolicy instantiates a new authenticationRequirementPolicy and sets the default values.
func NewAuthenticationRequirementPolicy()(*AuthenticationRequirementPolicy) {
    m := &AuthenticationRequirementPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationRequirementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationRequirementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationRequirementPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationRequirementPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDetail gets the detail property value. Provides additional detail on the feature identified in requirementProvider.
func (m *AuthenticationRequirementPolicy) GetDetail()(*string) {
    return m.detail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationRequirementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["requirementProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequirementProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequirementProvider(val.(*RequirementProvider))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationRequirementPolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetRequirementProvider gets the requirementProvider property value. Identifies what Azure AD feature requires MFA in this policy. Possible values are: user, request, servicePrincipal, v1ConditionalAccess, multiConditionalAccess, tenantSessionRiskPolicy, accountCompromisePolicies, v1ConditionalAccessDependency, v1ConditionalAccessPolicyIdRequested, mfaRegistrationRequiredByIdentityProtectionPolicy, baselineProtection, mfaRegistrationRequiredByBaselineProtection, mfaRegistrationRequiredByMultiConditionalAccess, enforcedForCspAdmins, securityDefaults, mfaRegistrationRequiredBySecurityDefaults, proofUpCodeRequest, crossTenantOutboundRule, gpsLocationCondition, riskBasedPolicy, unknownFutureValue.
func (m *AuthenticationRequirementPolicy) GetRequirementProvider()(*RequirementProvider) {
    return m.requirementProvider
}
// Serialize serializes information the current object
func (m *AuthenticationRequirementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
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
    if m.GetRequirementProvider() != nil {
        cast := (*m.GetRequirementProvider()).String()
        err := writer.WriteStringValue("requirementProvider", &cast)
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
func (m *AuthenticationRequirementPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDetail sets the detail property value. Provides additional detail on the feature identified in requirementProvider.
func (m *AuthenticationRequirementPolicy) SetDetail(value *string)() {
    m.detail = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationRequirementPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRequirementProvider sets the requirementProvider property value. Identifies what Azure AD feature requires MFA in this policy. Possible values are: user, request, servicePrincipal, v1ConditionalAccess, multiConditionalAccess, tenantSessionRiskPolicy, accountCompromisePolicies, v1ConditionalAccessDependency, v1ConditionalAccessPolicyIdRequested, mfaRegistrationRequiredByIdentityProtectionPolicy, baselineProtection, mfaRegistrationRequiredByBaselineProtection, mfaRegistrationRequiredByMultiConditionalAccess, enforcedForCspAdmins, securityDefaults, mfaRegistrationRequiredBySecurityDefaults, proofUpCodeRequest, crossTenantOutboundRule, gpsLocationCondition, riskBasedPolicy, unknownFutureValue.
func (m *AuthenticationRequirementPolicy) SetRequirementProvider(value *RequirementProvider)() {
    m.requirementProvider = value
}
