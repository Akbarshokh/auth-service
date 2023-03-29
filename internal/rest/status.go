package rest

const (
	ErrCodeNoError    = 0   // Success code.
	ErrCodeValidation = -10 // For http 422.

	ErrCodeDocumentNotFound = -40

	ErrCodeInternalErr = -50
)

const (
	StatusMsgSuccess = "Success"
	StatusMsgFailure = "Failure"
)
