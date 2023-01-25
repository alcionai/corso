package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationDetails provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UserRegistrationDetails struct {
    Entity
    // The method the user or admin selected as default for performing multi-factor authentication for the user. The possible values are: none, mobilePhone, alternateMobilePhone, officePhone, microsoftAuthenticatorPush, softwareOneTimePasscode, unknownFutureValue.
    defaultMfaMethod *DefaultMfaMethodType
    // Whether the user has an admin role in the tenant. This value can be used to check the authentication methods that privileged accounts are registered for and capable of.
    isAdmin *bool
    // Whether the user has registered a strong authentication method for multi-factor authentication. The method must be allowed by the authentication methods policy. Supports $filter (eq).
    isMfaCapable *bool
    // Whether the user has registered a strong authentication method for multi-factor authentication. The method may not necessarily be allowed by the authentication methods policy.  Supports $filter (eq).
    isMfaRegistered *bool
    // Whether the user has registered a passwordless strong authentication method (including FIDO2, Windows Hello for Business, and Microsoft Authenticator (Passwordless)) that is allowed by the authentication methods policy. Supports $filter (eq).
    isPasswordlessCapable *bool
    // Whether the user has registered the required number of authentication methods for self-service password reset and the user is allowed to perform self-service password reset by policy. Supports $filter (eq).
    isSsprCapable *bool
    // Whether the user is allowed to perform self-service password reset by policy. The user may not necessarily have registered the required number of authentication methods for self-service password reset. Supports $filter (eq).
    isSsprEnabled *bool
    // Whether the user has registered the required number of authentication methods for self-service password reset. The user may not necessarily be allowed to perform self-service password reset by policy. Supports $filter (eq).
    isSsprRegistered *bool
    // Collection of authentication methods registered, such as mobilePhone, email, fido2. Supports $filter (any with eq).
    methodsRegistered []string
    // The user display name, such as Adele Vance. Supports $filter (eq, startsWith) and $orderBy.
    userDisplayName *string
    // The user principal name, such as AdeleV@contoso.com. Supports $filter (eq, startsWith) and $orderBy.
    userPrincipalName *string
    // Identifies whether the user is a member or guest in the tenant. The possible values are: member, guest, unknownFutureValue.
    userType *SignInUserType
}
// NewUserRegistrationDetails instantiates a new userRegistrationDetails and sets the default values.
func NewUserRegistrationDetails()(*UserRegistrationDetails) {
    m := &UserRegistrationDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserRegistrationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationDetails(), nil
}
// GetDefaultMfaMethod gets the defaultMfaMethod property value. The method the user or admin selected as default for performing multi-factor authentication for the user. The possible values are: none, mobilePhone, alternateMobilePhone, officePhone, microsoftAuthenticatorPush, softwareOneTimePasscode, unknownFutureValue.
func (m *UserRegistrationDetails) GetDefaultMfaMethod()(*DefaultMfaMethodType) {
    return m.defaultMfaMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultMfaMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefaultMfaMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMfaMethod(val.(*DefaultMfaMethodType))
        }
        return nil
    }
    res["isAdmin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAdmin(val)
        }
        return nil
    }
    res["isMfaCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaCapable(val)
        }
        return nil
    }
    res["isMfaRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaRegistered(val)
        }
        return nil
    }
    res["isPasswordlessCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPasswordlessCapable(val)
        }
        return nil
    }
    res["isSsprCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprCapable(val)
        }
        return nil
    }
    res["isSsprEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprEnabled(val)
        }
        return nil
    }
    res["isSsprRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSsprRegistered(val)
        }
        return nil
    }
    res["methodsRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMethodsRegistered(res)
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
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSignInUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*SignInUserType))
        }
        return nil
    }
    return res
}
// GetIsAdmin gets the isAdmin property value. Whether the user has an admin role in the tenant. This value can be used to check the authentication methods that privileged accounts are registered for and capable of.
func (m *UserRegistrationDetails) GetIsAdmin()(*bool) {
    return m.isAdmin
}
// GetIsMfaCapable gets the isMfaCapable property value. Whether the user has registered a strong authentication method for multi-factor authentication. The method must be allowed by the authentication methods policy. Supports $filter (eq).
func (m *UserRegistrationDetails) GetIsMfaCapable()(*bool) {
    return m.isMfaCapable
}
// GetIsMfaRegistered gets the isMfaRegistered property value. Whether the user has registered a strong authentication method for multi-factor authentication. The method may not necessarily be allowed by the authentication methods policy.  Supports $filter (eq).
func (m *UserRegistrationDetails) GetIsMfaRegistered()(*bool) {
    return m.isMfaRegistered
}
// GetIsPasswordlessCapable gets the isPasswordlessCapable property value. Whether the user has registered a passwordless strong authentication method (including FIDO2, Windows Hello for Business, and Microsoft Authenticator (Passwordless)) that is allowed by the authentication methods policy. Supports $filter (eq).
func (m *UserRegistrationDetails) GetIsPasswordlessCapable()(*bool) {
    return m.isPasswordlessCapable
}
// GetIsSsprCapable gets the isSsprCapable property value. Whether the user has registered the required number of authentication methods for self-service password reset and the user is allowed to perform self-service password reset by policy. Supports $filter (eq).
func (m *UserRegistrationDetails) GetIsSsprCapable()(*bool) {
    return m.isSsprCapable
}
// GetIsSsprEnabled gets the isSsprEnabled property value. Whether the user is allowed to perform self-service password reset by policy. The user may not necessarily have registered the required number of authentication methods for self-service password reset. Supports $filter (eq).
func (m *UserRegistrationDetails) GetIsSsprEnabled()(*bool) {
    return m.isSsprEnabled
}
// GetIsSsprRegistered gets the isSsprRegistered property value. Whether the user has registered the required number of authentication methods for self-service password reset. The user may not necessarily be allowed to perform self-service password reset by policy. Supports $filter (eq).
func (m *UserRegistrationDetails) GetIsSsprRegistered()(*bool) {
    return m.isSsprRegistered
}
// GetMethodsRegistered gets the methodsRegistered property value. Collection of authentication methods registered, such as mobilePhone, email, fido2. Supports $filter (any with eq).
func (m *UserRegistrationDetails) GetMethodsRegistered()([]string) {
    return m.methodsRegistered
}
// GetUserDisplayName gets the userDisplayName property value. The user display name, such as Adele Vance. Supports $filter (eq, startsWith) and $orderBy.
func (m *UserRegistrationDetails) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name, such as AdeleV@contoso.com. Supports $filter (eq, startsWith) and $orderBy.
func (m *UserRegistrationDetails) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserType gets the userType property value. Identifies whether the user is a member or guest in the tenant. The possible values are: member, guest, unknownFutureValue.
func (m *UserRegistrationDetails) GetUserType()(*SignInUserType) {
    return m.userType
}
// Serialize serializes information the current object
func (m *UserRegistrationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefaultMfaMethod() != nil {
        cast := (*m.GetDefaultMfaMethod()).String()
        err = writer.WriteStringValue("defaultMfaMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAdmin", m.GetIsAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaCapable", m.GetIsMfaCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaRegistered", m.GetIsMfaRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPasswordlessCapable", m.GetIsPasswordlessCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprCapable", m.GetIsSsprCapable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprEnabled", m.GetIsSsprEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSsprRegistered", m.GetIsSsprRegistered())
        if err != nil {
            return err
        }
    }
    if m.GetMethodsRegistered() != nil {
        err = writer.WriteCollectionOfStringValues("methodsRegistered", m.GetMethodsRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err = writer.WriteStringValue("userType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultMfaMethod sets the defaultMfaMethod property value. The method the user or admin selected as default for performing multi-factor authentication for the user. The possible values are: none, mobilePhone, alternateMobilePhone, officePhone, microsoftAuthenticatorPush, softwareOneTimePasscode, unknownFutureValue.
func (m *UserRegistrationDetails) SetDefaultMfaMethod(value *DefaultMfaMethodType)() {
    m.defaultMfaMethod = value
}
// SetIsAdmin sets the isAdmin property value. Whether the user has an admin role in the tenant. This value can be used to check the authentication methods that privileged accounts are registered for and capable of.
func (m *UserRegistrationDetails) SetIsAdmin(value *bool)() {
    m.isAdmin = value
}
// SetIsMfaCapable sets the isMfaCapable property value. Whether the user has registered a strong authentication method for multi-factor authentication. The method must be allowed by the authentication methods policy. Supports $filter (eq).
func (m *UserRegistrationDetails) SetIsMfaCapable(value *bool)() {
    m.isMfaCapable = value
}
// SetIsMfaRegistered sets the isMfaRegistered property value. Whether the user has registered a strong authentication method for multi-factor authentication. The method may not necessarily be allowed by the authentication methods policy.  Supports $filter (eq).
func (m *UserRegistrationDetails) SetIsMfaRegistered(value *bool)() {
    m.isMfaRegistered = value
}
// SetIsPasswordlessCapable sets the isPasswordlessCapable property value. Whether the user has registered a passwordless strong authentication method (including FIDO2, Windows Hello for Business, and Microsoft Authenticator (Passwordless)) that is allowed by the authentication methods policy. Supports $filter (eq).
func (m *UserRegistrationDetails) SetIsPasswordlessCapable(value *bool)() {
    m.isPasswordlessCapable = value
}
// SetIsSsprCapable sets the isSsprCapable property value. Whether the user has registered the required number of authentication methods for self-service password reset and the user is allowed to perform self-service password reset by policy. Supports $filter (eq).
func (m *UserRegistrationDetails) SetIsSsprCapable(value *bool)() {
    m.isSsprCapable = value
}
// SetIsSsprEnabled sets the isSsprEnabled property value. Whether the user is allowed to perform self-service password reset by policy. The user may not necessarily have registered the required number of authentication methods for self-service password reset. Supports $filter (eq).
func (m *UserRegistrationDetails) SetIsSsprEnabled(value *bool)() {
    m.isSsprEnabled = value
}
// SetIsSsprRegistered sets the isSsprRegistered property value. Whether the user has registered the required number of authentication methods for self-service password reset. The user may not necessarily be allowed to perform self-service password reset by policy. Supports $filter (eq).
func (m *UserRegistrationDetails) SetIsSsprRegistered(value *bool)() {
    m.isSsprRegistered = value
}
// SetMethodsRegistered sets the methodsRegistered property value. Collection of authentication methods registered, such as mobilePhone, email, fido2. Supports $filter (any with eq).
func (m *UserRegistrationDetails) SetMethodsRegistered(value []string)() {
    m.methodsRegistered = value
}
// SetUserDisplayName sets the userDisplayName property value. The user display name, such as Adele Vance. Supports $filter (eq, startsWith) and $orderBy.
func (m *UserRegistrationDetails) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name, such as AdeleV@contoso.com. Supports $filter (eq, startsWith) and $orderBy.
func (m *UserRegistrationDetails) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserType sets the userType property value. Identifies whether the user is a member or guest in the tenant. The possible values are: member, guest, unknownFutureValue.
func (m *UserRegistrationDetails) SetUserType(value *SignInUserType)() {
    m.userType = value
}
