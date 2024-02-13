package apperrors

type BadRequest struct {
	ErrorMsg string
}

func (i BadRequest) Error() string {
	return i.ErrorMsg
}

type NoContent struct {
	ErrorMsg string
}

func (i NoContent) Error() string {
	return i.ErrorMsg
}

type UnAuthorizedAccess struct {
	ErrorMsg string
}

func (i UnAuthorizedAccess) Error() string {
	return i.ErrorMsg
}
