// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: eliza/v1/eliza.proto

package elizav1

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

// Validate checks the field values on SayRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SayRequestMultiError, or
// nil if none found.
func (m *SayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sentence

	if len(errors) > 0 {
		return SayRequestMultiError(errors)
	}

	return nil
}

// SayRequestMultiError is an error wrapping multiple validation errors
// returned by SayRequest.ValidateAll() if the designated constraints aren't met.
type SayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayRequestMultiError) AllErrors() []error { return m }

// SayRequestValidationError is the validation error returned by
// SayRequest.Validate if the designated constraints aren't met.
type SayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayRequestValidationError) ErrorName() string { return "SayRequestValidationError" }

// Error satisfies the builtin error interface
func (e SayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayRequestValidationError{}

// Validate checks the field values on SayResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SayResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SayResponseMultiError, or
// nil if none found.
func (m *SayResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SayResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sentence

	if len(errors) > 0 {
		return SayResponseMultiError(errors)
	}

	return nil
}

// SayResponseMultiError is an error wrapping multiple validation errors
// returned by SayResponse.ValidateAll() if the designated constraints aren't met.
type SayResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayResponseMultiError) AllErrors() []error { return m }

// SayResponseValidationError is the validation error returned by
// SayResponse.Validate if the designated constraints aren't met.
type SayResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayResponseValidationError) ErrorName() string { return "SayResponseValidationError" }

// Error satisfies the builtin error interface
func (e SayResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayResponseValidationError{}

// Validate checks the field values on ConverseRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConverseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConverseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConverseRequestMultiError, or nil if none found.
func (m *ConverseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConverseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sentence

	if len(errors) > 0 {
		return ConverseRequestMultiError(errors)
	}

	return nil
}

// ConverseRequestMultiError is an error wrapping multiple validation errors
// returned by ConverseRequest.ValidateAll() if the designated constraints
// aren't met.
type ConverseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConverseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConverseRequestMultiError) AllErrors() []error { return m }

// ConverseRequestValidationError is the validation error returned by
// ConverseRequest.Validate if the designated constraints aren't met.
type ConverseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConverseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConverseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConverseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConverseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConverseRequestValidationError) ErrorName() string { return "ConverseRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConverseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConverseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConverseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConverseRequestValidationError{}

// Validate checks the field values on ConverseResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConverseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConverseResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConverseResponseMultiError, or nil if none found.
func (m *ConverseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConverseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sentence

	if len(errors) > 0 {
		return ConverseResponseMultiError(errors)
	}

	return nil
}

// ConverseResponseMultiError is an error wrapping multiple validation errors
// returned by ConverseResponse.ValidateAll() if the designated constraints
// aren't met.
type ConverseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConverseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConverseResponseMultiError) AllErrors() []error { return m }

// ConverseResponseValidationError is the validation error returned by
// ConverseResponse.Validate if the designated constraints aren't met.
type ConverseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConverseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConverseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConverseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConverseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConverseResponseValidationError) ErrorName() string { return "ConverseResponseValidationError" }

// Error satisfies the builtin error interface
func (e ConverseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConverseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConverseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConverseResponseValidationError{}

// Validate checks the field values on IntroduceRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IntroduceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntroduceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntroduceRequestMultiError, or nil if none found.
func (m *IntroduceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IntroduceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return IntroduceRequestMultiError(errors)
	}

	return nil
}

// IntroduceRequestMultiError is an error wrapping multiple validation errors
// returned by IntroduceRequest.ValidateAll() if the designated constraints
// aren't met.
type IntroduceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntroduceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntroduceRequestMultiError) AllErrors() []error { return m }

// IntroduceRequestValidationError is the validation error returned by
// IntroduceRequest.Validate if the designated constraints aren't met.
type IntroduceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntroduceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntroduceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntroduceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntroduceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntroduceRequestValidationError) ErrorName() string { return "IntroduceRequestValidationError" }

// Error satisfies the builtin error interface
func (e IntroduceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntroduceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntroduceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntroduceRequestValidationError{}

// Validate checks the field values on IntroduceResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IntroduceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntroduceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntroduceResponseMultiError, or nil if none found.
func (m *IntroduceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IntroduceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sentence

	if len(errors) > 0 {
		return IntroduceResponseMultiError(errors)
	}

	return nil
}

// IntroduceResponseMultiError is an error wrapping multiple validation errors
// returned by IntroduceResponse.ValidateAll() if the designated constraints
// aren't met.
type IntroduceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntroduceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntroduceResponseMultiError) AllErrors() []error { return m }

// IntroduceResponseValidationError is the validation error returned by
// IntroduceResponse.Validate if the designated constraints aren't met.
type IntroduceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntroduceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntroduceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntroduceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntroduceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntroduceResponseValidationError) ErrorName() string {
	return "IntroduceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IntroduceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntroduceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntroduceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntroduceResponseValidationError{}
