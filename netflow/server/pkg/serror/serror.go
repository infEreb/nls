package serror

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

type ErrorMessage struct {
	Errs []error
}

func (e ErrorMessage) Error() string {
	s := ""
	for i, d := range e.Errs {
		if i == 0 {
			s += "Errors:\n"
		}

		s += d.Error()

		if i < len(e.Errs)-1 {
			s += "\n"
		}
	}

	return s
}