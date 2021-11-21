package response

import "github.com/garuda-hacks/go-entities-hackathon/entities/messageresponse/messages"

// Meta wrapper.
type Meta struct {
	messages.Message
	RequestID string `json:"request_id"`
}

type respMeta struct {
	// Teleconsultation response messages
	TelErrUserNotFound         messages.Message
	TelErrUserRoleExists       messages.Message
	TelErrNonActiveUser        messages.Message
	TelErrPassword             messages.Message
	TelErrUserSave             messages.Message
	TelErrCodeNotValid         messages.Message
	TelErrUserIsActive         messages.Message
	TelErrEmailAlreadyUsed     messages.Message
	TelErrUserNonActive        messages.Message
	TelErrRevocerRoute         messages.Message
	TelErrPageNotFound         messages.Message
	TelErrInvalidBind          messages.Message
	TelErrHandphoneAlreadyUsed messages.Message
	TelErrDataCantDecrypt      messages.Message
	TelErrShouldBindError      messages.Message
	TelErrTripleDes            messages.Message
	TelErrValidate             messages.Message
	TelErrVerifyEmail          messages.Message
	TelSuccessLogin            messages.Message
}

// RespMeta list of response message.
var RespMeta = respMeta{
	// Teleconsultation
	TelErrUserNotFound:         messages.TelErrUserNotFound,
	TelErrUserRoleExists:       messages.TelErrUserRoleExists,
	TelErrNonActiveUser:        messages.TelErrNonActiveUser,
	TelErrPassword:             messages.TelErrPassword,
	TelErrUserSave:             messages.TelErrPassword,
	TelErrCodeNotValid:         messages.TelErrCodeNotValid,
	TelErrUserIsActive:         messages.TelErrUserIsActive,
	TelErrEmailAlreadyUsed:     messages.TelErrEmailAlreadyUsed,
	TelErrUserNonActive:        messages.TelErrUserNonActive,
	TelErrRevocerRoute:         messages.TelErrRevocerRoute,
	TelErrPageNotFound:         messages.TelErrPageNotFound,
	TelErrInvalidBind:          messages.TelErrInvalidBind,
	TelErrHandphoneAlreadyUsed: messages.TelErrHandphoneAlreadyUsed,
	TelErrDataCantDecrypt:      messages.TelErrDataCantDecrypt,
	TelErrShouldBindError:      messages.TelErrShouldBindError,
	TelErrTripleDes:            messages.TelErrTripleDes,
	TelErrValidate:             messages.TelErrValidate,
	TelErrVerifyEmail:          messages.TelErrVerifyEmail,
	TelSuccessLogin:            messages.TelSuccessLogin,
}
