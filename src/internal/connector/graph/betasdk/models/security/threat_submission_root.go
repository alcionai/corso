package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ThreatSubmissionRoot 
type ThreatSubmissionRoot struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The emailThreats property
    emailThreats []EmailThreatSubmissionable
    // The emailThreatSubmissionPolicies property
    emailThreatSubmissionPolicies []EmailThreatSubmissionPolicyable
    // The fileThreats property
    fileThreats []FileThreatSubmissionable
    // The urlThreats property
    urlThreats []UrlThreatSubmissionable
}
// NewThreatSubmissionRoot instantiates a new threatSubmissionRoot and sets the default values.
func NewThreatSubmissionRoot()(*ThreatSubmissionRoot) {
    m := &ThreatSubmissionRoot{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateThreatSubmissionRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThreatSubmissionRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreatSubmissionRoot(), nil
}
// GetEmailThreats gets the emailThreats property value. The emailThreats property
func (m *ThreatSubmissionRoot) GetEmailThreats()([]EmailThreatSubmissionable) {
    return m.emailThreats
}
// GetEmailThreatSubmissionPolicies gets the emailThreatSubmissionPolicies property value. The emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRoot) GetEmailThreatSubmissionPolicies()([]EmailThreatSubmissionPolicyable) {
    return m.emailThreatSubmissionPolicies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatSubmissionRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailThreatSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailThreatSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(EmailThreatSubmissionable)
            }
            m.SetEmailThreats(res)
        }
        return nil
    }
    res["emailThreatSubmissionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailThreatSubmissionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailThreatSubmissionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(EmailThreatSubmissionPolicyable)
            }
            m.SetEmailThreatSubmissionPolicies(res)
        }
        return nil
    }
    res["fileThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileThreatSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileThreatSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(FileThreatSubmissionable)
            }
            m.SetFileThreats(res)
        }
        return nil
    }
    res["urlThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUrlThreatSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UrlThreatSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(UrlThreatSubmissionable)
            }
            m.SetUrlThreats(res)
        }
        return nil
    }
    return res
}
// GetFileThreats gets the fileThreats property value. The fileThreats property
func (m *ThreatSubmissionRoot) GetFileThreats()([]FileThreatSubmissionable) {
    return m.fileThreats
}
// GetUrlThreats gets the urlThreats property value. The urlThreats property
func (m *ThreatSubmissionRoot) GetUrlThreats()([]UrlThreatSubmissionable) {
    return m.urlThreats
}
// Serialize serializes information the current object
func (m *ThreatSubmissionRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmailThreats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailThreats()))
        for i, v := range m.GetEmailThreats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailThreats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailThreatSubmissionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailThreatSubmissionPolicies()))
        for i, v := range m.GetEmailThreatSubmissionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailThreatSubmissionPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileThreats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFileThreats()))
        for i, v := range m.GetFileThreats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fileThreats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUrlThreats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUrlThreats()))
        for i, v := range m.GetUrlThreats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("urlThreats", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailThreats sets the emailThreats property value. The emailThreats property
func (m *ThreatSubmissionRoot) SetEmailThreats(value []EmailThreatSubmissionable)() {
    m.emailThreats = value
}
// SetEmailThreatSubmissionPolicies sets the emailThreatSubmissionPolicies property value. The emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRoot) SetEmailThreatSubmissionPolicies(value []EmailThreatSubmissionPolicyable)() {
    m.emailThreatSubmissionPolicies = value
}
// SetFileThreats sets the fileThreats property value. The fileThreats property
func (m *ThreatSubmissionRoot) SetFileThreats(value []FileThreatSubmissionable)() {
    m.fileThreats = value
}
// SetUrlThreats sets the urlThreats property value. The urlThreats property
func (m *ThreatSubmissionRoot) SetUrlThreats(value []UrlThreatSubmissionable)() {
    m.urlThreats = value
}
