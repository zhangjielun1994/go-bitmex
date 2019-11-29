// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Notification Account Notifications
// swagger:model Notification
type Notification struct {

	// body
	// Required: true
	Body *string `json:"body"`

	// closable
	Closable *bool `json:"closable,omitempty"`

	// date
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// id
	ID int32 `json:"id,omitempty"`

	// persist
	Persist *bool `json:"persist,omitempty"`

	// sound
	Sound string `json:"sound,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// ttl
	// Required: true
	TTL *int32 `json:"ttl"`

	// type
	// Enum: [success error info]
	Type string `json:"type,omitempty"`

	// wait for visibility
	WaitForVisibility *bool `json:"waitForVisibility,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notification) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateTTL(formats strfmt.Registry) error {

	if err := validate.Required("ttl", "body", m.TTL); err != nil {
		return err
	}

	return nil
}

var notificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success","error","info"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notificationTypeTypePropEnum = append(notificationTypeTypePropEnum, v)
	}
}

const (

	// NotificationTypeSuccess captures enum value "success"
	NotificationTypeSuccess string = "success"

	// NotificationTypeError captures enum value "error"
	NotificationTypeError string = "error"

	// NotificationTypeInfo captures enum value "info"
	NotificationTypeInfo string = "info"
)

// prop value enum
func (m *Notification) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, notificationTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Notification) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}