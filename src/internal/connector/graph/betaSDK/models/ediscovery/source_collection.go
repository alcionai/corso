package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// SourceCollection 
type SourceCollection struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Adds an additional source to the sourceCollection.
    additionalSources []DataSourceable
    // Adds the results of the sourceCollection to the specified reviewSet.
    addToReviewSetOperation AddToReviewSetOperationable
    // The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
    contentQuery *string
    // The user who created the sourceCollection.
    createdBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // The date and time the sourceCollection was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Custodian sources that are included in the sourceCollection.
    custodianSources []DataSourceable
    // When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
    dataSourceScopes *DataSourceScopes
    // The description of the sourceCollection.
    description *string
    // The display name of the sourceCollection.
    displayName *string
    // The last estimate operation associated with the sourceCollection.
    lastEstimateStatisticsOperation EstimateStatisticsOperationable
    // The last user who modified the sourceCollection.
    lastModifiedBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // The last date and time the sourceCollection was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // noncustodialDataSource sources that are included in the sourceCollection
    noncustodialSources []NoncustodialDataSourceable
}
// NewSourceCollection instantiates a new sourceCollection and sets the default values.
func NewSourceCollection()(*SourceCollection) {
    m := &SourceCollection{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateSourceCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSourceCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSourceCollection(), nil
}
// GetAdditionalSources gets the additionalSources property value. Adds an additional source to the sourceCollection.
func (m *SourceCollection) GetAdditionalSources()([]DataSourceable) {
    return m.additionalSources
}
// GetAddToReviewSetOperation gets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) GetAddToReviewSetOperation()(AddToReviewSetOperationable) {
    return m.addToReviewSetOperation
}
// GetContentQuery gets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) GetContentQuery()(*string) {
    return m.contentQuery
}
// GetCreatedBy gets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustodianSources gets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) GetCustodianSources()([]DataSourceable) {
    return m.custodianSources
}
// GetDataSourceScopes gets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) GetDataSourceScopes()(*DataSourceScopes) {
    return m.dataSourceScopes
}
// GetDescription gets the description property value. The description of the sourceCollection.
func (m *SourceCollection) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SourceCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(DataSourceable)
            }
            m.SetAdditionalSources(res)
        }
        return nil
    }
    res["addToReviewSetOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAddToReviewSetOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToReviewSetOperation(val.(AddToReviewSetOperationable))
        }
        return nil
    }
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable))
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
    res["custodianSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(DataSourceable)
            }
            m.SetCustodianSources(res)
        }
        return nil
    }
    res["dataSourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceScopes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSourceScopes(val.(*DataSourceScopes))
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
    res["lastEstimateStatisticsOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEstimateStatisticsOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEstimateStatisticsOperation(val.(EstimateStatisticsOperationable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable))
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
    res["noncustodialSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNoncustodialDataSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NoncustodialDataSourceable, len(val))
            for i, v := range val {
                res[i] = v.(NoncustodialDataSourceable)
            }
            m.SetNoncustodialSources(res)
        }
        return nil
    }
    return res
}
// GetLastEstimateStatisticsOperation gets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) GetLastEstimateStatisticsOperation()(EstimateStatisticsOperationable) {
    return m.lastEstimateStatisticsOperation
}
// GetLastModifiedBy gets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNoncustodialSources gets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) GetNoncustodialSources()([]NoncustodialDataSourceable) {
    return m.noncustodialSources
}
// Serialize serializes information the current object
func (m *SourceCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalSources()))
        for i, v := range m.GetAdditionalSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("additionalSources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("addToReviewSetOperation", m.GetAddToReviewSetOperation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
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
    if m.GetCustodianSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodianSources()))
        for i, v := range m.GetCustodianSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("custodianSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDataSourceScopes() != nil {
        cast := (*m.GetDataSourceScopes()).String()
        err = writer.WriteStringValue("dataSourceScopes", &cast)
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
        err = writer.WriteObjectValue("lastEstimateStatisticsOperation", m.GetLastEstimateStatisticsOperation())
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
    if m.GetNoncustodialSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNoncustodialSources()))
        for i, v := range m.GetNoncustodialSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalSources sets the additionalSources property value. Adds an additional source to the sourceCollection.
func (m *SourceCollection) SetAdditionalSources(value []DataSourceable)() {
    m.additionalSources = value
}
// SetAddToReviewSetOperation sets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) SetAddToReviewSetOperation(value AddToReviewSetOperationable)() {
    m.addToReviewSetOperation = value
}
// SetContentQuery sets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// SetCreatedBy sets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustodianSources sets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) SetCustodianSources(value []DataSourceable)() {
    m.custodianSources = value
}
// SetDataSourceScopes sets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) SetDataSourceScopes(value *DataSourceScopes)() {
    m.dataSourceScopes = value
}
// SetDescription sets the description property value. The description of the sourceCollection.
func (m *SourceCollection) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastEstimateStatisticsOperation sets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) SetLastEstimateStatisticsOperation(value EstimateStatisticsOperationable)() {
    m.lastEstimateStatisticsOperation = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNoncustodialSources sets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) SetNoncustodialSources(value []NoncustodialDataSourceable)() {
    m.noncustodialSources = value
}
