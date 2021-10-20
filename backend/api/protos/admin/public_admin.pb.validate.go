// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/public_admin.proto

package admin

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
var _public_admin_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateUrlReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateUrlReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Username

	// no validation rules for Passwd

	if utf8.RuneCountInString(m.GetUri()) < 1 {
		return CreateUrlReqValidationError{
			field:  "Uri",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetSrvUri()) < 1 {
		return CreateUrlReqValidationError{
			field:  "SrvUri",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// CreateUrlReqValidationError is the validation error returned by
// CreateUrlReq.Validate if the designated constraints aren't met.
type CreateUrlReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUrlReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUrlReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUrlReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUrlReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUrlReqValidationError) ErrorName() string { return "CreateUrlReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateUrlReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUrlReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUrlReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUrlReqValidationError{}
