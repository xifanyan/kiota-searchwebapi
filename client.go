package kiotasearchwebapi

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i46acb77173f5e598a5625d17c50c7e1aab8a9da0dd4c2cc28a19d23dd2a8bd7b "github.com/xifanyan/kiota-searchwebapi/projects"
    i6c1725ab3a344cfd0ead45ef9dbc59d9517c57be52a90ff893abc9b997ca7e01 "github.com/xifanyan/kiota-searchwebapi/jobs"
    i83153a9a606a784b9ef8ed2be7d8ca4c36ccea90cd7242c26e657acd10419256 "github.com/xifanyan/kiota-searchwebapi/logout"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i9e134dc2e18e1b6c5033dd38552fd30678fd3bbef123ee94a22d9bd2fe49cdf7 "github.com/xifanyan/kiota-searchwebapi/login"
)

// Client the main entry point of the SDK, exposes the configuration and the fluent API.
type Client struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewClient instantiates a new Client and sets the default values.
func NewClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Client) {
    m := &Client{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://vm-rhauswirth2.otxlab.net:8443/searchWebApi")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Jobs the jobs property
// returns a *JobsRequestBuilder when successful
func (m *Client) Jobs()(*i6c1725ab3a344cfd0ead45ef9dbc59d9517c57be52a90ff893abc9b997ca7e01.JobsRequestBuilder) {
    return i6c1725ab3a344cfd0ead45ef9dbc59d9517c57be52a90ff893abc9b997ca7e01.NewJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Login the login property
// returns a *LoginRequestBuilder when successful
func (m *Client) Login()(*i9e134dc2e18e1b6c5033dd38552fd30678fd3bbef123ee94a22d9bd2fe49cdf7.LoginRequestBuilder) {
    return i9e134dc2e18e1b6c5033dd38552fd30678fd3bbef123ee94a22d9bd2fe49cdf7.NewLoginRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Logout the logout property
// returns a *LogoutRequestBuilder when successful
func (m *Client) Logout()(*i83153a9a606a784b9ef8ed2be7d8ca4c36ccea90cd7242c26e657acd10419256.LogoutRequestBuilder) {
    return i83153a9a606a784b9ef8ed2be7d8ca4c36ccea90cd7242c26e657acd10419256.NewLogoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ProjectsRequestBuilder when successful
func (m *Client) Projects()(*i46acb77173f5e598a5625d17c50c7e1aab8a9da0dd4c2cc28a19d23dd2a8bd7b.ProjectsRequestBuilder) {
    return i46acb77173f5e598a5625d17c50c7e1aab8a9da0dd4c2cc28a19d23dd2a8bd7b.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
