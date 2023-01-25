package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosSingleSignOnSettings iOS Kerberos authentication settings for single sign-on
type IosSingleSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of app identifiers that are allowed to use this login. If this field is omitted, the login applies to all applications on the device. This collection can contain a maximum of 500 elements.
    allowedAppsList []AppListItemable
    // List of HTTP URLs that must be matched in order to use this login. With iOS 9.0 or later, a wildcard characters may be used.
    allowedUrls []string
    // The display name of login settings shown on the receiving device.
    displayName *string
    // A Kerberos principal name. If not provided, the user is prompted for one during profile installation.
    kerberosPrincipalName *string
    // A Kerberos realm name. Case sensitive.
    kerberosRealm *string
    // The OdataType property
    odataType *string
}
// NewIosSingleSignOnSettings instantiates a new iosSingleSignOnSettings and sets the default values.
func NewIosSingleSignOnSettings()(*IosSingleSignOnSettings) {
    m := &IosSingleSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosSingleSignOnSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosSingleSignOnSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosSingleSignOnSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosSingleSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedAppsList gets the allowedAppsList property value. List of app identifiers that are allowed to use this login. If this field is omitted, the login applies to all applications on the device. This collection can contain a maximum of 500 elements.
func (m *IosSingleSignOnSettings) GetAllowedAppsList()([]AppListItemable) {
    return m.allowedAppsList
}
// GetAllowedUrls gets the allowedUrls property value. List of HTTP URLs that must be matched in order to use this login. With iOS 9.0 or later, a wildcard characters may be used.
func (m *IosSingleSignOnSettings) GetAllowedUrls()([]string) {
    return m.allowedUrls
}
// GetDisplayName gets the displayName property value. The display name of login settings shown on the receiving device.
func (m *IosSingleSignOnSettings) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosSingleSignOnSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedAppsList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetAllowedAppsList(res)
        }
        return nil
    }
    res["allowedUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedUrls(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["kerberosPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosPrincipalName(val)
        }
        return nil
    }
    res["kerberosRealm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosRealm(val)
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
// GetKerberosPrincipalName gets the kerberosPrincipalName property value. A Kerberos principal name. If not provided, the user is prompted for one during profile installation.
func (m *IosSingleSignOnSettings) GetKerberosPrincipalName()(*string) {
    return m.kerberosPrincipalName
}
// GetKerberosRealm gets the kerberosRealm property value. A Kerberos realm name. Case sensitive.
func (m *IosSingleSignOnSettings) GetKerberosRealm()(*string) {
    return m.kerberosRealm
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosSingleSignOnSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *IosSingleSignOnSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedAppsList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedAppsList()))
        for i, v := range m.GetAllowedAppsList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("allowedAppsList", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedUrls() != nil {
        err := writer.WriteCollectionOfStringValues("allowedUrls", m.GetAllowedUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kerberosPrincipalName", m.GetKerberosPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("kerberosRealm", m.GetKerberosRealm())
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
func (m *IosSingleSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedAppsList sets the allowedAppsList property value. List of app identifiers that are allowed to use this login. If this field is omitted, the login applies to all applications on the device. This collection can contain a maximum of 500 elements.
func (m *IosSingleSignOnSettings) SetAllowedAppsList(value []AppListItemable)() {
    m.allowedAppsList = value
}
// SetAllowedUrls sets the allowedUrls property value. List of HTTP URLs that must be matched in order to use this login. With iOS 9.0 or later, a wildcard characters may be used.
func (m *IosSingleSignOnSettings) SetAllowedUrls(value []string)() {
    m.allowedUrls = value
}
// SetDisplayName sets the displayName property value. The display name of login settings shown on the receiving device.
func (m *IosSingleSignOnSettings) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetKerberosPrincipalName sets the kerberosPrincipalName property value. A Kerberos principal name. If not provided, the user is prompted for one during profile installation.
func (m *IosSingleSignOnSettings) SetKerberosPrincipalName(value *string)() {
    m.kerberosPrincipalName = value
}
// SetKerberosRealm sets the kerberosRealm property value. A Kerberos realm name. Case sensitive.
func (m *IosSingleSignOnSettings) SetKerberosRealm(value *string)() {
    m.kerberosRealm = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosSingleSignOnSettings) SetOdataType(value *string)() {
    m.odataType = value
}
