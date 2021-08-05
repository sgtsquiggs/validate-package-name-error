// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: testsvc/test/v1/enums.proto

package testsvc_person_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on NestedEnumTestMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NestedEnumTestMessage) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	return nil
}

// NestedEnumTestMessageValidationError is the validation error returned by
// NestedEnumTestMessage.Validate if the designated constraints aren't met.
type NestedEnumTestMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NestedEnumTestMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NestedEnumTestMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NestedEnumTestMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NestedEnumTestMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NestedEnumTestMessageValidationError) ErrorName() string {
	return "NestedEnumTestMessageValidationError"
}

// Error satisfies the builtin error interface
func (e NestedEnumTestMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNestedEnumTestMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NestedEnumTestMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NestedEnumTestMessageValidationError{}

// Validate checks the field values on GetPersonTypeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPersonTypeRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PersonId

	return nil
}

// GetPersonTypeRequestValidationError is the validation error returned by
// GetPersonTypeRequest.Validate if the designated constraints aren't met.
type GetPersonTypeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPersonTypeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPersonTypeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPersonTypeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPersonTypeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPersonTypeRequestValidationError) ErrorName() string {
	return "GetPersonTypeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPersonTypeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPersonTypeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPersonTypeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPersonTypeRequestValidationError{}

// Validate checks the field values on GetPersonTypeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPersonTypeResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	return nil
}

// GetPersonTypeResponseValidationError is the validation error returned by
// GetPersonTypeResponse.Validate if the designated constraints aren't met.
type GetPersonTypeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPersonTypeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPersonTypeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPersonTypeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPersonTypeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPersonTypeResponseValidationError) ErrorName() string {
	return "GetPersonTypeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPersonTypeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPersonTypeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPersonTypeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPersonTypeResponseValidationError{}
