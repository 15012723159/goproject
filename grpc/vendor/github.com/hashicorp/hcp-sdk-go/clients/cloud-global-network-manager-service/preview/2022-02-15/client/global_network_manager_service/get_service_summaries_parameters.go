// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetServiceSummariesParams creates a new GetServiceSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceSummariesParams() *GetServiceSummariesParams {
	return &GetServiceSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceSummariesParamsWithTimeout creates a new GetServiceSummariesParams object
// with the ability to set a timeout on a request.
func NewGetServiceSummariesParamsWithTimeout(timeout time.Duration) *GetServiceSummariesParams {
	return &GetServiceSummariesParams{
		timeout: timeout,
	}
}

// NewGetServiceSummariesParamsWithContext creates a new GetServiceSummariesParams object
// with the ability to set a context for a request.
func NewGetServiceSummariesParamsWithContext(ctx context.Context) *GetServiceSummariesParams {
	return &GetServiceSummariesParams{
		Context: ctx,
	}
}

// NewGetServiceSummariesParamsWithHTTPClient creates a new GetServiceSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceSummariesParamsWithHTTPClient(client *http.Client) *GetServiceSummariesParams {
	return &GetServiceSummariesParams{
		HTTPClient: client,
	}
}

/*
GetServiceSummariesParams contains all the parameters to send to the API endpoint

	for the get service summaries operation.

	Typically these are written to a http.Request.
*/
type GetServiceSummariesParams struct {

	/* FilterCluster.

	   cluster matches summaries on the cluster name.
	*/
	FilterCluster []string

	/* FilterNamePrefix.

	   name_prefix matches summaries on the prefix of the service's name.
	*/
	FilterNamePrefix *string

	/* FilterNamespace.

	   namespace matches summaries on the namespace name.
	*/
	FilterNamespace []string

	/* FilterPartition.

	   partition matches summaries on the partition name.
	*/
	FilterPartition []string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	LocationRegionRegion *string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get service summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceSummariesParams) WithDefaults() *GetServiceSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service summaries params
