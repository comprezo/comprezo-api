package apperror

type AppError struct {
	Err       error
	HTTPCode  int
	PublicMsg string
}

func (err AppError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.PublicMsg
}

func (e AppError) Unwrap() error {
	return e.Err
}

func New(err error, httpCode int, publicMsg string) AppError {
	return AppError{
		Err:       err,
		HTTPCode:  httpCode,
		PublicMsg: publicMsg,
	}
}
