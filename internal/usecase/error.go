package usecase

const (
	ErrorBadRequest          int = 400
	ErrorUnauthorized        int = 403
	ErrorNotFound            int = 404
	ErrorDuplicate           int = 409
	ErrorUnprocessableEntity int = 422
)

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message string, code int) Error {
	return Error{
		Message: message,
		Code:    code,
	}
}

func NewParameterError(msg *string) Error {
	defaultMessage := "invalid parameter"
	if msg == nil {
		msg = &defaultMessage
	}
	return NewError(*msg, ErrorUnprocessableEntity)
}

func NewNotFoundError() Error {
	return NewError("resource not found", ErrorNotFound)
}

func NewDuplicateError() Error {
	return NewError("resource already exists", ErrorDuplicate)
}

func NewUnauthorizedError() Error {
	return NewError("unauthorized access", ErrorUnauthorized)
}

func NewLimitReachedError() Error {
	return NewError("reachout limit reached", ErrorBadRequest)
}

func IsDuplicateError(e error) bool {
	internalErr, isInternalErr := e.(Error)
	if !isInternalErr {
		return false
	}

	return internalErr.Code == ErrorDuplicate
}

func IsNotFoundError(e error) bool {
	internalErr, isInternalErr := e.(Error)
	if !isInternalErr {
		return false
	}

	return internalErr.Code == ErrorNotFound
}

func IsUnprocessableEntityError(e error) bool {
	internalErr, isInternalErr := e.(Error)
	if !isInternalErr {
		return false
	}

	return internalErr.Code == ErrorUnprocessableEntity
}
