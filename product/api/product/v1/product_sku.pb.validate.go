// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: product/api/product/v1/product_sku.proto

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

// Validate checks the field values on CreateProductSku with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateProductSku) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductSku with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductSkuMultiError, or nil if none found.
func (m *CreateProductSku) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductSku) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	if all {
		switch v := interface{}(m.GetMainPic()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateProductSkuValidationError{
					field:  "MainPic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateProductSkuValidationError{
					field:  "MainPic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMainPic()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProductSkuValidationError{
				field:  "MainPic",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetMedias() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateProductSkuValidationError{
						field:  fmt.Sprintf("Medias[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateProductSkuValidationError{
						field:  fmt.Sprintf("Medias[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateProductSkuValidationError{
					field:  fmt.Sprintf("Medias[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPrices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateProductSkuValidationError{
						field:  fmt.Sprintf("Prices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateProductSkuValidationError{
						field:  fmt.Sprintf("Prices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateProductSkuValidationError{
					field:  fmt.Sprintf("Prices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Barcode

	for idx, item := range m.GetStock() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateProductSkuValidationError{
						field:  fmt.Sprintf("Stock[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateProductSkuValidationError{
						field:  fmt.Sprintf("Stock[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateProductSkuValidationError{
					field:  fmt.Sprintf("Stock[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateProductSkuMultiError(errors)
	}

	return nil
}

// CreateProductSkuMultiError is an error wrapping multiple validation errors
// returned by CreateProductSku.ValidateAll() if the designated constraints
// aren't met.
type CreateProductSkuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductSkuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductSkuMultiError) AllErrors() []error { return m }

// CreateProductSkuValidationError is the validation error returned by
// CreateProductSku.Validate if the designated constraints aren't met.
type CreateProductSkuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductSkuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductSkuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductSkuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductSkuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductSkuValidationError) ErrorName() string { return "CreateProductSkuValidationError" }

// Error satisfies the builtin error interface
func (e CreateProductSkuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductSku.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductSkuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductSkuValidationError{}

// Validate checks the field values on ProductSku with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProductSku) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProductSku with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProductSkuMultiError, or
// nil if none found.
func (m *ProductSku) ValidateAll() error {
	return m.validate(true)
}

func (m *ProductSku) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	if all {
		switch v := interface{}(m.GetMainPic()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductSkuValidationError{
					field:  "MainPic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductSkuValidationError{
					field:  "MainPic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMainPic()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductSkuValidationError{
				field:  "MainPic",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetMedias() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProductSkuValidationError{
						field:  fmt.Sprintf("Medias[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProductSkuValidationError{
						field:  fmt.Sprintf("Medias[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductSkuValidationError{
					field:  fmt.Sprintf("Medias[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPrices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProductSkuValidationError{
						field:  fmt.Sprintf("Prices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProductSkuValidationError{
						field:  fmt.Sprintf("Prices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductSkuValidationError{
					field:  fmt.Sprintf("Prices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Barcode

	for idx, item := range m.GetStock() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProductSkuValidationError{
						field:  fmt.Sprintf("Stock[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProductSkuValidationError{
						field:  fmt.Sprintf("Stock[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProductSkuValidationError{
					field:  fmt.Sprintf("Stock[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ProductSkuMultiError(errors)
	}

	return nil
}

// ProductSkuMultiError is an error wrapping multiple validation errors
// returned by ProductSku.ValidateAll() if the designated constraints aren't met.
type ProductSkuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductSkuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductSkuMultiError) AllErrors() []error { return m }

// ProductSkuValidationError is the validation error returned by
// ProductSku.Validate if the designated constraints aren't met.
type ProductSkuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductSkuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductSkuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductSkuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductSkuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductSkuValidationError) ErrorName() string { return "ProductSkuValidationError" }

// Error satisfies the builtin error interface
func (e ProductSkuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProductSku.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductSkuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductSkuValidationError{}
