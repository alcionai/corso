package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyConfiguration the group policy configuration entity contains the configured values for one or more group policy definitions.
type GroupPolicyConfiguration struct {
    Entity
    // The list of group assignments for the configuration.
    assignments []GroupPolicyConfigurationAssignmentable
    // The date and time the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The list of enabled or disabled group policy definition values for the configuration.
    definitionValues []GroupPolicyDefinitionValueable
    // User provided description for the resource object.
    description *string
    // User provided name for the resource object.
    displayName *string
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Group Policy Configuration Ingestion Type
    policyConfigurationIngestionType *GroupPolicyConfigurationIngestionType
    // The list of scope tags for the configuration.
    roleScopeTagIds []string
}
// NewGroupPolicyConfiguration instantiates a new groupPolicyConfiguration and sets the default values.
func NewGroupPolicyConfiguration()(*GroupPolicyConfiguration) {
    m := &GroupPolicyConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the configuration.
func (m *GroupPolicyConfiguration) GetAssignments()([]GroupPolicyConfigurationAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefinitionValues gets the definitionValues property value. The list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyConfiguration) GetDefinitionValues()([]GroupPolicyDefinitionValueable) {
    return m.definitionValues
}
// GetDescription gets the description property value. User provided description for the resource object.
func (m *GroupPolicyConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. User provided name for the resource object.
func (m *GroupPolicyConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyConfigurationAssignmentable)
            }
            m.SetAssignments(res)
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
    res["definitionValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyDefinitionValueable)
            }
            m.SetDefinitionValues(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["policyConfigurationIngestionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyConfigurationIngestionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyConfigurationIngestionType(val.(*GroupPolicyConfigurationIngestionType))
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPolicyConfigurationIngestionType gets the policyConfigurationIngestionType property value. Group Policy Configuration Ingestion Type
func (m *GroupPolicyConfiguration) GetPolicyConfigurationIngestionType()(*GroupPolicyConfigurationIngestionType) {
    return m.policyConfigurationIngestionType
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyConfiguration) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *GroupPolicyConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
    if m.GetDefinitionValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefinitionValues()))
        for i, v := range m.GetDefinitionValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("definitionValues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyConfigurationIngestionType() != nil {
        cast := (*m.GetPolicyConfigurationIngestionType()).String()
        err = writer.WriteStringValue("policyConfigurationIngestionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the configuration.
func (m *GroupPolicyConfiguration) SetAssignments(value []GroupPolicyConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefinitionValues sets the definitionValues property value. The list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyConfiguration) SetDefinitionValues(value []GroupPolicyDefinitionValueable)() {
    m.definitionValues = value
}
// SetDescription sets the description property value. User provided description for the resource object.
func (m *GroupPolicyConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. User provided name for the resource object.
func (m *GroupPolicyConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPolicyConfigurationIngestionType sets the policyConfigurationIngestionType property value. Group Policy Configuration Ingestion Type
func (m *GroupPolicyConfiguration) SetPolicyConfigurationIngestionType(value *GroupPolicyConfigurationIngestionType)() {
    m.policyConfigurationIngestionType = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
