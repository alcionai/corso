package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenario 
type BusinessScenario struct {
    Entity
    // The identity of the user who created the scenario.
    createdBy IdentitySetable
    // The date and time when the scenario was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Display name of the scenario.
    displayName *string
    // The identity of the user who last modified the scenario.
    lastModifiedBy IdentitySetable
    // The date and time when the scenario was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identifiers of applications that are authorized to work with this scenario.
    ownerAppIds []string
    // Planner content related to the scenario.
    planner BusinessScenarioPlannerable
    // Unique name of the scenario. To avoid conflicts, the recommended value for the unique name is a reverse domain name format, owned by the author of the scenario. For example, a scenario authored by Contoso.com would have a unique name that starts with com.contoso.
    uniqueName *string
}
// NewBusinessScenario instantiates a new BusinessScenario and sets the default values.
func NewBusinessScenario()(*BusinessScenario) {
    m := &BusinessScenario{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBusinessScenarioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenario(), nil
}
// GetCreatedBy gets the createdBy property value. The identity of the user who created the scenario.
func (m *BusinessScenario) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the scenario was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BusinessScenario) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. Display name of the scenario.
func (m *BusinessScenario) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenario) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["ownerAppIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOwnerAppIds(res)
        }
        return nil
    }
    res["planner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBusinessScenarioPlannerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanner(val.(BusinessScenarioPlannerable))
        }
        return nil
    }
    res["uniqueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueName(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity of the user who last modified the scenario.
func (m *BusinessScenario) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the scenario was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BusinessScenario) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOwnerAppIds gets the ownerAppIds property value. Identifiers of applications that are authorized to work with this scenario.
func (m *BusinessScenario) GetOwnerAppIds()([]string) {
    return m.ownerAppIds
}
// GetPlanner gets the planner property value. Planner content related to the scenario.
func (m *BusinessScenario) GetPlanner()(BusinessScenarioPlannerable) {
    return m.planner
}
// GetUniqueName gets the uniqueName property value. Unique name of the scenario. To avoid conflicts, the recommended value for the unique name is a reverse domain name format, owned by the author of the scenario. For example, a scenario authored by Contoso.com would have a unique name that starts with com.contoso.
func (m *BusinessScenario) GetUniqueName()(*string) {
    return m.uniqueName
}
// Serialize serializes information the current object
func (m *BusinessScenario) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetOwnerAppIds() != nil {
        err = writer.WriteCollectionOfStringValues("ownerAppIds", m.GetOwnerAppIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("planner", m.GetPlanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uniqueName", m.GetUniqueName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The identity of the user who created the scenario.
func (m *BusinessScenario) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the scenario was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BusinessScenario) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. Display name of the scenario.
func (m *BusinessScenario) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity of the user who last modified the scenario.
func (m *BusinessScenario) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the scenario was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *BusinessScenario) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOwnerAppIds sets the ownerAppIds property value. Identifiers of applications that are authorized to work with this scenario.
func (m *BusinessScenario) SetOwnerAppIds(value []string)() {
    m.ownerAppIds = value
}
// SetPlanner sets the planner property value. Planner content related to the scenario.
func (m *BusinessScenario) SetPlanner(value BusinessScenarioPlannerable)() {
    m.planner = value
}
// SetUniqueName sets the uniqueName property value. Unique name of the scenario. To avoid conflicts, the recommended value for the unique name is a reverse domain name format, owned by the author of the scenario. For example, a scenario authored by Contoso.com would have a unique name that starts with com.contoso.
func (m *BusinessScenario) SetUniqueName(value *string)() {
    m.uniqueName = value
}
