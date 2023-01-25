package sites

import (
	"context"

	i7ad325c11fbf3db4d761c429267362d8b24daa1eda0081f914ebc3cdc85181a0 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder provides operations to call the copyToDefaultContentLocation method.
type ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// ItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal instantiates a new CopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder {
	m := &ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder{}
	m.urlTemplate = "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}/microsoft.graph.copyToDefaultContentLocation"
	urlTplParams := make(map[string]string)
	for idx, item := range pathParameters {
		urlTplParams[idx] = item
	}
	m.pathParameters = urlTplParams
	m.requestAdapter = requestAdapter
	return m
}

// NewItemContentTypesItemCopyToDefaultContentLocationRequestBuilder instantiates a new CopyToDefaultContentLocationRequestBuilder and sets the default values.
func NewItemContentTypesItemCopyToDefaultContentLocationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(urlParams, requestAdapter)
}

// CreatePostRequestInformation copy a file to a default content location in a [content type][contentType]. The file can then be added as a default file or template via a POST operation.
func (m *ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.urlTemplate
	requestInfo.PathParameters = m.pathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// Post copy a file to a default content location in a [content type][contentType]. The file can then be added as a default file or template via a POST operation.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-copytodefaultcontentlocation?view=graph-rest-1.0
func (m *ItemContentTypesItemCopyToDefaultContentLocationRequestBuilder) Post(ctx context.Context, body ItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyable, requestConfiguration *ItemContentTypesItemCopyToDefaultContentLocationRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration)
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
