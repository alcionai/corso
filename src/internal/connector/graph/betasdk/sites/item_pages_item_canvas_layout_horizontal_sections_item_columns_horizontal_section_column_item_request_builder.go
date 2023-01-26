package sites

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"

	ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters the set of vertical columns in this section.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters
}

// ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
//
//nolint:lll
type ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal instantiates a new HorizontalSectionColumnItemRequestBuilder and sets the default values.
//
//nolint:lll
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder {
	m := &ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/pages/{sitePage%2Did}/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}{?%24select,%24expand}"
	urlTplParams := make(map[string]string)

	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}

	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter

	return m
}

// NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder instantiates a new HorizontalSectionColumnItemRequestBuilder and sets the default values.
//
//nolint:lll, revive
func NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl

	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlParams, requestAdapter)
}

// CreateDeleteRequestInformation delete navigation property columns for sites
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE

	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}

	return requestInfo, nil
}

// CreateGetRequestInformation the set of vertical columns in this section.
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "application/json")

	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}

		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}

	return requestInfo, nil
}

// CreatePatchRequestInformation update the navigation property columns in sites
//
//nolint:lll,errcheck
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
	requestInfo.Headers.Add("Accept", "application/json")
	requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)

	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}

	return requestInfo, nil
}

// Delete delete navigation property columns for sites
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}

	err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return nil
}

// Get the set of vertical columns in this section.
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionColumnable, error) {
	requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}

	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionColumnable), nil
}

// Patch update the navigation property columns in sites
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Patch(ctx context.Context, body ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration) (ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionColumnable, error) {
	requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
		"5XX": i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0.CreateODataErrorFromDiscriminatorValue,
	}

	res, err := m.requestAdapter.Send(ctx, requestInfo, ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.HorizontalSectionColumnable), nil
}

// Webparts provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
//
//nolint:lll
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Webparts() *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder {
	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// WebpartsById provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
//
//nolint:lll,revive
func (m *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) WebpartsById(id string) *ItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}

	if id != "" {
		urlTplParams["webPart%2Did"] = id
	}

	return NewItemPagesItemCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
