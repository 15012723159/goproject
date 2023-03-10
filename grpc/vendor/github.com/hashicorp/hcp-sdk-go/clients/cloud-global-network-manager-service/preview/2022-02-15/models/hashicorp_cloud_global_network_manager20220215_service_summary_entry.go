// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry hashicorp cloud global network manager 20220215 service summary entry
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ServiceSummary.Entry
type HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry struct {

	// cluster is the ID of the cluster the summary is from
	Cluster string `json:"cluster,omitempty"`

	// critical
	Critical int32 `json:"critical,omitempty"`

	// namespace of the service
	Namespace string `json:"namespace,omitempty"`

	// partition of the service
	Partition string `json:"partition,omitempty"`

	// passing
	Passing int32 `json:"passing,omitempty"`

	// total
	Total int32 `json:"total,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// warning
	Warning int32 `json:"warning,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 service summary entry
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hashicorp cloud global network manager 20220215 service summary entry based on context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ServiceSummaryEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
