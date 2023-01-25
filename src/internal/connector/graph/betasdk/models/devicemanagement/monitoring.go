package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Monitoring 
type Monitoring struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The collection of records of alert events.
    alertRecords []AlertRecordable
    // The collection of alert rules.
    alertRules []AlertRuleable
}
// NewMonitoring instantiates a new monitoring and sets the default values.
func NewMonitoring()(*Monitoring) {
    m := &Monitoring{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateMonitoringFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonitoring(), nil
}
// GetAlertRecords gets the alertRecords property value. The collection of records of alert events.
func (m *Monitoring) GetAlertRecords()([]AlertRecordable) {
    return m.alertRecords
}
// GetAlertRules gets the alertRules property value. The collection of alert rules.
func (m *Monitoring) GetAlertRules()([]AlertRuleable) {
    return m.alertRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Monitoring) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertRecordable, len(val))
            for i, v := range val {
                res[i] = v.(AlertRecordable)
            }
            m.SetAlertRecords(res)
        }
        return nil
    }
    res["alertRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertRuleable, len(val))
            for i, v := range val {
                res[i] = v.(AlertRuleable)
            }
            m.SetAlertRules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Monitoring) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertRecords()))
        for i, v := range m.GetAlertRecords() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alertRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlertRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertRules()))
        for i, v := range m.GetAlertRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alertRules", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertRecords sets the alertRecords property value. The collection of records of alert events.
func (m *Monitoring) SetAlertRecords(value []AlertRecordable)() {
    m.alertRecords = value
}
// SetAlertRules sets the alertRules property value. The collection of alert rules.
func (m *Monitoring) SetAlertRules(value []AlertRuleable)() {
    m.alertRules = value
}
