package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"nso/logging"
	"os"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func RFunc(f func()) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				logging.Logger.Error("Error occur when run func ", zap.Error(err.(error)))
			case string:
				logging.Logger.Error("Error occur when run func ", zap.String("error", err.(string)))
			case fmt.Stringer:
				logging.Logger.Error("Error occur when run func ", zap.String("error", err.(fmt.Stringer).String()))
			}
			debug.PrintStack()
		}
	}()
	f()
}

func REFunc(f func() error, errorMessage string) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				logging.Logger.Error("Error occur when run func ", zap.Error(err.(error)))
			case string:
				logging.Logger.Error("Error occur when run func ", zap.String("error", err.(string)))
			case fmt.Stringer:
				logging.Logger.Error("Error occur when run func ", zap.String("error", err.(fmt.Stringer).String()))
			}
			debug.PrintStack()
		}
	}()
	err := f()
	if err != nil {
		logging.Logger.Error(errorMessage, zap.Error(err))
	}
}

func ReadAllBytes(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		logging.Logger.Panic("Error occur when read file ", zap.String("path", path), zap.Error(err))
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logging.Logger.Panic("Error occur when read file ", zap.String("path", path), zap.Error(err))
	}
	return data
}

func ReadAll(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Escape string avoid sql injection
func Escape(s string) string {
	return strings.Replace(s, "'", "''", -1)
}

func TimeDay(days int) int64 {
	return time.Now().Add(time.Hour * 24 * time.Duration(days)).UnixMilli()
}

func TimeHour(hours int) int64 {
	return time.Now().Add(time.Hour * time.Duration(hours)).UnixMilli()
}

func TimeSeconds(expires int64) int64 {
	return time.Now().Add(time.Second * time.Duration(expires)).UnixMilli()
}

func ToString(i int) string {
	return strconv.Itoa(i)
}

func IsNil(interf interface{}) bool {
	if interf == nil {
		return true
	}
	switch reflect.TypeOf(interf).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(interf).IsNil()
	}
	return false
}

func Compare(a, b []byte) bool {
	result := bytes.Compare(a, b)
	if result == 0 {
		return true
	}
	return false
}

func Hash(s interface{}) []byte {
	data, err := json.Marshal(s)
	if err != nil {
		logging.Logger.Panic("Error occur when encode ", zap.Error(err))
	}
	var r = sha256.Sum256(data)
	return r[:]
}
