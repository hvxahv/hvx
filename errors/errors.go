package errors

import (
	"fmt"
)

const (
	ErrTokenInvalid      = "INVALID_TOKEN"
	ErrTokenUnauthorized = "TOKEN_UNAUTHORIZED"
)

// Err is returned account errors.
const (
	ErrAccountVerification          = "ACCOUNT_VERIFICATION_FAILED"
	ErrAccountAlready               = "THE_USERNAME_OR_MAIL_ALREADY_EXISTS"
	ErrAccountUsernameAlreadyExists = "THE_USERNAME_ALREADY_EXISTS"
	ErrAccountCreate                = "FAILED_TO_CREATE_ACCOUNT"
	ErrActorDelete                  = "FAILED_TO_DELETE_ACTOR"
	ErrAccountGetByUsername         = "FAILED_TO_GET_THE_ACCOUNT_BY_USERNAME"
)

// Err channel errors.
const (
	ErrChannelAlready               = "THE_CHANNEL_EXISTS"
	ErrNotAchannelAdministrator     = "NOT_A_CHANNEL_ADMINISTRATOR"
	ErrNotAchannelOwner             = "NOT_A_CHANNEL_OWNER"
	ErrAlreadyAchannelAdministrator = "ADMINISTRATOR_ALREADY_EXISTS"
	ErrNotTheOwner                  = "NOT_THE_OWNER"
	ErrAlreadySubscribed            = "ALREADY_SUBSCRIBED"
	ErrNotSubscribed                = "NOT_SUBSCRIBED"
	ErrDeleteChannelActor           = "CHANNEL_ACTOR_DELETION_FAILED"
)

const (
	ErrMatrixAccountRegister = "MATRIX_ACCOUNT_REGISTRATION_FAILED"
)

const (
	ErrNoPermission = "NO_PERMISSION"
)

func New(err string) error {
	return fmt.Errorf(err)
}

func Newf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
