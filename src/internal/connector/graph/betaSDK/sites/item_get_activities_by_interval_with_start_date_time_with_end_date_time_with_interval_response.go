package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse provides operations to call the getActivitiesByInterval method.
type ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse struct {
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.BaseCollectionPaginationCountResponse
}
// NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse instantiates a new ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse and sets the default values.
func NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse()(*ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) {
    m := &ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse{
        BaseCollectionPaginationCountResponse: *ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateItemActivityStatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemActivityStat, len(val))
            for i, v := range val {
                res[i] = *(v.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemActivityStat))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) GetValue()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemActivityStatable) {
    return m.GetBackingStore().Get("value");
}
// Serialize serializes information the current object
func (m *ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            temp := v
            cast[i] = i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *ItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) SetValue(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemActivityStatable)() {
    m.GetBackingStore().Set("value", value)
}
