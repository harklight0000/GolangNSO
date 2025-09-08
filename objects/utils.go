package objects

import (
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"math/rand"
	"nso/logging"
	. "nso/utils"
	"regexp"
	"strings"
	"time"
)

func CurrentTimeMillis() int64 {
	return time.Now().UnixMilli()
}

func isCaveMap(mapId int) bool {
	return (mapId >= 114 && mapId <= 116) ||
		(mapId >= 91 && mapId <= 97) ||
		(mapId >= 105 && mapId <= 109) ||
		(mapId >= 125 && mapId <= 128) ||
		(mapId >= 157 && mapId <= 159) ||
		(mapId == 162)
}

func isLdgtMap(id int) bool {
	return id >= 80 && id <= 90
}

func isGtcMap(id int) bool {
	return id >= 118 && id <= 124
}

func isChienTruongKeo(id int) bool {
	return id >= 130 && id <= 133
}

func isNvMap(id int) bool {
	return id == 56 || id == 73 || id == 0
}

func combineErrors(errors []error) error {
	sb := strings.Builder{}
	var hasErr = false
	for _, err := range errors {
		if err != nil {
			sb.WriteString(err.Error())
			sb.WriteString("\n")
			hasErr = true
		}
	}
	if hasErr {
		return eris.New(sb.String())
	} else {
		return nil
	}
}

func Abs[N Number](x N) N {
	if x < 0 {
		return -x
	}
	return x
}

func nextTime(t time.Duration) int64 {
	return time.Now().Add(t).UnixMilli()
}

func cast[N Number](i interface{}) N {
	switch i.(type) {
	case int:
		return N(i.(int))
	case int64:
		return N(i.(int64))
	case int32:
		return N(i.(int32))
	case int16:
		return N(i.(int16))
	case int8:
		return N(i.(int8))
	case float32:
		return N(i.(float32))
	case float64:
		return N(i.(float64))
	case byte:
		return N(i.(byte))
	default:
		logging.Logger.Panic("Unsupport number type ", zap.Any("type", i))
	}
	return 0
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func checkRegex(str string, regex string) bool {
	match, err := regexp.Match(regex, []byte(str))
	if err != nil {
		logging.Logger.Warn("Unknown regex ", zap.String("regex", regex))
	}
	return match
}
