package utils

import (
	"go.uber.org/zap"
	"nso/logging"
	"strconv"
)

func ParseInt(str string) int {
	i, er := strconv.Atoi(str)
	if er != nil {
		logging.Logger.Panic("Error parser int of "+str, zap.Error(er))
	}
	return i
}
