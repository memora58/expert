package util

import (
	"errors"
	"gateway/common/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		return errors.New(removeTopStruct(errs.Translate(global.Trans)))
	}

	return nil
}

func removeTopStruct(fields map[string]string) (result string) {
	for _, err := range fields {
		result += err + ","
	}
	return result
}

func GetValidUriParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindUri(params); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		return errors.New(removeTopStruct(errs.Translate(global.Trans)))
	}

	return nil
}

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		LogrusObj.Info(err)
		panic(err)
	}
}
