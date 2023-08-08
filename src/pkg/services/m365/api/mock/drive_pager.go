package mock

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type PageLink struct {
	Link *string
}

func (pl *PageLink) GetOdataNextLink() *string {
	return pl.Link
}

type PagerResult struct {
	Drives   []models.Driveable
	NextLink *string
	Err      error
}

type DrivePager struct {
	ToReturn []PagerResult
	GetIdx   int
}

func (p *DrivePager) GetPage(context.Context) (api.PageLinker, error) {
	if len(p.ToReturn) <= p.GetIdx {
		return nil, clues.New("ToReturn index out of bounds")
	}

	idx := p.GetIdx
	p.GetIdx++

	return &PageLink{p.ToReturn[idx].NextLink}, p.ToReturn[idx].Err
}

func (p *DrivePager) SetNext(string) {}

func (p *DrivePager) ValuesIn(api.PageLinker) ([]models.Driveable, error) {
	idx := p.GetIdx
	if idx > 0 {
		// Return values lag by one since we increment in GetPage().
		idx--
	}

	if len(p.ToReturn) <= idx {
		return nil, clues.New("ToReturn index out of bounds")
	}

	return p.ToReturn[idx].Drives, nil
}
