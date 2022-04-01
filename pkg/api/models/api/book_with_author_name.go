// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BookWithAuthorName book with author name
//
// swagger:model BookWithAuthorName
type BookWithAuthorName struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this book with author name
func (m *BookWithAuthorName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this book with author name based on context it is used
func (m *BookWithAuthorName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BookWithAuthorName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookWithAuthorName) UnmarshalBinary(b []byte) error {
	var res BookWithAuthorName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
