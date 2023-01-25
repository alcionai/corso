package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleOwnerTypeEnrollmentType 
type AppleOwnerTypeEnrollmentType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enrollmentType property
    enrollmentType *AppleUserInitiatedEnrollmentType
    // The OdataType property
    odataType *string
    // Owner type of device.
    ownerType *ManagedDeviceOwnerType
}
// NewAppleOwnerTypeEnrollmentType instantiates a new appleOwnerTypeEnrollmentType and sets the default values.
func NewAppleOwnerTypeEnrollmentType()(*AppleOwnerTypeEnrollmentType) {
    m := &AppleOwnerTypeEnrollmentType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleOwnerTypeEnrollmentType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleOwnerTypeEnrollmentType) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnrollmentType gets the enrollmentType property value. The enrollmentType property
func (m *AppleOwnerTypeEnrollmentType) GetEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    return m.enrollmentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleOwnerTypeEnrollmentType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
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
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppleOwnerTypeEnrollmentType) GetOdataType()(*string) {
    return m.odataType
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *AppleOwnerTypeEnrollmentType) GetOwnerType()(*ManagedDeviceOwnerType) {
    return m.ownerType
}
// Serialize serializes information the current object
func (m *AppleOwnerTypeEnrollmentType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err := writer.WriteStringValue("enrollmentType", &cast)
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
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err := writer.WriteStringValue("ownerType", &cast)
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
func (m *AppleOwnerTypeEnrollmentType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnrollmentType sets the enrollmentType property value. The enrollmentType property
func (m *AppleOwnerTypeEnrollmentType) SetEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    m.enrollmentType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppleOwnerTypeEnrollmentType) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *AppleOwnerTypeEnrollmentType) SetOwnerType(value *ManagedDeviceOwnerType)() {
    m.ownerType = value
}
