// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sys/private/conf/conf.proto

package conf

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

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Bootstrap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bootstrap with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BootstrapMultiError, or nil
// if none found.
func (m *Bootstrap) ValidateAll() error {
	return m.validate(true)
}

func (m *Bootstrap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSecurity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Security",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Security",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSecurity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Security",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetServices()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Services",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Services",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServices()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Services",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLogging()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Logging",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Logging",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLogging()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Logging",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTracing()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Tracing",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Tracing",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTracing()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Tracing",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetApp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "App",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDev()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Dev",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Dev",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDev()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Dev",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSys()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Sys",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BootstrapValidationError{
					field:  "Sys",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSys()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Sys",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BootstrapMultiError(errors)
	}

	return nil
}

// BootstrapMultiError is an error wrapping multiple validation errors returned
// by Bootstrap.ValidateAll() if the designated constraints aren't met.
type BootstrapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BootstrapMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BootstrapMultiError) AllErrors() []error { return m }

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}

// Validate checks the field values on SysConf with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysConf with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SysConfMultiError, or nil if none found.
func (m *SysConf) ValidateAll() error {
	return m.validate(true)
}

func (m *SysConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetApisix()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SysConfValidationError{
					field:  "Apisix",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SysConfValidationError{
					field:  "Apisix",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApisix()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SysConfValidationError{
				field:  "Apisix",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SysConfMultiError(errors)
	}

	return nil
}

// SysConfMultiError is an error wrapping multiple validation errors returned
// by SysConf.ValidateAll() if the designated constraints aren't met.
type SysConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysConfMultiError) AllErrors() []error { return m }

// SysConfValidationError is the validation error returned by SysConf.Validate
// if the designated constraints aren't met.
type SysConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysConfValidationError) ErrorName() string { return "SysConfValidationError" }

// Error satisfies the builtin error interface
func (e SysConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysConfValidationError{}

// Validate checks the field values on SysConf_Apisix with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SysConf_Apisix) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SysConf_Apisix with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SysConf_ApisixMultiError,
// or nil if none found.
func (m *SysConf_Apisix) ValidateAll() error {
	return m.validate(true)
}

func (m *SysConf_Apisix) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Endpoint

	// no validation rules for ApiKey

	{
		sorted_keys := make([]string, len(m.GetRoutes()))
		i := 0
		for key := range m.GetRoutes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetRoutes()[key]
			_ = val

			// no validation rules for Routes[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("Routes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("Routes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SysConf_ApisixValidationError{
						field:  fmt.Sprintf("Routes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetGlobalRules()))
		i := 0
		for key := range m.GetGlobalRules() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetGlobalRules()[key]
			_ = val

			// no validation rules for GlobalRules[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("GlobalRules[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("GlobalRules[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SysConf_ApisixValidationError{
						field:  fmt.Sprintf("GlobalRules[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetUpstreams()))
		i := 0
		for key := range m.GetUpstreams() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetUpstreams()[key]
			_ = val

			// no validation rules for Upstreams[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("Upstreams[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("Upstreams[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SysConf_ApisixValidationError{
						field:  fmt.Sprintf("Upstreams[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	{
		sorted_keys := make([]string, len(m.GetStreamRoutes()))
		i := 0
		for key := range m.GetStreamRoutes() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetStreamRoutes()[key]
			_ = val

			// no validation rules for StreamRoutes[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("StreamRoutes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, SysConf_ApisixValidationError{
							field:  fmt.Sprintf("StreamRoutes[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SysConf_ApisixValidationError{
						field:  fmt.Sprintf("StreamRoutes[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return SysConf_ApisixMultiError(errors)
	}

	return nil
}

// SysConf_ApisixMultiError is an error wrapping multiple validation errors
// returned by SysConf_Apisix.ValidateAll() if the designated constraints
// aren't met.
type SysConf_ApisixMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SysConf_ApisixMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SysConf_ApisixMultiError) AllErrors() []error { return m }

// SysConf_ApisixValidationError is the validation error returned by
// SysConf_Apisix.Validate if the designated constraints aren't met.
type SysConf_ApisixValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SysConf_ApisixValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SysConf_ApisixValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SysConf_ApisixValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SysConf_ApisixValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SysConf_ApisixValidationError) ErrorName() string { return "SysConf_ApisixValidationError" }

// Error satisfies the builtin error interface
func (e SysConf_ApisixValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSysConf_Apisix.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SysConf_ApisixValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SysConf_ApisixValidationError{}
