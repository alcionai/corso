package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Device provides operations to manage the collection of site entities.
type Device struct {
    DirectoryObject
    // true if the account is enabled; otherwise, false. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
    accountEnabled *bool
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
    alternativeSecurityIds []AlternativeSecurityIdable
    // The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
    approximateLastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Set of commands sent to this device.
    commands []Commandable
    // The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    complianceExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
    deviceCategory *string
    // Unique Identifier set by Azure Device Registration Service at the time of registration. This is an alternate key that can be used to reference the device object. Also Supports $filter (eq, ne, not, startsWith).
    deviceId *string
    // For internal use only. Set to null.
    deviceMetadata *string
    // Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
    deviceOwnership *string
    // For internal use only.
    deviceVersion *int32
    // The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
    domainName *string
    // Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
    enrollmentProfileName *string
    // Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
    enrollmentType *string
    // Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
    extensionAttributes OnPremisesExtensionAttributesable
    // The collection of open extensions defined for the device. Read-only. Nullable.
    extensions []Extensionable
    // List of hostNames for the device.
    hostnames []string
    // true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isCompliant *bool
    // true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isManaged *bool
    // The isManagementRestricted property
    isManagementRestricted *bool
    // true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
    isRooted *bool
    // Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    kind *string
    // Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementType *string
    // Manufacturer of the device. Read-only.
    manufacturer *string
    // Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
    mdmAppId *string
    // Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // Model of the device. Read-only.
    model *string
    // Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    name *string
    // The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool
    // The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystem *string
    // Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystemVersion *string
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0.
    physicalIds []string
    // Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
    platform *string
    // The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
    profileType *string
    // The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
    registeredOwners []DirectoryObjectable
    // Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
    registeredUsers []DirectoryObjectable
    // Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
    status *string
    // List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
    systemLabels []string
    // Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
    transitiveMemberOf []DirectoryObjectable
    // Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
    trustType *string
    // Represents the usage rights a device has been granted.
    usageRights []UsageRightable
}
// NewDevice instantiates a new device and sets the default values.
func NewDevice()(*Device) {
    m := &Device{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.device";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevice(), nil
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetAlternativeSecurityIds gets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) GetAlternativeSecurityIds()([]AlternativeSecurityIdable) {
    return m.alternativeSecurityIds
}
// GetApproximateLastSignInDateTime gets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.approximateLastSignInDateTime
}
// GetCommands gets the commands property value. Set of commands sent to this device.
func (m *Device) GetCommands()([]Commandable) {
    return m.commands
}
// GetComplianceExpirationDateTime gets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceExpirationDateTime
}
// GetDeviceCategory gets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) GetDeviceCategory()(*string) {
    return m.deviceCategory
}
// GetDeviceId gets the deviceId property value. Unique Identifier set by Azure Device Registration Service at the time of registration. This is an alternate key that can be used to reference the device object. Also Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceMetadata gets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) GetDeviceMetadata()(*string) {
    return m.deviceMetadata
}
// GetDeviceOwnership gets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) GetDeviceOwnership()(*string) {
    return m.deviceOwnership
}
// GetDeviceVersion gets the deviceVersion property value. For internal use only.
func (m *Device) GetDeviceVersion()(*int32) {
    return m.deviceVersion
}
// GetDisplayName gets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) GetDisplayName()(*string) {
    return m.displayName
}
// GetDomainName gets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) GetDomainName()(*string) {
    return m.domainName
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) GetEnrollmentProfileName()(*string) {
    return m.enrollmentProfileName
}
// GetEnrollmentType gets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) GetEnrollmentType()(*string) {
    return m.enrollmentType
}
// GetExtensionAttributes gets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
func (m *Device) GetExtensionAttributes()(OnPremisesExtensionAttributesable) {
    return m.extensionAttributes
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Device) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["alternativeSecurityIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlternativeSecurityIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlternativeSecurityIdable, len(val))
            for i, v := range val {
                res[i] = v.(AlternativeSecurityIdable)
            }
            m.SetAlternativeSecurityIds(res)
        }
        return nil
    }
    res["approximateLastSignInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproximateLastSignInDateTime(val)
        }
        return nil
    }
    res["commands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Commandable, len(val))
            for i, v := range val {
                res[i] = v.(Commandable)
            }
            m.SetCommands(res)
        }
        return nil
    }
    res["complianceExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceExpirationDateTime(val)
        }
        return nil
    }
    res["deviceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMetadata(val)
        }
        return nil
    }
    res["deviceOwnership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnership(val)
        }
        return nil
    }
    res["deviceVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceVersion(val)
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
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["enrollmentProfileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfileName(val)
        }
        return nil
    }
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val)
        }
        return nil
    }
    res["extensionAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesExtensionAttributesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttributes(val.(OnPremisesExtensionAttributesable))
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["hostnames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHostnames(res)
        }
        return nil
    }
    res["isCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompliant(val)
        }
        return nil
    }
    res["isManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsManaged(val)
        }
        return nil
    }
    res["isManagementRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsManagementRestricted(val)
        }
        return nil
    }
    res["isRooted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRooted(val)
        }
        return nil
    }
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["managementType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementType(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["mdmAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmAppId(val)
        }
        return nil
    }
    res["memberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["operatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["physicalIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPhysicalIds(res)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val)
        }
        return nil
    }
    res["profileType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileType(val)
        }
        return nil
    }
    res["registeredOwners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetRegisteredOwners(res)
        }
        return nil
    }
    res["registeredUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetRegisteredUsers(res)
        }
        return nil
    }
    res["registrationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["systemLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSystemLabels(res)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["trustType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustType(val)
        }
        return nil
    }
    res["usageRights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUsageRightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UsageRightable, len(val))
            for i, v := range val {
                res[i] = v.(UsageRightable)
            }
            m.SetUsageRights(res)
        }
        return nil
    }
    return res
}
// GetHostnames gets the hostnames property value. List of hostNames for the device.
func (m *Device) GetHostnames()([]string) {
    return m.hostnames
}
// GetIsCompliant gets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsCompliant()(*bool) {
    return m.isCompliant
}
// GetIsManaged gets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsManaged()(*bool) {
    return m.isManaged
}
// GetIsManagementRestricted gets the isManagementRestricted property value. The isManagementRestricted property
func (m *Device) GetIsManagementRestricted()(*bool) {
    return m.isManagementRestricted
}
// GetIsRooted gets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) GetIsRooted()(*bool) {
    return m.isRooted
}
// GetKind gets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetKind()(*string) {
    return m.kind
}
// GetManagementType gets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) GetManagementType()(*string) {
    return m.managementType
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMdmAppId gets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetMdmAppId()(*string) {
    return m.mdmAppId
}
// GetMemberOf gets the memberOf property value. Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetModel gets the model property value. Model of the device. Read-only.
func (m *Device) GetModel()(*string) {
    return m.model
}
// GetName gets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetName()(*string) {
    return m.name
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetOperatingSystem gets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystemVersion()(*string) {
    return m.operatingSystemVersion
}
// GetPhysicalIds gets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0.
func (m *Device) GetPhysicalIds()([]string) {
    return m.physicalIds
}
// GetPlatform gets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetPlatform()(*string) {
    return m.platform
}
// GetProfileType gets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) GetProfileType()(*string) {
    return m.profileType
}
// GetRegisteredOwners gets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredOwners()([]DirectoryObjectable) {
    return m.registeredOwners
}
// GetRegisteredUsers gets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredUsers()([]DirectoryObjectable) {
    return m.registeredUsers
}
// GetRegistrationDateTime gets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.registrationDateTime
}
// GetStatus gets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetStatus()(*string) {
    return m.status
}
// GetSystemLabels gets the systemLabels property value. List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
func (m *Device) GetSystemLabels()([]string) {
    return m.systemLabels
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *Device) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetTrustType gets the trustType property value. Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) GetTrustType()(*string) {
    return m.trustType
}
// GetUsageRights gets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) GetUsageRights()([]UsageRightable) {
    return m.usageRights
}
// Serialize serializes information the current object
func (m *Device) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAlternativeSecurityIds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlternativeSecurityIds()))
        for i, v := range m.GetAlternativeSecurityIds() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alternativeSecurityIds", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("approximateLastSignInDateTime", m.GetApproximateLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCommands() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("commands", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("complianceExpirationDateTime", m.GetComplianceExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceMetadata", m.GetDeviceMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOwnership", m.GetDeviceOwnership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceVersion", m.GetDeviceVersion())
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
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentProfileName", m.GetEnrollmentProfileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentType", m.GetEnrollmentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("extensionAttributes", m.GetExtensionAttributes())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHostnames() != nil {
        err = writer.WriteCollectionOfStringValues("hostnames", m.GetHostnames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCompliant", m.GetIsCompliant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManaged", m.GetIsManaged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManagementRestricted", m.GetIsManagementRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRooted", m.GetIsRooted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementType", m.GetManagementType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmAppId", m.GetMdmAppId())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPhysicalIds() != nil {
        err = writer.WriteCollectionOfStringValues("physicalIds", m.GetPhysicalIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("platform", m.GetPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileType", m.GetProfileType())
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegisteredOwners()))
        for i, v := range m.GetRegisteredOwners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("registeredOwners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegisteredUsers()))
        for i, v := range m.GetRegisteredUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("registeredUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetSystemLabels() != nil {
        err = writer.WriteCollectionOfStringValues("systemLabels", m.GetSystemLabels())
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("trustType", m.GetTrustType())
        if err != nil {
            return err
        }
    }
    if m.GetUsageRights() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsageRights()))
        for i, v := range m.GetUsageRights() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("usageRights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetAlternativeSecurityIds sets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) SetAlternativeSecurityIds(value []AlternativeSecurityIdable)() {
    m.alternativeSecurityIds = value
}
// SetApproximateLastSignInDateTime sets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.approximateLastSignInDateTime = value
}
// SetCommands sets the commands property value. Set of commands sent to this device.
func (m *Device) SetCommands(value []Commandable)() {
    m.commands = value
}
// SetComplianceExpirationDateTime sets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceExpirationDateTime = value
}
// SetDeviceCategory sets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) SetDeviceCategory(value *string)() {
    m.deviceCategory = value
}
// SetDeviceId sets the deviceId property value. Unique Identifier set by Azure Device Registration Service at the time of registration. This is an alternate key that can be used to reference the device object. Also Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceMetadata sets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) SetDeviceMetadata(value *string)() {
    m.deviceMetadata = value
}
// SetDeviceOwnership sets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) SetDeviceOwnership(value *string)() {
    m.deviceOwnership = value
}
// SetDeviceVersion sets the deviceVersion property value. For internal use only.
func (m *Device) SetDeviceVersion(value *int32)() {
    m.deviceVersion = value
}
// SetDisplayName sets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDomainName sets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) SetDomainName(value *string)() {
    m.domainName = value
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) SetEnrollmentProfileName(value *string)() {
    m.enrollmentProfileName = value
}
// SetEnrollmentType sets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) SetEnrollmentType(value *string)() {
    m.enrollmentType = value
}
// SetExtensionAttributes sets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, not, startsWith, and eq on null values).
func (m *Device) SetExtensionAttributes(value OnPremisesExtensionAttributesable)() {
    m.extensionAttributes = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetHostnames sets the hostnames property value. List of hostNames for the device.
func (m *Device) SetHostnames(value []string)() {
    m.hostnames = value
}
// SetIsCompliant sets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsCompliant(value *bool)() {
    m.isCompliant = value
}
// SetIsManaged sets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsManaged(value *bool)() {
    m.isManaged = value
}
// SetIsManagementRestricted sets the isManagementRestricted property value. The isManagementRestricted property
func (m *Device) SetIsManagementRestricted(value *bool)() {
    m.isManagementRestricted = value
}
// SetIsRooted sets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) SetIsRooted(value *bool)() {
    m.isRooted = value
}
// SetKind sets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetKind(value *string)() {
    m.kind = value
}
// SetManagementType sets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) SetManagementType(value *string)() {
    m.managementType = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMdmAppId sets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetMdmAppId(value *string)() {
    m.mdmAppId = value
}
// SetMemberOf sets the memberOf property value. Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetModel sets the model property value. Model of the device. Read-only.
func (m *Device) SetModel(value *string)() {
    m.model = value
}
// SetName sets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetName(value *string)() {
    m.name = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOperatingSystem sets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. Operating system version of the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
// SetPhysicalIds sets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith, /$count eq 0, /$count ne 0.
func (m *Device) SetPhysicalIds(value []string)() {
    m.physicalIds = value
}
// SetPlatform sets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetPlatform(value *string)() {
    m.platform = value
}
// SetProfileType sets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) SetProfileType(value *string)() {
    m.profileType = value
}
// SetRegisteredOwners sets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredOwners(value []DirectoryObjectable)() {
    m.registeredOwners = value
}
// SetRegisteredUsers sets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredUsers(value []DirectoryObjectable)() {
    m.registeredUsers = value
}
// SetRegistrationDateTime sets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// SetStatus sets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) SetStatus(value *string)() {
    m.status = value
}
// SetSystemLabels sets the systemLabels property value. List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
func (m *Device) SetSystemLabels(value []string)() {
    m.systemLabels = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. Groups and administrative units that this device is a member of. This operation is transitive. Supports $expand.
func (m *Device) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetTrustType sets the trustType property value. Type of trust for the joined device. Read-only. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) SetTrustType(value *string)() {
    m.trustType = value
}
// SetUsageRights sets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) SetUsageRights(value []UsageRightable)() {
    m.usageRights = value
}
