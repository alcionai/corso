package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EasEmailProfileConfigurationBase 
type EasEmailProfileConfigurationBase struct {
    DeviceConfiguration
    // Custom domain name value used while generating an email profile before installing on the device.
    customDomainName *string
    // UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
    userDomainNameSource *DomainNameSource
    // Name of the AAD field, that will be used to retrieve UserName for email profile. Possible values are: userPrincipalName, primarySmtpAddress, samAccountName.
    usernameAADSource *UsernameSource
    // Possible values for username source or email source.
    usernameSource *UserEmailSource
}
// NewEasEmailProfileConfigurationBase instantiates a new EasEmailProfileConfigurationBase and sets the default values.
func NewEasEmailProfileConfigurationBase()(*EasEmailProfileConfigurationBase) {
    m := &EasEmailProfileConfigurationBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.easEmailProfileConfigurationBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEasEmailProfileConfigurationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEasEmailProfileConfigurationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.iosEasEmailProfileConfiguration":
                        return NewIosEasEmailProfileConfiguration(), nil
                    case "#microsoft.graph.windows10EasEmailProfileConfiguration":
                        return NewWindows10EasEmailProfileConfiguration(), nil
                    case "#microsoft.graph.windowsPhoneEASEmailProfileConfiguration":
                        return NewWindowsPhoneEASEmailProfileConfiguration(), nil
                }
            }
        }
    }
    return NewEasEmailProfileConfigurationBase(), nil
}
// GetCustomDomainName gets the customDomainName property value. Custom domain name value used while generating an email profile before installing on the device.
func (m *EasEmailProfileConfigurationBase) GetCustomDomainName()(*string) {
    return m.customDomainName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EasEmailProfileConfigurationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["customDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDomainName(val)
        }
        return nil
    }
    res["userDomainNameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDomainNameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDomainNameSource(val.(*DomainNameSource))
        }
        return nil
    }
    res["usernameAADSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsernameSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameAADSource(val.(*UsernameSource))
        }
        return nil
    }
    res["usernameSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserEmailSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameSource(val.(*UserEmailSource))
        }
        return nil
    }
    return res
}
// GetUserDomainNameSource gets the userDomainNameSource property value. UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
func (m *EasEmailProfileConfigurationBase) GetUserDomainNameSource()(*DomainNameSource) {
    return m.userDomainNameSource
}
// GetUsernameAADSource gets the usernameAADSource property value. Name of the AAD field, that will be used to retrieve UserName for email profile. Possible values are: userPrincipalName, primarySmtpAddress, samAccountName.
func (m *EasEmailProfileConfigurationBase) GetUsernameAADSource()(*UsernameSource) {
    return m.usernameAADSource
}
// GetUsernameSource gets the usernameSource property value. Possible values for username source or email source.
func (m *EasEmailProfileConfigurationBase) GetUsernameSource()(*UserEmailSource) {
    return m.usernameSource
}
// Serialize serializes information the current object
func (m *EasEmailProfileConfigurationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customDomainName", m.GetCustomDomainName())
        if err != nil {
            return err
        }
    }
    if m.GetUserDomainNameSource() != nil {
        cast := (*m.GetUserDomainNameSource()).String()
        err = writer.WriteStringValue("userDomainNameSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsernameAADSource() != nil {
        cast := (*m.GetUsernameAADSource()).String()
        err = writer.WriteStringValue("usernameAADSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsernameSource() != nil {
        cast := (*m.GetUsernameSource()).String()
        err = writer.WriteStringValue("usernameSource", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomDomainName sets the customDomainName property value. Custom domain name value used while generating an email profile before installing on the device.
func (m *EasEmailProfileConfigurationBase) SetCustomDomainName(value *string)() {
    m.customDomainName = value
}
// SetUserDomainNameSource sets the userDomainNameSource property value. UserDomainname attribute that is picked from AAD and injected into this profile before installing on the device. Possible values are: fullDomainName, netBiosDomainName.
func (m *EasEmailProfileConfigurationBase) SetUserDomainNameSource(value *DomainNameSource)() {
    m.userDomainNameSource = value
}
// SetUsernameAADSource sets the usernameAADSource property value. Name of the AAD field, that will be used to retrieve UserName for email profile. Possible values are: userPrincipalName, primarySmtpAddress, samAccountName.
func (m *EasEmailProfileConfigurationBase) SetUsernameAADSource(value *UsernameSource)() {
    m.usernameAADSource = value
}
// SetUsernameSource sets the usernameSource property value. Possible values for username source or email source.
func (m *EasEmailProfileConfigurationBase) SetUsernameSource(value *UserEmailSource)() {
    m.usernameSource = value
}
