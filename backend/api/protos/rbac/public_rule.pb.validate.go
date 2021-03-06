// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rbac/public_rule.proto

package rbac

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

// define the regex for a UUID once up-front
var _public_rule_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetRuleReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetRuleReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPage() <= 0 {
		return GetRuleReqValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
	}

	if m.GetPageSize() <= 0 {
		return GetRuleReqValidationError{
			field:  "PageSize",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// GetRuleReqValidationError is the validation error returned by
// GetRuleReq.Validate if the designated constraints aren't met.
type GetRuleReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRuleReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRuleReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRuleReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRuleReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRuleReqValidationError) ErrorName() string { return "GetRuleReqValidationError" }

// Error satisfies the builtin error interface
func (e GetRuleReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRuleReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRuleReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRuleReqValidationError{}

// Validate checks the field values on GetRuleResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetRuleResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetRuleRespValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetRuleRespValidationError is the validation error returned by
// GetRuleResp.Validate if the designated constraints aren't met.
type GetRuleRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRuleRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRuleRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRuleRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRuleRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRuleRespValidationError) ErrorName() string { return "GetRuleRespValidationError" }

// Error satisfies the builtin error interface
func (e GetRuleRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRuleResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRuleRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRuleRespValidationError{}

// Validate checks the field values on SetDetailReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SetDetailReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return SetDetailReqValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if l := utf8.RuneCountInString(m.GetDetail()); l < 1 || l > 255 {
		return SetDetailReqValidationError{
			field:  "Detail",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	return nil
}

// SetDetailReqValidationError is the validation error returned by
// SetDetailReq.Validate if the designated constraints aren't met.
type SetDetailReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetDetailReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetDetailReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetDetailReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetDetailReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetDetailReqValidationError) ErrorName() string { return "SetDetailReqValidationError" }

// Error satisfies the builtin error interface
func (e SetDetailReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetDetailReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetDetailReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetDetailReqValidationError{}
