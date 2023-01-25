package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HardwareInformation hardware information of a given device.
type HardwareInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
    batteryChargeCycles *int32
    // The device’s current battery’s health percentage. Valid values 0 to 100
    batteryHealthPercentage *int32
    // The battery level, between 0.0 and 100, or null if the battery level cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 5.0 and later, and is available only when Device Information access right is obtained. Valid values 0 to 100
    batteryLevelPercentage *float64
    // The serial number of the device’s current battery
    batterySerialNumber *string
    // Cellular technology of the device
    cellularTechnology *string
    // Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
    deviceFullQualifiedDomainName *string
    // The deviceGuardLocalSystemAuthorityCredentialGuardState property
    deviceGuardLocalSystemAuthorityCredentialGuardState *DeviceGuardLocalSystemAuthorityCredentialGuardState
    // The deviceGuardVirtualizationBasedSecurityHardwareRequirementState property
    deviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState
    // The deviceGuardVirtualizationBasedSecurityState property
    deviceGuardVirtualizationBasedSecurityState *DeviceGuardVirtualizationBasedSecurityState
    // A standard error code indicating the last error, or 0 indicating no error (default). The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing. Valid values 0 to 2147483647
    deviceLicensingLastErrorCode *int32
    // Error text message as a descripition for deviceLicensingLastErrorCode. The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing.
    deviceLicensingLastErrorDescription *string
    // Indicates the device licensing status after Windows device based subscription has been enabled.
    deviceLicensingStatus *DeviceLicensingStatus
    // eSIM identifier
    esimIdentifier *string
    // Free storage space of the device.
    freeStorageSpace *int64
    // IMEI
    imei *string
    // IPAddressV4
    ipAddressV4 *string
    // Encryption status of the device
    isEncrypted *bool
    // Shared iPad
    isSharedDevice *bool
    // Supervised mode of the device
    isSupervised *bool
    // Manufacturer of the device
    manufacturer *string
    // MEID
    meid *string
    // Model of the device
    model *string
    // The OdataType property
    odataType *string
    // String that specifies the OS edition.
    operatingSystemEdition *string
    // Operating system language of the device
    operatingSystemLanguage *string
    // Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
    operatingSystemProductType *int32
    // Operating System Build Number on Android device
    osBuildNumber *string
    // Phone number of the device
    phoneNumber *string
    // The product name, e.g. iPad8,12 etc. The update frequency of this property is weekly. Note this property is currently supported only on iOS/MacOS devices, and is available only when Device Information access right is obtained.
    productName *string
    // The number of users currently on this device, or null (default) if the value of this property cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 13.4 and later, and is available only when Device Information access right is obtained. Valid values 0 to 2147483647
    residentUsersCount *int32
    // Serial number.
    serialNumber *string
    // All users on the shared Apple device
    sharedDeviceCachedUsers []SharedAppleDeviceUserable
    // SubnetAddress
    subnetAddress *string
    // Subscriber carrier of the device
    subscriberCarrier *string
    // BIOS version as reported by SMBIOS
    systemManagementBIOSVersion *string
    // Total storage space of the device.
    totalStorageSpace *int64
    // The identifying information that uniquely names the TPM manufacturer
    tpmManufacturer *string
    // String that specifies the specification version.
    tpmSpecificationVersion *string
    // The version of the TPM, as specified by the manufacturer
    tpmVersion *string
    // WiFi MAC address of the device
    wifiMac *string
    // A list of wired IPv4 addresses. The update frequency (the maximum delay for the change of property value to be synchronized from the device to the cloud storage) of this property is daily. Note this property is currently supported only on devices running on Windows.
    wiredIPv4Addresses []string
}
// NewHardwareInformation instantiates a new hardwareInformation and sets the default values.
func NewHardwareInformation()(*HardwareInformation) {
    m := &HardwareInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateHardwareInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHardwareInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HardwareInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBatteryChargeCycles gets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
func (m *HardwareInformation) GetBatteryChargeCycles()(*int32) {
    return m.batteryChargeCycles
}
// GetBatteryHealthPercentage gets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
func (m *HardwareInformation) GetBatteryHealthPercentage()(*int32) {
    return m.batteryHealthPercentage
}
// GetBatteryLevelPercentage gets the batteryLevelPercentage property value. The battery level, between 0.0 and 100, or null if the battery level cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 5.0 and later, and is available only when Device Information access right is obtained. Valid values 0 to 100
func (m *HardwareInformation) GetBatteryLevelPercentage()(*float64) {
    return m.batteryLevelPercentage
}
// GetBatterySerialNumber gets the batterySerialNumber property value. The serial number of the device’s current battery
func (m *HardwareInformation) GetBatterySerialNumber()(*string) {
    return m.batterySerialNumber
}
// GetCellularTechnology gets the cellularTechnology property value. Cellular technology of the device
func (m *HardwareInformation) GetCellularTechnology()(*string) {
    return m.cellularTechnology
}
// GetDeviceFullQualifiedDomainName gets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
func (m *HardwareInformation) GetDeviceFullQualifiedDomainName()(*string) {
    return m.deviceFullQualifiedDomainName
}
// GetDeviceGuardLocalSystemAuthorityCredentialGuardState gets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. The deviceGuardLocalSystemAuthorityCredentialGuardState property
func (m *HardwareInformation) GetDeviceGuardLocalSystemAuthorityCredentialGuardState()(*DeviceGuardLocalSystemAuthorityCredentialGuardState) {
    return m.deviceGuardLocalSystemAuthorityCredentialGuardState
}
// GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState gets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. The deviceGuardVirtualizationBasedSecurityHardwareRequirementState property
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) {
    return m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState
}
// GetDeviceGuardVirtualizationBasedSecurityState gets the deviceGuardVirtualizationBasedSecurityState property value. The deviceGuardVirtualizationBasedSecurityState property
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityState()(*DeviceGuardVirtualizationBasedSecurityState) {
    return m.deviceGuardVirtualizationBasedSecurityState
}
// GetDeviceLicensingLastErrorCode gets the deviceLicensingLastErrorCode property value. A standard error code indicating the last error, or 0 indicating no error (default). The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing. Valid values 0 to 2147483647
func (m *HardwareInformation) GetDeviceLicensingLastErrorCode()(*int32) {
    return m.deviceLicensingLastErrorCode
}
// GetDeviceLicensingLastErrorDescription gets the deviceLicensingLastErrorDescription property value. Error text message as a descripition for deviceLicensingLastErrorCode. The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing.
func (m *HardwareInformation) GetDeviceLicensingLastErrorDescription()(*string) {
    return m.deviceLicensingLastErrorDescription
}
// GetDeviceLicensingStatus gets the deviceLicensingStatus property value. Indicates the device licensing status after Windows device based subscription has been enabled.
func (m *HardwareInformation) GetDeviceLicensingStatus()(*DeviceLicensingStatus) {
    return m.deviceLicensingStatus
}
// GetEsimIdentifier gets the esimIdentifier property value. eSIM identifier
func (m *HardwareInformation) GetEsimIdentifier()(*string) {
    return m.esimIdentifier
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwareInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["batteryChargeCycles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryChargeCycles(val)
        }
        return nil
    }
    res["batteryHealthPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthPercentage(val)
        }
        return nil
    }
    res["batteryLevelPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryLevelPercentage(val)
        }
        return nil
    }
    res["batterySerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatterySerialNumber(val)
        }
        return nil
    }
    res["cellularTechnology"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularTechnology(val)
        }
        return nil
    }
    res["deviceFullQualifiedDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFullQualifiedDomainName(val)
        }
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLocalSystemAuthorityCredentialGuardState(val.(*DeviceGuardLocalSystemAuthorityCredentialGuardState))
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityHardwareRequirementState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(val.(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState))
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardVirtualizationBasedSecurityState(val.(*DeviceGuardVirtualizationBasedSecurityState))
        }
        return nil
    }
    res["deviceLicensingLastErrorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLicensingLastErrorCode(val)
        }
        return nil
    }
    res["deviceLicensingLastErrorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLicensingLastErrorDescription(val)
        }
        return nil
    }
    res["deviceLicensingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceLicensingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLicensingStatus(val.(*DeviceLicensingStatus))
        }
        return nil
    }
    res["esimIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsimIdentifier(val)
        }
        return nil
    }
    res["freeStorageSpace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpace(val)
        }
        return nil
    }
    res["imei"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["ipAddressV4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressV4(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSharedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSharedDevice(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
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
    res["meid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
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
    res["operatingSystemEdition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemEdition(val)
        }
        return nil
    }
    res["operatingSystemLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemLanguage(val)
        }
        return nil
    }
    res["operatingSystemProductType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemProductType(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["residentUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidentUsersCount(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["sharedDeviceCachedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedAppleDeviceUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedAppleDeviceUserable, len(val))
            for i, v := range val {
                res[i] = v.(SharedAppleDeviceUserable)
            }
            m.SetSharedDeviceCachedUsers(res)
        }
        return nil
    }
    res["subnetAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetAddress(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["systemManagementBIOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemManagementBIOSVersion(val)
        }
        return nil
    }
    res["totalStorageSpace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpace(val)
        }
        return nil
    }
    res["tpmManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmManufacturer(val)
        }
        return nil
    }
    res["tpmSpecificationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmSpecificationVersion(val)
        }
        return nil
    }
    res["tpmVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmVersion(val)
        }
        return nil
    }
    res["wifiMac"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiMac(val)
        }
        return nil
    }
    res["wiredIPv4Addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetWiredIPv4Addresses(res)
        }
        return nil
    }
    return res
}
// GetFreeStorageSpace gets the freeStorageSpace property value. Free storage space of the device.
func (m *HardwareInformation) GetFreeStorageSpace()(*int64) {
    return m.freeStorageSpace
}
// GetImei gets the imei property value. IMEI
func (m *HardwareInformation) GetImei()(*string) {
    return m.imei
}
// GetIpAddressV4 gets the ipAddressV4 property value. IPAddressV4
func (m *HardwareInformation) GetIpAddressV4()(*string) {
    return m.ipAddressV4
}
// GetIsEncrypted gets the isEncrypted property value. Encryption status of the device
func (m *HardwareInformation) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetIsSharedDevice gets the isSharedDevice property value. Shared iPad
func (m *HardwareInformation) GetIsSharedDevice()(*bool) {
    return m.isSharedDevice
}
// GetIsSupervised gets the isSupervised property value. Supervised mode of the device
func (m *HardwareInformation) GetIsSupervised()(*bool) {
    return m.isSupervised
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device
func (m *HardwareInformation) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMeid gets the meid property value. MEID
func (m *HardwareInformation) GetMeid()(*string) {
    return m.meid
}
// GetModel gets the model property value. Model of the device
func (m *HardwareInformation) GetModel()(*string) {
    return m.model
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *HardwareInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatingSystemEdition gets the operatingSystemEdition property value. String that specifies the OS edition.
func (m *HardwareInformation) GetOperatingSystemEdition()(*string) {
    return m.operatingSystemEdition
}
// GetOperatingSystemLanguage gets the operatingSystemLanguage property value. Operating system language of the device
func (m *HardwareInformation) GetOperatingSystemLanguage()(*string) {
    return m.operatingSystemLanguage
}
// GetOperatingSystemProductType gets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
func (m *HardwareInformation) GetOperatingSystemProductType()(*int32) {
    return m.operatingSystemProductType
}
// GetOsBuildNumber gets the osBuildNumber property value. Operating System Build Number on Android device
func (m *HardwareInformation) GetOsBuildNumber()(*string) {
    return m.osBuildNumber
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device
func (m *HardwareInformation) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetProductName gets the productName property value. The product name, e.g. iPad8,12 etc. The update frequency of this property is weekly. Note this property is currently supported only on iOS/MacOS devices, and is available only when Device Information access right is obtained.
func (m *HardwareInformation) GetProductName()(*string) {
    return m.productName
}
// GetResidentUsersCount gets the residentUsersCount property value. The number of users currently on this device, or null (default) if the value of this property cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 13.4 and later, and is available only when Device Information access right is obtained. Valid values 0 to 2147483647
func (m *HardwareInformation) GetResidentUsersCount()(*int32) {
    return m.residentUsersCount
}
// GetSerialNumber gets the serialNumber property value. Serial number.
func (m *HardwareInformation) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSharedDeviceCachedUsers gets the sharedDeviceCachedUsers property value. All users on the shared Apple device
func (m *HardwareInformation) GetSharedDeviceCachedUsers()([]SharedAppleDeviceUserable) {
    return m.sharedDeviceCachedUsers
}
// GetSubnetAddress gets the subnetAddress property value. SubnetAddress
func (m *HardwareInformation) GetSubnetAddress()(*string) {
    return m.subnetAddress
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber carrier of the device
func (m *HardwareInformation) GetSubscriberCarrier()(*string) {
    return m.subscriberCarrier
}
// GetSystemManagementBIOSVersion gets the systemManagementBIOSVersion property value. BIOS version as reported by SMBIOS
func (m *HardwareInformation) GetSystemManagementBIOSVersion()(*string) {
    return m.systemManagementBIOSVersion
}
// GetTotalStorageSpace gets the totalStorageSpace property value. Total storage space of the device.
func (m *HardwareInformation) GetTotalStorageSpace()(*int64) {
    return m.totalStorageSpace
}
// GetTpmManufacturer gets the tpmManufacturer property value. The identifying information that uniquely names the TPM manufacturer
func (m *HardwareInformation) GetTpmManufacturer()(*string) {
    return m.tpmManufacturer
}
// GetTpmSpecificationVersion gets the tpmSpecificationVersion property value. String that specifies the specification version.
func (m *HardwareInformation) GetTpmSpecificationVersion()(*string) {
    return m.tpmSpecificationVersion
}
// GetTpmVersion gets the tpmVersion property value. The version of the TPM, as specified by the manufacturer
func (m *HardwareInformation) GetTpmVersion()(*string) {
    return m.tpmVersion
}
// GetWifiMac gets the wifiMac property value. WiFi MAC address of the device
func (m *HardwareInformation) GetWifiMac()(*string) {
    return m.wifiMac
}
// GetWiredIPv4Addresses gets the wiredIPv4Addresses property value. A list of wired IPv4 addresses. The update frequency (the maximum delay for the change of property value to be synchronized from the device to the cloud storage) of this property is daily. Note this property is currently supported only on devices running on Windows.
func (m *HardwareInformation) GetWiredIPv4Addresses()([]string) {
    return m.wiredIPv4Addresses
}
// Serialize serializes information the current object
func (m *HardwareInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("batteryChargeCycles", m.GetBatteryChargeCycles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("batteryHealthPercentage", m.GetBatteryHealthPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("batteryLevelPercentage", m.GetBatteryLevelPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("batterySerialNumber", m.GetBatterySerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cellularTechnology", m.GetCellularTechnology())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceFullQualifiedDomainName", m.GetDeviceFullQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState() != nil {
        cast := (*m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState()).String()
        err := writer.WriteStringValue("deviceGuardLocalSystemAuthorityCredentialGuardState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() != nil {
        cast := (*m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()).String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityHardwareRequirementState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityState() != nil {
        cast := (*m.GetDeviceGuardVirtualizationBasedSecurityState()).String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceLicensingLastErrorCode", m.GetDeviceLicensingLastErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceLicensingLastErrorDescription", m.GetDeviceLicensingLastErrorDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceLicensingStatus() != nil {
        cast := (*m.GetDeviceLicensingStatus()).String()
        err := writer.WriteStringValue("deviceLicensingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("esimIdentifier", m.GetEsimIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("freeStorageSpace", m.GetFreeStorageSpace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("imei", m.GetImei())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddressV4", m.GetIpAddressV4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSharedDevice", m.GetIsSharedDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("meid", m.GetMeid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
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
        err := writer.WriteStringValue("operatingSystemEdition", m.GetOperatingSystemEdition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemLanguage", m.GetOperatingSystemLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("operatingSystemProductType", m.GetOperatingSystemProductType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("residentUsersCount", m.GetResidentUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetSharedDeviceCachedUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedDeviceCachedUsers()))
        for i, v := range m.GetSharedDeviceCachedUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sharedDeviceCachedUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetAddress", m.GetSubnetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriberCarrier", m.GetSubscriberCarrier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("systemManagementBIOSVersion", m.GetSystemManagementBIOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalStorageSpace", m.GetTotalStorageSpace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmManufacturer", m.GetTpmManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmSpecificationVersion", m.GetTpmSpecificationVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmVersion", m.GetTpmVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMac", m.GetWifiMac())
        if err != nil {
            return err
        }
    }
    if m.GetWiredIPv4Addresses() != nil {
        err := writer.WriteCollectionOfStringValues("wiredIPv4Addresses", m.GetWiredIPv4Addresses())
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
func (m *HardwareInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBatteryChargeCycles sets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
func (m *HardwareInformation) SetBatteryChargeCycles(value *int32)() {
    m.batteryChargeCycles = value
}
// SetBatteryHealthPercentage sets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
func (m *HardwareInformation) SetBatteryHealthPercentage(value *int32)() {
    m.batteryHealthPercentage = value
}
// SetBatteryLevelPercentage sets the batteryLevelPercentage property value. The battery level, between 0.0 and 100, or null if the battery level cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 5.0 and later, and is available only when Device Information access right is obtained. Valid values 0 to 100
func (m *HardwareInformation) SetBatteryLevelPercentage(value *float64)() {
    m.batteryLevelPercentage = value
}
// SetBatterySerialNumber sets the batterySerialNumber property value. The serial number of the device’s current battery
func (m *HardwareInformation) SetBatterySerialNumber(value *string)() {
    m.batterySerialNumber = value
}
// SetCellularTechnology sets the cellularTechnology property value. Cellular technology of the device
func (m *HardwareInformation) SetCellularTechnology(value *string)() {
    m.cellularTechnology = value
}
// SetDeviceFullQualifiedDomainName sets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
func (m *HardwareInformation) SetDeviceFullQualifiedDomainName(value *string)() {
    m.deviceFullQualifiedDomainName = value
}
// SetDeviceGuardLocalSystemAuthorityCredentialGuardState sets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. The deviceGuardLocalSystemAuthorityCredentialGuardState property
func (m *HardwareInformation) SetDeviceGuardLocalSystemAuthorityCredentialGuardState(value *DeviceGuardLocalSystemAuthorityCredentialGuardState)() {
    m.deviceGuardLocalSystemAuthorityCredentialGuardState = value
}
// SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState sets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. The deviceGuardVirtualizationBasedSecurityHardwareRequirementState property
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(value *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)() {
    m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState = value
}
// SetDeviceGuardVirtualizationBasedSecurityState sets the deviceGuardVirtualizationBasedSecurityState property value. The deviceGuardVirtualizationBasedSecurityState property
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityState(value *DeviceGuardVirtualizationBasedSecurityState)() {
    m.deviceGuardVirtualizationBasedSecurityState = value
}
// SetDeviceLicensingLastErrorCode sets the deviceLicensingLastErrorCode property value. A standard error code indicating the last error, or 0 indicating no error (default). The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing. Valid values 0 to 2147483647
func (m *HardwareInformation) SetDeviceLicensingLastErrorCode(value *int32)() {
    m.deviceLicensingLastErrorCode = value
}
// SetDeviceLicensingLastErrorDescription sets the deviceLicensingLastErrorDescription property value. Error text message as a descripition for deviceLicensingLastErrorCode. The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing.
func (m *HardwareInformation) SetDeviceLicensingLastErrorDescription(value *string)() {
    m.deviceLicensingLastErrorDescription = value
}
// SetDeviceLicensingStatus sets the deviceLicensingStatus property value. Indicates the device licensing status after Windows device based subscription has been enabled.
func (m *HardwareInformation) SetDeviceLicensingStatus(value *DeviceLicensingStatus)() {
    m.deviceLicensingStatus = value
}
// SetEsimIdentifier sets the esimIdentifier property value. eSIM identifier
func (m *HardwareInformation) SetEsimIdentifier(value *string)() {
    m.esimIdentifier = value
}
// SetFreeStorageSpace sets the freeStorageSpace property value. Free storage space of the device.
func (m *HardwareInformation) SetFreeStorageSpace(value *int64)() {
    m.freeStorageSpace = value
}
// SetImei sets the imei property value. IMEI
func (m *HardwareInformation) SetImei(value *string)() {
    m.imei = value
}
// SetIpAddressV4 sets the ipAddressV4 property value. IPAddressV4
func (m *HardwareInformation) SetIpAddressV4(value *string)() {
    m.ipAddressV4 = value
}
// SetIsEncrypted sets the isEncrypted property value. Encryption status of the device
func (m *HardwareInformation) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetIsSharedDevice sets the isSharedDevice property value. Shared iPad
func (m *HardwareInformation) SetIsSharedDevice(value *bool)() {
    m.isSharedDevice = value
}
// SetIsSupervised sets the isSupervised property value. Supervised mode of the device
func (m *HardwareInformation) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device
func (m *HardwareInformation) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMeid sets the meid property value. MEID
func (m *HardwareInformation) SetMeid(value *string)() {
    m.meid = value
}
// SetModel sets the model property value. Model of the device
func (m *HardwareInformation) SetModel(value *string)() {
    m.model = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HardwareInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatingSystemEdition sets the operatingSystemEdition property value. String that specifies the OS edition.
func (m *HardwareInformation) SetOperatingSystemEdition(value *string)() {
    m.operatingSystemEdition = value
}
// SetOperatingSystemLanguage sets the operatingSystemLanguage property value. Operating system language of the device
func (m *HardwareInformation) SetOperatingSystemLanguage(value *string)() {
    m.operatingSystemLanguage = value
}
// SetOperatingSystemProductType sets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
func (m *HardwareInformation) SetOperatingSystemProductType(value *int32)() {
    m.operatingSystemProductType = value
}
// SetOsBuildNumber sets the osBuildNumber property value. Operating System Build Number on Android device
func (m *HardwareInformation) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device
func (m *HardwareInformation) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetProductName sets the productName property value. The product name, e.g. iPad8,12 etc. The update frequency of this property is weekly. Note this property is currently supported only on iOS/MacOS devices, and is available only when Device Information access right is obtained.
func (m *HardwareInformation) SetProductName(value *string)() {
    m.productName = value
}
// SetResidentUsersCount sets the residentUsersCount property value. The number of users currently on this device, or null (default) if the value of this property cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 13.4 and later, and is available only when Device Information access right is obtained. Valid values 0 to 2147483647
func (m *HardwareInformation) SetResidentUsersCount(value *int32)() {
    m.residentUsersCount = value
}
// SetSerialNumber sets the serialNumber property value. Serial number.
func (m *HardwareInformation) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSharedDeviceCachedUsers sets the sharedDeviceCachedUsers property value. All users on the shared Apple device
func (m *HardwareInformation) SetSharedDeviceCachedUsers(value []SharedAppleDeviceUserable)() {
    m.sharedDeviceCachedUsers = value
}
// SetSubnetAddress sets the subnetAddress property value. SubnetAddress
func (m *HardwareInformation) SetSubnetAddress(value *string)() {
    m.subnetAddress = value
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber carrier of the device
func (m *HardwareInformation) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// SetSystemManagementBIOSVersion sets the systemManagementBIOSVersion property value. BIOS version as reported by SMBIOS
func (m *HardwareInformation) SetSystemManagementBIOSVersion(value *string)() {
    m.systemManagementBIOSVersion = value
}
// SetTotalStorageSpace sets the totalStorageSpace property value. Total storage space of the device.
func (m *HardwareInformation) SetTotalStorageSpace(value *int64)() {
    m.totalStorageSpace = value
}
// SetTpmManufacturer sets the tpmManufacturer property value. The identifying information that uniquely names the TPM manufacturer
func (m *HardwareInformation) SetTpmManufacturer(value *string)() {
    m.tpmManufacturer = value
}
// SetTpmSpecificationVersion sets the tpmSpecificationVersion property value. String that specifies the specification version.
func (m *HardwareInformation) SetTpmSpecificationVersion(value *string)() {
    m.tpmSpecificationVersion = value
}
// SetTpmVersion sets the tpmVersion property value. The version of the TPM, as specified by the manufacturer
func (m *HardwareInformation) SetTpmVersion(value *string)() {
    m.tpmVersion = value
}
// SetWifiMac sets the wifiMac property value. WiFi MAC address of the device
func (m *HardwareInformation) SetWifiMac(value *string)() {
    m.wifiMac = value
}
// SetWiredIPv4Addresses sets the wiredIPv4Addresses property value. A list of wired IPv4 addresses. The update frequency (the maximum delay for the change of property value to be synchronized from the device to the cloud storage) of this property is daily. Note this property is currently supported only on devices running on Windows.
func (m *HardwareInformation) SetWiredIPv4Addresses(value []string)() {
    m.wiredIPv4Addresses = value
}
