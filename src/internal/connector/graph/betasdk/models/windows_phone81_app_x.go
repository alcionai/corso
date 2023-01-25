package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81AppX 
type WindowsPhone81AppX struct {
    MobileLobApp
    // Contains properties for Windows architecture.
    applicableArchitectures *WindowsArchitecture
    // The Identity Name.
    identityName *string
    // The Identity Publisher Hash.
    identityPublisherHash *string
    // The Identity Resource Identifier.
    identityResourceIdentifier *string
    // The identity version.
    identityVersion *string
    // The minimum operating system required for a Windows mobile app.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
    // The Phone Product Identifier.
    phoneProductIdentifier *string
    // The Phone Publisher Id.
    phonePublisherId *string
}
// NewWindowsPhone81AppX instantiates a new WindowsPhone81AppX and sets the default values.
func NewWindowsPhone81AppX()(*WindowsPhone81AppX) {
    m := &WindowsPhone81AppX{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81AppX";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhone81AppXFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81AppXFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsPhone81AppXBundle":
                        return NewWindowsPhone81AppXBundle(), nil
                }
            }
        }
    }
    return NewWindowsPhone81AppX(), nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsPhone81AppX) GetApplicableArchitectures()(*WindowsArchitecture) {
    return m.applicableArchitectures
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81AppX) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["identityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityName(val)
        }
        return nil
    }
    res["identityPublisherHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityPublisherHash(val)
        }
        return nil
    }
    res["identityResourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityResourceIdentifier(val)
        }
        return nil
    }
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    res["phoneProductIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneProductIdentifier(val)
        }
        return nil
    }
    res["phonePublisherId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhonePublisherId(val)
        }
        return nil
    }
    return res
}
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *WindowsPhone81AppX) GetIdentityName()(*string) {
    return m.identityName
}
// GetIdentityPublisherHash gets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsPhone81AppX) GetIdentityPublisherHash()(*string) {
    return m.identityPublisherHash
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPhone81AppX) GetIdentityResourceIdentifier()(*string) {
    return m.identityResourceIdentifier
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *WindowsPhone81AppX) GetIdentityVersion()(*string) {
    return m.identityVersion
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhone81AppX) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetPhoneProductIdentifier gets the phoneProductIdentifier property value. The Phone Product Identifier.
func (m *WindowsPhone81AppX) GetPhoneProductIdentifier()(*string) {
    return m.phoneProductIdentifier
}
// GetPhonePublisherId gets the phonePublisherId property value. The Phone Publisher Id.
func (m *WindowsPhone81AppX) GetPhonePublisherId()(*string) {
    return m.phonePublisherId
}
// Serialize serializes information the current object
func (m *WindowsPhone81AppX) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityPublisherHash", m.GetIdentityPublisherHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityResourceIdentifier", m.GetIdentityResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneProductIdentifier", m.GetPhoneProductIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phonePublisherId", m.GetPhonePublisherId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsPhone81AppX) SetApplicableArchitectures(value *WindowsArchitecture)() {
    m.applicableArchitectures = value
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *WindowsPhone81AppX) SetIdentityName(value *string)() {
    m.identityName = value
}
// SetIdentityPublisherHash sets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsPhone81AppX) SetIdentityPublisherHash(value *string)() {
    m.identityPublisherHash = value
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPhone81AppX) SetIdentityResourceIdentifier(value *string)() {
    m.identityResourceIdentifier = value
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsPhone81AppX) SetIdentityVersion(value *string)() {
    m.identityVersion = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhone81AppX) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetPhoneProductIdentifier sets the phoneProductIdentifier property value. The Phone Product Identifier.
func (m *WindowsPhone81AppX) SetPhoneProductIdentifier(value *string)() {
    m.phoneProductIdentifier = value
}
// SetPhonePublisherId sets the phonePublisherId property value. The Phone Publisher Id.
func (m *WindowsPhone81AppX) SetPhonePublisherId(value *string)() {
    m.phonePublisherId = value
}
