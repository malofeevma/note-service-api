// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: note.proto

package note_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on NoteInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NoteInfoMultiError, or nil
// if none found.
func (m *NoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *NoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 30 {
		err := NoteInfoValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 30 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Text

	if l := utf8.RuneCountInString(m.GetAuthor()); l < 1 || l > 20 {
		err := NoteInfoValidationError{
			field:  "Author",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NoteInfoMultiError(errors)
	}

	return nil
}

// NoteInfoMultiError is an error wrapping multiple validation errors returned
// by NoteInfo.ValidateAll() if the designated constraints aren't met.
type NoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteInfoMultiError) AllErrors() []error { return m }

// NoteInfoValidationError is the validation error returned by
// NoteInfo.Validate if the designated constraints aren't met.
type NoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteInfoValidationError) ErrorName() string { return "NoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e NoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteInfoValidationError{}

// Validate checks the field values on Note with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Note) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Note with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NoteMultiError, or nil if none found.
func (m *Note) ValidateAll() error {
	return m.validate(true)
}

func (m *Note) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "UpdateAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "UpdateAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "UpdateAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NoteMultiError(errors)
	}

	return nil
}

// NoteMultiError is an error wrapping multiple validation errors returned by
// Note.ValidateAll() if the designated constraints aren't met.
type NoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteMultiError) AllErrors() []error { return m }

// NoteValidationError is the validation error returned by Note.Validate if the
// designated constraints aren't met.
type NoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteValidationError) ErrorName() string { return "NoteValidationError" }

// Error satisfies the builtin error interface
func (e NoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteValidationError{}

// Validate checks the field values on UpdateNoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateNoteInfoMultiError,
// or nil if none found.
func (m *UpdateNoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTitle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Title",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Title",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTitle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Title",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetText()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetText()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Text",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNoteInfoMultiError(errors)
	}

	return nil
}

// UpdateNoteInfoMultiError is an error wrapping multiple validation errors
// returned by UpdateNoteInfo.ValidateAll() if the designated constraints
// aren't met.
type UpdateNoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteInfoMultiError) AllErrors() []error { return m }

// UpdateNoteInfoValidationError is the validation error returned by
// UpdateNoteInfo.Validate if the designated constraints aren't met.
type UpdateNoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteInfoValidationError) ErrorName() string { return "UpdateNoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e UpdateNoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteInfoValidationError{}

// Validate checks the field values on CreateNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteRequestMultiError, or nil if none found.
func (m *CreateNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNoteRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNoteRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNoteRequestValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNoteRequestMultiError(errors)
	}

	return nil
}

// CreateNoteRequestMultiError is an error wrapping multiple validation errors
// returned by CreateNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteRequestMultiError) AllErrors() []error { return m }

// CreateNoteRequestValidationError is the validation error returned by
// CreateNoteRequest.Validate if the designated constraints aren't met.
type CreateNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteRequestValidationError) ErrorName() string {
	return "CreateNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteRequestValidationError{}

// Validate checks the field values on CreateNoteResponce with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteResponce) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteResponce with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteResponceMultiError, or nil if none found.
func (m *CreateNoteResponce) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteResponce) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateNoteResponceMultiError(errors)
	}

	return nil
}

// CreateNoteResponceMultiError is an error wrapping multiple validation errors
// returned by CreateNoteResponce.ValidateAll() if the designated constraints
// aren't met.
type CreateNoteResponceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteResponceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteResponceMultiError) AllErrors() []error { return m }

// CreateNoteResponceValidationError is the validation error returned by
// CreateNoteResponce.Validate if the designated constraints aren't met.
type CreateNoteResponceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteResponceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteResponceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteResponceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteResponceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteResponceValidationError) ErrorName() string {
	return "CreateNoteResponceValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteResponceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteResponce.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteResponceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteResponceValidationError{}

// Validate checks the field values on GetNoteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetNoteRequestMultiError,
// or nil if none found.
func (m *GetNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetNoteRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetNoteRequestMultiError(errors)
	}

	return nil
}

