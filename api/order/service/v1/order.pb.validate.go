// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/order/service/v1/order.proto

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

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRequestMultiError, or nil if none found.
func (m *CreateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrderRequestMultiError(errors)
	}

	return nil
}

// CreateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRequestMultiError) AllErrors() []error { return m }

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on CreateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderReplyMultiError, or nil if none found.
func (m *CreateOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrderReplyMultiError(errors)
	}

	return nil
}

// CreateOrderReplyMultiError is an error wrapping multiple validation errors
// returned by CreateOrderReply.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderReplyMultiError) AllErrors() []error { return m }

// CreateOrderReplyValidationError is the validation error returned by
// CreateOrderReply.Validate if the designated constraints aren't met.
type CreateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyValidationError) ErrorName() string { return "CreateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyValidationError{}

// Validate checks the field values on UpdateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateOrderRequestMultiError, or nil if none found.
func (m *UpdateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateOrderRequestMultiError(errors)
	}

	return nil
}

// UpdateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateOrderRequestMultiError) AllErrors() []error { return m }

// UpdateOrderRequestValidationError is the validation error returned by
// UpdateOrderRequest.Validate if the designated constraints aren't met.
type UpdateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateOrderRequestValidationError) ErrorName() string {
	return "UpdateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateOrderRequestValidationError{}

// Validate checks the field values on UpdateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateOrderReplyMultiError, or nil if none found.
func (m *UpdateOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateOrderReplyMultiError(errors)
	}

	return nil
}

// UpdateOrderReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateOrderReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateOrderReplyMultiError) AllErrors() []error { return m }

// UpdateOrderReplyValidationError is the validation error returned by
// UpdateOrderReply.Validate if the designated constraints aren't met.
type UpdateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateOrderReplyValidationError) ErrorName() string { return "UpdateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateOrderReplyValidationError{}

// Validate checks the field values on DeleteOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteOrderRequestMultiError, or nil if none found.
func (m *DeleteOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteOrderRequestMultiError(errors)
	}

	return nil
}

// DeleteOrderRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteOrderRequestMultiError) AllErrors() []error { return m }

// DeleteOrderRequestValidationError is the validation error returned by
// DeleteOrderRequest.Validate if the designated constraints aren't met.
type DeleteOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrderRequestValidationError) ErrorName() string {
	return "DeleteOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrderRequestValidationError{}

// Validate checks the field values on DeleteOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteOrderReplyMultiError, or nil if none found.
func (m *DeleteOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteOrderReplyMultiError(errors)
	}

	return nil
}

// DeleteOrderReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteOrderReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteOrderReplyMultiError) AllErrors() []error { return m }

// DeleteOrderReplyValidationError is the validation error returned by
// DeleteOrderReply.Validate if the designated constraints aren't met.
type DeleteOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrderReplyValidationError) ErrorName() string { return "DeleteOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrderReplyValidationError{}

// Validate checks the field values on GetOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderRequestMultiError, or nil if none found.
func (m *GetOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetOrderRequestMultiError(errors)
	}

	return nil
}

// GetOrderRequestMultiError is an error wrapping multiple validation errors
// returned by GetOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type GetOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderRequestMultiError) AllErrors() []error { return m }

// GetOrderRequestValidationError is the validation error returned by
// GetOrderRequest.Validate if the designated constraints aren't met.
type GetOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderRequestValidationError) ErrorName() string { return "GetOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderRequestValidationError{}

// Validate checks the field values on GetOrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetOrderReplyMultiError, or
// nil if none found.
func (m *GetOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetOrderReplyMultiError(errors)
	}

	return nil
}

// GetOrderReplyMultiError is an error wrapping multiple validation errors
// returned by GetOrderReply.ValidateAll() if the designated constraints
// aren't met.
type GetOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderReplyMultiError) AllErrors() []error { return m }

// GetOrderReplyValidationError is the validation error returned by
// GetOrderReply.Validate if the designated constraints aren't met.
type GetOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderReplyValidationError) ErrorName() string { return "GetOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderReplyValidationError{}

// Validate checks the field values on ListOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrderRequestMultiError, or nil if none found.
func (m *ListOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListOrderRequestMultiError(errors)
	}

	return nil
}

// ListOrderRequestMultiError is an error wrapping multiple validation errors
// returned by ListOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type ListOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderRequestMultiError) AllErrors() []error { return m }

// ListOrderRequestValidationError is the validation error returned by
// ListOrderRequest.Validate if the designated constraints aren't met.
type ListOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderRequestValidationError) ErrorName() string { return "ListOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderRequestValidationError{}

// Validate checks the field values on ListOrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListOrderReplyMultiError,
// or nil if none found.
func (m *ListOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListOrderReplyMultiError(errors)
	}

	return nil
}

// ListOrderReplyMultiError is an error wrapping multiple validation errors
// returned by ListOrderReply.ValidateAll() if the designated constraints
// aren't met.
type ListOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderReplyMultiError) AllErrors() []error { return m }

// ListOrderReplyValidationError is the validation error returned by
// ListOrderReply.Validate if the designated constraints aren't met.
type ListOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderReplyValidationError) ErrorName() string { return "ListOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderReplyValidationError{}