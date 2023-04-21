package agent

type ErrorAlreadyExists struct {
	Err error
}

func (e ErrorAlreadyExists) Error() string {
	return e.Err.Error()
}

type ErrorNotFound struct {
	Err error
}

func (e ErrorNotFound) Error() string {
	return e.Err.Error()
}

type ErrorUnexpectedType struct {
	Err error
}

func (e ErrorUnexpectedType) Error() string {
	return e.Err.Error()
}

type ErrorUndefinedField struct {
	Err error
}

func (e ErrorUndefinedField) Error() string {
	return e.Err.Error()
}