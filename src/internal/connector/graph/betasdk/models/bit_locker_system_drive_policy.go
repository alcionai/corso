package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BitLockerSystemDrivePolicy bitLocker Encryption Base Policies.
type BitLockerSystemDrivePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Select the encryption method for operating system drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
    encryptionMethod *BitLockerEncryptionMethod
    // Indicates the minimum length of startup pin. Valid values 4 to 20
    minimumPinLength *int32
    // The OdataType property
    odataType *string
    // Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
    prebootRecoveryEnableMessageAndUrl *bool
    // Defines a custom recovery message.
    prebootRecoveryMessage *string
    // Defines a custom recovery URL.
    prebootRecoveryUrl *string
    // Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information. This policy setting is applied when you turn on BitLocker.
    recoveryOptions BitLockerRecoveryOptionsable
    // Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash drive).
    startupAuthenticationBlockWithoutTpmChip *bool
    // Require additional authentication at startup.
    startupAuthenticationRequired *bool
    // Possible values of the ConfigurationUsage list.
    startupAuthenticationTpmKeyUsage *ConfigurationUsage
    // Possible values of the ConfigurationUsage list.
    startupAuthenticationTpmPinAndKeyUsage *ConfigurationUsage
    // Possible values of the ConfigurationUsage list.
    startupAuthenticationTpmPinUsage *ConfigurationUsage
    // Possible values of the ConfigurationUsage list.
    startupAuthenticationTpmUsage *ConfigurationUsage
}
// NewBitLockerSystemDrivePolicy instantiates a new bitLockerSystemDrivePolicy and sets the default values.
func NewBitLockerSystemDrivePolicy()(*BitLockerSystemDrivePolicy) {
    m := &BitLockerSystemDrivePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBitLockerSystemDrivePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitLockerSystemDrivePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerSystemDrivePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerSystemDrivePolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEncryptionMethod gets the encryptionMethod property value. Select the encryption method for operating system drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerSystemDrivePolicy) GetEncryptionMethod()(*BitLockerEncryptionMethod) {
    return m.encryptionMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BitLockerSystemDrivePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encryptionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBitLockerEncryptionMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionMethod(val.(*BitLockerEncryptionMethod))
        }
        return nil
    }
    res["minimumPinLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumPinLength(val)
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
    res["prebootRecoveryEnableMessageAndUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebootRecoveryEnableMessageAndUrl(val)
        }
        return nil
    }
    res["prebootRecoveryMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebootRecoveryMessage(val)
        }
        return nil
    }
    res["prebootRecoveryUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrebootRecoveryUrl(val)
        }
        return nil
    }
    res["recoveryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerRecoveryOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecoveryOptions(val.(BitLockerRecoveryOptionsable))
        }
        return nil
    }
    res["startupAuthenticationBlockWithoutTpmChip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationBlockWithoutTpmChip(val)
        }
        return nil
    }
    res["startupAuthenticationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationRequired(val)
        }
        return nil
    }
    res["startupAuthenticationTpmKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmKeyUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["startupAuthenticationTpmPinAndKeyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmPinAndKeyUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["startupAuthenticationTpmPinUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmPinUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    res["startupAuthenticationTpmUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupAuthenticationTpmUsage(val.(*ConfigurationUsage))
        }
        return nil
    }
    return res
}
// GetMinimumPinLength gets the minimumPinLength property value. Indicates the minimum length of startup pin. Valid values 4 to 20
func (m *BitLockerSystemDrivePolicy) GetMinimumPinLength()(*int32) {
    return m.minimumPinLength
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BitLockerSystemDrivePolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetPrebootRecoveryEnableMessageAndUrl gets the prebootRecoveryEnableMessageAndUrl property value. Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
func (m *BitLockerSystemDrivePolicy) GetPrebootRecoveryEnableMessageAndUrl()(*bool) {
    return m.prebootRecoveryEnableMessageAndUrl
}
// GetPrebootRecoveryMessage gets the prebootRecoveryMessage property value. Defines a custom recovery message.
func (m *BitLockerSystemDrivePolicy) GetPrebootRecoveryMessage()(*string) {
    return m.prebootRecoveryMessage
}
// GetPrebootRecoveryUrl gets the prebootRecoveryUrl property value. Defines a custom recovery URL.
func (m *BitLockerSystemDrivePolicy) GetPrebootRecoveryUrl()(*string) {
    return m.prebootRecoveryUrl
}
// GetRecoveryOptions gets the recoveryOptions property value. Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerSystemDrivePolicy) GetRecoveryOptions()(BitLockerRecoveryOptionsable) {
    return m.recoveryOptions
}
// GetStartupAuthenticationBlockWithoutTpmChip gets the startupAuthenticationBlockWithoutTpmChip property value. Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash drive).
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationBlockWithoutTpmChip()(*bool) {
    return m.startupAuthenticationBlockWithoutTpmChip
}
// GetStartupAuthenticationRequired gets the startupAuthenticationRequired property value. Require additional authentication at startup.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationRequired()(*bool) {
    return m.startupAuthenticationRequired
}
// GetStartupAuthenticationTpmKeyUsage gets the startupAuthenticationTpmKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmKeyUsage()(*ConfigurationUsage) {
    return m.startupAuthenticationTpmKeyUsage
}
// GetStartupAuthenticationTpmPinAndKeyUsage gets the startupAuthenticationTpmPinAndKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmPinAndKeyUsage()(*ConfigurationUsage) {
    return m.startupAuthenticationTpmPinAndKeyUsage
}
// GetStartupAuthenticationTpmPinUsage gets the startupAuthenticationTpmPinUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmPinUsage()(*ConfigurationUsage) {
    return m.startupAuthenticationTpmPinUsage
}
// GetStartupAuthenticationTpmUsage gets the startupAuthenticationTpmUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) GetStartupAuthenticationTpmUsage()(*ConfigurationUsage) {
    return m.startupAuthenticationTpmUsage
}
// Serialize serializes information the current object
func (m *BitLockerSystemDrivePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumPinLength", m.GetMinimumPinLength())
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
        err := writer.WriteBoolValue("prebootRecoveryEnableMessageAndUrl", m.GetPrebootRecoveryEnableMessageAndUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prebootRecoveryMessage", m.GetPrebootRecoveryMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prebootRecoveryUrl", m.GetPrebootRecoveryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recoveryOptions", m.GetRecoveryOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("startupAuthenticationBlockWithoutTpmChip", m.GetStartupAuthenticationBlockWithoutTpmChip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("startupAuthenticationRequired", m.GetStartupAuthenticationRequired())
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmKeyUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmKeyUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmPinAndKeyUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmPinAndKeyUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmPinAndKeyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmPinUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmPinUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmPinUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartupAuthenticationTpmUsage() != nil {
        cast := (*m.GetStartupAuthenticationTpmUsage()).String()
        err := writer.WriteStringValue("startupAuthenticationTpmUsage", &cast)
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
func (m *BitLockerSystemDrivePolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEncryptionMethod sets the encryptionMethod property value. Select the encryption method for operating system drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerSystemDrivePolicy) SetEncryptionMethod(value *BitLockerEncryptionMethod)() {
    m.encryptionMethod = value
}
// SetMinimumPinLength sets the minimumPinLength property value. Indicates the minimum length of startup pin. Valid values 4 to 20
func (m *BitLockerSystemDrivePolicy) SetMinimumPinLength(value *int32)() {
    m.minimumPinLength = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerSystemDrivePolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrebootRecoveryEnableMessageAndUrl sets the prebootRecoveryEnableMessageAndUrl property value. Enable pre-boot recovery message and Url. If requireStartupAuthentication is false, this value does not affect.
func (m *BitLockerSystemDrivePolicy) SetPrebootRecoveryEnableMessageAndUrl(value *bool)() {
    m.prebootRecoveryEnableMessageAndUrl = value
}
// SetPrebootRecoveryMessage sets the prebootRecoveryMessage property value. Defines a custom recovery message.
func (m *BitLockerSystemDrivePolicy) SetPrebootRecoveryMessage(value *string)() {
    m.prebootRecoveryMessage = value
}
// SetPrebootRecoveryUrl sets the prebootRecoveryUrl property value. Defines a custom recovery URL.
func (m *BitLockerSystemDrivePolicy) SetPrebootRecoveryUrl(value *string)() {
    m.prebootRecoveryUrl = value
}
// SetRecoveryOptions sets the recoveryOptions property value. Allows to recover BitLocker encrypted operating system drives in the absence of the required startup key information. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerSystemDrivePolicy) SetRecoveryOptions(value BitLockerRecoveryOptionsable)() {
    m.recoveryOptions = value
}
// SetStartupAuthenticationBlockWithoutTpmChip sets the startupAuthenticationBlockWithoutTpmChip property value. Indicates whether to allow BitLocker without a compatible TPM (requires a password or a startup key on a USB flash drive).
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationBlockWithoutTpmChip(value *bool)() {
    m.startupAuthenticationBlockWithoutTpmChip = value
}
// SetStartupAuthenticationRequired sets the startupAuthenticationRequired property value. Require additional authentication at startup.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationRequired(value *bool)() {
    m.startupAuthenticationRequired = value
}
// SetStartupAuthenticationTpmKeyUsage sets the startupAuthenticationTpmKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmKeyUsage(value *ConfigurationUsage)() {
    m.startupAuthenticationTpmKeyUsage = value
}
// SetStartupAuthenticationTpmPinAndKeyUsage sets the startupAuthenticationTpmPinAndKeyUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmPinAndKeyUsage(value *ConfigurationUsage)() {
    m.startupAuthenticationTpmPinAndKeyUsage = value
}
// SetStartupAuthenticationTpmPinUsage sets the startupAuthenticationTpmPinUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmPinUsage(value *ConfigurationUsage)() {
    m.startupAuthenticationTpmPinUsage = value
}
// SetStartupAuthenticationTpmUsage sets the startupAuthenticationTpmUsage property value. Possible values of the ConfigurationUsage list.
func (m *BitLockerSystemDrivePolicy) SetStartupAuthenticationTpmUsage(value *ConfigurationUsage)() {
    m.startupAuthenticationTpmUsage = value
}
