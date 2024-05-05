package locales

import "embed"

//go:embed en.yaml zh.yaml

var Locales embed.FS

const (
	JWTMissingToken           = "jwt.token.missing"
	JWTTokenInvalid           = "jwt.token.invalid"
	JWTTokenExpired           = "jwt.token.expired"
	JWTTokenParseFail         = "jwt.token.parse.failed"
	JWTTokenSignFail          = "jwt.token.sign.failed"
	JWTUnSupportSigningMethod = "jwt.wrong.signing.method"
	IdempotentMissingToken    = "idempotent.token.missing"
	IdempotentTokenExpired    = "idempotent.token.invalid"
	UserListUnauthorized      = "user.list.unauthorized"
	UserOperationForbidden    = "user.operation.forbidden"
	UserAlreadyExists         = "user.exists"

	TooManyRequests    = "too.many.requests"
	DataNotChange      = "data.not.change"
	DuplicateField     = "duplicate.field"
	RecordNotFound     = "record.not.found"
	NoPermission       = "no.permission"
	IncorrectPassword  = "login.incorrect.password"
	SamePassword       = "login.same.password"
	InvalidCaptcha     = "login.invalid.captcha"
	LoginFailed        = "login.failed"
	UserLocked         = "login.user.locked"
	KeepLeastOntAction = "action.keep.least.one.action"
	DeleteYourself     = "user.delete.yourself"
)
