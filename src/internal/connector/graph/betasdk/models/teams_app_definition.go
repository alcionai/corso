package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppDefinition provides operations to manage the collection of site entities.
type TeamsAppDefinition struct {
    Entity
    // A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
    allowedInstallationScopes *TeamsAppInstallationScopes
    // The WebApplicationInfo.Id from the Teams app manifest.
    azureADAppId *string
    // The details of the bot specified in the Teams app manifest.
    bot TeamworkBotable
    // The color version of the Teams app's icon.
    colorIcon TeamsAppIconable
    // The createdBy property
    createdBy IdentitySetable
    // The description property
    description *string
    // The name of the app provided by the app developer.
    displayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The outline version of the Teams app's icon.
    outlineIcon TeamsAppIconable
    // The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
    publishingState *TeamsAppPublishingState
    // The shortdescription property
    shortdescription *string
    // The ID from the Teams app manifest.
    teamsAppId *string
    // The version number of the application.
    version *string
}
// NewTeamsAppDefinition instantiates a new teamsAppDefinition and sets the default values.
func NewTeamsAppDefinition()(*TeamsAppDefinition) {
    m := &TeamsAppDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppDefinition(), nil
}
// GetAllowedInstallationScopes gets the allowedInstallationScopes property value. A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
func (m *TeamsAppDefinition) GetAllowedInstallationScopes()(*TeamsAppInstallationScopes) {
    return m.allowedInstallationScopes
}
// GetAzureADAppId gets the azureADAppId property value. The WebApplicationInfo.Id from the Teams app manifest.
func (m *TeamsAppDefinition) GetAzureADAppId()(*string) {
    return m.azureADAppId
}
// GetBot gets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) GetBot()(TeamworkBotable) {
    return m.bot
}
// GetColorIcon gets the colorIcon property value. The color version of the Teams app's icon.
func (m *TeamsAppDefinition) GetColorIcon()(TeamsAppIconable) {
    return m.colorIcon
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *TeamsAppDefinition) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetDescription gets the description property value. The description property
func (m *TeamsAppDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedInstallationScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppInstallationScopes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInstallationScopes(val.(*TeamsAppInstallationScopes))
        }
        return nil
    }
    res["azureADAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADAppId(val)
        }
        return nil
    }
    res["bot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkBotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBot(val.(TeamworkBotable))
        }
        return nil
    }
    res["colorIcon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorIcon(val.(TeamsAppIconable))
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
    res["outlineIcon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlineIcon(val.(TeamsAppIconable))
        }
        return nil
    }
    res["publishingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppPublishingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(*TeamsAppPublishingState))
        }
        return nil
    }
    res["shortdescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortdescription(val)
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *TeamsAppDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOutlineIcon gets the outlineIcon property value. The outline version of the Teams app's icon.
func (m *TeamsAppDefinition) GetOutlineIcon()(TeamsAppIconable) {
    return m.outlineIcon
}
// GetPublishingState gets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
func (m *TeamsAppDefinition) GetPublishingState()(*TeamsAppPublishingState) {
    return m.publishingState
}
// GetShortdescription gets the shortdescription property value. The shortdescription property
func (m *TeamsAppDefinition) GetShortdescription()(*string) {
    return m.shortdescription
}
// GetTeamsAppId gets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) GetTeamsAppId()(*string) {
    return m.teamsAppId
}
// GetVersion gets the version property value. The version number of the application.
func (m *TeamsAppDefinition) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *TeamsAppDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedInstallationScopes() != nil {
        cast := (*m.GetAllowedInstallationScopes()).String()
        err = writer.WriteStringValue("allowedInstallationScopes", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureADAppId", m.GetAzureADAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bot", m.GetBot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("colorIcon", m.GetColorIcon())
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
    {
        err = writer.WriteObjectValue("outlineIcon", m.GetOutlineIcon())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingState() != nil {
        cast := (*m.GetPublishingState()).String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortdescription", m.GetShortdescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedInstallationScopes sets the allowedInstallationScopes property value. A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
func (m *TeamsAppDefinition) SetAllowedInstallationScopes(value *TeamsAppInstallationScopes)() {
    m.allowedInstallationScopes = value
}
// SetAzureADAppId sets the azureADAppId property value. The WebApplicationInfo.Id from the Teams app manifest.
func (m *TeamsAppDefinition) SetAzureADAppId(value *string)() {
    m.azureADAppId = value
}
// SetBot sets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) SetBot(value TeamworkBotable)() {
    m.bot = value
}
// SetColorIcon sets the colorIcon property value. The color version of the Teams app's icon.
func (m *TeamsAppDefinition) SetColorIcon(value TeamsAppIconable)() {
    m.colorIcon = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *TeamsAppDefinition) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetDescription sets the description property value. The description property
func (m *TeamsAppDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *TeamsAppDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOutlineIcon sets the outlineIcon property value. The outline version of the Teams app's icon.
func (m *TeamsAppDefinition) SetOutlineIcon(value TeamsAppIconable)() {
    m.outlineIcon = value
}
// SetPublishingState sets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
func (m *TeamsAppDefinition) SetPublishingState(value *TeamsAppPublishingState)() {
    m.publishingState = value
}
// SetShortdescription sets the shortdescription property value. The shortdescription property
func (m *TeamsAppDefinition) SetShortdescription(value *string)() {
    m.shortdescription = value
}
// SetTeamsAppId sets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// SetVersion sets the version property value. The version number of the application.
func (m *TeamsAppDefinition) SetVersion(value *string)() {
    m.version = value
}
