package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesDirectorySynchronizationConfiguration 
type OnPremisesDirectorySynchronizationConfiguration struct {
    // Contains the accidental deletion prevention configuration for a tenant.
    accidentalDeletionPrevention OnPremisesAccidentalDeletionPreventionable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Interval of time that the customer requested the sync client waits between sync cycles.
    customerRequestedSynchronizationInterval *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
    // Interval of time the sync client should honor between sync cycles
    synchronizationInterval *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewOnPremisesDirectorySynchronizationConfiguration instantiates a new onPremisesDirectorySynchronizationConfiguration and sets the default values.
func NewOnPremisesDirectorySynchronizationConfiguration()(*OnPremisesDirectorySynchronizationConfiguration) {
    m := &OnPremisesDirectorySynchronizationConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesDirectorySynchronizationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesDirectorySynchronizationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesDirectorySynchronizationConfiguration(), nil
}
// GetAccidentalDeletionPrevention gets the accidentalDeletionPrevention property value. Contains the accidental deletion prevention configuration for a tenant.
func (m *OnPremisesDirectorySynchronizationConfiguration) GetAccidentalDeletionPrevention()(OnPremisesAccidentalDeletionPreventionable) {
    return m.accidentalDeletionPrevention
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCustomerRequestedSynchronizationInterval gets the customerRequestedSynchronizationInterval property value. Interval of time that the customer requested the sync client waits between sync cycles.
func (m *OnPremisesDirectorySynchronizationConfiguration) GetCustomerRequestedSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.customerRequestedSynchronizationInterval
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesDirectorySynchronizationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accidentalDeletionPrevention"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesAccidentalDeletionPreventionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccidentalDeletionPrevention(val.(OnPremisesAccidentalDeletionPreventionable))
        }
        return nil
    }
    res["customerRequestedSynchronizationInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerRequestedSynchronizationInterval(val)
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
    res["synchronizationInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationInterval(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetSynchronizationInterval gets the synchronizationInterval property value. Interval of time the sync client should honor between sync cycles
func (m *OnPremisesDirectorySynchronizationConfiguration) GetSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.synchronizationInterval
}
// Serialize serializes information the current object
func (m *OnPremisesDirectorySynchronizationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accidentalDeletionPrevention", m.GetAccidentalDeletionPrevention())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("customerRequestedSynchronizationInterval", m.GetCustomerRequestedSynchronizationInterval())
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
        err := writer.WriteISODurationValue("synchronizationInterval", m.GetSynchronizationInterval())
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
// SetAccidentalDeletionPrevention sets the accidentalDeletionPrevention property value. Contains the accidental deletion prevention configuration for a tenant.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetAccidentalDeletionPrevention(value OnPremisesAccidentalDeletionPreventionable)() {
    m.accidentalDeletionPrevention = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCustomerRequestedSynchronizationInterval sets the customerRequestedSynchronizationInterval property value. Interval of time that the customer requested the sync client waits between sync cycles.
func (m *OnPremisesDirectorySynchronizationConfiguration) SetCustomerRequestedSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.customerRequestedSynchronizationInterval = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSynchronizationInterval sets the synchronizationInterval property value. Interval of time the sync client should honor between sync cycles
func (m *OnPremisesDirectorySynchronizationConfiguration) SetSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.synchronizationInterval = value
}
