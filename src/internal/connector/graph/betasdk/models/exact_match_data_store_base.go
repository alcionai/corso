package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchDataStoreBase provides operations to call the add method.
type ExactMatchDataStoreBase struct {
    Entity
    // The columns property
    columns []ExactDataMatchStoreColumnable
    // The dataLastUpdatedDateTime property
    dataLastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The displayName property
    displayName *string
}
// NewExactMatchDataStoreBase instantiates a new exactMatchDataStoreBase and sets the default values.
func NewExactMatchDataStoreBase()(*ExactMatchDataStoreBase) {
    m := &ExactMatchDataStoreBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExactMatchDataStoreBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDataStoreBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.exactMatchDataStore":
                        return NewExactMatchDataStore(), nil
                }
            }
        }
    }
    return NewExactMatchDataStoreBase(), nil
}
// GetColumns gets the columns property value. The columns property
func (m *ExactMatchDataStoreBase) GetColumns()([]ExactDataMatchStoreColumnable) {
    return m.columns
}
// GetDataLastUpdatedDateTime gets the dataLastUpdatedDateTime property value. The dataLastUpdatedDateTime property
func (m *ExactMatchDataStoreBase) GetDataLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dataLastUpdatedDateTime
}
// GetDescription gets the description property value. The description property
func (m *ExactMatchDataStoreBase) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ExactMatchDataStoreBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDataStoreBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["columns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactDataMatchStoreColumnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactDataMatchStoreColumnable, len(val))
            for i, v := range val {
                res[i] = v.(ExactDataMatchStoreColumnable)
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["dataLastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLastUpdatedDateTime(val)
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
    return res
}
// Serialize serializes information the current object
func (m *ExactMatchDataStoreBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColumns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dataLastUpdatedDateTime", m.GetDataLastUpdatedDateTime())
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
    return nil
}
// SetColumns sets the columns property value. The columns property
func (m *ExactMatchDataStoreBase) SetColumns(value []ExactDataMatchStoreColumnable)() {
    m.columns = value
}
// SetDataLastUpdatedDateTime sets the dataLastUpdatedDateTime property value. The dataLastUpdatedDateTime property
func (m *ExactMatchDataStoreBase) SetDataLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dataLastUpdatedDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *ExactMatchDataStoreBase) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ExactMatchDataStoreBase) SetDisplayName(value *string)() {
    m.displayName = value
}
