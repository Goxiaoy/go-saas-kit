// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order/event/v1/event.proto

package v1

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

// Validate checks the field values on OrderPaySuccessEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderPaySuccessEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderPaySuccessEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderPaySuccessEventMultiError, or nil if none found.
func (m *OrderPaySuccessEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderPaySuccessEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderPaySuccessEventValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderPaySuccessEventValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderPaySuccessEventValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderPaySuccessEventMultiError(errors)
	}

	return nil
}

// OrderPaySuccessEventMultiError is an error wrapping multiple validation
// errors returned by OrderPaySuccessEvent.ValidateAll() if the designated
// constraints aren't met.
type OrderPaySuccessEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderPaySuccessEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderPaySuccessEventMultiError) AllErrors() []error { return m }

// OrderPaySuccessEventValidationError is the validation error returned by
// OrderPaySuccessEvent.Validate if the designated constraints aren't met.
type OrderPaySuccessEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderPaySuccessEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderPaySuccessEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderPaySuccessEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderPaySuccessEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderPaySuccessEventValidationError) ErrorName() string {
	return "OrderPaySuccessEventValidationError"
}

// Error satisfies the builtin error interface
func (e OrderPaySuccessEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderPaySuccessEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderPaySuccessEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderPaySuccessEventValidationError{}

// Validate checks the field values on OrderRefundSuccessEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderRefundSuccessEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderRefundSuccessEvent with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderRefundSuccessEventMultiError, or nil if none found.
func (m *OrderRefundSuccessEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderRefundSuccessEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderRefundSuccessEventValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderRefundSuccessEventValidationError{
					field:  "Order",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderRefundSuccessEventValidationError{
				field:  "Order",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrderRefundSuccessEventMultiError(errors)
	}

	return nil
}

// OrderRefundSuccessEventMultiError is an error wrapping multiple validation
// errors returned by OrderRefundSuccessEvent.ValidateAll() if the designated
// constraints aren't met.
type OrderRefundSuccessEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderRefundSuccessEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderRefundSuccessEventMultiError) AllErrors() []error { return m }

// OrderRefundSuccessEventValidationError is the validation error returned by
// OrderRefundSuccessEvent.Validate if the designated constraints aren't met.
type OrderRefundSuccessEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderRefundSuccessEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderRefundSuccessEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderRefundSuccessEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderRefundSuccessEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderRefundSuccessEventValidationError) ErrorName() string {
	return "OrderRefundSuccessEventValidationError"
}

// Error satisfies the builtin error interface
func (e OrderRefundSuccessEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderRefundSuccessEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderRefundSuccessEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderRefundSuccessEventValidationError{}
