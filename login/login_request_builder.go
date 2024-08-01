package login

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b "github.com/xifanyan/kiota-searchwebapi/models"
)

// LoginRequestBuilder builds and executes requests for operations under \login
type LoginRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LoginRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LoginRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLoginRequestBuilderInternal instantiates a new LoginRequestBuilder and sets the default values.
func NewLoginRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoginRequestBuilder) {
    m := &LoginRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/login", pathParameters),
    }
    return m
}
// NewLoginRequestBuilder instantiates a new LoginRequestBuilder and sets the default values.
func NewLoginRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoginRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLoginRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a new session. NOTE: all other endpoints implicitly create new sessions based on request headers as well. This method merely allows an explicit login without using side effects. The session id is returned as response header SWA-SESSION and must be provided to all following requests.
// returns a LoginResultable when successful
func (m *LoginRequestBuilder) Post(ctx context.Context, requestConfiguration *LoginRequestBuilderPostRequestConfiguration)(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.LoginResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.CreateLoginResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i4eadb8137c655db471dd9174bc9e7e392dcdba3d1eda802a7656e87879605f9b.LoginResultable), nil
}
// ToPostRequestInformation creates a new session. NOTE: all other endpoints implicitly create new sessions based on request headers as well. This method merely allows an explicit login without using side effects. The session id is returned as response header SWA-SESSION and must be provided to all following requests.
// returns a *RequestInformation when successful
func (m *LoginRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *LoginRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LoginRequestBuilder when successful
func (m *LoginRequestBuilder) WithUrl(rawUrl string)(*LoginRequestBuilder) {
    return NewLoginRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
