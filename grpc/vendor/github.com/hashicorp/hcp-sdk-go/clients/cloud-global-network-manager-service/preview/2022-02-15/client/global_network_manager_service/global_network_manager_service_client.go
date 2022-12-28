// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new global network manager service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for global network manager service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AgentBootstrapConfig(params *AgentBootstrapConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentBootstrapConfigOK, error)

	AgentDiscover(params *AgentDiscoverParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentDiscoverOK, error)

	AgentPushServerState(params *AgentPushServerStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentPushServerStateOK, error)

	CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClusterOK, error)

	DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterOK, error)

	GetAggregateServiceSummary(params *GetAggregateServiceSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAggregateServiceSummaryOK, error)

	GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterOK, error)

	GetClusterAPIInfo(params *GetClusterAPIInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterAPIInfoOK, error)

	GetClusterSecrets(params *GetClusterSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterSecretsOK, error)

	GetServiceSummaries(params *GetServiceSummariesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceSummariesOK, error)

	ListClusterServers(params *ListClusterServersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClusterServersOK, error)

	ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClustersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AgentBootstrapConfig agent bootstrap config API
*/
func (a *Client) AgentBootstrapConfig(params *AgentBootstrapConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentBootstrapConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAgentBootstrapConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AgentBootstrapConfig",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/bootstrap_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AgentBootstrapConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AgentBootstrapConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AgentBootstrapConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AgentDiscover agent discover API
*/
func (a *Client) AgentDiscover(params *AgentDiscoverParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentDiscoverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAgentDiscoverParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AgentDiscover",
		Method:             "POST",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/discover",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AgentDiscoverReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AgentDiscoverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AgentDiscoverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AgentPushServerState agent push server state API
*/
func (a *Client) AgentPushServerState(params *AgentPushServerStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AgentPushServerStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAgentPushServerStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AgentPushServerState",
		Method:             "POST",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/agent/server-state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AgentPushServerStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AgentPushServerStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AgentPushServerStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateCluster create cluster API
*/
func (a *Client) CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCluster",
		Method:             "POST",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteCluster delete cluster API
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCluster",
		Method:             "DELETE",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAggregateServiceSummary get aggregate service summary API
*/
func (a *Client) GetAggregateServiceSummary(params *GetAggregateServiceSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAggregateServiceSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAggregateServiceSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAggregateServiceSummary",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/aggregate_service_summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAggregateServiceSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAggregateServiceSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAggregateServiceSummaryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCluster get cluster API
*/
func (a *Client) GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCluster",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterAPIInfo get cluster API info API
*/
func (a *Client) GetClusterAPIInfo(params *GetClusterAPIInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterAPIInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterAPIInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterAPIInfo",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/api_information",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterAPIInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterAPIInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterAPIInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterSecrets get cluster secrets API
*/
func (a *Client) GetClusterSecrets(params *GetClusterSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClusterSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterSecretsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterSecrets",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/secrets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterSecretsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterSecretsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetServiceSummaries get service summaries API
*/
func (a *Client) GetServiceSummaries(params *GetServiceSummariesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceSummariesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceSummariesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetServiceSummaries",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceSummariesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceSummariesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServiceSummariesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListClusterServers list cluster servers API
*/
func (a *Client) ListClusterServers(params *ListClusterServersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClusterServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClusterServersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListClusterServers",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{id}/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClusterServersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListClusterServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListClusterServersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListClusters list clusters API
*/
func (a *Client) ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListClusters",
		Method:             "GET",
		PathPattern:        "/global-network-manager/2022-02-15/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}