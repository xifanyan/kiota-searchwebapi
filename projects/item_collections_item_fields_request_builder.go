package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b "github.com/xifanyan/kiota-searchwebapi/models"
)

// ItemCollectionsItemFieldsRequestBuilder builds and executes requests for operations under \projects\{projectId}\collections\{collectionId}\fields
type ItemCollectionsItemFieldsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCollectionsItemFieldsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFieldsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCollectionsItemFieldsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCollectionsItemFieldsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCollectionsItemFieldsRequestBuilderInternal instantiates a new ItemCollectionsItemFieldsRequestBuilder and sets the default values.
func NewItemCollectionsItemFieldsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFieldsRequestBuilder) {
    m := &ItemCollectionsItemFieldsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}/collections/{collectionId}/fields", pathParameters),
    }
    return m
}
// NewItemCollectionsItemFieldsRequestBuilder instantiates a new ItemCollectionsItemFieldsRequestBuilder and sets the default values.
func NewItemCollectionsItemFieldsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCollectionsItemFieldsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCollectionsItemFieldsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the list of available fields and their properties. Note that a folder collection (i.e. a taxonomy) typically has no defined data model such that this end point is unsupported for folder collections. Folder collections always know the field rm_display_name; hierarchical folder collections also know rm_prop_parent (which resembles the folder id of the hierarchical parent).
// returns a FieldsResultable when successful
func (m *ItemCollectionsItemFieldsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCollectionsItemFieldsRequestBuilderGetRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FieldsResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateFieldsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FieldsResultable), nil
}
// Post returns the list of available fields and their properties. Note that a folder collection (i.e. a taxonomy) typically has no defined data model such that this end point is unsupported for folder collections. Folder collections always know the field rm_display_name; hierarchical folder collections also know rm_prop_parent (which resembles the folder id of the hierarchical parent).
// returns a FieldsResultable when successful
func (m *ItemCollectionsItemFieldsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCollectionsItemFieldsRequestBuilderPostRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FieldsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateFieldsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.FieldsResultable), nil
}
// ToGetRequestInformation returns the list of available fields and their properties. Note that a folder collection (i.e. a taxonomy) typically has no defined data model such that this end point is unsupported for folder collections. Folder collections always know the field rm_display_name; hierarchical folder collections also know rm_prop_parent (which resembles the folder id of the hierarchical parent).
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFieldsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFieldsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation returns the list of available fields and their properties. Note that a folder collection (i.e. a taxonomy) typically has no defined data model such that this end point is unsupported for folder collections. Folder collections always know the field rm_display_name; hierarchical folder collections also know rm_prop_parent (which resembles the folder id of the hierarchical parent).
// returns a *RequestInformation when successful
func (m *ItemCollectionsItemFieldsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCollectionsItemFieldsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCollectionsItemFieldsRequestBuilder when successful
func (m *ItemCollectionsItemFieldsRequestBuilder) WithUrl(rawUrl string)(*ItemCollectionsItemFieldsRequestBuilder) {
    return NewItemCollectionsItemFieldsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
