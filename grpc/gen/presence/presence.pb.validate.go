// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: presence/presence.proto

package presence

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

// Validate checks the field values on MessageA with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MessageA) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageA with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MessageAMultiError, or nil
// if none found.
func (m *MessageA) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageA) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for A

	if len(errors) > 0 {
		return MessageAMultiError(errors)
	}

	return nil
}

// MessageAMultiError is an error wrapping multiple validation errors returned
// by MessageA.ValidateAll() if the designated constraints aren't met.
type MessageAMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageAMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageAMultiError) AllErrors() []error { return m }

// MessageAValidationError is the validation error returned by
// MessageA.Validate if the designated constraints aren't met.
type MessageAValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageAValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageAValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageAValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageAValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageAValidationError) ErrorName() string { return "MessageAValidationError" }

// Error satisfies the builtin error interface
func (e MessageAValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageA.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageAValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageAValidationError{}

// Validate checks the field values on MessageB with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MessageB) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageB with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MessageBMultiError, or nil
// if none found.
func (m *MessageB) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageB) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for A

	if all {
		switch v := interface{}(m.GetB()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageBValidationError{
					field:  "B",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageBValidationError{
					field:  "B",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetB()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageBValidationError{
				field:  "B",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.C != nil {
		// no validation rules for C
	}

	if len(errors) > 0 {
		return MessageBMultiError(errors)
	}

	return nil
}

// MessageBMultiError is an error wrapping multiple validation errors returned
// by MessageB.ValidateAll() if the designated constraints aren't met.
type MessageBMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageBMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageBMultiError) AllErrors() []error { return m }

// MessageBValidationError is the validation error returned by
// MessageB.Validate if the designated constraints aren't met.
type MessageBValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageBValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageBValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageBValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageBValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageBValidationError) ErrorName() string { return "MessageBValidationError" }

// Error satisfies the builtin error interface
func (e MessageBValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageB.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageBValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageBValidationError{}