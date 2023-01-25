package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPackageInformation contains properties for the package information for a Windows line of business app.
type WindowsPackageInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Contains properties for Windows architecture.
    applicableArchitecture *WindowsArchitecture
    // The Display Name.
    displayName *string
    // The Identity Name.
    identityName *string
    // The Identity Publisher.
    identityPublisher *string
    // The Identity Resource Identifier.
    identityResourceIdentifier *string
    // The Identity Version.
    identityVersion *string
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
    // The OdataType property
    odataType *string
}
// NewWindowsPackageInformation instantiates a new windowsPackageInformation and sets the default values.
func NewWindowsPackageInformation()(*WindowsPackageInformation) {
    m := &WindowsPackageInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsPackageInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPackageInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPackageInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPackageInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicableArchitecture gets the applicableArchitecture property value. Contains properties for Windows architecture.
func (m *WindowsPackageInformation) GetApplicableArchitecture()(*WindowsArchitecture) {
    return m.applicableArchitecture
}
// GetDisplayName gets the displayName property value. The Display Name.
func (m *WindowsPackageInformation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPackageInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicableArchitecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitecture(val.(*WindowsArchitecture))
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
    res["identityPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityPublisher(val)
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
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *WindowsPackageInformation) GetIdentityName()(*string) {
    return m.identityName
}
// GetIdentityPublisher gets the identityPublisher property value. The Identity Publisher.
func (m *WindowsPackageInformation) GetIdentityPublisher()(*string) {
    return m.identityPublisher
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPackageInformation) GetIdentityResourceIdentifier()(*string) {
    return m.identityResourceIdentifier
}
// GetIdentityVersion gets the identityVersion property value. The Identity Version.
func (m *WindowsPackageInformation) GetIdentityVersion()(*string) {
    return m.identityVersion
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *WindowsPackageInformation) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsPackageInformation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WindowsPackageInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApplicableArchitecture() != nil {
        cast := (*m.GetApplicableArchitecture()).String()
        err := writer.WriteStringValue("applicableArchitecture", &cast)
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
        err := writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identityPublisher", m.GetIdentityPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identityResourceIdentifier", m.GetIdentityResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
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
func (m *WindowsPackageInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicableArchitecture sets the applicableArchitecture property value. Contains properties for Windows architecture.
func (m *WindowsPackageInformation) SetApplicableArchitecture(value *WindowsArchitecture)() {
    m.applicableArchitecture = value
}
// SetDisplayName sets the displayName property value. The Display Name.
func (m *WindowsPackageInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *WindowsPackageInformation) SetIdentityName(value *string)() {
    m.identityName = value
}
// SetIdentityPublisher sets the identityPublisher property value. The Identity Publisher.
func (m *WindowsPackageInformation) SetIdentityPublisher(value *string)() {
    m.identityPublisher = value
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPackageInformation) SetIdentityResourceIdentifier(value *string)() {
    m.identityResourceIdentifier = value
}
// SetIdentityVersion sets the identityVersion property value. The Identity Version.
func (m *WindowsPackageInformation) SetIdentityVersion(value *string)() {
    m.identityVersion = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *WindowsPackageInformation) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsPackageInformation) SetOdataType(value *string)() {
    m.odataType = value
}
