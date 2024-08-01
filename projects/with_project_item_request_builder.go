package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b "github.com/xifanyan/kiota-searchwebapi/models"
)

// WithProjectItemRequestBuilder builds and executes requests for operations under \projects\{projectId}
type WithProjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithProjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithProjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithProjectItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithProjectItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Collections the collections property
// returns a *ItemCollectionsRequestBuilder when successful
func (m *WithProjectItemRequestBuilder) Collections()(*ItemCollectionsRequestBuilder) {
    return NewItemCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithProjectItemRequestBuilderInternal instantiates a new WithProjectItemRequestBuilder and sets the default values.
func NewWithProjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithProjectItemRequestBuilder) {
    m := &WithProjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/{projectId}", pathParameters),
    }
    return m
}
// NewWithProjectItemRequestBuilder instantiates a new WithProjectItemRequestBuilder and sets the default values.
func NewWithProjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithProjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithProjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get discover available project-level resources.
// returns a ProjectResourcesResultable when successful
func (m *WithProjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithProjectItemRequestBuilderGetRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.ProjectResourcesResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateProjectResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.ProjectResourcesResultable), nil
}
// Post discover available project-level resources.
// returns a ProjectResourcesResultable when successful
func (m *WithProjectItemRequestBuilder) Post(ctx context.Context, requestConfiguration *WithProjectItemRequestBuilderPostRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.ProjectResourcesResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateProjectResourcesResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.ProjectResourcesResultable), nil
}
// ToGetRequestInformation discover available project-level resources.
// returns a *RequestInformation when successful
func (m *WithProjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithProjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation discover available project-level resources.
// returns a *RequestInformation when successful
func (m *WithProjectItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *WithProjectItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithProjectItemRequestBuilder when successful
func (m *WithProjectItemRequestBuilder) WithUrl(rawUrl string)(*WithProjectItemRequestBuilder) {
    return NewWithProjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
