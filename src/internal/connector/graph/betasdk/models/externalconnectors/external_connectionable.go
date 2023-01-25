package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ExternalConnectionable 
type ExternalConnectionable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivitySettings()(ActivitySettingsable)
    GetComplianceSettings()(ComplianceSettingsable)
    GetConfiguration()(Configurationable)
    GetConnectorId()(*string)
    GetDescription()(*string)
    GetEnabledContentExperiences()(*ContentExperienceType)
    GetGroups()([]ExternalGroupable)
    GetIngestedItemsCount()(*int64)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetQuota()(ConnectionQuotaable)
    GetSchema()(Schemaable)
    GetSearchSettings()(SearchSettingsable)
    GetState()(*ConnectionState)
    SetActivitySettings(value ActivitySettingsable)()
    SetComplianceSettings(value ComplianceSettingsable)()
    SetConfiguration(value Configurationable)()
    SetConnectorId(value *string)()
    SetDescription(value *string)()
    SetEnabledContentExperiences(value *ContentExperienceType)()
    SetGroups(value []ExternalGroupable)()
    SetIngestedItemsCount(value *int64)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetQuota(value ConnectionQuotaable)()
    SetSchema(value Schemaable)()
    SetSearchSettings(value SearchSettingsable)()
    SetState(value *ConnectionState)()
}
