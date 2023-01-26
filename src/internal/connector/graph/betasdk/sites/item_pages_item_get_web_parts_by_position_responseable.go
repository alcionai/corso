package sites

import (
	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemPagesItemGetWebPartsByPositionResponseable 
type ItemPagesItemGetWebPartsByPositionResponseable interface {
    msmodel.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable)
    SetValue(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.WebPartable)()
}
