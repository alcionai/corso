package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagementTemplateStepVersion 
type ManagementTemplateStepVersion struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The acceptedFor property
    acceptedFor ManagementTemplateStepable
    // The contentMarkdown property
    contentMarkdown *string
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deployments property
    deployments []ManagementTemplateStepDeploymentable
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name property
    name *string
    // The templateStep property
    templateStep ManagementTemplateStepable
    // The version property
    version *int32
    // The versionInformation property
    versionInformation *string
}
// NewManagementTemplateStepVersion instantiates a new managementTemplateStepVersion and sets the default values.
func NewManagementTemplateStepVersion()(*ManagementTemplateStepVersion) {
    m := &ManagementTemplateStepVersion{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagementTemplateStepVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStepVersion(), nil
}
// GetAcceptedFor gets the acceptedFor property value. The acceptedFor property
func (m *ManagementTemplateStepVersion) GetAcceptedFor()(ManagementTemplateStepable) {
    return m.acceptedFor
}
// GetContentMarkdown gets the contentMarkdown property value. The contentMarkdown property
func (m *ManagementTemplateStepVersion) GetContentMarkdown()(*string) {
    return m.contentMarkdown
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepVersion) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepVersion) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeployments gets the deployments property value. The deployments property
func (m *ManagementTemplateStepVersion) GetDeployments()([]ManagementTemplateStepDeploymentable) {
    return m.deployments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedFor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedFor(val.(ManagementTemplateStepable))
        }
        return nil
    }
    res["contentMarkdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentMarkdown(val)
        }
        return nil
    }
    res["createdByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
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
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepDeploymentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepDeploymentable)
            }
            m.SetDeployments(res)
        }
        return nil
    }
    res["lastActionByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["templateStep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateStep(val.(ManagementTemplateStepable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["versionInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInformation(val)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepVersion) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepVersion) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetName gets the name property value. The name property
func (m *ManagementTemplateStepVersion) GetName()(*string) {
    return m.name
}
// GetTemplateStep gets the templateStep property value. The templateStep property
func (m *ManagementTemplateStepVersion) GetTemplateStep()(ManagementTemplateStepable) {
    return m.templateStep
}
// GetVersion gets the version property value. The version property
func (m *ManagementTemplateStepVersion) GetVersion()(*int32) {
    return m.version
}
// GetVersionInformation gets the versionInformation property value. The versionInformation property
func (m *ManagementTemplateStepVersion) GetVersionInformation()(*string) {
    return m.versionInformation
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("acceptedFor", m.GetAcceptedFor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentMarkdown", m.GetContentMarkdown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
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
    if m.GetDeployments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStep", m.GetTemplateStep())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionInformation", m.GetVersionInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedFor sets the acceptedFor property value. The acceptedFor property
func (m *ManagementTemplateStepVersion) SetAcceptedFor(value ManagementTemplateStepable)() {
    m.acceptedFor = value
}
// SetContentMarkdown sets the contentMarkdown property value. The contentMarkdown property
func (m *ManagementTemplateStepVersion) SetContentMarkdown(value *string)() {
    m.contentMarkdown = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepVersion) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepVersion) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeployments sets the deployments property value. The deployments property
func (m *ManagementTemplateStepVersion) SetDeployments(value []ManagementTemplateStepDeploymentable)() {
    m.deployments = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepVersion) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepVersion) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetName sets the name property value. The name property
func (m *ManagementTemplateStepVersion) SetName(value *string)() {
    m.name = value
}
// SetTemplateStep sets the templateStep property value. The templateStep property
func (m *ManagementTemplateStepVersion) SetTemplateStep(value ManagementTemplateStepable)() {
    m.templateStep = value
}
// SetVersion sets the version property value. The version property
func (m *ManagementTemplateStepVersion) SetVersion(value *int32)() {
    m.version = value
}
// SetVersionInformation sets the versionInformation property value. The versionInformation property
func (m *ManagementTemplateStepVersion) SetVersionInformation(value *string)() {
    m.versionInformation = value
}
