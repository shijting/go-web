package common

import (
	"github.com/go-playground/validator/v10"
	myValidator "github.com/shijting/go-web/libs/validator"
	"go.uber.org/zap"
)

type Response struct {
	Code int
	Msg  interface{}
	Data interface{}
}

func ValidateError(err error) (resp *Response) {
	zap.L().Error("SignUp with invalid param", zap.Error(err))
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		resp = &Response{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		}
		return
	}
	resp = &Response{
		Code: 400,
		Msg:  myValidator.RemoveTopStruct(errs.Translate(myValidator.Trans)),
		Data: nil,
	}
	return
}
