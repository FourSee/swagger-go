// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PairingRequestCreateResponse pairing request create response
// swagger:model PairingRequestCreateResponse
type PairingRequestCreateResponse struct {

	// accept url
	AcceptURL string `json:"accept_url,omitempty"`

	// show url
	ShowURL string `json:"show_url,omitempty"`
}

// Validate validates this pairing request create response
func (m *PairingRequestCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PairingRequestCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PairingRequestCreateResponse) UnmarshalBinary(b []byte) error {
	var res PairingRequestCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}