func (o *GetServiceSummariesParams) WithTimeout(timeout time.Duration) *GetServiceSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service summaries params
func (o *GetServiceSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service summaries params
func (o *GetServiceSummariesParams) WithContext(ctx context.Context) *GetServiceSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service summaries params
func (o *GetServiceSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service summaries params
func (o *GetServiceSummariesParams) WithHTTPClient(client *http.Client) *GetServiceSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service summaries params
func (o *GetServiceSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterCluster adds the filterCluster to the get service summaries params
func (o *GetServiceSummariesParams) WithFilterCluster(filterCluster []string) *GetServiceSummariesParams {
	o.SetFilterCluster(filterCluster)
	return o
}

// SetFilterCluster adds the filterCluster to the get service summaries params
func (o *GetServiceSummariesParams) SetFilterCluster(filterCluster []string) {
	o.FilterCluster = filterCluster
}

// WithFilterNamePrefix adds the filterNamePrefix to the get service summaries params
func (o *GetServiceSummariesParams) WithFilterNamePrefix(filterNamePrefix *string) *GetServiceSummariesParams {
	o.SetFilterNamePrefix(filterNamePrefix)
	return o
}

// SetFilterNamePrefix adds the filterNamePrefix to the get service summaries params
func (o *GetServiceSummariesParams) SetFilterNamePrefix(filterNamePrefix *string) {
	o.FilterNamePrefix = filterNamePrefix
}

// WithFilterNamespace adds the filterNamespace to the get service summaries params
func (o *GetServiceSummariesParams) WithFilterNamespace(filterNamespace []string) *GetServiceSummariesParams {
	o.SetFilterNamespace(filterNamespace)
	return o
}

// SetFilterNamespace adds the filterNamespace to the get service summaries params
func (o *GetServiceSummariesParams) SetFilterNamespace(filterNamespace []string) {
	o.FilterNamespace = filterNamespace
}

// WithFilterPartition adds the filterPartition to the get service summaries params
func (o *GetServiceSummariesParams) WithFilterPartition(filterPartition []string) *GetServiceSummariesParams {
	o.SetFilterPartition(filterPartition)
	return o
}

// SetFilterPartition adds the filterPartition to the get service summaries params
func (o *GetServiceSummariesParams) SetFilterPartition(filterPartition []string) {
	o.FilterPartition = filterPartition
}

// WithLocationOrganizationID adds the locationOrganizationID to the get service summaries params
func (o *GetServiceSummariesParams) WithLocationOrganizationID(locationOrganizationID string) *GetServiceSummariesParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get service summaries params
func (o *GetServiceSummariesParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get service summaries params
func (o *GetServiceSummariesParams) WithLocationProjectID(locationProjectID string) *GetServiceSummariesParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get service summaries params
func (o *GetServiceSummariesParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get service summaries params
func (o *GetServiceSummariesParams) WithLocationRegionProvider(locationRegionProvider *string) *GetServiceSummariesParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get service summaries params
func (o *GetServiceSummariesParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get service summaries params
func (o *GetServiceSummariesParams) WithLocationRegionRegion(locationRegionRegion *string) *GetServiceSummariesParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get service summaries params
func (o *GetServiceSummariesParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the get service summaries params
func (o *GetServiceSummariesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *GetServiceSummariesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the get service summaries params
func (o *GetServiceSummariesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the get service summaries params
func (o *GetServiceSummariesParams) WithPaginationPageSize(paginationPageSize *int64) *GetServiceSummariesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the get service summaries params
func (o *GetServiceSummariesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the get service summaries params
func (o *GetServiceSummariesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *GetServiceSummariesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the get service summaries params
func (o *GetServiceSummariesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCluster != nil {

		// binding items for filter.cluster
		joinedFilterCluster := o.bindParamFilterCluster(reg)

		// query array param filter.cluster
		if err := r.SetQueryParam("filter.cluster", joinedFilterCluster...); err != nil {
			return err
		}
	}

	if o.FilterNamePrefix != nil {

		// query param filter.name_prefix
		var qrFilterNamePrefix string

		if o.FilterNamePrefix != nil {
			qrFilterNamePrefix = *o.FilterNamePrefix
		}
		qFilterNamePrefix := qrFilterNamePrefix
		if qFilterNamePrefix != "" {

			if err := r.SetQueryParam("filter.name_prefix", qFilterNamePrefix); err != nil {
				return err
			}
		}
	}

	if o.FilterNamespace != nil {

		// binding items for filter.namespace
		joinedFilterNamespace := o.bindParamFilterNamespace(reg)

		// query array param filter.namespace
		if err := r.SetQueryParam("filter.namespace", joinedFilterNamespace...); err != nil {
			return err
		}
	}

	if o.FilterPartition != nil {

		// binding items for filter.partition
		joinedFilterPartition := o.bindParamFilterPartition(reg)

		// query array param filter.partition
		if err := r.SetQueryParam("filter.partition", joinedFilterPartition...); err != nil {
			return err
		}
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string

		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {

			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string

		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {

			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetServiceSummaries binds the parameter filter.cluster
func (o *GetServiceSummariesParams) bindParamFilterCluster(formats strfmt.Registry) []string {
	filterClusterIR := o.FilterCluster

	var filterClusterIC []string
	for _, filterClusterIIR := range filterClusterIR { // explode []string

		filterClusterIIV := filterClusterIIR // string as string
		filterClusterIC = append(filterClusterIC, filterClusterIIV)
	}

	// items.CollectionFormat: "multi"
	filterClusterIS := swag.JoinByFormat(filterClusterIC, "multi")

	return filterClusterIS
}

// bindParamGetServiceSummaries binds the parameter filter.namespace
func (o *GetServiceSummariesParams) bindParamFilterNamespace(formats strfmt.Registry) []string {
	filterNamespaceIR := o.FilterNamespace

	var filterNamespaceIC []string
	for _, filterNamespaceIIR := range filterNamespaceIR { // explode []string

		filterNamespaceIIV := filterNamespaceIIR // string as string
		filterNamespaceIC = append(filterNamespaceIC, filterNamespaceIIV)
	}

	// items.CollectionFormat: "multi"
	filterNamespaceIS := swag.JoinByFormat(filterNamespaceIC, "multi")

	return filterNamespaceIS
}

// bindParamGetServiceSummaries binds the parameter filter.partition
func (o *GetServiceSummariesParams) bindParamFilterPartition(formats strfmt.Registry) []string {
	filterPartitionIR := o.FilterPartition

	var filterPartitionIC []string
	for _, filterPartitionIIR := range filterPartitionIR { // explode []string

		filterPartitionIIV := filterPartitionIIR // string as string
		filterPartitionIC = append(filterPartitionIC, filterPartitionIIV)
	}

	// items.CollectionFormat: "multi"
	filterPartitionIS := swag.JoinByFormat(filterPartitionIC, "multi")

	return filterPartitionIS
}
