package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesAccidentalDeletionPrevention 
type OnPremisesAccidentalDeletionPrevention struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Threshold value which triggers accidental deletion prevention. The threshold is either an absolute number of objects or a percentage number of objects.
    alertThreshold *int32
    // The OdataType property
    odataType *string
    // The status of the accidental deletion prevention feature. The possible values are: disabled, enabledForCount, enabledForPercentage.
    synchronizationPreventionType *OnPremisesDirectorySynchronizationDeletionPreventionType
}
// NewOnPremisesAccidentalDeletionPrevention instantiates a new onPremisesAccidentalDeletionPrevention and sets the default values.
func NewOnPremisesAccidentalDeletionPrevention()(*OnPremisesAccidentalDeletionPrevention) {
    m := &OnPremisesAccidentalDeletionPrevention{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesAccidentalDeletionPreventionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesAccidentalDeletionPreventionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesAccidentalDeletionPrevention(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesAccidentalDeletionPrevention) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAlertThreshold gets the alertThreshold property value. Threshold value which triggers accidental deletion prevention. The threshold is either an absolute number of objects or a percentage number of objects.
func (m *OnPremisesAccidentalDeletionPrevention) GetAlertThreshold()(*int32) {
    return m.alertThreshold
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesAccidentalDeletionPrevention) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertThreshold(val)
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
    res["synchronizationPreventionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnPremisesDirectorySynchronizationDeletionPreventionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationPreventionType(val.(*OnPremisesDirectorySynchronizationDeletionPreventionType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesAccidentalDeletionPrevention) GetOdataType()(*string) {
    return m.odataType
}
// GetSynchronizationPreventionType gets the synchronizationPreventionType property value. The status of the accidental deletion prevention feature. The possible values are: disabled, enabledForCount, enabledForPercentage.
func (m *OnPremisesAccidentalDeletionPrevention) GetSynchronizationPreventionType()(*OnPremisesDirectorySynchronizationDeletionPreventionType) {
    return m.synchronizationPreventionType
}
// Serialize serializes information the current object
func (m *OnPremisesAccidentalDeletionPrevention) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("alertThreshold", m.GetAlertThreshold())
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
    if m.GetSynchronizationPreventionType() != nil {
        cast := (*m.GetSynchronizationPreventionType()).String()
        err := writer.WriteStringValue("synchronizationPreventionType", &cast)
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
func (m *OnPremisesAccidentalDeletionPrevention) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAlertThreshold sets the alertThreshold property value. Threshold value which triggers accidental deletion prevention. The threshold is either an absolute number of objects or a percentage number of objects.
func (m *OnPremisesAccidentalDeletionPrevention) SetAlertThreshold(value *int32)() {
    m.alertThreshold = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesAccidentalDeletionPrevention) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSynchronizationPreventionType sets the synchronizationPreventionType property value. The status of the accidental deletion prevention feature. The possible values are: disabled, enabledForCount, enabledForPercentage.
func (m *OnPremisesAccidentalDeletionPrevention) SetSynchronizationPreventionType(value *OnPremisesDirectorySynchronizationDeletionPreventionType)() {
    m.synchronizationPreventionType = value
}
