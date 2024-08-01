package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b "github.com/xifanyan/kiota-searchwebapi/models"
)

// ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder builds and executes requests for operations under \jobs\projects\{projectId}\collections\{collectionId}\registeredJobTypes
type ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderInternal instantiates a new ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) {
    m := &ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/projects/{projectId}/collections/{collectionId}/registeredJobTypes", pathParameters),
    }
    return m
}
// NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder instantiates a new ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder and sets the default values.
func NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the result is the (complete) list of jobs which can be submitted using a POST on /jobs.
// returns a []GenericJobDescriptionable when successful
func (m *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderGetRequestConfiguration)([]i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.GenericJobDescriptionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateGenericJobDescriptionFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.GenericJobDescriptionable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.GenericJobDescriptionable)
        }
    }
    return val, nil
}
// ToGetRequestInformation the result is the (complete) list of jobs which can be submitted using a POST on /jobs.
// returns a *RequestInformation when successful
func (m *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder when successful
func (m *ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) WithUrl(rawUrl string)(*ProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder) {
    return NewProjectsItemCollectionsItemRegisteredJobTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}