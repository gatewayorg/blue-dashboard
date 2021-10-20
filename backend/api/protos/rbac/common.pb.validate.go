// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rbac/common.proto

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
var _common_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Role) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Detail

	if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Enable

	return nil
}

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on Rule with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Rule) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Method

	// no validation rules for Path

	return nil
}

// RuleValidationError is the validation error returned by Rule.Validate if the
// designated constraints aren't met.
type RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuleValidationError) ErrorName() string { return "RuleValidationError" }

// Error satisfies the builtin error interface
func (e RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuleValidationError{}

// Validate checks the field values on RoleRule with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RoleRule) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoleId

	return nil
}

// RoleRuleValidationError is the validation error returned by
// RoleRule.Validate if the designated constraints aren't met.
type RoleRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleRuleValidationError) ErrorName() string { return "RoleRuleValidationError" }

// Error satisfies the builtin error interface
func (e RoleRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoleRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleRuleValidationError{}
