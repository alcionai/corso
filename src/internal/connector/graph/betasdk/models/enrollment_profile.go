package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentProfile 
type EnrollmentProfile struct {
    Entity
    // Configuration endpoint url to use for Enrollment
    configurationEndpointUrl *string
    // Description of the profile
    description *string
    // Name of the profile
    displayName *string
    // Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
    enableAuthenticationViaCompanyPortal *bool
    // Indicates that Company Portal is required on setup assistant enrolled devices
    requireCompanyPortalOnSetupAssistantEnrolledDevices *bool
    // Indicates if the profile requires user authentication
    requiresUserAuthentication *bool
}
// NewEnrollmentProfile instantiates a new enrollmentProfile and sets the default values.
func NewEnrollmentProfile()(*EnrollmentProfile) {
    m := &EnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.depEnrollmentBaseProfile":
                        return NewDepEnrollmentBaseProfile(), nil
                    case "#microsoft.graph.depEnrollmentProfile":
                        return NewDepEnrollmentProfile(), nil
                    case "#microsoft.graph.depIOSEnrollmentProfile":
                        return NewDepIOSEnrollmentProfile(), nil
                    case "#microsoft.graph.depMacOSEnrollmentProfile":
                        return NewDepMacOSEnrollmentProfile(), nil
                }
            }
        }
    }
    return NewEnrollmentProfile(), nil
}
// GetConfigurationEndpointUrl gets the configurationEndpointUrl property value. Configuration endpoint url to use for Enrollment
func (m *EnrollmentProfile) GetConfigurationEndpointUrl()(*string) {
    return m.configurationEndpointUrl
}
// GetDescription gets the description property value. Description of the profile
func (m *EnrollmentProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *EnrollmentProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnableAuthenticationViaCompanyPortal gets the enableAuthenticationViaCompanyPortal property value. Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
func (m *EnrollmentProfile) GetEnableAuthenticationViaCompanyPortal()(*bool) {
    return m.enableAuthenticationViaCompanyPortal
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationEndpointUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationEndpointUrl(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["enableAuthenticationViaCompanyPortal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAuthenticationViaCompanyPortal(val)
        }
        return nil
    }
    res["requireCompanyPortalOnSetupAssistantEnrolledDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(val)
        }
        return nil
    }
    res["requiresUserAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresUserAuthentication(val)
        }
        return nil
    }
    return res
}
// GetRequireCompanyPortalOnSetupAssistantEnrolledDevices gets the requireCompanyPortalOnSetupAssistantEnrolledDevices property value. Indicates that Company Portal is required on setup assistant enrolled devices
func (m *EnrollmentProfile) GetRequireCompanyPortalOnSetupAssistantEnrolledDevices()(*bool) {
    return m.requireCompanyPortalOnSetupAssistantEnrolledDevices
}
// GetRequiresUserAuthentication gets the requiresUserAuthentication property value. Indicates if the profile requires user authentication
func (m *EnrollmentProfile) GetRequiresUserAuthentication()(*bool) {
    return m.requiresUserAuthentication
}
// Serialize serializes information the current object
func (m *EnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("configurationEndpointUrl", m.GetConfigurationEndpointUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableAuthenticationViaCompanyPortal", m.GetEnableAuthenticationViaCompanyPortal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireCompanyPortalOnSetupAssistantEnrolledDevices", m.GetRequireCompanyPortalOnSetupAssistantEnrolledDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresUserAuthentication", m.GetRequiresUserAuthentication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationEndpointUrl sets the configurationEndpointUrl property value. Configuration endpoint url to use for Enrollment
func (m *EnrollmentProfile) SetConfigurationEndpointUrl(value *string)() {
    m.configurationEndpointUrl = value
}
// SetDescription sets the description property value. Description of the profile
func (m *EnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *EnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnableAuthenticationViaCompanyPortal sets the enableAuthenticationViaCompanyPortal property value. Indicates to authenticate with Apple Setup Assistant instead of Company Portal.
func (m *EnrollmentProfile) SetEnableAuthenticationViaCompanyPortal(value *bool)() {
    m.enableAuthenticationViaCompanyPortal = value
}
// SetRequireCompanyPortalOnSetupAssistantEnrolledDevices sets the requireCompanyPortalOnSetupAssistantEnrolledDevices property value. Indicates that Company Portal is required on setup assistant enrolled devices
func (m *EnrollmentProfile) SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(value *bool)() {
    m.requireCompanyPortalOnSetupAssistantEnrolledDevices = value
}
// SetRequiresUserAuthentication sets the requiresUserAuthentication property value. Indicates if the profile requires user authentication
func (m *EnrollmentProfile) SetRequiresUserAuthentication(value *bool)() {
    m.requiresUserAuthentication = value
}