// GetNoteRequestMultiError is an error wrapping multiple validation errors
// returned by GetNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteRequestMultiError) AllErrors() []error { return m }

// GetNoteRequestValidationError is the validation error returned by
// GetNoteRequest.Validate if the designated constraints aren't met.
type GetNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteRequestValidationError) ErrorName() string { return "GetNoteRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteRequestValidationError{}

// Validate checks the field values on GetNoteResponce with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNoteResponce) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteResponce with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNoteResponceMultiError, or nil if none found.
func (m *GetNoteResponce) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteResponce) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNoteResponceValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNoteResponceValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNoteResponceValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNoteResponceMultiError(errors)
	}

	return nil
}

// GetNoteResponceMultiError is an error wrapping multiple validation errors
// returned by GetNoteResponce.ValidateAll() if the designated constraints
// aren't met.
type GetNoteResponceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteResponceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteResponceMultiError) AllErrors() []error { return m }

// GetNoteResponceValidationError is the validation error returned by
// GetNoteResponce.Validate if the designated constraints aren't met.
type GetNoteResponceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteResponceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteResponceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteResponceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteResponceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteResponceValidationError) ErrorName() string { return "GetNoteResponceValidationError" }

// Error satisfies the builtin error interface
func (e GetNoteResponceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteResponce.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteResponceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteResponceValidationError{}

// Validate checks the field values on GetListNoteResponce with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListNoteResponce) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListNoteResponce with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListNoteResponceMultiError, or nil if none found.
func (m *GetListNoteResponce) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListNoteResponce) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetListNoteResponceValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListNoteResponceValidationError{
						field:  fmt.Sprintf("Notes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListNoteResponceValidationError{
					field:  fmt.Sprintf("Notes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetListNoteResponceMultiError(errors)
	}

	return nil
}

// GetListNoteResponceMultiError is an error wrapping multiple validation
// errors returned by GetListNoteResponce.ValidateAll() if the designated
// constraints aren't met.
type GetListNoteResponceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListNoteResponceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListNoteResponceMultiError) AllErrors() []error { return m }

// GetListNoteResponceValidationError is the validation error returned by
// GetListNoteResponce.Validate if the designated constraints aren't met.
type GetListNoteResponceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListNoteResponceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListNoteResponceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListNoteResponceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListNoteResponceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListNoteResponceValidationError) ErrorName() string {
	return "GetListNoteResponceValidationError"
}

// Error satisfies the builtin error interface
func (e GetListNoteResponceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListNoteResponce.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListNoteResponceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListNoteResponceValidationError{}

// Validate checks the field values on UpdateNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNoteRequestMultiError, or nil if none found.
func (m *UpdateNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetUpdateInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteRequestValidationError{
					field:  "UpdateInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteRequestValidationError{
					field:  "UpdateInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteRequestValidationError{
				field:  "UpdateInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNoteRequestMultiError(errors)
	}

	return nil
}

// UpdateNoteRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteRequestMultiError) AllErrors() []error { return m }

// UpdateNoteRequestValidationError is the validation error returned by
// UpdateNoteRequest.Validate if the designated constraints aren't met.
type UpdateNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteRequestValidationError) ErrorName() string {
	return "UpdateNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteRequestValidationError{}

// Validate checks the field values on DeleteNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNoteRequestMultiError, or nil if none found.
func (m *DeleteNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteNoteRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteNoteRequestMultiError(errors)
	}

	return nil
}

// DeleteNoteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNoteRequestMultiError) AllErrors() []error { return m }

// DeleteNoteRequestValidationError is the validation error returned by
// DeleteNoteRequest.Validate if the designated constraints aren't met.
type DeleteNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNoteRequestValidationError) ErrorName() string {
	return "DeleteNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNoteRequestValidationError{}
