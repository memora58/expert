package service

import (
	"errors"
	"gateway/common/util"
)

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}
