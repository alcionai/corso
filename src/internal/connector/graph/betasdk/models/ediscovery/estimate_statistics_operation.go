package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EstimateStatisticsOperation 
type EstimateStatisticsOperation struct {
    CaseOperation
    // The estimated count of items for the sourceCollection that matched the content query.
    indexedItemCount *int64
    // The estimated size of items for the sourceCollection that matched the content query.
    indexedItemsSize *int64
    // The number of mailboxes that had search hits.
    mailboxCount *int32
    // The number of mailboxes that had search hits.
    siteCount *int32
    // eDiscovery collection, commonly known as a search.
    sourceCollection SourceCollectionable
    // The estimated count of unindexed items for the collection.
    unindexedItemCount *int64
    // The estimated size of unindexed items for the collection.
    unindexedItemsSize *int64
}
// NewEstimateStatisticsOperation instantiates a new estimateStatisticsOperation and sets the default values.
func NewEstimateStatisticsOperation()(*EstimateStatisticsOperation) {
    m := &EstimateStatisticsOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEstimateStatisticsOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEstimateStatisticsOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEstimateStatisticsOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EstimateStatisticsOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["indexedItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexedItemCount(val)
        }
        return nil
    }
    res["indexedItemsSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexedItemsSize(val)
        }
        return nil
    }
    res["mailboxCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxCount(val)
        }
        return nil
    }
    res["siteCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCount(val)
        }
        return nil
    }
    res["sourceCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSourceCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCollection(val.(SourceCollectionable))
        }
        return nil
    }
    res["unindexedItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnindexedItemCount(val)
        }
        return nil
    }
    res["unindexedItemsSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnindexedItemsSize(val)
        }
        return nil
    }
    return res
}
// GetIndexedItemCount gets the indexedItemCount property value. The estimated count of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) GetIndexedItemCount()(*int64) {
    return m.indexedItemCount
}
// GetIndexedItemsSize gets the indexedItemsSize property value. The estimated size of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) GetIndexedItemsSize()(*int64) {
    return m.indexedItemsSize
}
// GetMailboxCount gets the mailboxCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) GetMailboxCount()(*int32) {
    return m.mailboxCount
}
// GetSiteCount gets the siteCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) GetSiteCount()(*int32) {
    return m.siteCount
}
// GetSourceCollection gets the sourceCollection property value. eDiscovery collection, commonly known as a search.
func (m *EstimateStatisticsOperation) GetSourceCollection()(SourceCollectionable) {
    return m.sourceCollection
}
// GetUnindexedItemCount gets the unindexedItemCount property value. The estimated count of unindexed items for the collection.
func (m *EstimateStatisticsOperation) GetUnindexedItemCount()(*int64) {
    return m.unindexedItemCount
}
// GetUnindexedItemsSize gets the unindexedItemsSize property value. The estimated size of unindexed items for the collection.
func (m *EstimateStatisticsOperation) GetUnindexedItemsSize()(*int64) {
    return m.unindexedItemsSize
}
// Serialize serializes information the current object
func (m *EstimateStatisticsOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("indexedItemCount", m.GetIndexedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("indexedItemsSize", m.GetIndexedItemsSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mailboxCount", m.GetMailboxCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("siteCount", m.GetSiteCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceCollection", m.GetSourceCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unindexedItemCount", m.GetUnindexedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unindexedItemsSize", m.GetUnindexedItemsSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIndexedItemCount sets the indexedItemCount property value. The estimated count of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) SetIndexedItemCount(value *int64)() {
    m.indexedItemCount = value
}
// SetIndexedItemsSize sets the indexedItemsSize property value. The estimated size of items for the sourceCollection that matched the content query.
func (m *EstimateStatisticsOperation) SetIndexedItemsSize(value *int64)() {
    m.indexedItemsSize = value
}
// SetMailboxCount sets the mailboxCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) SetMailboxCount(value *int32)() {
    m.mailboxCount = value
}
// SetSiteCount sets the siteCount property value. The number of mailboxes that had search hits.
func (m *EstimateStatisticsOperation) SetSiteCount(value *int32)() {
    m.siteCount = value
}
// SetSourceCollection sets the sourceCollection property value. eDiscovery collection, commonly known as a search.
func (m *EstimateStatisticsOperation) SetSourceCollection(value SourceCollectionable)() {
    m.sourceCollection = value
}
// SetUnindexedItemCount sets the unindexedItemCount property value. The estimated count of unindexed items for the collection.
func (m *EstimateStatisticsOperation) SetUnindexedItemCount(value *int64)() {
    m.unindexedItemCount = value
}
// SetUnindexedItemsSize sets the unindexedItemsSize property value. The estimated size of unindexed items for the collection.
func (m *EstimateStatisticsOperation) SetUnindexedItemsSize(value *int64)() {
    m.unindexedItemsSize = value
}
