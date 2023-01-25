package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationCustomization 
type EducationSynchronizationCustomization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the display name of the resource can be overwritten by the sync.
    allowDisplayNameUpdate *bool
    // Indicates whether synchronization of the parent entity is deferred to a later date.
    isSyncDeferred *bool
    // The OdataType property
    odataType *string
    // The collection of property names to sync. If set to null, all properties will be synchronized. Does not apply to Student Enrollments or Teacher Rosters
    optionalPropertiesToSync []string
    // The date that the synchronization should start. This value should be set to a future date. If set to null, the resource will be synchronized when the profile setup completes. Only applies to Student Enrollments
    synchronizationStartDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewEducationSynchronizationCustomization instantiates a new educationSynchronizationCustomization and sets the default values.
func NewEducationSynchronizationCustomization()(*EducationSynchronizationCustomization) {
    m := &EducationSynchronizationCustomization{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationSynchronizationCustomizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationCustomizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationCustomization(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationCustomization) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowDisplayNameUpdate gets the allowDisplayNameUpdate property value. Indicates whether the display name of the resource can be overwritten by the sync.
func (m *EducationSynchronizationCustomization) GetAllowDisplayNameUpdate()(*bool) {
    return m.allowDisplayNameUpdate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationCustomization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowDisplayNameUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDisplayNameUpdate(val)
        }
        return nil
    }
    res["isSyncDeferred"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSyncDeferred(val)
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
    res["optionalPropertiesToSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOptionalPropertiesToSync(res)
        }
        return nil
    }
    res["synchronizationStartDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationStartDate(val)
        }
        return nil
    }
    return res
}
// GetIsSyncDeferred gets the isSyncDeferred property value. Indicates whether synchronization of the parent entity is deferred to a later date.
func (m *EducationSynchronizationCustomization) GetIsSyncDeferred()(*bool) {
    return m.isSyncDeferred
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationCustomization) GetOdataType()(*string) {
    return m.odataType
}
// GetOptionalPropertiesToSync gets the optionalPropertiesToSync property value. The collection of property names to sync. If set to null, all properties will be synchronized. Does not apply to Student Enrollments or Teacher Rosters
func (m *EducationSynchronizationCustomization) GetOptionalPropertiesToSync()([]string) {
    return m.optionalPropertiesToSync
}
// GetSynchronizationStartDate gets the synchronizationStartDate property value. The date that the synchronization should start. This value should be set to a future date. If set to null, the resource will be synchronized when the profile setup completes. Only applies to Student Enrollments
func (m *EducationSynchronizationCustomization) GetSynchronizationStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.synchronizationStartDate
}
// Serialize serializes information the current object
func (m *EducationSynchronizationCustomization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowDisplayNameUpdate", m.GetAllowDisplayNameUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSyncDeferred", m.GetIsSyncDeferred())
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
    if m.GetOptionalPropertiesToSync() != nil {
        err := writer.WriteCollectionOfStringValues("optionalPropertiesToSync", m.GetOptionalPropertiesToSync())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("synchronizationStartDate", m.GetSynchronizationStartDate())
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
func (m *EducationSynchronizationCustomization) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowDisplayNameUpdate sets the allowDisplayNameUpdate property value. Indicates whether the display name of the resource can be overwritten by the sync.
func (m *EducationSynchronizationCustomization) SetAllowDisplayNameUpdate(value *bool)() {
    m.allowDisplayNameUpdate = value
}
// SetIsSyncDeferred sets the isSyncDeferred property value. Indicates whether synchronization of the parent entity is deferred to a later date.
func (m *EducationSynchronizationCustomization) SetIsSyncDeferred(value *bool)() {
    m.isSyncDeferred = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationCustomization) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOptionalPropertiesToSync sets the optionalPropertiesToSync property value. The collection of property names to sync. If set to null, all properties will be synchronized. Does not apply to Student Enrollments or Teacher Rosters
func (m *EducationSynchronizationCustomization) SetOptionalPropertiesToSync(value []string)() {
    m.optionalPropertiesToSync = value
}
// SetSynchronizationStartDate sets the synchronizationStartDate property value. The date that the synchronization should start. This value should be set to a future date. If set to null, the resource will be synchronized when the profile setup completes. Only applies to Student Enrollments
func (m *EducationSynchronizationCustomization) SetSynchronizationStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.synchronizationStartDate = value
}
