// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: woa.proto

package woa

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

// Validate checks the field values on CheckSignatureReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CheckSignatureReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckSignatureReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckSignatureReqMultiError, or nil if none found.
func (m *CheckSignatureReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckSignatureReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Signature

	// no validation rules for Timestamp

	// no validation rules for Nonce

	// no validation rules for Echostr

	if len(errors) > 0 {
		return CheckSignatureReqMultiError(errors)
	}

	return nil
}

// CheckSignatureReqMultiError is an error wrapping multiple validation errors
// returned by CheckSignatureReq.ValidateAll() if the designated constraints
// aren't met.
type CheckSignatureReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckSignatureReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckSignatureReqMultiError) AllErrors() []error { return m }

// CheckSignatureReqValidationError is the validation error returned by
// CheckSignatureReq.Validate if the designated constraints aren't met.
type CheckSignatureReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckSignatureReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckSignatureReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckSignatureReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckSignatureReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckSignatureReqValidationError) ErrorName() string {
	return "CheckSignatureReqValidationError"
}

// Error satisfies the builtin error interface
func (e CheckSignatureReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckSignatureReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckSignatureReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckSignatureReqValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Message) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MessageMultiError, or nil if none found.
func (m *Message) ValidateAll() error {
	return m.validate(true)
}

func (m *Message) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Text

	if len(errors) > 0 {
		return MessageMultiError(errors)
	}

	return nil
}

// MessageMultiError is an error wrapping multiple validation errors returned
// by Message.ValidateAll() if the designated constraints aren't met.
type MessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageMultiError) AllErrors() []error { return m }

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on ReceiveMessageReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReceiveMessageReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveMessageReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveMessageReqMultiError, or nil if none found.
func (m *ReceiveMessageReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveMessageReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ToUserName

	// no validation rules for FromUserName

	// no validation rules for CreateTime

	// no validation rules for MsgType

	// no validation rules for Content

	// no validation rules for MsgId

	// no validation rules for MsgDataId

	// no validation rules for Idx

	if len(errors) > 0 {
		return ReceiveMessageReqMultiError(errors)
	}

	return nil
}

// ReceiveMessageReqMultiError is an error wrapping multiple validation errors
// returned by ReceiveMessageReq.ValidateAll() if the designated constraints
// aren't met.
type ReceiveMessageReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveMessageReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveMessageReqMultiError) AllErrors() []error { return m }

// ReceiveMessageReqValidationError is the validation error returned by
// ReceiveMessageReq.Validate if the designated constraints aren't met.
type ReceiveMessageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveMessageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveMessageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveMessageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveMessageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveMessageReqValidationError) ErrorName() string {
	return "ReceiveMessageReqValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveMessageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveMessageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveMessageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveMessageReqValidationError{}

// Validate checks the field values on ReceiveMessageRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReceiveMessageRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveMessageRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveMessageRspMultiError, or nil if none found.
func (m *ReceiveMessageRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveMessageRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ToUserName

	// no validation rules for FromUserName

	// no validation rules for CreateTime

	// no validation rules for MsgType

	// no validation rules for Content

	if len(errors) > 0 {
		return ReceiveMessageRspMultiError(errors)
	}

	return nil
}

// ReceiveMessageRspMultiError is an error wrapping multiple validation errors
// returned by ReceiveMessageRsp.ValidateAll() if the designated constraints
// aren't met.
type ReceiveMessageRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveMessageRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveMessageRspMultiError) AllErrors() []error { return m }

// ReceiveMessageRspValidationError is the validation error returned by
// ReceiveMessageRsp.Validate if the designated constraints aren't met.
type ReceiveMessageRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveMessageRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveMessageRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveMessageRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveMessageRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveMessageRspValidationError) ErrorName() string {
	return "ReceiveMessageRspValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveMessageRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveMessageRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveMessageRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveMessageRspValidationError{}
