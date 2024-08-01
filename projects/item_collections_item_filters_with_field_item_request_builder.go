package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b "github.com/xifanyan/kiota-searchwebapi/models"
)

// ItemCollectionsItemFiltersWithFieldItemRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\filters\{fieldId}
type ItemCollectionsItemFiltersWithFieldItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemFiltersWithFieldItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFiltersWithFieldItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsItemFiltersWithFieldItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFiltersWithFieldItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCollectionsItemFiltersWithFieldItemRequestBuilderInternal instantiates a new ItemCollectionsItemFiltersWithFieldItemRequestBuilder and sets the default values.
func NewItemCollectionsItemFiltersWithFieldItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFiltersWithFieldItemRequestBuilder) {
    m := &ItemCollectionsItemFiltersWithFieldItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/filters/{fieldId}", pathParameters),
    }
    return m
}
// NewItemCollectionsItemFiltersWithFieldItemRequestBuilder instantiates a new ItemCollectionsItemFiltersWithFieldItemRequestBuilder and sets the default values.
func NewItemCollectionsItemFiltersWithFieldItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFiltersWithFieldItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemFiltersWithFieldItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get discover available resources for a folder field (taxonomy, smart filter).
// returns a FolderFieldResourcesResultable when successful
func (m *ItemCollectionsItemFiltersWithFieldItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersWithFieldItemRequestBuilderGetRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FolderFieldResourcesResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateFolderFieldResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FolderFieldResourcesResultable), nil
}
// Post discover available resources for a folder field (taxonomy, smart filter).
// returns a FolderFieldResourcesResultable when successful
func (m *ItemCollectionsItemFiltersWithFieldItemRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersWithFieldItemRequestBuilderPostRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FolderFieldResourcesResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateFolderFieldResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FolderFieldResourcesResultable), nil
}
// ToGetRequestInformation discover available resources for a folder field (taxonomy, smart filter).
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFiltersWithFieldItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersWithFieldItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation discover available resources for a folder field (taxonomy, smart filter).
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFiltersWithFieldItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFiltersWithFieldItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// Values the values property
// returns a *ItemCollectionsItemFiltersItemValuesRequestBuilder when successful
func (m *ItemCollectionsItemFiltersWithFieldItemRequestBuilder) Values()(*ItemCollectionsItemFiltersItemValuesRequestBuilder) {
    return NewItemCollectionsItemFiltersItemValuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemFiltersWithFieldItemRequestBuilder when successful
func (m *ItemCollectionsItemFiltersWithFieldItemRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemFiltersWithFieldItemRequestBuilder) {
    return NewItemCollectionsItemFiltersWithFieldItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
