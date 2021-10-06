package v1

// DO NOT EDIT: this file is automatically generated by ./tools/generator in this repo

import (
	"encoding/json"
	"fmt"

	client_protocol "gitlab.com/signald/signald-go/signald/client-protocol"
)

func mkerr(response client_protocol.BasicResponse) error {
	switch response.ErrorType {
	case "AccountAlreadyVerifiedError":
		result := AccountAlreadyVerifiedError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "AccountHasNoKeysError":
		result := AccountHasNoKeysError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "AccountLockedError":
		result := AccountLockedError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "CaptchaRequiredError":
		result := CaptchaRequiredError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "FingerprintVersionMismatchError":
		result := FingerprintVersionMismatchError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "GroupLinkNotActiveError":
		result := GroupLinkNotActiveError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "GroupNotActiveError":
		result := GroupNotActiveError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "GroupVerificationError":
		result := GroupVerificationError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InternalError":
		result := InternalError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidAttachmentError":
		result := InvalidAttachmentError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidBase64Error":
		result := InvalidBase64Error{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidFingerprintError":
		result := InvalidFingerprintError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidGroupError":
		result := InvalidGroupError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidGroupStateError":
		result := InvalidGroupStateError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidInviteURIError":
		result := InvalidInviteURIError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidProxyError":
		result := InvalidProxyError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidRecipientError":
		result := InvalidRecipientError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "InvalidRequestError":
		result := InvalidRequestError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "NoKnownUUIDError":
		result := NoKnownUUIDError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "NoSendPermissionError":
		result := NoSendPermissionError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "NoSuchAccountError":
		result := NoSuchAccountError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "NoSuchSessionError":
		result := NoSuchSessionError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "OwnProfileKeyDoesNotExistError":
		result := OwnProfileKeyDoesNotExistError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "ProfileUnavailableError":
		result := ProfileUnavailableError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "ServerNotFoundError":
		result := ServerNotFoundError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "UnknownGroupError":
		result := UnknownGroupError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "UnknownIdentityKeyError":
		result := UnknownIdentityKeyError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "UntrustedIdentityError":
		result := UntrustedIdentityError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	case "UserAlreadyExistsError":
		result := UserAlreadyExistsError{}
		err := json.Unmarshal(response.Error, &result)
		if err != nil {
			return err
		}
		return result
	default:
		return fmt.Errorf("unexpected response type from signald: %s: %s", response.ErrorType, string(response.Error))
	}
}

type AccountAlreadyVerifiedError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e AccountAlreadyVerifiedError) Error() string {
	return e.Message
}

type AccountHasNoKeysError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e AccountHasNoKeysError) Error() string {
	return e.Message
}

type AccountLockedError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	More    string `json:"more,omitempty" yaml:"more,omitempty"`
}

func (e AccountLockedError) Error() string {
	return e.Message
}

type CaptchaRequiredError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	More    string `json:"more,omitempty" yaml:"more,omitempty"`
}

func (e CaptchaRequiredError) Error() string {
	return e.Message
}

type FingerprintVersionMismatchError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e FingerprintVersionMismatchError) Error() string {
	return e.Message
}

type GroupLinkNotActiveError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e GroupLinkNotActiveError) Error() string {
	return e.Message
}

type GroupNotActiveError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e GroupNotActiveError) Error() string {
	return e.Message
}

type GroupVerificationError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e GroupVerificationError) Error() string {
	return e.Message
}

// InternalError: an internal error in signald has occured.
type InternalError struct {
	Exceptions []string `json:"exceptions,omitempty" yaml:"exceptions,omitempty"`
	Message    string   `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InternalError) Error() string {
	return e.Message
}

type InvalidAttachmentError struct {
	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`
	Message  string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidAttachmentError) Error() string {
	return e.Message
}

type InvalidBase64Error struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidBase64Error) Error() string {
	return e.Message
}

type InvalidFingerprintError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidFingerprintError) Error() string {
	return e.Message
}

type InvalidGroupError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidGroupError) Error() string {
	return e.Message
}

type InvalidGroupStateError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidGroupStateError) Error() string {
	return e.Message
}

type InvalidInviteURIError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidInviteURIError) Error() string {
	return e.Message
}

type InvalidProxyError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidProxyError) Error() string {
	return e.Message
}

type InvalidRecipientError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidRecipientError) Error() string {
	return e.Message
}

type InvalidRequestError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e InvalidRequestError) Error() string {
	return e.Message
}

type NoKnownUUIDError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e NoKnownUUIDError) Error() string {
	return e.Message
}

type NoSendPermissionError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e NoSendPermissionError) Error() string {
	return e.Message
}

type NoSuchAccountError struct {
	Account string `json:"account,omitempty" yaml:"account,omitempty"`
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e NoSuchAccountError) Error() string {
	return e.Message
}

type NoSuchSessionError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e NoSuchSessionError) Error() string {
	return e.Message
}

type OwnProfileKeyDoesNotExistError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e OwnProfileKeyDoesNotExistError) Error() string {
	return e.Message
}

type ProfileUnavailableError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e ProfileUnavailableError) Error() string {
	return e.Message
}

type ServerNotFoundError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

func (e ServerNotFoundError) Error() string {
	return e.Message
}

type UnknownGroupError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e UnknownGroupError) Error() string {
	return e.Message
}

type UnknownIdentityKeyError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e UnknownIdentityKeyError) Error() string {
	return e.Message
}

type UntrustedIdentityError struct {
	Identifier  string        `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	IdentityKey **IdentityKey `json:"identity_key,omitempty" yaml:"identity_key,omitempty"`
	Message     string        `json:"message,omitempty" yaml:"message,omitempty"`
}

func (e UntrustedIdentityError) Error() string {
	return e.Message
}

type UserAlreadyExistsError struct {
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

func (e UserAlreadyExistsError) Error() string {
	return e.Message
}
