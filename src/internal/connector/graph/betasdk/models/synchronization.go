package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Synchronization 
type Synchronization struct {
    Entity
    // Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
    jobs []SynchronizationJobable
    // Represents a collection of credentials to access provisioned cloud applications.
    secrets []SynchronizationSecretKeyStringValuePairable
    // Pre-configured synchronization settings for a particular application.
    templates []SynchronizationTemplateable
}
// NewSynchronization instantiates a new synchronization and sets the default values.
func NewSynchronization()(*Synchronization) {
    m := &Synchronization{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronization(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Synchronization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobable)
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["secrets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationSecretKeyStringValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationSecretKeyStringValuePairable)
            }
            m.SetSecrets(res)
        }
        return nil
    }
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationTemplateable)
            }
            m.SetTemplates(res)
        }
        return nil
    }
    return res
}
// GetJobs gets the jobs property value. Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *Synchronization) GetJobs()([]SynchronizationJobable) {
    return m.jobs
}
// GetSecrets gets the secrets property value. Represents a collection of credentials to access provisioned cloud applications.
func (m *Synchronization) GetSecrets()([]SynchronizationSecretKeyStringValuePairable) {
    return m.secrets
}
// GetTemplates gets the templates property value. Pre-configured synchronization settings for a particular application.
func (m *Synchronization) GetTemplates()([]SynchronizationTemplateable) {
    return m.templates
}
// Serialize serializes information the current object
func (m *Synchronization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecrets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecrets()))
        for i, v := range m.GetSecrets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secrets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetJobs sets the jobs property value. Performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *Synchronization) SetJobs(value []SynchronizationJobable)() {
    m.jobs = value
}
// SetSecrets sets the secrets property value. Represents a collection of credentials to access provisioned cloud applications.
func (m *Synchronization) SetSecrets(value []SynchronizationSecretKeyStringValuePairable)() {
    m.secrets = value
}
// SetTemplates sets the templates property value. Pre-configured synchronization settings for a particular application.
func (m *Synchronization) SetTemplates(value []SynchronizationTemplateable)() {
    m.templates = value
}
