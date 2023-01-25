package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizationPolicy 
type AuthorizationPolicy struct {
    PolicyBase
    // Indicates whether users can sign up for email based subscriptions.
    allowedToSignUpEmailBasedSubscriptions *bool
    // Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
    allowedToUseSSPR *bool
    // Indicates whether a user can join the tenant by email validation.
    allowEmailVerifiedUsersToJoinOrganization *bool
    // Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone. everyone is the default setting for all cloud environments except US Government. See more in the table below.
    allowInvitesFrom *AllowInvitesFrom
    // Indicates whether user consent for risky apps is allowed. We recommend to keep this as false.
    allowUserConsentForRiskyApps *bool
    // To disable the use of the MSOnline PowerShell module set this property to true. This will also disable user-based access to the legacy service endpoint used by the MSOnline PowerShell module. This does not affect Azure AD Connect or Microsoft Graph.
    blockMsolPowerShell *bool
    // The defaultUserRoleOverrides property
    defaultUserRoleOverrides []DefaultUserRoleOverrideable
    // The defaultUserRolePermissions property
    defaultUserRolePermissions DefaultUserRolePermissionsable
    // List of features enabled for private preview on the tenant.
    enabledPreviewFeatures []string
    // Represents role templateId for the role that should be granted to guest user. Refer to List unifiedRoleDefinitions to find the list of available role templates. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
    guestUserRoleId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Indicates if user consent to apps is allowed, and if it is, which app consent policy (permissionGrantPolicy) governs the permission for users to grant consent. Values should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
    permissionGrantPolicyIdsAssignedToDefaultUserRole []string
}
// NewAuthorizationPolicy instantiates a new AuthorizationPolicy and sets the default values.
func NewAuthorizationPolicy()(*AuthorizationPolicy) {
    m := &AuthorizationPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.authorizationPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthorizationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthorizationPolicy(), nil
}
// GetAllowedToSignUpEmailBasedSubscriptions gets the allowedToSignUpEmailBasedSubscriptions property value. Indicates whether users can sign up for email based subscriptions.
func (m *AuthorizationPolicy) GetAllowedToSignUpEmailBasedSubscriptions()(*bool) {
    return m.allowedToSignUpEmailBasedSubscriptions
}
// GetAllowedToUseSSPR gets the allowedToUseSSPR property value. Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
func (m *AuthorizationPolicy) GetAllowedToUseSSPR()(*bool) {
    return m.allowedToUseSSPR
}
// GetAllowEmailVerifiedUsersToJoinOrganization gets the allowEmailVerifiedUsersToJoinOrganization property value. Indicates whether a user can join the tenant by email validation.
func (m *AuthorizationPolicy) GetAllowEmailVerifiedUsersToJoinOrganization()(*bool) {
    return m.allowEmailVerifiedUsersToJoinOrganization
}
// GetAllowInvitesFrom gets the allowInvitesFrom property value. Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone. everyone is the default setting for all cloud environments except US Government. See more in the table below.
func (m *AuthorizationPolicy) GetAllowInvitesFrom()(*AllowInvitesFrom) {
    return m.allowInvitesFrom
}
// GetAllowUserConsentForRiskyApps gets the allowUserConsentForRiskyApps property value. Indicates whether user consent for risky apps is allowed. We recommend to keep this as false.
func (m *AuthorizationPolicy) GetAllowUserConsentForRiskyApps()(*bool) {
    return m.allowUserConsentForRiskyApps
}
// GetBlockMsolPowerShell gets the blockMsolPowerShell property value. To disable the use of the MSOnline PowerShell module set this property to true. This will also disable user-based access to the legacy service endpoint used by the MSOnline PowerShell module. This does not affect Azure AD Connect or Microsoft Graph.
func (m *AuthorizationPolicy) GetBlockMsolPowerShell()(*bool) {
    return m.blockMsolPowerShell
}
// GetDefaultUserRoleOverrides gets the defaultUserRoleOverrides property value. The defaultUserRoleOverrides property
func (m *AuthorizationPolicy) GetDefaultUserRoleOverrides()([]DefaultUserRoleOverrideable) {
    return m.defaultUserRoleOverrides
}
// GetDefaultUserRolePermissions gets the defaultUserRolePermissions property value. The defaultUserRolePermissions property
func (m *AuthorizationPolicy) GetDefaultUserRolePermissions()(DefaultUserRolePermissionsable) {
    return m.defaultUserRolePermissions
}
// GetEnabledPreviewFeatures gets the enabledPreviewFeatures property value. List of features enabled for private preview on the tenant.
func (m *AuthorizationPolicy) GetEnabledPreviewFeatures()([]string) {
    return m.enabledPreviewFeatures
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["allowedToSignUpEmailBasedSubscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToSignUpEmailBasedSubscriptions(val)
        }
        return nil
    }
    res["allowedToUseSSPR"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToUseSSPR(val)
        }
        return nil
    }
    res["allowEmailVerifiedUsersToJoinOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEmailVerifiedUsersToJoinOrganization(val)
        }
        return nil
    }
    res["allowInvitesFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowInvitesFrom)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInvitesFrom(val.(*AllowInvitesFrom))
        }
        return nil
    }
    res["allowUserConsentForRiskyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUserConsentForRiskyApps(val)
        }
        return nil
    }
    res["blockMsolPowerShell"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockMsolPowerShell(val)
        }
        return nil
    }
    res["defaultUserRoleOverrides"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDefaultUserRoleOverrideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultUserRoleOverrideable, len(val))
            for i, v := range val {
                res[i] = v.(DefaultUserRoleOverrideable)
            }
            m.SetDefaultUserRoleOverrides(res)
        }
        return nil
    }
    res["defaultUserRolePermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDefaultUserRolePermissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultUserRolePermissions(val.(DefaultUserRolePermissionsable))
        }
        return nil
    }
    res["enabledPreviewFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnabledPreviewFeatures(res)
        }
        return nil
    }
    res["guestUserRoleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestUserRoleId(val)
        }
        return nil
    }
    res["permissionGrantPolicyIdsAssignedToDefaultUserRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPermissionGrantPolicyIdsAssignedToDefaultUserRole(res)
        }
        return nil
    }
    return res
}
// GetGuestUserRoleId gets the guestUserRoleId property value. Represents role templateId for the role that should be granted to guest user. Refer to List unifiedRoleDefinitions to find the list of available role templates. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
func (m *AuthorizationPolicy) GetGuestUserRoleId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.guestUserRoleId
}
// GetPermissionGrantPolicyIdsAssignedToDefaultUserRole gets the permissionGrantPolicyIdsAssignedToDefaultUserRole property value. Indicates if user consent to apps is allowed, and if it is, which app consent policy (permissionGrantPolicy) governs the permission for users to grant consent. Values should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
func (m *AuthorizationPolicy) GetPermissionGrantPolicyIdsAssignedToDefaultUserRole()([]string) {
    return m.permissionGrantPolicyIdsAssignedToDefaultUserRole
}
// Serialize serializes information the current object
func (m *AuthorizationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowedToSignUpEmailBasedSubscriptions", m.GetAllowedToSignUpEmailBasedSubscriptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowedToUseSSPR", m.GetAllowedToUseSSPR())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowEmailVerifiedUsersToJoinOrganization", m.GetAllowEmailVerifiedUsersToJoinOrganization())
        if err != nil {
            return err
        }
    }
    if m.GetAllowInvitesFrom() != nil {
        cast := (*m.GetAllowInvitesFrom()).String()
        err = writer.WriteStringValue("allowInvitesFrom", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowUserConsentForRiskyApps", m.GetAllowUserConsentForRiskyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockMsolPowerShell", m.GetBlockMsolPowerShell())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultUserRoleOverrides() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultUserRoleOverrides()))
        for i, v := range m.GetDefaultUserRoleOverrides() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("defaultUserRoleOverrides", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultUserRolePermissions", m.GetDefaultUserRolePermissions())
        if err != nil {
            return err
        }
    }
    if m.GetEnabledPreviewFeatures() != nil {
        err = writer.WriteCollectionOfStringValues("enabledPreviewFeatures", m.GetEnabledPreviewFeatures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("guestUserRoleId", m.GetGuestUserRoleId())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionGrantPolicyIdsAssignedToDefaultUserRole() != nil {
        err = writer.WriteCollectionOfStringValues("permissionGrantPolicyIdsAssignedToDefaultUserRole", m.GetPermissionGrantPolicyIdsAssignedToDefaultUserRole())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedToSignUpEmailBasedSubscriptions sets the allowedToSignUpEmailBasedSubscriptions property value. Indicates whether users can sign up for email based subscriptions.
func (m *AuthorizationPolicy) SetAllowedToSignUpEmailBasedSubscriptions(value *bool)() {
    m.allowedToSignUpEmailBasedSubscriptions = value
}
// SetAllowedToUseSSPR sets the allowedToUseSSPR property value. Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.
func (m *AuthorizationPolicy) SetAllowedToUseSSPR(value *bool)() {
    m.allowedToUseSSPR = value
}
// SetAllowEmailVerifiedUsersToJoinOrganization sets the allowEmailVerifiedUsersToJoinOrganization property value. Indicates whether a user can join the tenant by email validation.
func (m *AuthorizationPolicy) SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)() {
    m.allowEmailVerifiedUsersToJoinOrganization = value
}
// SetAllowInvitesFrom sets the allowInvitesFrom property value. Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone. everyone is the default setting for all cloud environments except US Government. See more in the table below.
func (m *AuthorizationPolicy) SetAllowInvitesFrom(value *AllowInvitesFrom)() {
    m.allowInvitesFrom = value
}
// SetAllowUserConsentForRiskyApps sets the allowUserConsentForRiskyApps property value. Indicates whether user consent for risky apps is allowed. We recommend to keep this as false.
func (m *AuthorizationPolicy) SetAllowUserConsentForRiskyApps(value *bool)() {
    m.allowUserConsentForRiskyApps = value
}
// SetBlockMsolPowerShell sets the blockMsolPowerShell property value. To disable the use of the MSOnline PowerShell module set this property to true. This will also disable user-based access to the legacy service endpoint used by the MSOnline PowerShell module. This does not affect Azure AD Connect or Microsoft Graph.
func (m *AuthorizationPolicy) SetBlockMsolPowerShell(value *bool)() {
    m.blockMsolPowerShell = value
}
// SetDefaultUserRoleOverrides sets the defaultUserRoleOverrides property value. The defaultUserRoleOverrides property
func (m *AuthorizationPolicy) SetDefaultUserRoleOverrides(value []DefaultUserRoleOverrideable)() {
    m.defaultUserRoleOverrides = value
}
// SetDefaultUserRolePermissions sets the defaultUserRolePermissions property value. The defaultUserRolePermissions property
func (m *AuthorizationPolicy) SetDefaultUserRolePermissions(value DefaultUserRolePermissionsable)() {
    m.defaultUserRolePermissions = value
}
// SetEnabledPreviewFeatures sets the enabledPreviewFeatures property value. List of features enabled for private preview on the tenant.
func (m *AuthorizationPolicy) SetEnabledPreviewFeatures(value []string)() {
    m.enabledPreviewFeatures = value
}
// SetGuestUserRoleId sets the guestUserRoleId property value. Represents role templateId for the role that should be granted to guest user. Refer to List unifiedRoleDefinitions to find the list of available role templates. Currently following roles are supported:  User (a0b1b346-4d3e-4e8b-98f8-753987be4970), Guest User (10dae51f-b6af-4016-8d66-8c2a99b929b3), and Restricted Guest User (2af84b1e-32c8-42b7-82bc-daa82404023b).
func (m *AuthorizationPolicy) SetGuestUserRoleId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.guestUserRoleId = value
}
// SetPermissionGrantPolicyIdsAssignedToDefaultUserRole sets the permissionGrantPolicyIdsAssignedToDefaultUserRole property value. Indicates if user consent to apps is allowed, and if it is, which app consent policy (permissionGrantPolicy) governs the permission for users to grant consent. Values should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
func (m *AuthorizationPolicy) SetPermissionGrantPolicyIdsAssignedToDefaultUserRole(value []string)() {
    m.permissionGrantPolicyIdsAssignedToDefaultUserRole = value
}
