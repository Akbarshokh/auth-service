package errs

var (
	ErrValidation    = New("validation error")
	ErrInternal      = New("internal error")
	ErrNotFound      = New("entity not found")
	ErrAuthorization = New("authorization failed")

	ErrTokenEmpty = Errf(ErrValidation, "empty token")
)
