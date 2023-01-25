package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfigurationLocalization provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PlannerPlanConfigurationLocalization struct {
    Entity
    // Localized names for configured buckets in the plan configuration.
    buckets []PlannerPlanConfigurationBucketLocalizationable
    // The language code associated with the localized names in this object.
    languageTag *string
    // Localized title of the plan.
    planTitle *string
}
// NewPlannerPlanConfigurationLocalization instantiates a new plannerPlanConfigurationLocalization and sets the default values.
func NewPlannerPlanConfigurationLocalization()(*PlannerPlanConfigurationLocalization) {
    m := &PlannerPlanConfigurationLocalization{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanConfigurationLocalization(), nil
}
// GetBuckets gets the buckets property value. Localized names for configured buckets in the plan configuration.
func (m *PlannerPlanConfigurationLocalization) GetBuckets()([]PlannerPlanConfigurationBucketLocalizationable) {
    return m.buckets
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanConfigurationLocalization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanConfigurationBucketLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanConfigurationBucketLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanConfigurationBucketLocalizationable)
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["planTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanTitle(val)
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. The language code associated with the localized names in this object.
func (m *PlannerPlanConfigurationLocalization) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetPlanTitle gets the planTitle property value. Localized title of the plan.
func (m *PlannerPlanConfigurationLocalization) GetPlanTitle()(*string) {
    return m.planTitle
}
// Serialize serializes information the current object
func (m *PlannerPlanConfigurationLocalization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planTitle", m.GetPlanTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. Localized names for configured buckets in the plan configuration.
func (m *PlannerPlanConfigurationLocalization) SetBuckets(value []PlannerPlanConfigurationBucketLocalizationable)() {
    m.buckets = value
}
// SetLanguageTag sets the languageTag property value. The language code associated with the localized names in this object.
func (m *PlannerPlanConfigurationLocalization) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetPlanTitle sets the planTitle property value. Localized title of the plan.
func (m *PlannerPlanConfigurationLocalization) SetPlanTitle(value *string)() {
    m.planTitle = value
}
