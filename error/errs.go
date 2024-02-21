package error

func NewServerError(err error) *Err {
	return &Err{ServerError, err}
}

func NewBadRequestError(err error) *Err {
	return &Err{BadRequest, err}
}
