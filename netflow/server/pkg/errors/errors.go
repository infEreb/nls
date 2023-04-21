package errors

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