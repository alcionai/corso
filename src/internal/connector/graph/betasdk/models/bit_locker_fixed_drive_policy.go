package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BitLockerFixedDrivePolicy bitLocker Fixed Drive Policies.
type BitLockerFixedDrivePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Select the encryption method for fixed drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
    encryptionMethod *BitLockerEncryptionMethod
    // The OdataType property
    odataType *string
    // This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of the required credentials. This policy setting is applied when you turn on BitLocker.
    recoveryOptions BitLockerRecoveryOptionsable
    // This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a computer.
    requireEncryptionForWriteAccess *bool
}
// NewBitLockerFixedDrivePolicy instantiates a new bitLockerFixedDrivePolicy and sets the default values.
func NewBitLockerFixedDrivePolicy()(*BitLockerFixedDrivePolicy) {
    m := &BitLockerFixedDrivePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBitLockerFixedDrivePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBitLockerFixedDrivePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBitLockerFixedDrivePolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BitLockerFixedDrivePolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEncryptionMethod gets the encryptionMethod property value. Select the encryption method for fixed drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerFixedDrivePolicy) GetEncryptionMethod()(*BitLockerEncryptionMethod) {
    return m.encryptionMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BitLockerFixedDrivePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["requireEncryptionForWriteAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireEncryptionForWriteAccess(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BitLockerFixedDrivePolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetRecoveryOptions gets the recoveryOptions property value. This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of the required credentials. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerFixedDrivePolicy) GetRecoveryOptions()(BitLockerRecoveryOptionsable) {
    return m.recoveryOptions
}
// GetRequireEncryptionForWriteAccess gets the requireEncryptionForWriteAccess property value. This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a computer.
func (m *BitLockerFixedDrivePolicy) GetRequireEncryptionForWriteAccess()(*bool) {
    return m.requireEncryptionForWriteAccess
}
// Serialize serializes information the current object
func (m *BitLockerFixedDrivePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
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
        err := writer.WriteObjectValue("recoveryOptions", m.GetRecoveryOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireEncryptionForWriteAccess", m.GetRequireEncryptionForWriteAccess())
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
func (m *BitLockerFixedDrivePolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEncryptionMethod sets the encryptionMethod property value. Select the encryption method for fixed drives. Possible values are: aesCbc128, aesCbc256, xtsAes128, xtsAes256.
func (m *BitLockerFixedDrivePolicy) SetEncryptionMethod(value *BitLockerEncryptionMethod)() {
    m.encryptionMethod = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BitLockerFixedDrivePolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecoveryOptions sets the recoveryOptions property value. This policy setting allows you to control how BitLocker-protected fixed data drives are recovered in the absence of the required credentials. This policy setting is applied when you turn on BitLocker.
func (m *BitLockerFixedDrivePolicy) SetRecoveryOptions(value BitLockerRecoveryOptionsable)() {
    m.recoveryOptions = value
}
// SetRequireEncryptionForWriteAccess sets the requireEncryptionForWriteAccess property value. This policy setting determines whether BitLocker protection is required for fixed data drives to be writable on a computer.
func (m *BitLockerFixedDrivePolicy) SetRequireEncryptionForWriteAccess(value *bool)() {
    m.requireEncryptionForWriteAccess = value
}
