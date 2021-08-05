// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: testsvc/person/v1/person.proto

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

// Validate checks the field values on Person with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Person) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	return nil
}

// PersonValidationError is the validation error returned by Person.Validate if
// the designated constraints aren't met.
type PersonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersonValidationError) ErrorName() string { return "PersonValidationError" }

// Error satisfies the builtin error interface
func (e PersonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerson.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersonValidationError{}

// Validate checks the field values on CreatePersonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePersonRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return CreatePersonRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetEmail() != "" {

		if err := m._validateEmail(m.GetEmail()); err != nil {
			return CreatePersonRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *CreatePersonRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreatePersonRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreatePersonRequestValidationError is the validation error returned by
// CreatePersonRequest.Validate if the designated constraints aren't met.
type CreatePersonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePersonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePersonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePersonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePersonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePersonRequestValidationError) ErrorName() string {
	return "CreatePersonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePersonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePersonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePersonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePersonRequestValidationError{}

// Validate checks the field values on CreatePersonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePersonResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPerson()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePersonResponseValidationError{
				field:  "Person",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreatePersonResponseValidationError is the validation error returned by
// CreatePersonResponse.Validate if the designated constraints aren't met.
type CreatePersonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePersonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePersonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePersonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePersonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePersonResponseValidationError) ErrorName() string {
	return "CreatePersonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePersonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePersonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePersonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePersonResponseValidationError{}

// Validate checks the field values on GetPersonRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetPersonRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PersonId

	return nil
}

// GetPersonRequestValidationError is the validation error returned by
// GetPersonRequest.Validate if the designated constraints aren't met.
type GetPersonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPersonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPersonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPersonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPersonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPersonRequestValidationError) ErrorName() string { return "GetPersonRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPersonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPersonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPersonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPersonRequestValidationError{}

// Validate checks the field values on GetPersonResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetPersonResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPerson()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPersonResponseValidationError{
				field:  "Person",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetPersonResponseValidationError is the validation error returned by
// GetPersonResponse.Validate if the designated constraints aren't met.
type GetPersonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPersonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPersonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPersonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPersonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPersonResponseValidationError) ErrorName() string {
	return "GetPersonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPersonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPersonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPersonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPersonResponseValidationError{}