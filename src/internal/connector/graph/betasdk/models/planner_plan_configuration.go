package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfiguration 
type PlannerPlanConfiguration struct {
    Entity
    // List the buckets that should be created in the plan.
    buckets []PlannerPlanConfigurationBucketDefinitionable
    // The identity of the creator of the plan configuration.
    createdBy IdentitySetable
    // The date and time when the plan configuration was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The language code for the default language to be used for the names of the objects created for the plan.
    defaultLanguage *string
    // The identity of the user who last modified the plan configuration.
    lastModifiedBy IdentitySetable
    // The date and time when the plan configuration was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Localized names for the plan configuration.
    localizations []PlannerPlanConfigurationLocalizationable
}
// NewPlannerPlanConfiguration instantiates a new plannerPlanConfiguration and sets the default values.
func NewPlannerPlanConfiguration()(*PlannerPlanConfiguration) {
    m := &PlannerPlanConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerPlanConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanConfiguration(), nil
}
// GetBuckets gets the buckets property value. List the buckets that should be created in the plan.
func (m *PlannerPlanConfiguration) GetBuckets()([]PlannerPlanConfigurationBucketDefinitionable) {
    return m.buckets
}
// GetCreatedBy gets the createdBy property value. The identity of the creator of the plan configuration.
func (m *PlannerPlanConfiguration) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the plan configuration was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PlannerPlanConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefaultLanguage gets the defaultLanguage property value. The language code for the default language to be used for the names of the objects created for the plan.
func (m *PlannerPlanConfiguration) GetDefaultLanguage()(*string) {
    return m.defaultLanguage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanConfigurationBucketDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanConfigurationBucketDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanConfigurationBucketDefinitionable)
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["defaultLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguage(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["localizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanConfigurationLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanConfigurationLocalizationable)
            }
            m.SetLocalizations(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity of the user who last modified the plan configuration.
func (m *PlannerPlanConfiguration) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the plan configuration was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PlannerPlanConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLocalizations gets the localizations property value. Localized names for the plan configuration.
func (m *PlannerPlanConfiguration) GetLocalizations()([]PlannerPlanConfigurationLocalizationable) {
    return m.localizations
}
// Serialize serializes information the current object
func (m *PlannerPlanConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBuckets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. List the buckets that should be created in the plan.
func (m *PlannerPlanConfiguration) SetBuckets(value []PlannerPlanConfigurationBucketDefinitionable)() {
    m.buckets = value
}
// SetCreatedBy sets the createdBy property value. The identity of the creator of the plan configuration.
func (m *PlannerPlanConfiguration) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the plan configuration was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PlannerPlanConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefaultLanguage sets the defaultLanguage property value. The language code for the default language to be used for the names of the objects created for the plan.
func (m *PlannerPlanConfiguration) SetDefaultLanguage(value *string)() {
    m.defaultLanguage = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity of the user who last modified the plan configuration.
func (m *PlannerPlanConfiguration) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the plan configuration was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *PlannerPlanConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLocalizations sets the localizations property value. Localized names for the plan configuration.
func (m *PlannerPlanConfiguration) SetLocalizations(value []PlannerPlanConfigurationLocalizationable)() {
    m.localizations = value
}
