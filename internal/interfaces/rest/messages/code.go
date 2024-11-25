package messages

const (
	InternalServerError = 15002
	RedisError          = 15001

	ResourceNotFound    = 14000
	InactiveAccount     = 14001
	TooManyOTPRequests  = 14002
	InvalidOTPData      = 14003
	InvalidCredentials  = 14004
	IncorrectOTP        = 14005
	LoginDataIncorrect  = 14006
	PasswordIncorrect   = 14007
	Locked              = 14008
	TooManyRequests     = 14009
	MissingMobileNumber = 14010
	InvalidMobileNumber = 14011
	WrongData           = 14012
	MissedToken         = 14013
	InvalidToken        = 14014
	NoUser              = 14015
	InvalidUserData     = 14016
	InvalidNationalId   = 14017
	InvalidInputData    = 14018
	InvalidEmail        = 14019
	PasswordMismatch    = 14020
	InvalidNameOrFamily = 14021
	Unauthorized        = 14022
	DuplicatePermission = 14023
	ExistNationalId     = 14024
	BadRequest          = 14025

	SuccessAuth         = 12000
	UserExists          = 12001
	OtpSuccessfullySent = 12002
	Success             = 12003
	ProfileUpdated      = 12004
	ResetPassword       = 12005
	PermissionCreated   = 12006
)

type Data struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// Text creates a new Data instance with the given code and error message.
func Text(code int, err string) *Data {
	return &Data{
		Code: code,
		Text: err,
	}
}
