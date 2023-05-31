package api

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type DriveItemPager interface {
	GetPage(context.Context) (DeltaPageLinker, error)
	SetNext(nextLink string)
	Reset()
	ValuesIn(DeltaPageLinker) ([]models.DriveItemable, error)
}
