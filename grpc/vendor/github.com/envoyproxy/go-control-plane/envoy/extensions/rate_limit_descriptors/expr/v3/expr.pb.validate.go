// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/rate_limit_descriptors/expr/v3/expr.proto

package exprv3

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

// Validate checks the field values on Descriptor with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Descriptor) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Descriptor with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DescriptorMultiError, or
// nil if none found.
func (m *Descriptor) ValidateAll() error {
	return m.validate(true)
}

func (m *Descriptor) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDescriptorKey()) < 1 {
		err := DescriptorValidationError{
			field:  "DescriptorKey",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SkipIfError

	switch m.ExprSpecifier.(type) {

	case *Descriptor_Text:

		if utf8.RuneCountInString(m.GetText()) < 1 {
			err := DescriptorValidationError{
				field:  "Text",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *Descriptor_Parsed:

		if all {
			switch v := interface{}(m.GetParsed()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DescriptorValidationError{
						field:  "Parsed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DescriptorValidationError{
						field:  "Parsed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetParsed()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DescriptorValidationError{
					field:  "Parsed",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DescriptorMultiError(errors)
	}

	return nil
}

// DescriptorMultiError is an error wrapping multiple validation errors
// returned by Descriptor.ValidateAll() if the designated constraints aren't met.
type DescriptorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DescriptorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DescriptorMultiError) AllErrors() []error { return m }

// DescriptorValidationError is the validation error returned by
// Descriptor.Validate if the designated constraints aren't met.
type DescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescriptorValidationError) ErrorName() string { return "DescriptorValidationError" }

// Error satisfies the builtin error interface
func (e DescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescriptorValidationError{}
