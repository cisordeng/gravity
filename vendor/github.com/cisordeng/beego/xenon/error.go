package xenon

type Error struct {
	Business BusinessError
	Inner    error
}

type BusinessError struct {
	ErrCode string
	ErrMsg  string
}

func NewBusinessError(ErrCode string, ErrMsg string) BusinessError {
	return BusinessError{
		ErrCode: ErrCode,
		ErrMsg:  ErrMsg,
	}
}

func RaiseError(ctx *Ctx, err error, businessError ...BusinessError) {
	business := NewBusinessError("", "")
	if len(businessError) > 0 {
		business = businessError[0]
	}
	ctx.Errors = append(ctx.Errors, Error{
		Business: business,
		Inner:    err,
	})
}
