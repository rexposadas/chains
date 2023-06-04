// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Block block
//
// swagger:model block
type Block struct {

	// bpm
	Bpm int64 `json:"bpm,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// index
	Index int64 `json:"index,omitempty"`

	// prev hash
	PrevHash string `json:"prevHash,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this block
func (m *Block) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block based on context it is used
func (m *Block) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Block) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Block) UnmarshalBinary(b []byte) error {
	var res Block
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}