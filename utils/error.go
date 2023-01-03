package utils

import (
	"chat/constant"
)

func Error(err error, info string) {
	if err != nil {
		Logger(constant.Error, map[string]string{
			"mes": err.Error(),
		})
		panic(info + err.Error())
	}
}
