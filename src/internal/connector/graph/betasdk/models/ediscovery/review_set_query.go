package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ReviewSetQuery provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ReviewSetQuery struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The user who created the query.
    createdBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // The time and date when the query was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the query.
    displayName *string
    // The user who last modified the query.
    lastModifiedBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // The date and time the query was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The query string in KQL (Keyword Query Language) query. For details, see Document metadata fields in Advanced eDiscovery.  This field maps directly to the keywords condition.  You can refine searches by using fields listed in the searchable field name paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
    query *string
}
// NewReviewSetQuery instantiates a new reviewSetQuery and sets the default values.
func NewReviewSetQuery()(*ReviewSetQuery) {
    m := &ReviewSetQuery{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateReviewSetQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReviewSetQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewSetQuery(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created the query.
func (m *ReviewSetQuery) GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The time and date when the query was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the query.
func (m *ReviewSetQuery) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReviewSetQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the query.
func (m *ReviewSetQuery) GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the query was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetQuery gets the query property value. The query string in KQL (Keyword Query Language) query. For details, see Document metadata fields in Advanced eDiscovery.  This field maps directly to the keywords condition.  You can refine searches by using fields listed in the searchable field name paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *ReviewSetQuery) GetQuery()(*string) {
    return m.query
}
// Serialize serializes information the current object
func (m *ReviewSetQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created the query.
func (m *ReviewSetQuery) SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time and date when the query was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the query.
func (m *ReviewSetQuery) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the query.
func (m *ReviewSetQuery) SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the query was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetQuery sets the query property value. The query string in KQL (Keyword Query Language) query. For details, see Document metadata fields in Advanced eDiscovery.  This field maps directly to the keywords condition.  You can refine searches by using fields listed in the searchable field name paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *ReviewSetQuery) SetQuery(value *string)() {
    m.query = value
}
