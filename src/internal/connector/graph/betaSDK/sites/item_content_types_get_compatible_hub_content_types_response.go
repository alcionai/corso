package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContentTypesGetCompatibleHubContentTypesResponse provides operations to call the getCompatibleHubContentTypes method.
type ItemContentTypesGetCompatibleHubContentTypesResponse struct {
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.BaseCollectionPaginationCountResponse
}
// NewItemContentTypesGetCompatibleHubContentTypesResponse instantiates a new ItemContentTypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewItemContentTypesGetCompatibleHubContentTypesResponse()(*ItemContentTypesGetCompatibleHubContentTypesResponse) {
    m := &ItemContentTypesGetCompatibleHubContentTypesResponse{
        BaseCollectionPaginationCountResponse: *ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContentTypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContentTypesGetCompatibleHubContentTypesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemContentTypesGetCompatibleHubContentTypesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateContentTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentType))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemContentTypesGetCompatibleHubContentTypesResponse) GetValue()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentTypeable) {
    return m.GetBackingStore().Get("value");
}
// Serialize serializes information the current object
func (m *ItemContentTypesGetCompatibleHubContentTypesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemContentTypesGetCompatibleHubContentTypesResponse) SetValue(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentTypeable)() {
    m.GetBackingStore().Set("value", value)
}
